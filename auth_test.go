package proxyhat

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAuth_Register(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body RegisterParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("email = %q, want %q", body.Email, "test@example.com")
		}
		writePayload(w, RegisterResponse{
			Message:     "Account created",
			AccessToken: "tok_123",
			TokenType:   "bearer",
		})
	})

	resp, err := client.Auth.Register(context.Background(), RegisterParams{
		Name:                 "Test",
		Email:                "test@example.com",
		Password:             "password123",
		PasswordConfirmation: "password123",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.AccessToken != "tok_123" {
		t.Errorf("AccessToken = %q, want %q", resp.AccessToken, "tok_123")
	}
}

func TestAuth_Login(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, LoginResponse{
			AccessToken: "tok_456",
			TokenType:   "bearer",
		})
	})

	resp, err := client.Auth.Login(context.Background(), LoginParams{
		Email:    "test@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.AccessToken != "tok_456" {
		t.Errorf("AccessToken = %q, want %q", resp.AccessToken, "tok_456")
	}
}

func TestAuth_User(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writePayload(w, User{
			UUID:  "uuid-123",
			Name:  "Test User",
			Email: "test@example.com",
		})
	})

	user, err := client.Auth.User(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if user.UUID != "uuid-123" {
		t.Errorf("UUID = %q, want %q", user.UUID, "uuid-123")
	}
}

func TestAuth_Logout(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.Auth.Logout(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuth_SupportedProviders(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/supported-providers", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []SupportedProvider{{Name: "Google", Slug: "google"}})
	})

	providers, err := client.Auth.SupportedProviders(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(providers) != 1 || providers[0].Slug != "google" {
		t.Errorf("unexpected providers: %v", providers)
	}
}

func TestAuth_SocialAccounts(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/social-accounts", func(w http.ResponseWriter, r *http.Request) {
		writeData(w, []SocialAccount{{Provider: "google"}})
	})

	accounts, err := client.Auth.SocialAccounts(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) != 1 || accounts[0].Provider != "google" {
		t.Errorf("unexpected accounts: %v", accounts)
	}
}

func TestAuth_DisconnectSocial(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/social-accounts/google", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.Auth.DisconnectSocial(context.Background(), "google")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuth_OAuthRedirect(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/auth/google/redirect", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, map[string]string{"url": "https://accounts.google.com/..."})
	})

	result, err := client.Auth.OAuthRedirect(context.Background(), "google")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Error("expected non-nil result")
	}
}
