package proxyhat

import (
	"context"
	"net/http"
	"testing"
)

func TestProfile_GetPreferences(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writePayload(w, Preferences{Data: map[string]any{"theme": "dark"}})
	})

	prefs, err := client.Profile.GetPreferences(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if prefs.Data["theme"] != "dark" {
		t.Errorf("theme = %v, want %q", prefs.Data["theme"], "dark")
	}
}

func TestProfile_UpdatePreferences(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		writePayload(w, Preferences{Data: map[string]any{"theme": "light"}})
	})

	prefs, err := client.Profile.UpdatePreferences(context.Background(), map[string]any{"theme": "light"})
	if err != nil {
		t.Fatal(err)
	}
	if prefs.Data["theme"] != "light" {
		t.Errorf("theme = %v, want %q", prefs.Data["theme"], "light")
	}
}

func TestProfile_ListAPIKeys(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/api-keys", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []APIKey{{ID: "key-1"}})
	})

	keys, err := client.Profile.ListAPIKeys(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 1 || keys[0].ID != "key-1" {
		t.Errorf("unexpected keys: %v", keys)
	}
}

func TestProfile_CreateAPIKey(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/api-keys", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		token := "sk_new_123"
		writePayload(w, APIKey{ID: "key-new", PlainTextToken: &token})
	})

	name := "My Key"
	key, err := client.Profile.CreateAPIKey(context.Background(), &name)
	if err != nil {
		t.Fatal(err)
	}
	if key.ID != "key-new" {
		t.Errorf("ID = %q, want %q", key.ID, "key-new")
	}
}

func TestProfile_DeleteAPIKey(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/api-keys/key-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.Profile.DeleteAPIKey(context.Background(), "key-1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestProfile_RegenerateAPIKey(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/profile/api-keys/key-1/regenerate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		token := "sk_regen_456"
		writePayload(w, APIKey{ID: "key-1", PlainTextToken: &token})
	})

	key, err := client.Profile.RegenerateAPIKey(context.Background(), "key-1")
	if err != nil {
		t.Fatal(err)
	}
	if key.PlainTextToken == nil || *key.PlainTextToken != "sk_regen_456" {
		t.Errorf("PlainTextToken = %v, want %q", key.PlainTextToken, "sk_regen_456")
	}
}
