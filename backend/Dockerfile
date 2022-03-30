FROM golang:1.17.6-buster as builder

RUN apt update
RUN apt upgrade -y

WORKDIR /go/src/tsukuyomi

COPY . .
RUN go mod download
RUN go build tsukuyomi.go

FROM ubuntu:latest

WORKDIR /usr/local
COPY --from=builder /go/src/tsukuyomi .

ENV PORT=8080
EXPOSE ${PORT}

CMD ["./tsukuyomi"]


