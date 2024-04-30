FROM golang:1.21.0-alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o films cmd/films/main.go

EXPOSE 8020

CMD ["./films"]