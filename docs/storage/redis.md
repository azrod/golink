---
hide:
  - toc
---

# Redis storage backend

Use Redis as the storage backend for Golink. Golink persists the data in Redis, so that it can be used across multiple instances of the Golink server.

## Configuration

The Redis storage backend is configured using environment variables. The following environment variables are available:

| Environment Variable | Default | Description |
| --- | --- | --- |
| `GOLINK_SB_REDIS_ADDRESS` | `localhost:6379` | The address of the Redis server. |
| `GOLINK_SB_REDIS_USERNAME` | _None_ | The username for the Redis server. |
| `GOLINK_SB_REDIS_PASSWORD` | _None_ | The password for the Redis server. |
| `GOLINK_SB_REDIS_DB` | `0` | The Redis database number. |
| `GOLINK_SB_REDIS_MAX_RETRIES` | `3` | The maximum number of retries when connecting to the Redis server. |
| `GOLINK_SB_REDIS_DIAL_TIMEOUT` | `5` | The timeout in seconds for connecting to the Redis server. |
| `GOLINK_SB_REDIS_READ_TIMEOUT` | `3` | The timeout in seconds for reading from the Redis server. |
| `GOLINK_SB_REDIS_WRITE_TIMEOUT` | `3` | The timeout in seconds for writing to the Redis server. |
| `GOLINK_SB_REDIS_CERT_FILE` | _None_ | The path to the certificate file for the Redis server. |
| `GOLINK_SB_REDIS_KEY_FILE` | _None_ | The path to the key file for the Redis server. |
| `GOLINK_SB_REDIS_CA_FILE` | _None_ | The path to the CA file for the Redis server. |

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
          GOLINK_SB_REDIS_ADDRESS=redis:6379
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
