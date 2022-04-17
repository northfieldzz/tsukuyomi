FROM golang:1.17.6-buster as builder

RUN apt update
RUN apt upgrade -y

WORKDIR /go/src/tsukuyomi

COPY . .
RUN go mod download
RUN go build tsukuyomi.go

FROM ubuntu:latest

RUN apt update && apt install -y ca-certificates

WORKDIR /usr/local
COPY --from=builder /go/src/tsukuyomi .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV ENV_MODE=release
COPY .env.${ENV_MODE} .env.${ENV_MODE}
EXPOSE ${PORT}

CMD ["./tsukuyomi"]


