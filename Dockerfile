FROM golang:1.14.3-alpine3.11 as builder
RUN \
    cd / && \
    apk update && \
    apk add --no-cache git ca-certificates make tzdata && \
    mkdir /prometheus_bot
COPY . /prometheus_bot
RUN cd /prometheus_bot && \
    GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o prometheus_bot

FROM alpine:3.11
COPY --from=builder /prometheus_bot/prometheus_bot /
RUN apk add --no-cache ca-certificates tzdata tini
USER nobody
EXPOSE 9087
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/prometheus_bot"]
