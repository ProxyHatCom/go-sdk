package proxyhat

import "context"

// PlansService handles plan endpoints.
type PlansService struct {
	client *Client
}

// ListRegular returns all regular (one-time) plans.
func (s *PlansService) ListRegular(ctx context.Context) ([]RegularPlan, error) {
	var result []RegularPlan
	err := s.client.doRequest(ctx, "GET", "regular-options", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSubscriptions returns all subscription plans.
func (s *PlansService) ListSubscriptions(ctx context.Context) ([]SubscriptionPlan, error) {
	var result []SubscriptionPlan
	err := s.client.doRequest(ctx, "GET", "subscription-plans", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRegular returns a regular plan by name.
func (s *PlansService) GetRegular(ctx context.Context, name string) (*RegularPlan, error) {
	var result RegularPlan
	err := s.client.doRequest(ctx, "GET", "plans/regular/"+name, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSubscription returns a subscription plan by name.
func (s *PlansService) GetSubscription(ctx context.Context, name string) (*SubscriptionPlan, error) {
	var result SubscriptionPlan
	err := s.client.doRequest(ctx, "GET", "plans/subscription/"+name, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PricingRegular returns regular plan pricing.
func (s *PlansService) PricingRegular(ctx context.Context) ([]any, error) {
	var result []any
	err := s.client.doRequest(ctx, "GET", "pricing/regular", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PricingSubscriptions returns subscription plan pricing.
func (s *PlansService) PricingSubscriptions(ctx context.Context) ([]any, error) {
	var result []any
	err := s.client.doRequest(ctx, "GET", "pricing/subscriptions", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
