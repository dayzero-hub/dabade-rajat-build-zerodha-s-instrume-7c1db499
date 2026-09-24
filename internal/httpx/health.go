// Package httpx holds HTTP handlers.
package httpx

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Health answers GET /healthz with a JSON body.
//
// It reports the DATABASE, not just the process. A health check that returns ok
// while its database is unreachable tells you the one thing you already knew —
// that the process is running — at the moment you needed the other answer.
func Health(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := map[string]string{"status": "ok", "database": "ok"}
		code := http.StatusOK
		if err := db.PingContext(r.Context()); err != nil {
			body["status"] = "degraded"
			body["database"] = err.Error()
			code = http.StatusServiceUnavailable
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(body)
	}
}
