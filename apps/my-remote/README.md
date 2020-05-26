# my-remote

## deploy

```shell
SLACK_TOKEN=<my-slack-token>; eval $(docker-machine env raspberrypi) && make build-arm && SLACK_TOKEN=$SLACK_TOKEN make docker-run-arm && eval $(docker-machine env -u)
```