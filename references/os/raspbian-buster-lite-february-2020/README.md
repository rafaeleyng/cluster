# raspbian buster lite February 2020

- version: February 2020
- release date: 2020-02-13
- kernel version: 4.19
- size: 434 MB

http://downloads.raspberrypi.org/raspbian/release_notes.txt

## user and default password

- default: `pi / raspberry` (https://www.raspberrypi.org/documentation/linux/usage/users.md)

This is changed in the setup script.

## custom image

I'm creating my own image (with `ssh` and `wpa_supplicant.conf` already configured), following these steps:

- download the base image:
  - direct link: https://downloads.raspberrypi.org/raspbian_lite_latest
  - https://www.raspberrypi.org/downloads/raspbian/

- download Raspberry Pi Imager to burn the image:
  - direct link: https://downloads.raspberrypi.org/imager/imager.dmg
  - https://github.com/raspberrypi/rpi-imager
  - https://www.raspberrypi.org/downloads/

- burn the base image to the SD card.

- go to the folder where this current README is.
  - create a `.wifi` file containing:
    ```sh
    export ssid="TODO" # your wifi SSID
    export psk="TODO"  # your wifi password
    ```

  - run `./raspbian-customize-image.sh`.

- run `diskutil list` an locate the SD card (the one with the `Windows_FAT_32 boot` partition). It was `/dev/disk2` when I did this.

- run `sudo fdisk /dev/disk2`, locate and sum the the `start` (532480) and `size` (3080192) of the last non-empty partition, and sum 100 to have some margin for error. For me, it was 3612673.

- create the image with `sudo dd if=/dev/disk2 of=/Users/rafael/Desktop/2020-02-13-raspbian-buster-lite-rafael.img bs=512 count=3612772`

## installation

- burn the custom image to the SD with Raspberry Pi Imager
- startup the raspberry pi with the SD card
- ssh into the device: `ssh -o "UserKnownHostsFile /dev/null" pi@raspberrypi.local`
- run the setup script:
  ```sh
  # NOTE: set the variables in the last command
  curl https://raw.githubusercontent.com/rafaeleyng/cluster/master/references/os/raspbian-buster-lite-february-2020/raspbian-setup-device.sh --output raspbian-setup-device.sh
  chmod 755 raspbian-setup-device.sh
  sudo DEVICE_NAME=<TODO> PASSWORD=<TODO> ./raspbian-setup-device.sh
  sudo reboot # to apply hostname changes
  ```
- before rebooting, you can check whether everyting is fine with:
  ```sh
  hostname
  cat /etc/hostname
  cat /etc/hosts
  cat /etc/ssh/sshd_config | grep PasswordAuthentication
  cat ~/.ssh/authorized_keys
  ```

## references

- https://www.raspberrypi.org/documentation/configuration/wireless/headless.md
- https://medium.com/@decrocksam/building-your-custom-raspbian-image-8b54a24f814e
- https://desertbot.io/blog/headless-raspberry-pi-3-bplus-ssh-wifi-setup

---
