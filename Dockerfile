FROM golang:1.24.2-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN apk add --no-cache git \
&& rm -f go.work go.work.sum \
&& CGO_ENABLED=0 go build -o npulse-watcher main.go
    

FROM alpine:latest

RUN mkdir -p /app_n/bin
COPY --from=builder /build/npulse-watcher /app_n/bin/npulse-watcher

EXPOSE 8080

#ENTRYPOINT ["sleep", "infinity"]
ENTRYPOINT ["/app_n/bin/npulse-watcher"]
