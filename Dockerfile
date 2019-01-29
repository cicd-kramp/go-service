FROM golang:1.10.5-stretch
WORKDIR /go/src/github.com/kramp/go-service
COPY . .
RUN GOOS=linux GOARCH=amd64 go install -ldflags "-linkmode=external" .
FROM debian:stretch-slim
COPY --from=0 /go/bin/go-service /usr/bin
ENTRYPOINT ["/usr/bin/go-service"]
