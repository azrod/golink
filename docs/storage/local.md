---
hide:
  - toc
---

# Local storage backend

Use the local filesystem as the storage backend for Golink. Golink persists the data in the local filesystem. This is the default storage backend.

!!! warning "Not recommended for production"
    The local storage backend is not recommended for production. It is recommended to use a storage backend that can be shared across multiple instances of the Golink server.

## Configuration

The default configuration for the local storage backend is:

``` yaml
storage:
  type: local
  local:
    path: ./
```

!!! example "Set configuration of the local storage backend"

    === "YAML"

        ``` yaml
        storage:
          type: local
          local:
            path: /data/golink
        ```

    === "Flag"

        ``` sh
        --storage.type=local
        --storage.local.path=/data/golink
        ```

    === "Environment variable"

        ``` sh
        export GOLINK_STORAGE_TYPE=local
        export GOLINK_STORAGE_LOCAL_PATH=/data/golink
        ```
