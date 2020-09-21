# prometheus-node-scan

This is a setup of [Prometheus](https://prometheus.io/) and the small app `node-scan` (present in this folder), that will scan periodically my network to find my nodes that expose metrics with [node_exporter](https://github.com/prometheus/node_exporter) on http://<host>:9100/metrics.

## deploy

1. [on the node] create a volume (only once):
  ```sh
  docker volume create prometheus
  ```

2. [on the control plane] ensure I have the updated prometheus-node-scan image published to Docker Hub:
  ```sh
  make docker-build-and-push
  ```

3. [on the node] run:
  ```sh
  docker run \
    --name prometheus-node-scan \
    --net host \
    --restart=unless-stopped \
    --volume prometheus:/prometheus \
    -d \
    rafaeleyng/prometheus-node-scan
    ```

---
