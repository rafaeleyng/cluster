#!/bin/sh

set -e

SYSTEM=$(uname -s)
DARWIN="Darwin"
if [ "$SYSTEM" = "$DARWIN" ]; then
  echo "this script is not mean to be run on your my own machine, but in the raspberry pi instead"
  exit 1
fi

if [ -z "$DEVICE_NAME" ]; then
  echo "DEVICE_NAME env is required"
  exit 1
fi

if [ -z "$PASSWORD" ]; then
  echo "PASSWORD env is required"
  exit 1
fi

# setup hostname
echo "$DEVICE_NAME" | sudo tee /etc/hostname
sudo hostname "$DEVICE_NAME"

# setup hosts
# thanks:
# - https://www.raspberrypi.org/forums/viewtopic.php?t=15315
# - https://linuxhandbook.com/sudo-unable-resolve-host/
# 1 - cleanup in case of reruning
sudo sed -i -e 's/^.*hostname-setter.*$//g' /etc/hosts
# 2 - add the new device name, so we don't get errors
echo "127.0.1.1       $DEVICE_NAME ### Set by hostname-setter"  | sudo tee -a /etc/hosts
# 3 - cleanup old hostname
sudo sed -i -e 's/^.*raspberrypi.*$//g' /etc/hosts

# generate ~/.ssh folder (after setup hosts, because uses that value)
ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa 2>/dev/null >/dev/null

# set my personal public key as authorized key
curl https://github.com/rafaeleyng.keys > ~/.ssh/authorized_keys

# change password - thanks https://askubuntu.com/a/80447/384952
usermod --password "$(echo $PASSWORD | openssl passwd -1 -stdin)" pi

# disable ssh password authentication
sudo sed 's/#PasswordAuthentication yes/PasswordAuthentication no/g' /etc/ssh/sshd_config

# reboot to apply hostname changes
sudo reboot
