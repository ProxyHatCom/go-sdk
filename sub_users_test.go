package proxyhat

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSubUsers_List(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []SubUser{{UUID: "su-1", ProxyUsername: "user1", CreatedAt: "2024-01-01"}})
	})

	users, err := client.SubUsers.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 || users[0].UUID != "su-1" {
		t.Errorf("unexpected users: %v", users)
	}
}

func TestSubUsers_Create(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body CreateSubUserParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.ProxyPassword != "pass123" {
			t.Errorf("proxy_password = %q, want %q", body.ProxyPassword, "pass123")
		}
		writePayload(w, SubUser{UUID: "su-new", ProxyUsername: "newuser", CreatedAt: "2024-01-01"})
	})

	user, err := client.SubUsers.Create(context.Background(), CreateSubUserParams{
		ProxyPassword: "pass123",
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.UUID != "su-new" {
		t.Errorf("UUID = %q, want %q", user.UUID, "su-new")
	}
}

func TestSubUsers_Get(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/su-1", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, SubUser{UUID: "su-1", ProxyUsername: "user1", CreatedAt: "2024-01-01"})
	})

	user, err := client.SubUsers.Get(context.Background(), "su-1")
	if err != nil {
		t.Fatal(err)
	}
	if user.UUID != "su-1" {
		t.Errorf("UUID = %q, want %q", user.UUID, "su-1")
	}
}

func TestSubUsers_Update(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/su-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		name := "updated"
		writePayload(w, SubUser{UUID: "su-1", Name: &name, CreatedAt: "2024-01-01"})
	})

	name := "updated"
	user, err := client.SubUsers.Update(context.Background(), "su-1", UpdateSubUserParams{
		Name: &name,
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.Name == nil || *user.Name != "updated" {
		t.Errorf("Name = %v, want %q", user.Name, "updated")
	}
}

func TestSubUsers_Delete(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/su-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.SubUsers.Delete(context.Background(), "su-1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubUsers_ResetUsage(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/reset/usage", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, ResetUsageResponse{Reset: 2})
	})

	resp, err := client.SubUsers.ResetUsage(context.Background(), []string{"su-1", "su-2"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Reset != 2 {
		t.Errorf("Reset = %d, want 2", resp.Reset)
	}
}

func TestSubUsers_BulkDelete(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/bulk-delete", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, BulkDeleteResponse{Requested: 2, Deleted: 2})
	})

	resp, err := client.SubUsers.BulkDelete(context.Background(), []string{"su-1", "su-2"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Deleted != 2 {
		t.Errorf("Deleted = %d, want 2", resp.Deleted)
	}
}

func TestSubUsers_BulkMoveToGroup(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-users/bulk-move-to-group", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		writePayload(w, map[string]string{"message": "ok"})
	})

	groupID := "grp-1"
	_, err := client.SubUsers.BulkMoveToGroup(context.Background(), []string{"su-1"}, &groupID)
	if err != nil {
		t.Fatal(err)
	}
}
