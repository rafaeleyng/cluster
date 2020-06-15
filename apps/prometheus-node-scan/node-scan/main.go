package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ullaakut/nmap"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

type Options struct {
	dnsZone      string
	jobPrefix    string
	outputPath   string
	targetLabel  string
	targetMetric string
}

type Node struct {
	Name string
}

type Job struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func buildMetricsURL(host string) string {
	return fmt.Sprintf("http://%s:9100/metrics", host)
}

func buildScanner(ctx context.Context) (*nmap.Scanner, error) {
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("192.168.1.0/24"),
		nmap.WithPorts("9100"),
		nmap.WithContext(ctx),

		nmap.WithFilterHost(func(host nmap.Host) bool {
			port := host.Ports[0]
			if port.ID != 9100 || port.State.State != "open" {
				return false
			}

			metricsURL := buildMetricsURL(host.Addresses[0].Addr)
			resp, err := http.Get(metricsURL)
			if err != nil {
				return false
			}
			defer resp.Body.Close()

			return resp.StatusCode == 200
		}),
	)

	if err != nil {
		return nil, err
	}

	return scanner, nil
}

func runScanner(scanner *nmap.Scanner) (*nmap.Run, error) {
	result, warnings, err := scanner.Run()
	if err != nil {
		return nil, err
	}

	if warnings != nil {
		fmt.Printf("Warnings: \n %v", warnings)
	}

	return result, nil
}

func scanNetwork() (*nmap.Run, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	scanner, err := buildScanner(ctx)
	if err != nil {
		return nil, err
	}

	result, err := runScanner(scanner)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getNodeName(metricFamiliesMap map[string]*dto.MetricFamily, options Options) (string, error) {
	metricFamily, ok := metricFamiliesMap[options.targetMetric]
	if !ok {
		return "", fmt.Errorf("metric %s (which contains the nodename label) was not found", options.targetMetric)
	}

	metric := metricFamily.GetMetric()[0]

	for _, labelPair := range metric.GetLabel() {
		if *labelPair.Name == options.targetLabel {
			return *labelPair.Value, nil
		}
	}

	return "", fmt.Errorf("label %s was not found", options.targetLabel)
}

func buildNodes(scanResult *nmap.Run, options Options) ([]Node, error) {
	nodes := []Node{}

	for _, host := range scanResult.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		metricsURL := buildMetricsURL(host.Addresses[0].Addr)
		resp, err := http.Get(metricsURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("failed to read metrics for %s", metricsURL)
		}

		textParser := expfmt.TextParser{}
		metricFamiliesMap, err := textParser.TextToMetricFamilies(resp.Body)
		if err != nil {
			return nil, err
		}
		nodeName, err := getNodeName(metricFamiliesMap, options)

		nodes = append(nodes, Node{Name: nodeName})
	}

	return nodes, nil
}

func buildZonedHost(hostname string, options Options) string {
	return fmt.Sprintf("%s.%s:9100", hostname, options.dnsZone)
}

func buildJobName(hostname string, options Options) string {
	return fmt.Sprintf("%s-%s", options.jobPrefix, hostname)
}

func createJobsFromNodes(nodes []Node, options Options) []Job {
	jobs := []Job{}

	for _, node := range nodes {
		job := Job{
			Targets: []string{buildZonedHost(node.Name, options)},
			Labels: map[string]string{
				"job": buildJobName(node.Name, options),
			},
		}

		jobs = append(jobs, job)
	}

	return jobs
}

func writeFileSD(nodes []Node, options Options) error {
	jobs := createJobsFromNodes(nodes, options)

	file, err := os.Create(options.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jobs)
	if err != nil {
		return err
	}

	return nil
}

func watch(options Options) {
	for {
		fmt.Println("[node-scan] watch")
		scanResult, err := scanNetwork()
		if err != nil {
			fmt.Printf("failed to scan network: %v", err)
			time.Sleep(30 * time.Second)
			continue
		}

		nodes, err := buildNodes(scanResult, options)
		if err != nil {
			fmt.Printf("failed to build nodes: %v", err)
			time.Sleep(30 * time.Second)
			continue
		}

		err = writeFileSD(nodes, options)
		if err != nil {
			fmt.Printf("failed to write file SD: %v", err)
			time.Sleep(30 * time.Second)
			continue
		}

		fmt.Printf("did write file SD: %d nodes found (%v) in %3f seconds\n", len(scanResult.Hosts), nodes, scanResult.Stats.Finished.Elapsed)
		time.Sleep(60 * time.Second)
	}
}

func main() {
	fmt.Println("[node-scan] starting")
	options := Options{}

	flag.StringVar(&options.dnsZone, "zone", "cluster.rafael", "")
	flag.StringVar(&options.jobPrefix, "job-prefix", "node-exporter", "")
	flag.StringVar(&options.outputPath, "output", "/tmp/nodes.json", "")
	flag.StringVar(&options.targetLabel, "label", "nodename", "")
	flag.StringVar(&options.targetMetric, "metric", "node_uname_info", "")
	flag.Parse()

	watch(options)
}
