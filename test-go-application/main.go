package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	appName         = "test-go-application"
	defaultPort     = "8080"
	shutdownTimeout = 10 * time.Second
)

type app struct {
	logger *log.Logger
}

type responseEnvelope struct {
	Application string `json:"application,omitempty"`
	Status      string `json:"status"`
	Message     string `json:"message,omitempty"`
	Version     string `json:"version,omitempty"`
	Timestamp   string `json:"timestamp"`
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
	server := &http.Server{
		Addr:         ":" + envOrDefault("PORT", defaultPort),
		Handler:      newApp(logger).routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Printf("starting %s on %s", appName, server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("server failed: %v", err)
		}
	}()

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-shutdownCtx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	logger.Printf("shutting down %s", appName)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("graceful shutdown failed: %v", err)
	}
	logger.Printf("%s stopped", appName)
}

func newApp(logger *log.Logger) *app {
	return &app{logger: logger}
}

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.handleHome)
	mux.HandleFunc("/health", a.handleHealth)
	mux.HandleFunc("/healthz", a.handleHealth)
	mux.HandleFunc("/readyz", a.handleReadiness)

	return a.logRequests(mux)
}

func (a *app) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.writeJSON(w, http.StatusNotFound, responseEnvelope{
			Status:    "error",
			Message:   "route not found",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		a.writeJSON(w, http.StatusMethodNotAllowed, responseEnvelope{
			Status:    "error",
			Message:   "method not allowed",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	a.writeJSON(w, http.StatusOK, responseEnvelope{
		Application: appName,
		Status:      "ok",
		Message:     "Hello from test-go-application!",
		Version:     envOrDefault("APP_VERSION", "dev"),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	})
}

func (a *app) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", strings.Join([]string{http.MethodGet, http.MethodHead}, ", "))
		a.writeJSON(w, http.StatusMethodNotAllowed, responseEnvelope{
			Status:    "error",
			Message:   "method not allowed",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	a.writeJSON(w, http.StatusOK, responseEnvelope{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func (a *app) handleReadiness(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", strings.Join([]string{http.MethodGet, http.MethodHead}, ", "))
		a.writeJSON(w, http.StatusMethodNotAllowed, responseEnvelope{
			Status:    "error",
			Message:   "method not allowed",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	a.writeJSON(w, http.StatusOK, responseEnvelope{
		Status:    "ready",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func (a *app) writeJSON(w http.ResponseWriter, status int, body responseEnvelope) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		a.logger.Printf("failed to write response: %v", err)
	}
}

func (a *app) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		a.logger.Printf("%s %s from %s completed in %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Round(time.Millisecond))
	})
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
