FROM golang:1.21.6-alpine AS builder

WORKDIR /usr/local/src

EXPOSE 8080

RUN apk --no-cache add bash git make gcc gettext libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY sso/config ./sso/config

COPY ./config /root/.aws/config
COPY ./credentials /root/.aws/credentials

RUN go build -o ./bin/app sso/cmd/sso/main.go
ENTRYPOINT [ "./bin/app" ]