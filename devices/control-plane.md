# control-plane

This is the machine I actually use to interact and build the cluster.

## docker setup

```shell
docker context create pi0 --docker host=tcp://pi0.cluster.rafael:2375
docker context create pi1 --docker host=tcp://pi1.cluster.rafael:2375
docker context create pi2 --docker host=tcp://pi2.cluster.rafael:2375
```

---
