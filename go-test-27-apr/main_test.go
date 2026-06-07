package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	newMux().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	expected := "Hello from go-test-27-apr!\n"
	if res.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, res.Body.String())
	}

	if contentType := res.Header().Get("Content-Type"); contentType != "text/plain; charset=utf-8" {
		t.Fatalf("expected text content type, got %q", contentType)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	res := httptest.NewRecorder()

	newMux().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	var payload healthResponse
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		t.Fatalf("failed to decode health response: %v", err)
	}

	if payload.Status != "ok" || payload.Service != "go-test-27-apr" {
		t.Fatalf("unexpected health response: %+v", payload)
	}
}

func TestUnknownPathReturnsNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	res := httptest.NewRecorder()

	newMux().ServeHTTP(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, res.Code)
	}
}

func TestUnsupportedMethodsReturnMethodNotAllowed(t *testing.T) {
	for _, path := range []string{"/", "/health"} {
		req := httptest.NewRequest(http.MethodPost, path, nil)
		res := httptest.NewRecorder()

		newMux().ServeHTTP(res, req)

		if res.Code != http.StatusMethodNotAllowed {
			t.Fatalf("%s: expected status %d, got %d", path, http.StatusMethodNotAllowed, res.Code)
		}

		if allow := res.Header().Get("Allow"); allow != http.MethodGet {
			t.Fatalf("%s: expected Allow header %q, got %q", path, http.MethodGet, allow)
		}
	}
}

func TestSecurityHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	newMux().ServeHTTP(res, req)

	if got := res.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Fatalf("expected X-Content-Type-Options header, got %q", got)
	}

	if got := res.Header().Get("X-Frame-Options"); got != "DENY" {
		t.Fatalf("expected X-Frame-Options header, got %q", got)
	}
}
