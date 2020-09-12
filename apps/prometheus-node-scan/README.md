# prometheus-node-scan

This is a setup of [Prometheus](https://prometheus.io/) and the small app `node-scan` (present in this folder), that will scan periodically my network to find my nodes that expose metrics with [node_exporter](https://github.com/prometheus/node_exporter) on http://<host>:9100/metrics.

## deploy

TODO: add volume

1. ensure I have the updated coredns image published to Docker Hub:
  ```sh
  make docker-build-and-push
  ```

2. run on the node:
  ```sh
  docker run \
    --name prometheus-node-scan \
    --net host \
    --restart=unless-stopped \
    -d \
    rafaeleyng/prometheus-node-scan
    ```

---
