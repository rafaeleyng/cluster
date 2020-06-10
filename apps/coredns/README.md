# coredns

```sh
docker create --name coredns -p 8080:8080 -p 53:53/udp coredns/coredns
curl -H 'Cache-Control: no-cache' -O https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/coredns/Corefile
curl -H 'Cache-Control: no-cache' -O https://raw.githubusercontent.com/rafaeleyng/cluster/master/apps/coredns/db.cluster.rafael
docker cp Corefile coredns:/Corefile
docker cp db.cluster.rafael coredns:/db.cluster.rafael
docker start coredns
```

---
