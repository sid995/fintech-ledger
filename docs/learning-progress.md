# Learning Progress

## Current status

- Current phase: Phase 1
- Current lesson: Not started
- Last completed checkpoint: Phase 0 — Go module and tooling smoke check
- Current blockers: None

## Phase tracker

| Phase | Topic | Status | Notes |
|---|---|---|---|
| 0 | Environment and learning discipline | Complete | Go module, cmd/ledger-check, internal/version test |
| 1 | Go foundations | Not started | |
| 2 | Fintech domain modeling | Not started | |
| 3 | Clean Architecture and ports | Not started | |
| 4 | CQRS handlers | Not started | |
| 5 | PostgreSQL schema design | Not started | |
| 6 | pgx and repositories | Not started | |
| 7 | Transaction safety | Not started | |
| 8 | Idempotency | Not started | |
| 9 | HTTP API | Not started | |
| 10 | Transactional outbox | Not started | |
| 11 | Testing strategy | Not started | |
| 12 | Docker Compose | Not started | |
| 13 | OpenAPI and CI | Not started | |
| 14 | Production hardening | Not started | |
| 15 | Parallel technology comparisons | Not started | |
| 16 | Capstone extensions | Not started | |

## Completed lessons

### 2026-06-04: Phase 0, Developer environment and learning discipline

- Concepts understood: Go module path, package vs command (`package main`), `cmd/` layout, `go fmt` / `go vet` / `go test`
- Code completed: `go.mod`, `README.md`, `cmd/ledger-check`, `internal/version` with table-style unit test
- Tests completed: `go test ./...` (version banner test)
- Commands run: `go mod init`, `go fmt ./...`, `go vet ./...`, `go test ./...`, `go run ./cmd/ledger-check`
- Design decisions: module path `github.com/siddharthkundu/fintech-ledger`; smoke binary named `ledger-check`
- Unresolved questions: none for Phase 0
- Next step: Phase 1 — Go foundations exercises under `learning/`

Add one entry after each lesson:

### YYYY-MM-DD: Phase X, Lesson title

- Concepts understood:
- Code completed:
- Tests completed:
- Commands run:
- Design decisions:
- Unresolved questions:
- Next step:
