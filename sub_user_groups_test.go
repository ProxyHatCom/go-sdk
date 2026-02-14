package proxyhat

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSubUserGroups_List(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-user-groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []SubUserGroup{{ID: "grp-1", Name: "Group 1", CreatedAt: "2024-01-01"}})
	})

	groups, err := client.SubUserGroups.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) != 1 || groups[0].ID != "grp-1" {
		t.Errorf("unexpected groups: %v", groups)
	}
}

func TestSubUserGroups_Create(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-user-groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body CreateSubUserGroupParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "New Group" {
			t.Errorf("name = %q, want %q", body.Name, "New Group")
		}
		writePayload(w, SubUserGroup{ID: "grp-new", Name: "New Group", CreatedAt: "2024-01-01"})
	})

	group, err := client.SubUserGroups.Create(context.Background(), CreateSubUserGroupParams{
		Name: "New Group",
	})
	if err != nil {
		t.Fatal(err)
	}
	if group.ID != "grp-new" {
		t.Errorf("ID = %q, want %q", group.ID, "grp-new")
	}
}

func TestSubUserGroups_Get(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-user-groups/grp-1", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, SubUserGroup{ID: "grp-1", Name: "Group 1", CreatedAt: "2024-01-01"})
	})

	group, err := client.SubUserGroups.Get(context.Background(), "grp-1")
	if err != nil {
		t.Fatal(err)
	}
	if group.Name != "Group 1" {
		t.Errorf("Name = %q, want %q", group.Name, "Group 1")
	}
}

func TestSubUserGroups_Update(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-user-groups/grp-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		writePayload(w, SubUserGroup{ID: "grp-1", Name: "Updated", CreatedAt: "2024-01-01"})
	})

	name := "Updated"
	group, err := client.SubUserGroups.Update(context.Background(), "grp-1", UpdateSubUserGroupParams{
		Name: &name,
	})
	if err != nil {
		t.Fatal(err)
	}
	if group.Name != "Updated" {
		t.Errorf("Name = %q, want %q", group.Name, "Updated")
	}
}

func TestSubUserGroups_Delete(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/sub-user-groups/grp-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.SubUserGroups.Delete(context.Background(), "grp-1")
	if err != nil {
		t.Fatal(err)
	}
}
