package proxyhat

import "context"

// ProfileService handles profile endpoints.
type ProfileService struct {
	client *Client
}

// GetPreferences returns user preferences.
func (s *ProfileService) GetPreferences(ctx context.Context) (*Preferences, error) {
	var result Preferences
	err := s.client.doRequest(ctx, "GET", "profile/preferences", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePreferences updates user preferences.
func (s *ProfileService) UpdatePreferences(ctx context.Context, preferences map[string]any) (*Preferences, error) {
	var result Preferences
	err := s.client.doRequest(ctx, "PUT", "profile/preferences", preferences, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAPIKeys returns all API keys.
func (s *ProfileService) ListAPIKeys(ctx context.Context) ([]APIKey, error) {
	var result []APIKey
	err := s.client.doRequest(ctx, "GET", "profile/api-keys", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAPIKey creates a new API key.
func (s *ProfileService) CreateAPIKey(ctx context.Context, name *string) (*APIKey, error) {
	var body any
	if name != nil {
		body = map[string]string{"name": *name}
	}
	var result APIKey
	err := s.client.doRequest(ctx, "POST", "profile/api-keys", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAPIKey deletes an API key by ID.
func (s *ProfileService) DeleteAPIKey(ctx context.Context, id string) error {
	return s.client.doRequest(ctx, "DELETE", "profile/api-keys/"+id, nil, nil)
}

// RegenerateAPIKey regenerates an API key.
func (s *ProfileService) RegenerateAPIKey(ctx context.Context, id string) (*APIKey, error) {
	var result APIKey
	err := s.client.doRequest(ctx, "POST", "profile/api-keys/"+id+"/regenerate", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
