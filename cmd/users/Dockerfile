FROM golang:1.21.0-alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

WORKDIR /app
COPY . .

<<<<<<<< HEAD:cmd/users/Dockerfile
RUN go mod tidy
========
RUN go mod download
>>>>>>>> ed945f869dcea4edeadb645e222b10658ca51b2a:cmd/users/users.dockerfile
RUN go build -o users cmd/users/main.go

EXPOSE 8030

CMD ["./users"]