package proxyhat

import "context"

// SubUserGroupsService handles sub-user group endpoints.
type SubUserGroupsService struct {
	client *Client
}

type CreateSubUserGroupParams struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateSubUserGroupParams struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// List returns all sub-user groups.
func (s *SubUserGroupsService) List(ctx context.Context) ([]SubUserGroup, error) {
	var result []SubUserGroup
	err := s.client.doRequest(ctx, "GET", "sub-user-groups", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new sub-user group.
func (s *SubUserGroupsService) Create(ctx context.Context, params CreateSubUserGroupParams) (*SubUserGroup, error) {
	var result SubUserGroup
	err := s.client.doRequest(ctx, "POST", "sub-user-groups", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a sub-user group by ID.
func (s *SubUserGroupsService) Get(ctx context.Context, id string) (*SubUserGroup, error) {
	var result SubUserGroup
	err := s.client.doRequest(ctx, "GET", "sub-user-groups/"+id, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a sub-user group.
func (s *SubUserGroupsService) Update(ctx context.Context, id string, params UpdateSubUserGroupParams) (*SubUserGroup, error) {
	var result SubUserGroup
	err := s.client.doRequest(ctx, "PUT", "sub-user-groups/"+id, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a sub-user group.
func (s *SubUserGroupsService) Delete(ctx context.Context, id string) error {
	return s.client.doRequest(ctx, "DELETE", "sub-user-groups/"+id, nil, nil)
}
