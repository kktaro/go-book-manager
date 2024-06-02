# syntax=docker/dockerfile:1
FROM golang:1.22.3
WORKDIR /app
COPY /app .
ENTRYPOINT [ "go", "run", "main.go" ]
EXPOSE 3000