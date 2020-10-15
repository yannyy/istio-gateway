FROM golang:1.14.4-alpine3.12 

ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/bin/

ADD istio-user istio-user 

CMD ["/go/bin/istio-user"]
