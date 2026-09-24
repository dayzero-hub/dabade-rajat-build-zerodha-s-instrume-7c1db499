package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"dayzer0/be-go-sqlite-api/internal/store"
)

// This test is the one the setup ticket asks you to run. It is deliberately a
// REAL test rather than a placeholder: a green suite that asserts nothing is
// indistinguishable from a broken one, and the ticket tells you green means the
// skeleton is fine.
func TestHealthzReportsTheDatabase(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer db.Close()

	rec := httptest.NewRecorder()
	Health(db)(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", rec.Code, rec.Body.String())
	}
	var got map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("body is not JSON (%v): %s", err, rec.Body.String())
	}
	if got["status"] != "ok" || got["database"] != "ok" {
		t.Errorf("got %v, want status=ok database=ok", got)
	}
}

// The degraded path matters more than the healthy one: it is what tells you the
// check is actually reading the database rather than returning a constant.
func TestHealthzSaysDegradedWhenTheDatabaseIsGone(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	db.Close() // pings now fail

	rec := httptest.NewRecorder()
	Health(db)(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", rec.Code)
	}
	var got map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("body is not JSON: %s", rec.Body.String())
	}
	if got["status"] != "degraded" {
		t.Errorf("status = %q, want degraded", got["status"])
	}
}
