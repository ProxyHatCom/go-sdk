package proxyhat

import (
	"context"
	"fmt"
	"net/url"
)

// LocationsService handles location endpoints.
type LocationsService struct {
	client *Client
}

// LocationParams are common query parameters for location endpoints.
type LocationParams struct {
	Limit          *int
	Offset         *int
	Name           *string
	ConnectionType *string
}

func (p *LocationParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Limit != nil {
		v.Set("limit", fmt.Sprintf("%d", *p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", fmt.Sprintf("%d", *p.Offset))
	}
	if p.Name != nil {
		v.Set("name", *p.Name)
	}
	if p.ConnectionType != nil {
		v.Set("connection_type", *p.ConnectionType)
	}
	return v
}

// Countries returns the list of available countries.
func (s *LocationsService) Countries(ctx context.Context, params *LocationParams) ([]Country, error) {
	var result []Country
	err := s.client.doRequestWithParams(ctx, "GET", "locations/countries", params.values(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RegionParams extends LocationParams with a country code filter.
type RegionParams struct {
	LocationParams
	CountryCode *string
}

func (p *RegionParams) values() url.Values {
	v := p.LocationParams.values()
	if p.CountryCode != nil {
		v.Set("country__code", *p.CountryCode)
	}
	return v
}

// Regions returns the list of available regions.
func (s *LocationsService) Regions(ctx context.Context, params *RegionParams) ([]Region, error) {
	if params == nil {
		params = &RegionParams{}
	}
	var result []Region
	err := s.client.doRequestWithParams(ctx, "GET", "locations/regions", params.values(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CityParams extends RegionParams with a region code filter.
type CityParams struct {
	RegionParams
	RegionCode *string
}

func (p *CityParams) values() url.Values {
	v := p.RegionParams.values()
	if p.RegionCode != nil {
		v.Set("region__code", *p.RegionCode)
	}
	return v
}

// Cities returns the list of available cities.
func (s *LocationsService) Cities(ctx context.Context, params *CityParams) ([]City, error) {
	if params == nil {
		params = &CityParams{}
	}
	var result []City
	err := s.client.doRequestWithParams(ctx, "GET", "locations/cities", params.values(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ISPs returns the list of available ISPs.
func (s *LocationsService) ISPs(ctx context.Context, params *RegionParams) ([]ISP, error) {
	if params == nil {
		params = &RegionParams{}
	}
	var result []ISP
	err := s.client.doRequestWithParams(ctx, "GET", "locations/isps", params.values(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZipcodeParams extends LocationParams with country and city code filters.
type ZipcodeParams struct {
	LocationParams
	CountryCode *string
	CityCode    *string
}

func (p *ZipcodeParams) values() url.Values {
	v := p.LocationParams.values()
	if p.CountryCode != nil {
		v.Set("country__code", *p.CountryCode)
	}
	if p.CityCode != nil {
		v.Set("city__code", *p.CityCode)
	}
	return v
}

// Zipcodes returns the list of available zipcodes.
func (s *LocationsService) Zipcodes(ctx context.Context, params *ZipcodeParams) ([]Zipcode, error) {
	if params == nil {
		params = &ZipcodeParams{}
	}
	var result []Zipcode
	err := s.client.doRequestWithParams(ctx, "GET", "locations/zipcodes", params.values(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
