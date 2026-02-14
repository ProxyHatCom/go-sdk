package proxyhat

import "context"

// SubUsersService handles sub-user endpoints.
type SubUsersService struct {
	client *Client
}

type CreateSubUserParams struct {
	ProxyPassword    string  `json:"proxy_password"`
	IsTrafficLimited bool    `json:"is_traffic_limited"`
	TrafficLimit     *string `json:"traffic_limit,omitempty"`
	Name             *string `json:"name,omitempty"`
	Notes            *string `json:"notes,omitempty"`
	SubUserGroupID   *string `json:"sub_user_group_id,omitempty"`
}

type UpdateSubUserParams struct {
	ProxyPassword    *string `json:"proxy_password,omitempty"`
	IsTrafficLimited *bool   `json:"is_traffic_limited,omitempty"`
	TrafficLimit     *string `json:"traffic_limit,omitempty"`
	Name             *string `json:"name,omitempty"`
	Notes            *string `json:"notes,omitempty"`
}

// List returns all sub-users.
func (s *SubUsersService) List(ctx context.Context) ([]SubUser, error) {
	var result []SubUser
	err := s.client.doRequest(ctx, "GET", "sub-users", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new sub-user.
func (s *SubUsersService) Create(ctx context.Context, params CreateSubUserParams) (*SubUser, error) {
	var result SubUser
	err := s.client.doRequest(ctx, "POST", "sub-users", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a sub-user by ID.
func (s *SubUsersService) Get(ctx context.Context, id string) (*SubUser, error) {
	var result SubUser
	err := s.client.doRequest(ctx, "GET", "sub-users/"+id, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a sub-user.
func (s *SubUsersService) Update(ctx context.Context, id string, params UpdateSubUserParams) (*SubUser, error) {
	var result SubUser
	err := s.client.doRequest(ctx, "PUT", "sub-users/"+id, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a sub-user.
func (s *SubUsersService) Delete(ctx context.Context, id string) error {
	return s.client.doRequest(ctx, "DELETE", "sub-users/"+id, nil, nil)
}

// ResetUsage resets traffic usage for the given sub-user IDs.
func (s *SubUsersService) ResetUsage(ctx context.Context, ids []string) (*ResetUsageResponse, error) {
	var result ResetUsageResponse
	body := map[string]any{"ids": ids}
	err := s.client.doRequest(ctx, "POST", "sub-users/reset/usage", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkDelete deletes multiple sub-users.
func (s *SubUsersService) BulkDelete(ctx context.Context, ids []string) (*BulkDeleteResponse, error) {
	var result BulkDeleteResponse
	body := map[string]any{"ids": ids}
	err := s.client.doRequest(ctx, "POST", "sub-users/bulk-delete", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkMoveToGroup moves multiple sub-users to a group.
func (s *SubUsersService) BulkMoveToGroup(ctx context.Context, ids []string, groupID *string) (any, error) {
	var result any
	body := map[string]any{"ids": ids}
	if groupID != nil {
		body["group_id"] = *groupID
	}
	err := s.client.doRequest(ctx, "POST", "sub-users/bulk-move-to-group", body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
