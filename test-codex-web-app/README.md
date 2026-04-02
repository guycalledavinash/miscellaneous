# test-codex-web-app

Generic dynamic e-commerce web app built with Node.js + Express + EJS.

## Run locally

```bash
npm install
npm start
```

Open `http://localhost:3000`.

## Build and run with Docker

```bash
docker build -t test-codex-web-app .
docker run --rm -p 3000:3000 test-codex-web-app
```

Then open `http://localhost:3000`.
