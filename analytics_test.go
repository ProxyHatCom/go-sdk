package proxyhat

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAnalytics_Traffic(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/traffic", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body AnalyticsParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.Period != "24h" {
			t.Errorf("period = %q, want %q", body.Period, "24h")
		}
		writePayload(w, TimeSeriesResponse{
			Labels: []string{"00:00", "01:00"},
			Data:   []int{100, 200},
		})
	})

	resp, err := client.Analytics.Traffic(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Labels) != 2 {
		t.Errorf("Labels length = %d, want 2", len(resp.Labels))
	}
	if resp.Data[1] != 200 {
		t.Errorf("Data[1] = %d, want 200", resp.Data[1])
	}
}

func TestAnalytics_TrafficTotal(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/traffic/period-total", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, TotalResponse{Total: 5000})
	})

	resp, err := client.Analytics.TrafficTotal(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Total != 5000 {
		t.Errorf("Total = %d, want 5000", resp.Total)
	}
}

func TestAnalytics_Requests(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, TimeSeriesResponse{
			Labels: []string{"00:00"},
			Data:   []int{42},
		})
	})

	resp, err := client.Analytics.Requests(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Data[0] != 42 {
		t.Errorf("Data[0] = %d, want 42", resp.Data[0])
	}
}

func TestAnalytics_RequestsTotal(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/requests/period-total", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, TotalResponse{Total: 1234})
	})

	resp, err := client.Analytics.RequestsTotal(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Total != 1234 {
		t.Errorf("Total = %d, want 1234", resp.Total)
	}
}

func TestAnalytics_DomainBreakdown(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/domain-breakdown", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, DomainBreakdownResponse{
			Items: []DomainBreakdownItem{
				{Domain: "example.com", Bandwidth: 1000, Requests: 50},
			},
		})
	})

	resp, err := client.Analytics.DomainBreakdown(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 1 || resp.Items[0].Domain != "example.com" {
		t.Errorf("unexpected items: %v", resp.Items)
	}
}
