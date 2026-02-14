package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestTwoFactor_Status(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writePayload(w, TwoFactorStatus{Enabled: true})
	})

	status, err := client.TwoFactor.Status(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !status.Enabled {
		t.Error("Enabled = false, want true")
	}
}

func TestTwoFactor_Enable(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/enable", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, TwoFactorEnableResponse{
			QR:            "data:image/png;base64,...",
			Secret:        "JBSWY3DPEHPK3PXP",
			RecoveryCodes: []string{"code1", "code2"},
		})
	})

	resp, err := client.TwoFactor.Enable(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Secret != "JBSWY3DPEHPK3PXP" {
		t.Errorf("Secret = %q, want %q", resp.Secret, "JBSWY3DPEHPK3PXP")
	}
}

func TestTwoFactor_Confirm(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/confirm", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, map[string]string{"message": "confirmed"})
	})

	_, err := client.TwoFactor.Confirm(context.Background(), "123456")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTwoFactor_Disable(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/disable", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, map[string]string{"message": "disabled"})
	})

	_, err := client.TwoFactor.Disable(context.Background(), "123456")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTwoFactor_QRCode(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/qr-code", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, TwoFactorEnableResponse{
			QR:     "data:image/png;base64,...",
			Secret: "SECRET",
		})
	})

	resp, err := client.TwoFactor.QRCode(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Secret != "SECRET" {
		t.Errorf("Secret = %q, want %q", resp.Secret, "SECRET")
	}
}

func TestTwoFactor_RecoveryCodes(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/recovery-codes", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, RecoveryCodes{Codes: []string{"abc", "def"}})
	})

	codes, err := client.TwoFactor.RecoveryCodes(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(codes.Codes) != 2 {
		t.Errorf("len(Codes) = %d, want 2", len(codes.Codes))
	}
}

func TestTwoFactor_DisableByRecovery(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/2fa/disable-by-recovery-code", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, map[string]string{"message": "disabled"})
	})

	_, err := client.TwoFactor.DisableByRecovery(context.Background(), "recovery-code-1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTwoFactor_ChangePassword(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/password", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, map[string]string{"message": "changed"})
	})

	_, err := client.TwoFactor.ChangePassword(context.Background(), ChangePasswordParams{
		CurrentPassword:      "old123",
		Password:             "new456",
		PasswordConfirmation: "new456",
	})
	if err != nil {
		t.Fatal(err)
	}
}
