---
hide:
  - toc
---

# Version

A version command returns the version of glctl and the version of the golink server.

## Usages

### Get version

``` sh
$> glctl get version
Client informations:
  Version: x.x.x
  Commit: Commit hash
  Build Date: 1970-01-01T00:00:00Z

Server informations:
  Version: x.x.x
```

!!! note "The server version is only available if the server is reachable"
    If the server is not reachable, only the client version is displayed.
