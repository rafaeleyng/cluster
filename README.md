# raspberry-pi
My raspberry-pi setup


## hardware

Modelo: raspberry pi zero w
Architecture: ARMv6

Review (from https://www.androidauthority.com/raspberry-pi-zero-w-review-756498/):

It is a fully working 32-bit computer with a 1 GHz ARMv6 single core microprocessor (ARM1176) , a VideoCore 4 GPU, and 512 MB of memory. The GPU is capable of driving a full HD display at 60 fps

The Zero W uses the Cypress CYW43438 wireless chip as Raspberry Pi 3 Model B meaning it has 802.11n wireless LAN and Bluetooth 4.0 connectivity


## OS

Raspbian Stretch Lite - Minimal image based on Debian Stretch
Version: November 2018
Release date: 2018-11-13
Kernel version: 4.14

Installation steps:

https://desertbot.io/blog/headless-raspberry-pi-3-bplus-ssh-wifi-setup

```
pi@raspberrypi:~ $ uname -a
Linux raspberrypi 4.14.79+ #1159 Sun Nov 4 17:28:08 GMT 2018 armv6l GNU/Linux
```

## hostname

Hostname: raspberrypi.local


## user

- user: pi
- password: raspberry

https://www.raspberrypi.org/documentation/linux/usage/users.md

## ssh

```shell
ssh pi@raspberrypi.local
```

### passwordless ssh

https://www.raspberrypi.org/documentation/remote-access/ssh/passwordless.md


## static IP

Set to 192.168.0.50/24.

From https://raspberrypi.stackexchange.com/a/74428/77623

```shell
cat >> /etc/dhcpcd.conf << EOL
interface wlan0
static ip_address=192.168.0.50/24
static routers=192.168.0.1
static domain_name_servers=8.8.8.8 8.8.4.4
EOL
```

## docker-machine

### on Raspberry Pi

1. install docker
  ```shell
  sudo apt-get install docker-ce=18.06.1~ce~3-0~raspbian
  ```

1. change the OS's ID:
  ```shell
  sudo sed -i 's/ID=raspbian/ID=debian/g' /etc/os-release
  ```

1. add my user to `docker` group so I don't have to `sudo` the commands:
  ```shell
  sudo usermod -aG docker pi
  ```

### on main machine

1. create the machine in docker-machine
  ```shell
  docker-machine create \
    --driver generic \
    --generic-ip-address 192.168.0.50 \
    --generic-ssh-key ~/.ssh/id_rsa \
    --generic-ssh-user pi \
    --engine-storage-driver overlay2 \
    raspberrypi
  ```

1. point commands to the new machine
  ```shell
  eval $(docker-machine env raspberrypi)
  ```

1. test the machine
  ```
  docker run --rm arm32v6/busybox echo "hello world"
  ```

1. point commands back to main machine
  ```
  eval $(docker-machine env -u)
  ```

## running apps

### my-remote

https://github.com/rafaeleyng/my-remote


### pihole

http://pi.local/admin/

```shell
#!/bin/bash

# Just hard code these to your docker server's LAN IP if lookups aren't working
IP=192.168.0.50

# Default of directory you run this from, update to where ever.
DOCKER_CONFIGS="$(pwd)"

echo "### Make sure your IPs are correct, hard code ServerIP ENV VARs if necessary\nIP: ${IP}\nIPv6: ${IPv6}"

# Default ports + daemonized docker container
docker run -d \
    --name pihole \
    -p 53:53/tcp \
    -p 53:53/udp \
    -p 80:80 \
    -p 443:443 \
    --cap-add=NET_ADMIN \
    -v "${DOCKER_CONFIGS}/pihole/:/etc/pihole/" \
    -v "${DOCKER_CONFIGS}/dnsmasq.d/:/etc/dnsmasq.d/" \
    -e ServerIP="${IP}" \
    --restart=unless-stopped \
    --dns=127.0.0.1 --dns=8.8.8.8 \
    diginc/pi-hole-multiarch:debian_armel_prerelease

echo -n "Your password for https://${IP}/admin/ is "
docker logs pihole 2> /dev/null | grep 'password:'
```
