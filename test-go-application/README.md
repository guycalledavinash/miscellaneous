# test-go-application

A tiny Go HTTP service with JSON responses, health probes, structured timeouts, and graceful shutdown.

## Run locally

```bash
go run .
```

The server listens on port `8080` by default. Override it with `PORT`:

```bash
PORT=3000 APP_VERSION=local go run .
```

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/` | Application metadata and greeting. |
| `GET`/`HEAD` | `/health` | Backward-compatible liveness probe. |
| `GET`/`HEAD` | `/healthz` | Liveness probe. |
| `GET`/`HEAD` | `/readyz` | Readiness probe. |

## Test

```bash
go test ./...
```

## Build the container

```bash
docker build -t test-go-application .
docker run --rm -p 8080:8080 test-go-application
```
