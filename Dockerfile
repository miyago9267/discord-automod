FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

ARG TOKEN
ENV ENV=prd \
    PRIFIX="!" \
    DEBUG=false
ENV TOKEN=${TOKEN}

COPY --from=builder /app/main .
COPY /config/banned_words.txt /app/config/banned_words.txt

RUN chmod +x ./main

CMD ["./main"]