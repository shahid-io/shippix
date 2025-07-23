# SHIPPIX

## Deployment

# Clean up previous builds
docker-compose -f configs/docker-compose.yml down

# Rebuild with clean cache
docker-compose -f configs/docker-compose.yml build --no-cache auth-service

# Start the service
docker-compose -f configs/docker-compose.yml up -d auth-service