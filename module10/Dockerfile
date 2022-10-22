FROM golang:1.18.3 as build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go build -o httpserver

FROM alpine:3.15
RUN apk add --no-cache tini
COPY --from=build ./build/httpserver /httpserver

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/httpserver"]
