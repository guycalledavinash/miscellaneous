# go-test-27-apr

A generic deployable Go web application.

## Run locally

```bash
go run .
```

The service starts on `PORT` (default `8080`) with:
- `GET /` returns plain text
- `GET /health` returns JSON health status

## Build and run with Docker

```bash
docker build -t go-test-27-apr .
docker run --rm -p 8080:8080 go-test-27-apr
```
