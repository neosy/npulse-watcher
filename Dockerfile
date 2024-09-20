FROM golang:1.23.1-alpine3.20 AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 go build -o npulse-watcher ./cmd/main.go

FROM alpine:latest

RUN mkdir -p /app_n/bin
COPY --from=builder /build/npulse-watcher /app_n/bin/npulse-watcher

EXPOSE 8080

ENTRYPOINT ["/app_n/bin/npulse-watcher"]
