# Stage 1: Build
# FROM golang:1.22-alpine AS builder
FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/tracking-service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/tracking-service

# Stage 2: Run
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/tracking-service /usr/local/bin/tracking-service
EXPOSE 8082
ENTRYPOINT ["tracking-service"]
