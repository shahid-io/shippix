# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/courier-service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/courier-service

# Stage 2: Run
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/courier-service /usr/local/bin/courier-service
EXPOSE 8083
ENTRYPOINT ["courier-service"]
