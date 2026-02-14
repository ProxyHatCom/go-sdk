package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestPlans_ListRegular(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/regular-options", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []RegularPlan{{ID: "plan-1", Name: "Basic", GB: 10, PricePerGB: 5.0, PriceTotal: 50.0, Currency: "USD"}})
	})

	plans, err := client.Plans.ListRegular(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(plans) != 1 || plans[0].Name != "Basic" {
		t.Errorf("unexpected plans: %v", plans)
	}
}

func TestPlans_ListSubscriptions(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/subscription-plans", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []SubscriptionPlan{{ID: "sub-1", Name: "Pro", GB: 100, Period: "monthly"}})
	})

	plans, err := client.Plans.ListSubscriptions(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(plans) != 1 || plans[0].Name != "Pro" {
		t.Errorf("unexpected plans: %v", plans)
	}
}

func TestPlans_GetRegular(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/plans/regular/basic", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, RegularPlan{ID: "plan-1", Name: "basic", GB: 10, Currency: "USD"})
	})

	plan, err := client.Plans.GetRegular(context.Background(), "basic")
	if err != nil {
		t.Fatal(err)
	}
	if plan.Name != "basic" {
		t.Errorf("Name = %q, want %q", plan.Name, "basic")
	}
}

func TestPlans_GetSubscription(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/plans/subscription/pro", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, SubscriptionPlan{ID: "sub-1", Name: "pro", Period: "monthly"})
	})

	plan, err := client.Plans.GetSubscription(context.Background(), "pro")
	if err != nil {
		t.Fatal(err)
	}
	if plan.Period != "monthly" {
		t.Errorf("Period = %q, want %q", plan.Period, "monthly")
	}
}

func TestPlans_PricingRegular(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/pricing/regular", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []map[string]any{{"tier": "basic"}})
	})

	pricing, err := client.Plans.PricingRegular(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(pricing) != 1 {
		t.Errorf("len(pricing) = %d, want 1", len(pricing))
	}
}

func TestPlans_PricingSubscriptions(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/pricing/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []map[string]any{{"tier": "pro"}})
	})

	pricing, err := client.Plans.PricingSubscriptions(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(pricing) != 1 {
		t.Errorf("len(pricing) = %d, want 1", len(pricing))
	}
}
