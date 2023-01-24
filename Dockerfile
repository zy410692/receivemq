FROM golang:1.19

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 80

ENTRYPOINT ["./receivemq"]