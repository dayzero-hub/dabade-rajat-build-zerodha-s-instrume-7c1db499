# Go + SQLite API — starting point

A runnable Go service with one endpoint. **Everything your brief asks you to build is missing on
purpose** — this is a starting point, not a partial answer.

## Run it

You need **git** and **Go 1.22 or newer**. Nothing else — no Docker, no database server.

```bash
go mod download
go test ./...          # green on the untouched skeleton
go run ./cmd/api       # then, in another terminal:
curl localhost:8080/healthz
```

`/healthz` answers JSON. If it does, you are ready to start ticket 2.

If `go test ./...` is **red before you have changed anything**, that is worth telling Raj about
rather than working around — a skeleton that arrives broken is our bug, not yours.

## What is here

```
cmd/api/main.go          the entry point: opens the database, serves /healthz
internal/store/store.go  opens (and creates) the SQLite file
internal/httpx/health.go the health handler
```

Two decisions in here that are worth knowing rather than discovering:

- **The SQLite driver is `modernc.org/sqlite`, which is pure Go.** The better-known
  `mattn/go-sqlite3` needs cgo and a C compiler, so it fails to build on a machine that has
  none — and that failure reads as "the skeleton is broken" on somebody's first day.
- **`/healthz` reports the database, not just the process.** A health check that says `ok` while
  its database is unreachable answers the one question you already knew the answer to.

The database file defaults to `data/app.db` and is git-ignored. Override with `DB_PATH`; override
the listen address with `ADDR`.

## What is NOT here

**Your brief's seeded data is not in this template, and neither is any harness it names.** This
repository is shared by several briefs that each need different data — different tables, different
CSVs — so it carries the parts that are the same for all of them and none of the parts that are
not.

If your brief's "What's provided" section describes a seeded database or a named helper you cannot
find here, **that is a real gap and not something you have missed**. Say so on the ticket: Raj
would rather hear it early than have you build against data you had to invent.

## Adding what you need

Migrations, tables and endpoints are yours. A reasonable order:

1. write the schema your brief needs, as SQL you can re-run
2. load the seed data your brief describes
3. add handlers under `internal/httpx`, wiring them in `cmd/api/main.go`
4. test each one — `internal/httpx/health_test.go` shows the shape

Commit on a branch and open a pull request; do not commit on `main`.
