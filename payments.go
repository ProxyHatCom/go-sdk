package proxyhat

import (
	"context"
	"net/http"
	"net/url"
)

// PaymentsService handles payment endpoints.
type PaymentsService struct {
	client *Client
}

type CreatePaymentParams struct {
	Type               string  `json:"type"`
	PlanID             string  `json:"plan_id"`
	Gate               string  `json:"gate"`
	CryptocurrencyCode string  `json:"cryptocurrency_code"`
	CouponCode         *string `json:"coupon_code,omitempty"`
}

// List returns all payments.
func (s *PaymentsService) List(ctx context.Context) ([]Payment, error) {
	var result []Payment
	err := s.client.doRequest(ctx, "GET", "payments", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new payment.
func (s *PaymentsService) Create(ctx context.Context, params CreatePaymentParams) (*PaymentCreateResponse, error) {
	if params.Gate == "" {
		params.Gate = "crypto"
	}
	var result PaymentCreateResponse
	err := s.client.doRequest(ctx, "POST", "payments", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a payment by ID.
func (s *PaymentsService) Get(ctx context.Context, id string) (*PaymentDetails, error) {
	var result PaymentDetails
	err := s.client.doRequest(ctx, "GET", "payments/"+id, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Check checks the status of a payment.
func (s *PaymentsService) Check(ctx context.Context, id string) (*PaymentDetails, error) {
	var result PaymentDetails
	err := s.client.doRequest(ctx, "GET", "payments/"+id+"/check", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Invoice downloads a payment invoice. The caller must close the response body.
func (s *PaymentsService) Invoice(ctx context.Context, id string, format string) (*http.Response, error) {
	if format == "" {
		format = "pdf"
	}
	params := url.Values{}
	params.Set("format", format)
	return s.client.doRequestRaw(ctx, "GET", "payments/"+id+"/invoice", params)
}

// Cryptocurrencies returns the list of supported cryptocurrencies.
func (s *PaymentsService) Cryptocurrencies(ctx context.Context) ([]Cryptocurrency, error) {
	var result []Cryptocurrency
	err := s.client.doRequest(ctx, "GET", "payments/cryptocurrencies", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
