// Package store opens the SQLite database the service reads.
//
// The driver is modernc.org/sqlite, which is PURE GO — deliberately, not by
// accident. The common alternative (mattn/go-sqlite3) needs cgo and a C
// toolchain, so it fails to build on a machine with no compiler installed.
// That failure looks like "the skeleton is broken" to somebody on their first
// day, which is the one thing a starting point must never do.
package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// Open opens (and creates if absent) the SQLite database at path.
func Open(path string) (*sql.DB, error) {
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("create %s: %w", dir, err)
		}
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// sql.Open is lazy — it does not touch the file. Ping here so a bad path is
	// an error at startup rather than on the first request.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping %s: %w", path, err)
	}
	return db, nil
}
