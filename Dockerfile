# Dockerfile
FROM golang:1.21.1 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

