Here is a **comprehensive, categorized list of all important Docker, Docker Compose, and migration-related commands** discussed in this thread. These cover common workflows for Go microservices development with Docker Compose—including build, up/down, logs, migrations, database inspection, cleanup, and troubleshooting.

# Docker & Docker Compose Command Reference

## 1. **Core Project Workflow**

**Navigate to project root:**
```sh
cd ~/Shippix
```

### **Stop All Containers and Networks**
```sh
docker compose -f configs/docker-compose.yml down
```

### **Prune Unused Networks (optional, safe for dev)**
```sh
docker network prune
# Confirm with 'y'
```

### **Rebuild All Service Images (No Cache)**
```sh
docker compose -f configs/docker-compose.yml build --no-cache
```

### **Start All Services in Background (Detached)**
```sh
docker compose -f configs/docker-compose.yml up -d
```

### **View Status of All Containers**
```sh
docker compose -f configs/docker-compose.yml ps
```

### **Tail Logs (All Services)**
```sh
docker compose -f configs/docker-compose.yml logs -f
```

### **Tail Logs (Single Service, e.g. auth-service)**
```sh
docker compose -f configs/docker-compose.yml logs -f auth-service
```

## 2. **Database Migration Commands**

### **Run Migrations**
```sh
docker compose -f configs/docker-compose.yml run --rm migrate
```
- Applies all new `.up.sql` files in your /migrations directory.

### **Check Migration Files in Container (ensure mapping)**
```sh
docker compose -f configs/docker-compose.yml run --rm migrate ls /migrations
```

### **List Applied Migrations (inside Postgres)**
```sh
docker exec -it configs-postgres-1 psql -U shippix -d shippix_db -c 'SELECT * FROM schema_migrations;'
```

### **Reset Migration State (dev only, use with caution!)**
```sh
docker exec -it configs-postgres-1 psql -U shippix -d shippix_db
# At the psql prompt:
DELETE FROM schema_migrations;
\q
```

## 3. **Database/Admin Utilities**

### **Connect to Postgres DB Container with psql**
```sh
docker exec -it configs-postgres-1 psql -U shippix -d shippix_db
```
- Enter Postgres shell for manual queries.

### **List Tables in Database**
```sh
docker exec -it configs-postgres-1 psql -U shippix -d shippix_db -c '\dt'
```

## 4. **Container & Network Inspection / Debugging**

### **List All Docker Networks**
```sh
docker network ls
```

### **Inspect a Specific Network**
```sh
docker network inspect configs_default
```

### **List All Containers (running and stopped)**
```sh
docker ps -a
```
- Only running containers: `docker ps`

### **Remove a Custom Network**
```sh
docker network rm configs_default
```
- OK if it says "not found".

### **Full System Prune (CAUTION: removes ALL stopped containers/images/volumes)**
```sh
docker system prune -a --volumes
# Confirm with 'y'
```

## 5. **Miscellaneous/Service Utilities**

### **View Docker Compose Version (plugin)**
```sh
docker compose version
```

### **View Docker Engine Version**
```sh
docker version
```

### **Upgrade Docker Compose (if using apt)**
```sh
sudo apt update
sudo apt install docker-compose
```
- _But prefer using the plugin: `docker compose`_

### **Restart Docker Daemon (Linux)**
```sh
sudo systemctl restart docker
```

## 6. **Sample Minimal Compose File for Isolated Troubleshooting**
Create in a fresh directory as `test-compose.yml`:
```yaml
services:
  hello:
    image: alpine
    command: sleep 60
```
Start it:
```sh
docker compose -f test-compose.yml up
```

## 7. **Good Practices**

- **Always use docker compose CLI v2 (with a space):**  
  e.g. `docker compose -f ...` rather than the old `docker-compose` binary.
- **Delete the version field from top of your docker-compose.yml** for modern Compose.
- **Monitor health and logs after up:** Always check with `docker compose ... ps` and `logs -f`.

## 8. **Migration File Troubleshooting**

- Remove a migration version from history (dev only!):
  ```sql
  DELETE FROM schema_migrations;
  ```
  (inside psql shell, then `\q` to exit)

- Re-run migration after clearing history:
  ```sh
  docker compose -f configs/docker-compose.yml run --rm migrate
  ```

## 9. **Quick Reference: Compose Lifecycle**

| Action                    | Command                                                            |
|---------------------------|---------------------------------------------------------------------|
| Down & clean              | `docker compose -f configs/docker-compose.yml down`                |
| Prune networks            | `docker network prune`                                             |
| Build all (no cache)      | `docker compose -f configs/docker-compose.yml build --no-cache`    |
| Up (all services)         | `docker compose -f configs/docker-compose.yml up -d`               |
| Logs (all)                | `docker compose -f configs/docker-compose.yml logs -f`             |
| Logs (one service)        | `docker compose -f configs/docker-compose.yml logs -f servicename` |
| Status                    | `docker compose -f configs/docker-compose.yml ps`                  |
| Run migrations            | `docker compose -f configs/docker-compose.yml run --rm migrate`    |

**Feel free to copy this as your personal cheat-sheet for daily development and debugging!** If you want anything else on volumes, environment management, or specific advanced Compose/Docker patterns, just ask.