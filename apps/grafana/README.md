# grafana

Default credentials are `admin`/`admin`.

## deploy

1. [on the node] create a volume (only once):
  ```sh
  docker volume create grafana
  ```

2. [on the node] run (note that it uses the `dns-net` network, so the configuration for coredns must be already done):
  ```sh
  docker run \
    --dns=172.18.0.3 \
    --ip 172.18.0.10 \
    --name grafana \
    --network dns-net \
    --restart=unless-stopped \
    --volume grafana:/var/lib/grafana \
    -d \
    -e "GF_INSTALL_PLUGINS=grafana-piechart-panel" \
    -p 3000:3000 \
    grafana/grafana:7.1.5
  ```

3. configure Prometheus datasource

4. configure dashboards:
  - https://grafana.com/grafana/dashboards/1860
  - https://grafana.com/grafana/dashboards/12539
  - https://grafana.com/grafana/dashboards/10176 (depends on `grafana-piechart-panel`)

---
