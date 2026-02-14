package proxyhat

import "context"

// TwoFactorService handles two-factor authentication endpoints.
type TwoFactorService struct {
	client *Client
}

// Status returns the 2FA status for the authenticated user.
func (s *TwoFactorService) Status(ctx context.Context) (*TwoFactorStatus, error) {
	var result TwoFactorStatus
	err := s.client.doRequest(ctx, "GET", "profile/2fa/status", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Enable initiates 2FA setup.
func (s *TwoFactorService) Enable(ctx context.Context) (*TwoFactorEnableResponse, error) {
	var result TwoFactorEnableResponse
	err := s.client.doRequest(ctx, "POST", "profile/2fa/enable", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Confirm confirms 2FA setup with a verification code.
func (s *TwoFactorService) Confirm(ctx context.Context, code string) (any, error) {
	var result any
	body := map[string]string{"code": code}
	err := s.client.doRequest(ctx, "POST", "profile/2fa/confirm", body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Disable disables 2FA with a verification code.
func (s *TwoFactorService) Disable(ctx context.Context, twofaCode string) (any, error) {
	var result any
	body := map[string]string{"twofa_code": twofaCode}
	err := s.client.doRequest(ctx, "POST", "profile/2fa/disable", body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QRCode returns the QR code for 2FA setup.
func (s *TwoFactorService) QRCode(ctx context.Context) (*TwoFactorEnableResponse, error) {
	var result TwoFactorEnableResponse
	err := s.client.doRequest(ctx, "GET", "profile/2fa/qr-code", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RecoveryCodes returns the 2FA recovery codes.
func (s *TwoFactorService) RecoveryCodes(ctx context.Context) (*RecoveryCodes, error) {
	var result RecoveryCodes
	err := s.client.doRequest(ctx, "GET", "profile/2fa/recovery-codes", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DisableByRecovery disables 2FA using a recovery code.
func (s *TwoFactorService) DisableByRecovery(ctx context.Context, recoveryCode string) (any, error) {
	var result any
	body := map[string]string{"recovery_code": recoveryCode}
	err := s.client.doRequest(ctx, "POST", "profile/2fa/disable-by-recovery-code", body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ChangePasswordParams struct {
	CurrentPassword      string  `json:"current_password"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
	TwofaCode            *string `json:"twofa_code,omitempty"`
}

// ChangePassword changes the user's password.
func (s *TwoFactorService) ChangePassword(ctx context.Context, params ChangePasswordParams) (any, error) {
	var result any
	err := s.client.doRequest(ctx, "POST", "profile/password", params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
