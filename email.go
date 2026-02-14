package proxyhat

import "context"

// EmailService handles email change endpoints.
type EmailService struct {
	client *Client
}

type RequestEmailChangeParams struct {
	Email     string  `json:"email"`
	TwofaCode *string `json:"twofa_code,omitempty"`
}

// RequestChange initiates an email change.
func (s *EmailService) RequestChange(ctx context.Context, params RequestEmailChangeParams) (*EmailChangeResponse, error) {
	var result EmailChangeResponse
	err := s.client.doRequest(ctx, "POST", "profile/email/request-change", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmChange confirms an email change with a token.
func (s *EmailService) ConfirmChange(ctx context.Context, token string) (*EmailChangeResponse, error) {
	var result EmailChangeResponse
	body := map[string]string{"token": token}
	err := s.client.doRequest(ctx, "POST", "profile/email/confirm-change", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelChange cancels a pending email change.
func (s *EmailService) CancelChange(ctx context.Context) (*EmailChangeResponse, error) {
	var result EmailChangeResponse
	err := s.client.doRequest(ctx, "POST", "profile/email/cancel-change", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ResendVerification resends the email verification.
func (s *EmailService) ResendVerification(ctx context.Context) (*EmailChangeResponse, error) {
	var result EmailChangeResponse
	err := s.client.doRequest(ctx, "POST", "profile/email/resend-verification", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
