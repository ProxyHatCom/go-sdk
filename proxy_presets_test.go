package proxyhat

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestProxyPresets_List(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/proxy-presets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		writeData(w, []ProxyPreset{{ID: "pp-1", Name: "Preset 1", Data: map[string]any{"key": "val"}}})
	})

	presets, err := client.ProxyPresets.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(presets) != 1 || presets[0].ID != "pp-1" {
		t.Errorf("unexpected presets: %v", presets)
	}
}

func TestProxyPresets_Create(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/proxy-presets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body CreateProxyPresetParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "New Preset" {
			t.Errorf("name = %q, want %q", body.Name, "New Preset")
		}
		writePayload(w, ProxyPreset{ID: "pp-new", Name: "New Preset", Data: body.Data})
	})

	preset, err := client.ProxyPresets.Create(context.Background(), CreateProxyPresetParams{
		Name: "New Preset",
		Data: map[string]any{"country": "US"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if preset.ID != "pp-new" {
		t.Errorf("ID = %q, want %q", preset.ID, "pp-new")
	}
}

func TestProxyPresets_Get(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/proxy-presets/pp-1", func(w http.ResponseWriter, r *http.Request) {
		writePayload(w, ProxyPreset{ID: "pp-1", Name: "Preset 1", Data: map[string]any{}})
	})

	preset, err := client.ProxyPresets.Get(context.Background(), "pp-1")
	if err != nil {
		t.Fatal(err)
	}
	if preset.Name != "Preset 1" {
		t.Errorf("Name = %q, want %q", preset.Name, "Preset 1")
	}
}

func TestProxyPresets_Update(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/proxy-presets/pp-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		writePayload(w, ProxyPreset{ID: "pp-1", Name: "Updated", Data: map[string]any{}})
	})

	name := "Updated"
	preset, err := client.ProxyPresets.Update(context.Background(), "pp-1", UpdateProxyPresetParams{
		Name: &name,
	})
	if err != nil {
		t.Fatal(err)
	}
	if preset.Name != "Updated" {
		t.Errorf("Name = %q, want %q", preset.Name, "Updated")
	}
}

func TestProxyPresets_Delete(t *testing.T) {
	client, mux, cleanup := setupTest()
	defer cleanup()

	mux.HandleFunc("/proxy-presets/pp-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	})

	err := client.ProxyPresets.Delete(context.Background(), "pp-1")
	if err != nil {
		t.Fatal(err)
	}
}
