package proxyhat

import "context"

// CouponsService handles coupon endpoints.
type CouponsService struct {
	client *Client
}

type CouponParams struct {
	Code     string   `json:"code"`
	PlanID   *string  `json:"plan_id,omitempty"`
	OrderSum *float64 `json:"order_sum,omitempty"`
	Currency *string  `json:"currency,omitempty"`
}

// Validate validates a coupon code.
func (s *CouponsService) Validate(ctx context.Context, params CouponParams) (*CouponResponse, error) {
	var result CouponResponse
	err := s.client.doRequest(ctx, "POST", "coupon/validate", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Apply applies a coupon code.
func (s *CouponsService) Apply(ctx context.Context, params CouponParams) (*CouponResponse, error) {
	var result CouponResponse
	err := s.client.doRequest(ctx, "POST", "coupon/apply", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Redeem redeems a coupon code.
func (s *CouponsService) Redeem(ctx context.Context, code string) (*CouponResponse, error) {
	var result CouponResponse
	body := map[string]string{"code": code}
	err := s.client.doRequest(ctx, "POST", "coupon/redeem", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
