FROM golang:1.21.0-alpine AS builder

COPY go.mod go.sum /github.com/go-park-mail-ru/2024_1_TeaStealers/
WORKDIR /github.com/go-park-mail-ru/2024_1_TeaStealers/

COPY . .

RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/app/main.go
