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
echo "START setup hostname"
echo "$DEVICE_NAME" | tee /etc/hostname
hostname "$DEVICE_NAME"
echo "DONE  setup hostname"

# setup hosts
# thanks:
# - https://www.raspberrypi.org/forums/viewtopic.php?t=15315
# - https://linuxhandbook.com/sudo-unable-resolve-host/
# 1 - cleanup in case of reruning
echo "START setup hosts 1"
set +e
sed -i -e 's/^.*hostname-setter.*$//g' /etc/hosts
set -e
echo "DONE  setup hosts 1"

# 2 - add the new device name, so we don't get errors
echo "START setup hosts 2"
echo "127.0.1.1       $DEVICE_NAME ### Set by hostname-setter"  | tee -a /etc/hosts
echo "DONE  setup hosts 2"

# 3 - cleanup old hostname
echo "START setup hosts 1"
sed -i -e 's/^.*raspberrypi.*$//g' /etc/hosts
sed -i '/^[[:space:]]*$/d' /etc/hosts
echo "DONE  setup hosts 3"

# generate .ssh folder (after setup hosts, because uses that value)
echo "START generate ssh folder"
rm -fr /home/pi/.ssh
mkdir /home/pi/.ssh
echo "DONE  generate ssh folder"

# set my personal public key as authorized key
echo "START add public key"
curl https://github.com/rafaeleyng.keys > /home/pi/.ssh/authorized_keys
echo "DONE  add public key"

# change password - thanks https://askubuntu.com/a/80447/384952
echo "START set password"
usermod --password "$(echo $PASSWORD | openssl passwd -1 -stdin)" pi
echo "DONE  set password"

# disable ssh password authentication
echo "START disable ssh password authentication"
sed -i 's/#PasswordAuthentication yes/PasswordAuthentication no/g' /etc/ssh/sshd_config
echo "DONE  disable ssh password authentication"
