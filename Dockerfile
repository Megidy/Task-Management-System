FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev
# dependencies
COPY ["go.mod","go.sum","./"]
RUN go mod download
# build
COPY ["./pkj","./cmd", "./"]

RUN go build -o ./bin/app cmd/main.go


