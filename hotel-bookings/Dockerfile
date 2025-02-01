FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod tidy
RUN go build -o hotel-booking
RUN chmod +x hotel-booking

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/hotel-booking .

EXPOSE 8082

CMD ["./hotel-booking"]
