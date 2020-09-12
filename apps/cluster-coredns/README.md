# cluster-coredns

My home cluster's coredns image with configuration.

It runs:
- a coredns, that contains my local zone and forwards to the pihole
- a pihole

## setup

1. first, I've setup static IPs, by the MAC addresses of my cluster devices according to https://gist.github.com/rafaeleyng/d3fabc5c09636016e1c9c09ad0f19c70

2. create a docker network:
  ```
  docker network create -d bridge --subnet=172.18.0.0/16 dns-net
  ```

3. run the pihole:
  ```sh
  docker run \
    --dns=8.8.4.4 \
    --dns=8.8.8.8 \
    --hostname pi.hole \
    --ip 172.18.0.2 \
    --name pihole \
    --network dns-net \
    --restart=unless-stopped \
    -d \
    -e PROXY_LOCATION="pi.hole" \
    -e ServerIP="127.0.0.1" \
    -e TZ="America/Sao_Paulo" \
    -e VIRTUAL_HOST="pi.hole" \
    -p 443:443 \
    -p 80:80 \
    -v "$(pwd)/etc-dnsmasq.d/:/etc/dnsmasq.d/" \
    -v "$(pwd)/etc-pihole/:/etc/pihole/" \
    pihole/pihole:v5.0
  ```

4. run the pihole exporter (first obtain the token, check https://github.com/eko/pihole-exporter):
  ```sh
  API_TOKEN=$(docker exec pihole awk -F= -v key="WEBPASSWORD" '$1==key {print $2}' /etc/pihole/setupVars.conf)
  docker run \
    --ip 172.18.0.4 \
    --name pihole-exporter \
    --network dns-net \
    --restart=unless-stopped \
    -d \
    -e 'INTERVAL=30s' \
    -e 'PIHOLE_HOSTNAME=172.18.0.2' \
    -e 'PORT=9617' \
    -e "PIHOLE_API_TOKEN=$API_TOKEN" \
    -p 9617:9617 \
    ekofr/pihole-exporter:0.0.9
  ```

5. ensure I have the updated coredns image published to Docker Hub:
  ```sh
  make docker-build-and-push
  ```

6. run on `pi2` (because it is configured to be the name server on my router):
  ```sh
  docker run \
    --ip 172.18.0.3 \
    --name cluster-coredns \
    --network dns-net \
    --restart=unless-stopped \
    -d \
    -p 53:53 \
    -p 53:53/udp \
    -p 8080:8080 \
    -p 9153:9153 \
    rafaeleyng/cluster-coredns
  ```

## configure the router to use the DNS server

My router is an Archer C60

Configure in the "DHCP" menu, not in the "Internet" menu, like this (just use the appropriate IP):

- ![configuration](https://i.imgur.com/Dng3IiV.png)
- ![configuration](https://user-images.githubusercontent.com/4842605/87379839-8a406200-c567-11ea-98ba-d1857651f908.png)

> TP-Link router don't allow you to have DNS server on the same sub domain. But you can fix it. Turn of DHCP and put in your ip to Pihole in the field for DNS. Turn the DHCP on and you are ready to go.

## references

- https://blog.idempotent.ca/2018/04/18/run-your-own-home-dns-on-coredns/#comment-4942809218
- https://www.reddit.com/r/pihole/comments/cz2d7l/tplink_dns_rebind_protection_solution/
- https://www.reddit.com/r/pihole/comments/90f45v/dns_server_ip_address_and_lan_ip_address_cannot/
- https://en.wikipedia.org/wiki/Zone_file
- https://en.wikipedia.org/wiki/SOA_record
- https://support.rackspace.com/how-to/what-is-an-soa-record/

---
