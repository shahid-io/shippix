# SHIPPIX - Shipment Tracking Microservices Platform

## Overview

**Shippix** is a backend-oriented, modern microservices platform for courier and shipment tracking. It leverages Go (Golang), Docker Compose, PostgreSQL, Redis, Kafka, and follows a modular monorepo architecture. The system is composed of independently deployable services, including authentication, shipment management, tracking, courier management, notifications, and route optimization.

## Features

- **User Authentication**: Secure login/signup via phone or email using JWT tokens.
- **Shipment Management**: Creation, update, and retrieval of shipments.
- **Real-time Tracking**: Track shipment statuses and updates through dedicated tracking service.
- **Courier Management**: Register, assign, and monitor couriers.
- **Notification Service**: Send automated messages to users on shipment events.
- **Route Optimization**: Suggest optimal routes for deliveries (optional/enhanced feature).
- **API Gateway**: Unified entry point for all APIs, routing client requests to appropriate services.
- **Scalable & Modular**: Each microservice can be developed, tested, and scaled independently.

## Directory Structure

```
Shippix/
├── cmd/                     # Service entrypoints (main.go per service)
├── internal/                # Domain/business logic per service
├── pkg/                     # Shared code and libraries
├── deployments/docker/      # Dockerfiles for each service
├── configs/                 # docker-compose.yml and configuration files
├── proto/                   # Protobuf (gRPC) definitions
├── migrations/              # Database migrations
└── ...                      # Other supporting directories
```

## Deployment

### 1. Clean Up Previous Builds

```bash
docker-compose -f configs/docker-compose.yml down
```

### 2. Rebuild with Clean Cache

```bash
docker-compose -f configs/docker-compose.yml build --no-cache auth-service
```

### 3. Start the Service

```bash
docker-compose -f configs/docker-compose.yml up -d auth-service
```

> **Tip:** To build and start the full stack (all services and dependencies), omit the specific service:
> 
> ```bash
> docker-compose -f configs/docker-compose.yml up -d --build
> ```

## Usage

- API endpoints are exposed per service (see code and OpenAPI specs under `/api/openapi.yaml`).
- Healthcheck: Each service implements a `/health` endpoint for monitoring.
- You can invoke signup/login APIs on `auth-service`, shipment creation on `shipment-service`, and so on. See code comments and handlers for available endpoints.

## Requirements

- Docker and Docker Compose (recommended via WSL2 or Linux/macOS)
- Go (for development or running outside containers)
- Make (optional, for dev utility commands)

## Contributing

- Branch from `main` using feature or bugfix branch naming.
- Run all linters and code generators using `make deps`.
- Submit PRs with clear descriptions.

## License

This project is open-sourced; see the `LICENSE` file for details.

**For questions or setup issues, review the deployment steps or inspect logs via**  
```bash
docker logs 
```
**to troubleshoot service startup problems.**

Happy shipping!


- rough
logs : docker-compose -f configs/docker-compose.yml logs -f

