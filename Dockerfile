FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]