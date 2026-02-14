package proxyhat

import (
	"context"
	"io"
	"net/http"
	"testing"
)

func TestPayments_List(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		amount := 50.0
		writeData(w, []Payment{{ID: "pay-1", Type: "regular", Status: "completed", Amount: &amount}})
	})

	payments, err := client.Payments.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(payments) != 1 || payments[0].ID != "pay-1" {
		t.Errorf("unexpected payments: %v", payments)
	}
}

func TestPayments_Create(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, PaymentCreateResponse{Success: true, PaymentID: "pay-new"})
	})

	resp, err := client.Payments.Create(context.Background(), CreatePaymentParams{
		Type:               "regular",
		PlanID:             "plan-1",
		CryptocurrencyCode: "BTC",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.PaymentID != "pay-new" {
		t.Errorf("PaymentID = %q, want %q", resp.PaymentID, "pay-new")
	}
}

func TestPayments_Get(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments/pay-1", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, PaymentDetails{
			PayAddress:   "1A1zP1...",
			CryptoAmount: 0.001,
			AmountUSD:    50.0,
			Status:       "pending",
			Crypto:       CryptoInfo{Code: "BTC", Currency: "Bitcoin", Network: "mainnet"},
		})
	})

	details, err := client.Payments.Get(context.Background(), "pay-1")
	if err != nil {
		t.Fatal(err)
	}
	if details.Status != "pending" {
		t.Errorf("Status = %q, want %q", details.Status, "pending")
	}
}

func TestPayments_Check(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments/pay-1/check", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, PaymentDetails{
			PayAddress:   "1A1zP1...",
			CryptoAmount: 0.001,
			AmountUSD:    50.0,
			Status:       "completed",
			Crypto:       CryptoInfo{Code: "BTC", Currency: "Bitcoin", Network: "mainnet"},
		})
	})

	details, err := client.Payments.Check(context.Background(), "pay-1")
	if err != nil {
		t.Fatal(err)
	}
	if details.Status != "completed" {
		t.Errorf("Status = %q, want %q", details.Status, "completed")
	}
}

func TestPayments_Invoice(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments/pay-1/invoice", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("format"); got != "pdf" {
			t.Errorf("format = %q, want %q", got, "pdf")
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Write([]byte("%PDF-1.4 fake content"))
	})

	resp, err := client.Payments.Invoice(context.Background(), "pay-1", "pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "%PDF-1.4 fake content" {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestPayments_Cryptocurrencies(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/payments/cryptocurrencies", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []Cryptocurrency{{Code: "BTC", Currency: "Bitcoin", Network: "mainnet"}})
	})

	cryptos, err := client.Payments.Cryptocurrencies(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(cryptos) != 1 || cryptos[0].Code != "BTC" {
		t.Errorf("unexpected cryptos: %v", cryptos)
	}
}
