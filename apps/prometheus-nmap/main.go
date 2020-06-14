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

const dnsZone = "cluster.rafael"
const nodeExporterJob = "node-exporter"
const targetMetric = "node_uname_info"
const targetLabel = "nodename"

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

func getNodeName(metricFamiliesMap map[string]*dto.MetricFamily) (string, error) {
	metricFamily, ok := metricFamiliesMap[targetMetric]
	if !ok {
		return "", fmt.Errorf("metric %s (which contains the nodename label) was not found", targetMetric)
	}

	metric := metricFamily.GetMetric()[0]

	for _, labelPair := range metric.GetLabel() {
		if *labelPair.Name == targetLabel {
			return *labelPair.Value, nil
		}
	}

	return "", fmt.Errorf("label %s was not found", targetLabel)
}

func buildNodes(scanResult *nmap.Run) ([]Node, error) {
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
		nodeName, err := getNodeName(metricFamiliesMap)

		nodes = append(nodes, Node{Name: nodeName})
	}

	return nodes, nil
}

func buildZonedHost(hostname string) string {
	return fmt.Sprintf("%s.%s", hostname, dnsZone)
}

func buildJobName(hostname string) string {
	return fmt.Sprintf("%s-%s", hostname, nodeExporterJob)
}

func createJobsFromNodes(nodes []Node) []Job {
	jobs := []Job{}

	for _, node := range nodes {
		job := Job{
			Targets: []string{buildZonedHost(node.Name)},
			Labels: map[string]string{
				"job": buildJobName(node.Name),
			},
		}

		jobs = append(jobs, job)
	}

	return jobs
}

func writeFileSD(nodes []Node, outputPath string) error {
	jobs := createJobsFromNodes(nodes)

	file, err := os.Create(outputPath)
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

func main() {
	var outputPath = ""
	flag.StringVar(&outputPath, "output", "/tmp/nodes.json", "")
	flag.Parse()

	scanResult, err := scanNetwork()
	if err != nil {
		fmt.Printf("failed to scan network: %v", err)
		return
	}

	nodes, err := buildNodes(scanResult)
	if err != nil {
		fmt.Printf("failed to build nodes: %v", err)
		return
	}

	err = writeFileSD(nodes, outputPath)
	if err != nil {
		fmt.Printf("failed to write file SD: %v", err)
		return
	}

	fmt.Printf("did write file SD: %d nodes found (%v) in %3f seconds\n", len(scanResult.Hosts), nodes, scanResult.Stats.Finished.Elapsed)
}
