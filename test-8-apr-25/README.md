# test-8-apr-25

A minimal functional Go web application with a Dockerfile.

## Run locally

```bash
go run .
```

Then open `http://localhost:8080`.

## Build and run with Docker

```bash
docker build -t test-8-apr-25 .
docker run --rm -p 8080:8080 test-8-apr-25
```
