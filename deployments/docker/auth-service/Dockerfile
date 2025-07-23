# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Set up module-aware mode
ENV GO111MODULE=on

# Create and set working directory
WORKDIR /go/src/github.com/shahid-io/shippix

# Copy module files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the service
WORKDIR /go/src/github.com/shahid-io/shippix/cmd/auth-service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/auth-service

# Stage 2: Run
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/auth-service /usr/local/bin/auth-service
EXPOSE 8081
ENTRYPOINT ["auth-service"]