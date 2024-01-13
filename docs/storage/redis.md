---
hide:
  - toc
---

# Redis storage backend

Use Redis as the storage backend for Golink. Golink persists the data in Redis, so that it can be used across multiple instances of the Golink server.

## Configuration

The default configuration for the Redis storage backend is:

``` yaml
  redis:
    address: localhost:6379
    database: 0
    max_retries: 3
    dial_timeout: 5
    read_timeout: 3
    write_timeout: 3
```

!!! example "Set configuration of the Redis storage backend"

    === "YAML"

        ``` yaml
        storage:
          type: redis
          redis:
            address: localhost:6379
            username: ""
            password: ""
            database: 0
            max_retries: 3
            dial_timeout: 5
            read_timeout: 3
            write_timeout: 3
        ```

    === "Flag"

        ``` sh
        --storage.type=redis
        --storage.redis.address=localhost:6379
        --storage.redis.username=""
        --storage.redis.password=""
        --storage.redis.database=0
        --storage.redis.max.retries=3
        --storage.redis.dial.timeout=5s
        --storage.redis.read.timeout=3s
        --storage.redis.write.timeout=3s
        ```

    === "Environment variable"

        ``` sh
        export GOLINK_STORAGE_TYPE=redis
        export GOLINK_STORAGE_REDIS_ADDRESS=localhost:6379
        export GOLINK_STORAGE_REDIS_USERNAME=""
        export GOLINK_STORAGE_REDIS_PASSWORD=""
        export GOLINK_STORAGE_REDIS_DATABASE=0
        export GOLINK_STORAGE_REDIS_MAX_RETRIES=3
        export GOLINK_STORAGE_REDIS_DIAL_TIMEOUT=5s
        export GOLINK_STORAGE_REDIS_READ_TIMEOUT=3s
        export GOLINK_STORAGE_REDIS_WRITE_TIMEOUT=3s
        ```

## Example

To use Redis as the storage backend for Golink, you can run a Redis server in a Docker container and then run the Golink server in a Docker container. The following example shows how to do this.

1. Create a directory for the Redis server:

    ```bash
    mkdir -p ~/redis/data
    ```

2. Create a file named `docker-compose.yml` with the following content:

    ```yaml
    version: '3.7'

    services:
      redis:
        image: redis:latest
        container_name: redis
        volumes:
          - ~/redis/data:/data
      golink:
        image: ghcr.io/azrod/golink:latest
        container_name: golink
        environment:
          GOLINK_STORAGE_TYPE=redis
          GOLINK_STORAGE_REDIS_ADDRESS=redis:6379
        ports:
          - "8081:8081"
    ```

3. Run the Redis and Golink server:

    ```bash
    docker-compose up -d
    ```

4. Use glctl to add a golink:

    ```bash
    glctl --host localhost:8081 add link MyApp /myapp https://myapp.example.com
    ```

5. Open the golink web ui in a browser:

    <http://localhost:8081/u/>

6. Follow the link to MyApp:

    <http://localhost:8081/myapp>
