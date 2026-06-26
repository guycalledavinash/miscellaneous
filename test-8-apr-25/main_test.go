package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	newServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Hello from test-8-apr-25") {
		t.Fatalf("home page did not include app greeting: %q", rec.Body.String())
	}
	if contentType := rec.Header().Get("Content-Type"); !strings.Contains(contentType, "text/html") {
		t.Fatalf("expected HTML content type, got %q", contentType)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	newServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if strings.TrimSpace(rec.Body.String()) != "ok" {
		t.Fatalf("expected ok response, got %q", rec.Body.String())
	}
}

func TestInfoHandler(t *testing.T) {
	t.Setenv("APP_VERSION", "v1.2.3")
	req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
	rec := httptest.NewRecorder()

	newServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var info appInfo
	if err := json.NewDecoder(rec.Body).Decode(&info); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
	if info.Name != appName || info.Status != "ready" || info.Version != "v1.2.3" {
		t.Fatalf("unexpected app info: %+v", info)
	}
}

func TestNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	newServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}
