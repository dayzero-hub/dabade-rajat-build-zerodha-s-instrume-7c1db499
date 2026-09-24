// Command api is the HTTP entry point.
//
// It starts, answers /healthz, and does nothing else on purpose — the endpoints
// your brief asks for are yours to add. See README.md.
package main

import (
	"log"
	"net/http"
	"os"

	"dayzer0/be-go-sqlite-api/internal/httpx"
	"dayzer0/be-go-sqlite-api/internal/store"
)

func main() {
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "data/app.db"
	}
	db, err := store.Open(path)
	if err != nil {
		// Fatal on purpose: a service that starts without its database answers
		// every request wrongly, which is harder to diagnose than not starting.
		log.Fatalf("open %s: %v", path, err)
	}
	defer db.Close()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", httpx.Health(db))

	log.Printf("listening on %s (db %s)", addr, path)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
