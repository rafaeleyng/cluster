# prometheus node-exporter

## install on raspbian nodes

```sh
curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/node-exporter/install-raspbian.sh --output node-exporter-install.sh
chmod 755 node-exporter-install.sh
sudo ./node-exporter-install.sh
rm ./node-exporter-install.sh
```

## install on mac nodes

```sh
curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/node-exporter/install-mac.sh --output node-exporter-install.sh
chmod 755 node-exporter-install.sh
sudo ./node-exporter-install.sh
rm ./node-exporter-install.sh
```


## references

- https://prometheus.io/docs/guides/node-exporter/#monitoring-linux-host-metrics-with-the-node-exporter
- https://prometheus.io/download/#node_exporter
- https://prometheus.io/docs/guides/node-exporter/#installing-and-running-the-node-exporter
- https://linuxhit.com/prometheus-node-exporter-on-raspberry-pi-how-to-install/
- https://github.com/galarzafrancisco/home-monitoring/tree/master/client

---
