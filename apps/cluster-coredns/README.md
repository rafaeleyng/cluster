# cluster-coredns

My home cluster's coredns image with configuration.

It runs:
- a coredns, that contains my local zone and forwards to the pihole
- a pihole

## setup

1. first, I've setup static IPs, by the MAC addresses of my cluster devices

| hostname | ip            |
| -------- | ------------- |
| pi0      | 192.168.1.100 |
| pi1      | 192.168.1.101 |
| pi2      | 192.168.1.102 |
| rafael1  | 192.168.1.131 |

2. create a docker network:
  ```
  docker network create -d bridge --subnet=172.18.0.0/16 dns-net
  ```

3. run the pihole:
  ```sh
  docker run -d \
    --name pihole \
    --network dns-net \
    --ip 172.18.0.2 \
    -p 80:80 \
    -p 443:443 \
    -e TZ="America/Sao_Paulo" \
    -v "$(pwd)/etc-pihole/:/etc/pihole/" \
    -v "$(pwd)/etc-dnsmasq.d/:/etc/dnsmasq.d/" \
    --dns=8.8.8.8 --dns=8.8.4.4 \
    --restart=unless-stopped \
    --hostname pi.hole \
    -e VIRTUAL_HOST="pi.hole" \
    -e PROXY_LOCATION="pi.hole" \
    -e ServerIP="127.0.0.1" \
    pihole/pihole:v5.0
  ```

4. run on `pi2` (because it is configured to be the name server on my router):
  ```sh
  # ensure I have the latest config published
  make docker-build-and-push
  docker run --network dns-net --ip 172.18.0.3 -d -p 53:53 -p 53:53/udp -p 8080:8080 -p 9153:9153 --name cluster-coredns --restart=unless-stopped rafaeleyng/cluster-coredns
  ```

## configure the router to use the DNS server

My router is an Archer C60

Configure in the "DHCP" menu, not in the "Internet" menu, like this (just use the appropriate IP):

- [configuration](https://i.imgur.com/Dng3IiV.png)
- ![configuration](https://user-images.githubusercontent.com/4842605/87379839-8a406200-c567-11ea-98ba-d1857651f908.png)

> TP-Link router don´'t allow you to have dns server on the same sub domain. But you can fix it. Turn of DHCP and put in your ip to Pihole in the field for DNS. Turn the DHCP on and you are ready to go.

2 Comments

## references

- https://blog.idempotent.ca/2018/04/18/run-your-own-home-dns-on-coredns/#comment-4942809218
- https://www.reddit.com/r/pihole/comments/cz2d7l/tplink_dns_rebind_protection_solution/
- https://www.reddit.com/r/pihole/comments/90f45v/dns_server_ip_address_and_lan_ip_address_cannot/
- https://en.wikipedia.org/wiki/Zone_file
- https://en.wikipedia.org/wiki/SOA_record
- https://support.rackspace.com/how-to/what-is-an-soa-record/

---
