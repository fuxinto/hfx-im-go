FROM golang:alpine as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\
   CGO_ENABLED=0

WORKDIR /build
COPY internal/config .

RUN  go mod tidy \
    && go build -ldflags="-s -w" -o gateGo .

FROM alpine:latest

LABEL MAINTAINER="fuxinto@qq.com"

ENV TZ Asia/Shanghai
WORKDIR /gate

COPY --from=0 /build/service/gate/config/config.docker.yaml ./gate/
COPY --from=0 /build/gate ./

EXPOSE 8090
ENTRYPOINT ./gateGo -c config.docker.yaml
