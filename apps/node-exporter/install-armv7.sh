#!/bin/sh

wget https://github.com/prometheus/node_exporter/releases/download/v1.0.1/node_exporter-1.0.1.linux-armv7.tar.gz
tar xvfz node_exporter-1.0.1.linux-armv7.tar.gz
mv ./node_exporter-1.0.1.linux-armv7/node_exporter /usr/local/bin
rm -fr node_exporter-1.0.1.linux-armv7*

# https://linuxhit.com/prometheus-node-exporter-on-raspberry-pi-how-to-install/
useradd -m -s /bin/bash node_exporter
mkdir /var/lib/node_exporter
chown -R node_exporter:node_exporter /var/lib/node_exporter

cat <<EOT >> /etc/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter

[Service]
# Provide a text file location for https://github.com/fahlke/raspberrypi_exporter data with the
# --collector.textfile.directory parameter.
ExecStart=/usr/local/bin/node_exporter --collector.textfile.directory /var/lib/node_exporter/textfile_collector

[Install]
WantedBy=multi-user.target
EOT

# https://github.com/fahlke/raspberrypi_exporter
curl -fsSL "https://raw.githubusercontent.com/fahlke/raspberrypi_exporter/master/installer.sh" | sudo bash

systemctl daemon-reload
systemctl enable node_exporter.service
systemctl start node_exporter.service

systemctl status node_exporter
