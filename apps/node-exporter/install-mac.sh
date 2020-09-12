#!/bin/sh

wget https://github.com/prometheus/node_exporter/releases/download/v1.0.1/node_exporter-1.0.1.darwin-amd64.tar.gz
tar xvfz node_exporter-1.0.1.darwin-amd64.tar.gz
mv ./node_exporter-1.0.1.darwin-amd64/node_exporter /usr/local/bin
rm -fr node_exporter-1.0.1.darwin-amd64*

cat <<EOT >> /Library/LaunchAgents/local.node_exporter.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
      <key>Label</key>
          <string>local.node_exporter</string>

      <key>Program</key>
          <string>/usr/local/bin/node_exporter</string>

      <key>RunAtLoad</key>
          <true/>

      <key>KeepAlive</key>
          <true/>

  </dict>
</plist>
EOT

launchctl load /Library/LaunchAgents/local.node_exporter.plist
