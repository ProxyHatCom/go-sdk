package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestEmail_RequestChange(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/email/request-change", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, EmailChangeResponse{Message: "Verification email sent"})
	})

	resp, err := client.Email.RequestChange(context.Background(), RequestEmailChangeParams{
		Email: "new@example.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "Verification email sent" {
		t.Errorf("Message = %q, want %q", resp.Message, "Verification email sent")
	}
}

func TestEmail_ConfirmChange(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/email/confirm-change", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, EmailChangeResponse{Message: "Email changed"})
	})

	resp, err := client.Email.ConfirmChange(context.Background(), "token-123")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "Email changed" {
		t.Errorf("Message = %q, want %q", resp.Message, "Email changed")
	}
}

func TestEmail_CancelChange(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/email/cancel-change", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, EmailChangeResponse{Message: "Cancelled"})
	})

	resp, err := client.Email.CancelChange(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "Cancelled" {
		t.Errorf("Message = %q, want %q", resp.Message, "Cancelled")
	}
}

func TestEmail_ResendVerification(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/email/resend-verification", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, EmailChangeResponse{Message: "Resent"})
	})

	resp, err := client.Email.ResendVerification(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "Resent" {
		t.Errorf("Message = %q, want %q", resp.Message, "Resent")
	}
}
