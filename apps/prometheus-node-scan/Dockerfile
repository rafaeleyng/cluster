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

RUN curl -LO https://github.com/prometheus/prometheus/releases/download/v2.19.0/prometheus-2.19.0.linux-amd64.tar.gz \
    && tar -zxvf prometheus-2.19.0.linux-amd64.tar.gz \
    && cp prometheus-2.19.0.linux-amd64/promtool /usr/local/bin/ \
    && cp prometheus-2.19.0.linux-amd64/prometheus /usr/local/bin/ \
    && cp -R prometheus-2.19.0.linux-amd64/console_libraries/ /etc/prometheus/ \
    && cp -R prometheus-2.19.0.linux-amd64/consoles/ /etc/prometheus/ \
    && rm -fr prometheus-2.19.0.linux-amd64*

RUN curl -LO https://bin.equinox.io/c/ekMN3bCZFUn/forego-stable-linux-amd64.tgz \
    && tar -zxvf forego-stable-linux-amd64.tgz \
    && mv forego /usr/bin \
    && rm -fr forego-stable-linux-amd64*

COPY --from=build /src/node-scan/node-scan /usr/local/bin

ADD ./Procfile ./Procfile
ADD ./prometheus.yml /etc/prometheus/prometheus.yml

ENTRYPOINT forego start
