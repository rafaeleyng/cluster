# coredns

## setup

1. first, I've setup static IPs, by the MAC addresses of my cluster devices

| hostname | ip            |
| -------- | ------------- |
| pi0      | 192.168.1.100 |
| pi1      | 192.168.1.101 |
| pi2      | 192.168.1.102 |
| rafael1  | 192.168.1.131 |

2. run on `pi2` (because it is configured to be the name server on my router):
  ```sh
  curl -H 'Cache-Control: no-cache' -O https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/coredns/Corefile
  curl -H 'Cache-Control: no-cache' -O https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/coredns/db.cluster.rafael
  docker create --name coredns -p 53:53 -p 53:53/udp -p 8080:8080 -p 9153:9153 --restart=always coredns/coredns
  docker cp Corefile coredns:/Corefile
  docker cp db.cluster.rafael coredns:/db.cluster.rafael
  docker restart coredns
  rm Corefile db.cluster.rafael
  ```

## configure the router to use the DNS server

My router is an Archer C60

Configure in the "DHCP" menu, not in the "Internet" menu, like this (just use the appropriate IP):

[configuration](https://i.imgur.com/Dng3IiV.png)

## references

- https://blog.idempotent.ca/2018/04/18/run-your-own-home-dns-on-coredns/#comment-4942809218
- https://www.reddit.com/r/pihole/comments/cz2d7l/tplink_dns_rebind_protection_solution/
- https://www.reddit.com/r/pihole/comments/90f45v/dns_server_ip_address_and_lan_ip_address_cannot/
- https://en.wikipedia.org/wiki/Zone_file
- https://en.wikipedia.org/wiki/SOA_record
- https://support.rackspace.com/how-to/what-is-an-soa-record/

---
