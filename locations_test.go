package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestLocations_Countries(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/countries", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []Country{{Code: "US", Name: "United States", Availability: "high", ConnectionType: "residential"}})
	})

	countries, err := client.Locations.Countries(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(countries) != 1 || countries[0].Code != "US" {
		t.Errorf("unexpected countries: %v", countries)
	}
}

func TestLocations_Countries_WithParams(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/countries", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("limit"); got != "10" {
			t.Errorf("limit = %q, want %q", got, "10")
		}
		if got := r.URL.Query().Get("name"); got != "United" {
			t.Errorf("name = %q, want %q", got, "United")
		}
		writeData(w, []Country{})
	})

	limit := 10
	name := "United"
	_, err := client.Locations.Countries(context.Background(), &LocationParams{
		Limit: &limit,
		Name:  &name,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocations_Regions(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/regions", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []Region{{Code: "CA", Name: "California", CountryCode: "US"}})
	})

	regions, err := client.Locations.Regions(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(regions) != 1 || regions[0].Code != "CA" {
		t.Errorf("unexpected regions: %v", regions)
	}
}

func TestLocations_Cities(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/cities", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []City{{Code: "LA", Name: "Los Angeles", CountryCode: "US"}})
	})

	cities, err := client.Locations.Cities(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(cities) != 1 || cities[0].Code != "LA" {
		t.Errorf("unexpected cities: %v", cities)
	}
}

func TestLocations_ISPs(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/isps", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []ISP{{Code: "comcast", Name: "Comcast", CountryCode: "US"}})
	})

	isps, err := client.Locations.ISPs(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(isps) != 1 || isps[0].Code != "comcast" {
		t.Errorf("unexpected ISPs: %v", isps)
	}
}

func TestLocations_Zipcodes(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/locations/zipcodes", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []Zipcode{{Code: "90001", Name: "90001", CountryCode: "US"}})
	})

	zipcodes, err := client.Locations.Zipcodes(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(zipcodes) != 1 || zipcodes[0].Code != "90001" {
		t.Errorf("unexpected zipcodes: %v", zipcodes)
	}
}
