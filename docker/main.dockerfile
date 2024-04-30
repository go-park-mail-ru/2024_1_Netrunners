FROM golang:1.21.0-alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app cmd/app/main.go

EXPOSE 8081

CMD ["./app"]