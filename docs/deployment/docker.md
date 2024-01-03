---
hide:
  - toc
---

# Docker deployment

Run GoLink in a Docker container. This is the easiest way to get started with GoLink.
See the [backend configuration](../configuration/backend.md) documentation and choose a storage backend.

This example uses Redis as storage backend.

```bash
docker run -itd \
    --env "SERVER_HOST=0.0.0.0" \ 
    --env "DB_REDIS_ADDRESS=redis:6379" \ 
    -p 8081:8081 \ 
    --name golink \
    --restart always \ 
    ghcr.io/azrod/golink:latest
```
