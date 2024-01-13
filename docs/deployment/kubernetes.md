---
hide:
  - toc
---

# Kubernetes deployment

Run Golink in a Kubernetes cluster. Golink is stateless, so you can scale it horizontally if your use storage backend that supports it.
See the [backend configuration](../storage/index.md) documentation and choose a storage backend.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: golink
  name: golink
spec:
  replicas: 3
  selector:
    matchLabels:
      app: golink
  template:
    metadata:
      labels:
        app: golink
    spec:
      containers:
      - name: golink
        image: ghcr.io/azrod/golink:latest
        args:
        - --storage.type=redis
        - --storage.redis.address=redis.default.svc.cluster.local:6379
        - --storage.redis.db=10
        - --app.address=0.0.0.0
        - --health.address=0.0.0.0
        - --health.port=8082
        env:
        - name: GOLINK_STORAGE_REDIS_PASSWORD
          valueFrom:
            secretKeyRef:
              name: redis-auth
              key: password
        ports:
        - containerPort: 8081
        - name: healthcheck
          containerPort: 8082
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /health
            port: healthcheck
          initialDelaySeconds: 3
          periodSeconds: 3
```
