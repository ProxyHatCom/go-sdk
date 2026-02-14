package proxyhat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTest() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient("test-api-key", WithBaseURL(server.URL))
	return client, mux, server.Close
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writePayload(w http.ResponseWriter, v any) {
	writeJSON(w, http.StatusOK, map[string]any{"payload": v})
}

func writeData(w http.ResponseWriter, v any) {
	writeJSON(w, http.StatusOK, map[string]any{"data": v})
}

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient("my-key")
	if c.apiKey != "my-key" {
		t.Errorf("apiKey = %q, want %q", c.apiKey, "my-key")
	}
	if c.baseURL != DefaultBaseURL {
		t.Errorf("baseURL = %q, want %q", c.baseURL, DefaultBaseURL)
	}
	if c.timeout != DefaultTimeout {
		t.Errorf("timeout = %v, want %v", c.timeout, DefaultTimeout)
	}
}

func TestNewClient_WithOptions(t *testing.T) {
	hc := &http.Client{}
	c := NewClient("key",
		WithBaseURL("https://custom.api.com"),
		WithHTTPClient(hc),
		WithTimeout(60*time.Second),
	)
	if c.baseURL != "https://custom.api.com" {
		t.Errorf("baseURL = %q, want %q", c.baseURL, "https://custom.api.com")
	}
	if c.httpClient != hc {
		t.Error("httpClient not set correctly")
	}
	if c.timeout != 60*time.Second {
		t.Errorf("timeout = %v, want %v", c.timeout, 60*time.Second)
	}
}

func TestNewClient_ServicesInitialized(t *testing.T) {
	c := NewClient("key")
	if c.Auth == nil {
		t.Error("Auth service is nil")
	}
	if c.SubUsers == nil {
		t.Error("SubUsers service is nil")
	}
	if c.SubUserGroups == nil {
		t.Error("SubUserGroups service is nil")
	}
	if c.Locations == nil {
		t.Error("Locations service is nil")
	}
	if c.Analytics == nil {
		t.Error("Analytics service is nil")
	}
	if c.ProxyPresets == nil {
		t.Error("ProxyPresets service is nil")
	}
	if c.Profile == nil {
		t.Error("Profile service is nil")
	}
	if c.TwoFactor == nil {
		t.Error("TwoFactor service is nil")
	}
	if c.Email == nil {
		t.Error("Email service is nil")
	}
	if c.Coupons == nil {
		t.Error("Coupons service is nil")
	}
	if c.Plans == nil {
		t.Error("Plans service is nil")
	}
	if c.Payments == nil {
		t.Error("Payments service is nil")
	}
}

func TestDoRequest_SetsHeaders(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer test-api-key" {
			t.Errorf("Authorization = %q, want %q", got, "Bearer test-api-key")
		}
		if got := r.Header.Get("Accept"); got != "application/json" {
			t.Errorf("Accept = %q, want %q", got, "application/json")
		}
		if got := r.Header.Get("User-Agent"); got != userAgent {
			t.Errorf("User-Agent = %q, want %q", got, userAgent)
		}
		writeJSON(w, http.StatusOK, map[string]string{"ok": "true"})
	})

	var result map[string]string
	err := client.doRequest(context.Background(), "GET", "test", nil, &result)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoRequest_EnvelopePayload(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, map[string]string{"name": "alice"})
	})

	var result map[string]string
	err := client.doRequest(context.Background(), "GET", "test", nil, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result["name"] != "alice" {
		t.Errorf("name = %q, want %q", result["name"], "alice")
	}
}

func TestDoRequest_EnvelopeData(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []map[string]string{{"id": "1"}})
	})

	var result []map[string]string
	err := client.doRequest(context.Background(), "GET", "test", nil, &result)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 || result[0]["id"] != "1" {
		t.Errorf("unexpected result: %v", result)
	}
}

func TestCheckResponse_401(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsAuthenticationError(err) {
		t.Errorf("expected authentication error, got %v", err)
	}
}

func TestCheckResponse_403(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusForbidden, map[string]string{"message": "Forbidden"})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if !IsPermissionError(err) {
		t.Errorf("expected permission error, got %v", err)
	}
}

func TestCheckResponse_404(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, map[string]string{"message": "Not found"})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if !IsNotFoundError(err) {
		t.Errorf("expected not found error, got %v", err)
	}
}

func TestCheckResponse_422(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
			"message": "Validation failed",
			"errors":  map[string][]string{"email": {"required"}},
		})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if !IsValidationError(err) {
		t.Errorf("expected validation error, got %v", err)
	}
}

func TestCheckResponse_429(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "30")
		writeJSON(w, http.StatusTooManyRequests, map[string]string{"message": "Rate limited"})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if !IsRateLimitError(err) {
		t.Errorf("expected rate limit error, got %v", err)
	}

	rle, ok := AsRateLimitError(err)
	if !ok {
		t.Fatal("AsRateLimitError returned false")
	}
	if rle.RetryAfter != 30 {
		t.Errorf("RetryAfter = %d, want 30", rle.RetryAfter)
	}
}

func TestCheckResponse_500(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "Internal error"})
	})

	err := client.doRequest(context.Background(), "GET", "test", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *Error
	if !isError(err, &apiErr) {
		t.Fatal("expected *Error")
	}
	if apiErr.StatusCode != 500 {
		t.Errorf("StatusCode = %d, want 500", apiErr.StatusCode)
	}
}

func TestError_ErrorString(t *testing.T) {
	e := &Error{Message: "test error", StatusCode: 400}
	expected := "proxyhat: 400 test error"
	if e.Error() != expected {
		t.Errorf("Error() = %q, want %q", e.Error(), expected)
	}
}

func TestError_ErrorStringEmpty(t *testing.T) {
	e := &Error{StatusCode: 400}
	expected := "proxyhat: 400 Bad Request"
	if e.Error() != expected {
		t.Errorf("Error() = %q, want %q", e.Error(), expected)
	}
}

func isError(err error, target any) bool {
	if ptr, ok := target.(**Error); ok {
		var e *Error
		if errors.As(err, &e) {
			*ptr = e
			return true
		}
	}
	return false
}
