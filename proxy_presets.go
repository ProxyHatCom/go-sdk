package proxyhat

import "context"

// ProxyPresetsService handles proxy preset endpoints.
type ProxyPresetsService struct {
	client *Client
}

type CreateProxyPresetParams struct {
	Name string         `json:"name"`
	Data map[string]any `json:"data"`
}

type UpdateProxyPresetParams struct {
	Name *string        `json:"name,omitempty"`
	Data map[string]any `json:"data,omitempty"`
}

// List returns all proxy presets.
func (s *ProxyPresetsService) List(ctx context.Context) ([]ProxyPreset, error) {
	var result []ProxyPreset
	err := s.client.doRequest(ctx, "GET", "proxy-presets", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new proxy preset.
func (s *ProxyPresetsService) Create(ctx context.Context, params CreateProxyPresetParams) (*ProxyPreset, error) {
	var result ProxyPreset
	err := s.client.doRequest(ctx, "POST", "proxy-presets", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a proxy preset by ID.
func (s *ProxyPresetsService) Get(ctx context.Context, id string) (*ProxyPreset, error) {
	var result ProxyPreset
	err := s.client.doRequest(ctx, "GET", "proxy-presets/"+id, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a proxy preset.
func (s *ProxyPresetsService) Update(ctx context.Context, id string, params UpdateProxyPresetParams) (*ProxyPreset, error) {
	var result ProxyPreset
	err := s.client.doRequest(ctx, "PUT", "proxy-presets/"+id, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a proxy preset.
func (s *ProxyPresetsService) Delete(ctx context.Context, id string) error {
	return s.client.doRequest(ctx, "DELETE", "proxy-presets/"+id, nil, nil)
}
