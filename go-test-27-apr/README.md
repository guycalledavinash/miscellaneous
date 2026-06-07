# go-test-27-apr

A small, production-friendly Go web service that is ready to run locally or in a container.

## Features

- `GET /` returns a plain-text greeting.
- `GET /health` returns a JSON health payload for load balancers and uptime checks.
- Unknown paths return `404 Not Found` instead of falling through to the home page.
- Unsupported methods return `405 Method Not Allowed` with an `Allow: GET` header.
- Basic hardening headers are added to every response.
- The HTTP server includes read, write, idle, and header timeouts.
- Graceful shutdown is handled for `SIGINT` and `SIGTERM`.

## Run locally

```bash
go run .
```

The service starts on `PORT` if it is set, otherwise it defaults to `8080`.

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
```

## Test

```bash
go test ./...
```

## Build and run with Docker

```bash
docker build -t go-test-27-apr .
docker run --rm -p 8080:8080 go-test-27-apr
```

To use a different port inside the container:

```bash
docker run --rm -e PORT=9090 -p 9090:9090 go-test-27-apr
```
