# raspbian buster lite February 2020

version: February 2020
release date: 2020-02-13
kernel version: 4.19
size: 434 MB

http://downloads.raspberrypi.org/raspbian/release_notes.txt

## custom image

I'm creating my own image (with `ssh` and `wpa_supplicant.conf` already configured), following these steps

1. download the base image
  - reference:
    - direct link: https://downloads.raspberrypi.org/raspbian_lite_latest
    - https://www.raspberrypi.org/downloads/raspbian/

1. download Raspberry Pi Imager to burn the image
  - reference:
    - direct link: https://downloads.raspberrypi.org/imager/imager.dmg
    - https://github.com/raspberrypi/rpi-imager
    - https://www.raspberrypi.org/downloads/

1. burn the base image to the SD card

1. go to the folder where this current README is

  1. create a `.wifi` file containing:
    ```sh
    ssid="TODO" # your wifi SSID
    psk="TODO"  # your wifi password
    ```

  1. run `./customize-raspbian.sh`

1. run `diskutil list` an locate the SD card (the one with the `Windows_FAT_32 boot` partition). It was `/dev/disk2` when I did this.

1. run `sudo fdisk /dev/disk2`, locate and sum the the `start` (532480) and `size` (3080192) of the last non-empty partition, and sum 100 to have some margin for error. For me, it was 3612673.

1. create the image with `sudo dd if=/dev/disk2 of=/Users/rafael/Desktop/2020-02-13-raspbian-buster-lite-rafael.img bs=512 count=3612772`

## installation

1. burn the custom image to the SSD with Raspberry Pi Imager

## references

- https://www.raspberrypi.org/documentation/configuration/wireless/headless.md
- https://medium.com/@decrocksam/building-your-custom-raspbian-image-8b54a24f814e
- https://desertbot.io/blog/headless-raspberry-pi-3-bplus-ssh-wifi-setup

---
