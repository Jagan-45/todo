FROM golang:1.16 AS builder

WORKDIR /app

RUN git clone https://github.com/Jagan-45/todo /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/main ./

USER nonroot

EXPOSE 8080

CMD ["/app/main"]
