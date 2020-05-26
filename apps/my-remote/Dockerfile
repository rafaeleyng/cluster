FROM golang:1.11

ARG slack_token

WORKDIR /go/src/my-remote
COPY . .

ENV SLACK_TOKEN $slack_token

CMD ["./my-remote"]
