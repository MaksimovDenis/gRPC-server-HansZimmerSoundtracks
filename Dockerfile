FROM golang:1.21.6-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

#dependencies
COPY ["go.mod", "go.sum", "./" ] 
RUN go mod download


#build
COPY sso ./
RUN go build -o ./bin/app cmd/sso/main.go 

COPY sso/config/local.yaml ./bin
