# Fintech Ledger (learning project)

Incremental Go backend for a wallet-style ledger: double-entry bookkeeping, PostgreSQL, CQRS, idempotency, and transactional outbox—built phase by phase for understanding, not speed.

This repository follows the curriculum in [AGENTS.md](AGENTS.md). Phase 0 only verifies the Go toolchain.

## Prerequisites

- Go 1.22+ (1.21+ minimum)
- Git

Docker is optional until Phase 5.

## Run the Phase 0 smoke binary

```bash
go run ./cmd/ledger-check
```

## Verify formatting, vet, and tests

```bash
go fmt ./...
go vet ./...
go test ./...
```

## Learning docs

- [docs/learning-progress.md](docs/learning-progress.md)
- [docs/learning-journal.md](docs/learning-journal.md)
- [docs/architecture-decisions/](docs/architecture-decisions/)