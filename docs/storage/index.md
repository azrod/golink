---
hide: 
  - toc
---

# Storage backend

The golink server has a pluggable storage backend. The storage backend is responsible for storing and retrieving the golinks.

## List of storage backends

* [Local](local.md) (***default***)
* [Redis](redis.md)

## Configuration

!!! example "Set configuration of the storage backend"

    === "YAML"

        ``` yaml
        storage:
          type: local
        ```

    === "Flag"

        ``` sh
        --storage.type=local
        ```

    === "Environment variable"

        ``` sh
        export GOLINK_STORAGE_TYPE=local
        ```
