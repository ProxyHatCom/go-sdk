package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestCoupons_Validate(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/coupon/validate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		discount := 10.0
		writePayload(w, CouponResponse{
			Success: true,
			Coupon: &Coupon{
				ID:       "c-1",
				Code:     "SAVE10",
				Type:     "percentage",
				Discount: &discount,
			},
		})
	})

	resp, err := client.Coupons.Validate(context.Background(), CouponParams{Code: "SAVE10"})
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Error("Success = false, want true")
	}
	if resp.Coupon == nil || resp.Coupon.Code != "SAVE10" {
		t.Errorf("unexpected coupon: %v", resp.Coupon)
	}
}

func TestCoupons_Apply(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/coupon/apply", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, CouponResponse{Success: true})
	})

	resp, err := client.Coupons.Apply(context.Background(), CouponParams{Code: "SAVE10"})
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Error("Success = false, want true")
	}
}

func TestCoupons_Redeem(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/coupon/redeem", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, CouponResponse{Success: true})
	})

	resp, err := client.Coupons.Redeem(context.Background(), "FREECODE")
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Error("Success = false, want true")
	}
}
