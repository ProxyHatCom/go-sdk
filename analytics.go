package proxyhat

import "context"

// AnalyticsService handles analytics endpoints.
type AnalyticsService struct {
	client *Client
}

// AnalyticsParams are common parameters for analytics endpoints.
type AnalyticsParams struct {
	Period    string  `json:"period"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate   *string `json:"end_date,omitempty"`
}

func defaultAnalyticsParams(params *AnalyticsParams) *AnalyticsParams {
	if params == nil {
		return &AnalyticsParams{Period: "24h"}
	}
	if params.Period == "" {
		params.Period = "24h"
	}
	return params
}

// Traffic returns traffic time series data.
func (s *AnalyticsService) Traffic(ctx context.Context, params *AnalyticsParams) (*TimeSeriesResponse, error) {
	params = defaultAnalyticsParams(params)
	var result TimeSeriesResponse
	err := s.client.doRequest(ctx, "POST", "traffic", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TrafficTotal returns total traffic for the period.
func (s *AnalyticsService) TrafficTotal(ctx context.Context, params *AnalyticsParams) (*TotalResponse, error) {
	params = defaultAnalyticsParams(params)
	var result TotalResponse
	err := s.client.doRequest(ctx, "POST", "traffic/period-total", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Requests returns request time series data.
func (s *AnalyticsService) Requests(ctx context.Context, params *AnalyticsParams) (*TimeSeriesResponse, error) {
	params = defaultAnalyticsParams(params)
	var result TimeSeriesResponse
	err := s.client.doRequest(ctx, "POST", "requests", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestsTotal returns total requests for the period.
func (s *AnalyticsService) RequestsTotal(ctx context.Context, params *AnalyticsParams) (*TotalResponse, error) {
	params = defaultAnalyticsParams(params)
	var result TotalResponse
	err := s.client.doRequest(ctx, "POST", "requests/period-total", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DomainBreakdown returns domain breakdown analytics.
func (s *AnalyticsService) DomainBreakdown(ctx context.Context, params *AnalyticsParams) (*DomainBreakdownResponse, error) {
	params = defaultAnalyticsParams(params)
	var result DomainBreakdownResponse
	err := s.client.doRequest(ctx, "POST", "domain-breakdown", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
