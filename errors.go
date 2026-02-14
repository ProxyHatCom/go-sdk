package proxyhat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Error represents an API error response.
type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Errors     any    `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("proxyhat: %d %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("proxyhat: %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}

// RateLimitError is returned when the API rate limit is exceeded (HTTP 429).
type RateLimitError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Errors     any    `json:"errors,omitempty"`
	RetryAfter int    `json:"retry_after"`
}

func (e *RateLimitError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("proxyhat: %d %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("proxyhat: %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}

// IsAuthenticationError returns true if the error is a 401 Unauthorized.
func IsAuthenticationError(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode == http.StatusUnauthorized
	}
	var rle *RateLimitError
	if errors.As(err, &rle) {
		return rle.StatusCode == http.StatusUnauthorized
	}
	return false
}

// IsPermissionError returns true if the error is a 403 Forbidden.
func IsPermissionError(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode == http.StatusForbidden
	}
	return false
}

// IsNotFoundError returns true if the error is a 404 Not Found.
func IsNotFoundError(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode == http.StatusNotFound
	}
	return false
}

// IsValidationError returns true if the error is a 422 Unprocessable Entity.
func IsValidationError(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode == http.StatusUnprocessableEntity
	}
	return false
}

// IsRateLimitError returns true if the error is a 429 Too Many Requests.
func IsRateLimitError(err error) bool {
	var rle *RateLimitError
	if errors.As(err, &rle) {
		return true
	}
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode == http.StatusTooManyRequests
	}
	return false
}

// AsRateLimitError extracts a RateLimitError from err if present.
func AsRateLimitError(err error) (*RateLimitError, bool) {
	var rle *RateLimitError
	if errors.As(err, &rle) {
		return rle, true
	}
	return nil, false
}

func checkResponse(resp *http.Response, body []byte) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	msg := parseErrorMessage(body)

	var errs any
	var raw map[string]json.RawMessage
	if json.Unmarshal(body, &raw) == nil {
		if e, ok := raw["errors"]; ok {
			var parsed any
			if json.Unmarshal(e, &parsed) == nil {
				errs = parsed
			}
		}
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		rle := &RateLimitError{
			Message:    msg,
			StatusCode: resp.StatusCode,
			Errors:     errs,
		}
		if ra := resp.Header.Get("Retry-After"); ra != "" {
			if v, err := strconv.Atoi(ra); err == nil {
				rle.RetryAfter = v
			}
		}
		return rle
	}

	return &Error{
		Message:    msg,
		StatusCode: resp.StatusCode,
		Errors:     errs,
	}
}

func parseErrorMessage(body []byte) string {
	var raw map[string]json.RawMessage
	if json.Unmarshal(body, &raw) != nil {
		return ""
	}

	// Try "description" first, then "message"
	for _, key := range []string{"description", "message"} {
		if v, ok := raw[key]; ok {
			var s string
			if json.Unmarshal(v, &s) == nil {
				return s
			}
		}
	}
	return ""
}
