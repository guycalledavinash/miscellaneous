# test-8-apr-25

A compact Go web service with a polished landing page, a health check, JSON metadata, and Docker support.

## Features

- Responsive HTML landing page at `/`
- Plain-text readiness check at `/healthz`
- JSON service metadata at `/api/info`
- Configurable `PORT` and `APP_VERSION` environment variables
- Multi-stage Dockerfile for small runtime images

## Run locally

```bash
go run .
```

Then open <http://localhost:8080>.

## API endpoints

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/api/info
```

## Test

```bash
go test ./...
```

## Build and run with Docker

```bash
docker build -t test-8-apr-25 .
docker run --rm -p 8080:8080 test-8-apr-25
```

For a development-oriented single-stage image, use `Dockerfile_single`:

```bash
docker build -f Dockerfile_single -t test-8-apr-25:single .
```
