FROM golang:1.21.0-alpine AS builder

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o films cmd/sessions/main.go

EXPOSE 8030

CMD ["./users"]