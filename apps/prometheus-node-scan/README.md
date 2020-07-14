# prometheus-node-scan

This is a setup of [Prometheus](https://prometheus.io/) and the small app `node-scan` (present in this folder), that will scan periodically my network to find my nodes that expose metrics with [node_exporter](https://github.com/prometheus/node_exporter) on http://<host>:9100/metrics.

## deploy

TODO: add volume

```sh
docker run -d --net host --name prometheus-node-scan --restart=unless-stopped rafaeleyng/prometheus-node-scan
```

---
