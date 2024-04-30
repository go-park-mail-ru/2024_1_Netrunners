FROM golang:1.21.0-alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o sessions cmd/sessions/main.go

EXPOSE 8010

CMD ["./sessions"]