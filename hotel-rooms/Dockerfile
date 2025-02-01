FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod tidy
RUN go build -o hotel-room
RUN chmod +x hotel-room

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/hotel-room .

EXPOSE 8081

CMD ["./hotel-room"]
