FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod tidy
RUN go build -o hotel-ums
RUN chmod +x hotel-ums

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/hotel-ums .

EXPOSE 8080

CMD ["./hotel-ums"]
