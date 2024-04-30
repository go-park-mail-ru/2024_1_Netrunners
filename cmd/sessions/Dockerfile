FROM golang:1.21.0-alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

<<<<<<<< HEAD:cmd/sessions/Dockerfile
RUN go mod tidy
========
RUN go mod download
>>>>>>>> ed945f869dcea4edeadb645e222b10658ca51b2a:cmd/sessions/sessions.dockerfile
RUN go build -o sessions cmd/sessions/main.go

EXPOSE 8010

CMD ["./sessions"]