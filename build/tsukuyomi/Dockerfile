FROM golang:1.17.6-buster as builder

RUN apt update && apt upgrade -y

WORKDIR /go/src/tsukuyomi

COPY  .. .
RUN go mod download
RUN go build -o tsukuyomi ./cmd/tsukuyomi/main.go

FROM ubuntu:latest

RUN apt update && apt upgrade -y

WORKDIR /usr/local
COPY --from=builder /go/src/tsukuyomi .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV ENV_MODE=release
COPY --from=builder /go/src/tsukuyomi/configs/.env.${ENV_MODE} ./configs/.env.${ENV_MODE}
EXPOSE ${PORT}

CMD ["./tsukuyomi"]


