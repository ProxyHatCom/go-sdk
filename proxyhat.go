// Package proxyhat provides a Go client for the ProxyHat API.
//
// Usage:
//
//	client := proxyhat.NewClient("your-api-key")
//	user, err := client.Auth.User(context.Background())
package proxyhat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultBaseURL = "https://api.proxyhat.com/v1"
	DefaultTimeout = 30 * time.Second
	userAgent      = "proxyhat-go/0.1.0"
)

// Client manages communication with the ProxyHat API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration

	Auth         *AuthService
	SubUsers     *SubUsersService
	SubUserGroups *SubUserGroupsService
	Locations    *LocationsService
	Analytics    *AnalyticsService
	ProxyPresets *ProxyPresetsService
	Profile      *ProfileService
	TwoFactor    *TwoFactorService
	Email        *EmailService
	Coupons      *CouponsService
	Plans        *PlansService
	Payments     *PaymentsService
}

// NewClient creates a new ProxyHat API client.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:     apiKey,
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{},
		timeout:    DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.httpClient.Timeout = c.timeout

	c.Auth = &AuthService{client: c}
	c.SubUsers = &SubUsersService{client: c}
	c.SubUserGroups = &SubUserGroupsService{client: c}
	c.Locations = &LocationsService{client: c}
	c.Analytics = &AnalyticsService{client: c}
	c.ProxyPresets = &ProxyPresetsService{client: c}
	c.Profile = &ProfileService{client: c}
	c.TwoFactor = &TwoFactorService{client: c}
	c.Email = &EmailService{client: c}
	c.Coupons = &CouponsService{client: c}
	c.Plans = &PlansService{client: c}
	c.Payments = &PaymentsService{client: c}

	return c
}

func (c *Client) doRequest(ctx context.Context, method, path string, body any, result any) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	reqURL := strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp, respBody); err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	// Try envelope: look for "payload" or "data" key
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(respBody, &envelope); err == nil {
		if payload, ok := envelope["payload"]; ok {
			return json.Unmarshal(payload, result)
		}
		if data, ok := envelope["data"]; ok {
			return json.Unmarshal(data, result)
		}
	}

	return json.Unmarshal(respBody, result)
}

func (c *Client) doRequestWithParams(ctx context.Context, method, path string, params url.Values, result any) error {
	reqURL := strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp, respBody); err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(respBody, &envelope); err == nil {
		if payload, ok := envelope["payload"]; ok {
			return json.Unmarshal(payload, result)
		}
		if data, ok := envelope["data"]; ok {
			return json.Unmarshal(data, result)
		}
	}

	return json.Unmarshal(respBody, result)
}

func (c *Client) doRequestRaw(ctx context.Context, method, path string, params url.Values) (*http.Response, error) {
	reqURL := strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, checkResponse(resp, body)
	}

	return resp, nil
}
