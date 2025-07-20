# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/notification-service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/notification-service

# Stage 2: Run
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/notification-service /usr/local/bin/notification-service
EXPOSE 8084
ENTRYPOINT ["notification-service"]
