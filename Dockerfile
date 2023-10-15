FROM golang:1.16 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8080

CMD ["/app/main"]
