---
hide:
    - toc
---

# Docker Compose

Docker Compose is a tool for defining and running multi-container Docker applications. It is a good choice for small deployments.

## How to run

1. Install Docker Compose. See the [Docker Compose installation instructions](https://docs.docker.com/compose/install/).
2. Create a file named `docker-compose.yml` with the following content:

    ```yaml
    version: '3.7'

    services:
      golink:
        image: ghcr.io/azrod/golink:latest
        container_name: golink
        restart: unless-stopped
        healthcheck:
          test: ["CMD", "curl -f http://localhost:8082 || exit 1"]
          timeout: 30s
          interval: 1m
          retries: 3
        ports:
          - "127.0.0.1:8081:8081"
    ```

3. Run the Golink server:

    ```bash
    docker compose up -d
    ```
