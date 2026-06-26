package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const appName = "test-8-apr-25"

var startedAt = time.Now().UTC()

var homeTemplate = template.Must(template.New("home").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Name}}</title>
  <style>
    :root { color-scheme: light dark; font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; }
    body { margin: 0; min-height: 100vh; display: grid; place-items: center; background: radial-gradient(circle at top left, #60a5fa, transparent 30%), linear-gradient(135deg, #0f172a 0%, #1e293b 55%, #334155 100%); color: #f8fafc; }
    main { width: min(92vw, 720px); padding: 3rem; border: 1px solid rgb(255 255 255 / 18%); border-radius: 28px; background: rgb(15 23 42 / 72%); box-shadow: 0 24px 80px rgb(0 0 0 / 35%); backdrop-filter: blur(16px); }
    .eyebrow { color: #93c5fd; font-weight: 700; letter-spacing: .16em; text-transform: uppercase; }
    h1 { margin: .6rem 0 1rem; font-size: clamp(2.3rem, 8vw, 5rem); line-height: .95; }
    p { color: #cbd5e1; font-size: 1.1rem; line-height: 1.7; }
    .actions { display: flex; flex-wrap: wrap; gap: 1rem; margin-top: 2rem; }
    a { color: #0f172a; background: #93c5fd; border-radius: 999px; padding: .85rem 1.1rem; text-decoration: none; font-weight: 800; }
    a.secondary { color: #f8fafc; background: rgb(255 255 255 / 12%); border: 1px solid rgb(255 255 255 / 18%); }
    dl { display: grid; gap: .75rem; grid-template-columns: max-content 1fr; margin-top: 2rem; color: #cbd5e1; }
    dt { color: #f8fafc; font-weight: 800; }
    dd { margin: 0; }
  </style>
</head>
<body>
  <main>
    <div class="eyebrow">Go web service</div>
    <h1>Hello from {{.Name}}.</h1>
    <p>This small service is container-ready, exposes machine-readable status endpoints, and keeps a clean single-binary deployment path.</p>
    <div class="actions">
      <a href="/healthz">Health check</a>
      <a class="secondary" href="/api/info">API info</a>
    </div>
    <dl>
      <dt>Status</dt><dd>{{.Status}}</dd>
      <dt>Started</dt><dd>{{.StartedAt}}</dd>
    </dl>
  </main>
</body>
</html>`))

type homeData struct {
	Name      string
	Status    string
	StartedAt string
}

type appInfo struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

func newServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /healthz", healthHandler)
	mux.HandleFunc("GET /api/info", infoHandler)
	return requestLogger(mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := homeData{
		Name:      appName,
		Status:    "ready",
		StartedAt: startedAt.Format(time.RFC3339),
	}
	if err := homeTemplate.Execute(w, data); err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, "ok")
}

func infoHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(appInfo{
		Name:      appName,
		Status:    "ready",
		StartedAt: startedAt.Format(time.RFC3339),
		Version:   version(),
	})
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func version() string {
	value := strings.TrimSpace(os.Getenv("APP_VERSION"))
	if value == "" {
		return "dev"
	}
	return value
}

func main() {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	server := &http.Server{
		Addr:              addr,
		Handler:           newServer(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("%s listening on %s", appName, addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
