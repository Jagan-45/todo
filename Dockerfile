FROM golang:1.17.0-alpine AS BUILDER
WORKDIR /app
RUN apt-get update && apt-get install -y git && \
    git clone "https://github.com/Jagan-45/todo" . && \
    go mod download && \
    go build -o main main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=BUILDER main .
CMD ["./main"]