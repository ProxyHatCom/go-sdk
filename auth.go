package proxyhat

import "context"

// AuthService handles authentication endpoints.
type AuthService struct {
	client *Client
}

type RegisterParams struct {
	Name                 string  `json:"name"`
	Email                string  `json:"email"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
	ReferralCode         *string `json:"referral_code,omitempty"`
	UTMSource            *string `json:"utm_source,omitempty"`
	UTMMedium            *string `json:"utm_medium,omitempty"`
	UTMCampaign          *string `json:"utm_campaign,omitempty"`
}

type LoginParams struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	TwofaCode *string `json:"twofa_code,omitempty"`
}

// Register creates a new account.
func (s *AuthService) Register(ctx context.Context, params RegisterParams) (*RegisterResponse, error) {
	var result RegisterResponse
	err := s.client.doRequest(ctx, "POST", "auth/register", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Login authenticates a user.
func (s *AuthService) Login(ctx context.Context, params LoginParams) (*LoginResponse, error) {
	var result LoginResponse
	err := s.client.doRequest(ctx, "POST", "auth/login", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// User returns the authenticated user.
func (s *AuthService) User(ctx context.Context) (*User, error) {
	var result User
	err := s.client.doRequest(ctx, "GET", "auth/user", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Logout invalidates the current session.
func (s *AuthService) Logout(ctx context.Context) error {
	return s.client.doRequest(ctx, "POST", "auth/logout", nil, nil)
}

// SupportedProviders returns the list of supported OAuth providers.
func (s *AuthService) SupportedProviders(ctx context.Context) ([]SupportedProvider, error) {
	var result []SupportedProvider
	err := s.client.doRequest(ctx, "GET", "auth/supported-providers", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SocialAccounts returns the user's connected social accounts.
func (s *AuthService) SocialAccounts(ctx context.Context) ([]SocialAccount, error) {
	var result []SocialAccount
	err := s.client.doRequest(ctx, "GET", "auth/social-accounts", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DisconnectSocial removes a connected social account.
func (s *AuthService) DisconnectSocial(ctx context.Context, provider string) error {
	return s.client.doRequest(ctx, "DELETE", "auth/social-accounts/"+provider, nil, nil)
}

// OAuthRedirect returns the OAuth redirect URL for a provider.
func (s *AuthService) OAuthRedirect(ctx context.Context, provider string) (any, error) {
	var result any
	err := s.client.doRequest(ctx, "GET", "auth/"+provider+"/redirect", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
