FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

ARG TOKEN
ENV PRIFIX="!" \
    DEBUG=false
ENV TOKEN=${TOKEN}

COPY --from=builder /app/main .
COPY .env /app/.env

RUN chmod +x ./main

CMD ["./main"]