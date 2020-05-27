#!/bin/sh

set -e

source .wifi

cd /Volumes/boot

# https://www.raspberrypi.org/documentation/configuration/wireless/headless.md
# https://www.raspberrypi.org/documentation/configuration/wireless/wireless-cli.md
echo "country=BR
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1

network={
  ssid=\"$ssid\"
  psk=\"$psk\"
}
" > wpa_supplicant.conf

# https://www.raspberrypi.org/documentation/remote-access/ssh/README.md
touch ssh
