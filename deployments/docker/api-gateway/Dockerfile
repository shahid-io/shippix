# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/api-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/api-gateway

# Stage 2: Run
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/api-gateway /usr/local/bin/api-gateway
EXPOSE 8080
ENTRYPOINT ["api-gateway"]
