###################
# build image
###################
FROM golang:1.14-alpine3.12 AS build

ADD ./node-scan /src/node-scan
RUN cd /src/node-scan && go build

###################
# execution image
###################
FROM alpine:3.12

EXPOSE 9090

RUN apk update
RUN apk add curl bash nmap

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ARG PROMETHEUS_VERSION=2.19.0

RUN curl -L -o prometheus.tar.gz https://github.com/prometheus/prometheus/releases/download/v${PROMETHEUS_VERSION}/prometheus-${PROMETHEUS_VERSION}.${TARGETOS}-${TARGETARCH}${TARGETVARIANT}.tar.gz \
    && mkdir prometheus \
    && tar -zxvf prometheus.tar.gz -C prometheus --strip-components 1 \
    && cp prometheus/promtool /usr/local/bin/ \
    && cp prometheus/prometheus /usr/local/bin/ \
    && cp -R prometheus/console_libraries/ /etc/prometheus/ \
    && cp -R prometheus/consoles/ /etc/prometheus/ \
    && rm -fr prometheus*

RUN curl -L -o forego.tgz https://bin.equinox.io/c/ekMN3bCZFUn/forego-stable-${TARGETOS}-${TARGETARCH}.tgz \
    && tar -zxvf forego.tgz \
    && mv forego /usr/bin \
    && rm -fr forego.tgz*

COPY --from=build /src/node-scan/node-scan /usr/local/bin

ADD ./Procfile ./Procfile
ADD ./prometheus.yml /etc/prometheus/prometheus.yml

ENTRYPOINT forego start
