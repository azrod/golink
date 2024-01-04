---
hide:
  - toc
---

# Deployment

* [Docker](docker.md)
* [Docker Compose](docker-compose.md)
* Kubernetes (coming soon)
* Helm (coming soon)
* [Manual](manual.md)

## How to run

1. Choose a deployment method.
2. Follow the instructions for installing glctl.
3. Use glctl to add a link:

    ```bash
    glctl --host localhost:8081 add link MyApp /myapp https://myapp.example.com
    ```

4. Open the golink web ui in a browser:

    <http://localhost:8081/u/>

5. Follow the link to MyApp:

    <http://localhost:8081/myapp>

!!! note "Note"
    Replace `localhost:8081` with the host and port for your Golink server.
