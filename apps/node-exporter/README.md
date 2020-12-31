# prometheus node-exporter

## install on armv6 with raspbian

```sh
curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/node-exporter/install-armv6.sh --output node-exporter-install.sh
chmod 755 node-exporter-install.sh
sudo ./node-exporter-install.sh
rm ./node-exporter-install.sh
```

## install on armv7 with raspbian

```sh
curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/node-exporter/install-armv7.sh --output node-exporter-install.sh
chmod 755 node-exporter-install.sh
sudo ./node-exporter-install.sh
rm ./node-exporter-install.sh
```

## install on mac

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
