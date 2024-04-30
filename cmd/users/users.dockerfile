FROM golang:1.21.0-alpine AS builder

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o users cmd/users/main.go

EXPOSE 8030

CMD ["./users"]