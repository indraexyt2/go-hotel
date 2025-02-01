FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod tidy
RUN go build -o hotel-payment
RUN chmod +x hotel-payment

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/hotel-payment .

EXPOSE 8083

CMD ["./hotel-payment"]
