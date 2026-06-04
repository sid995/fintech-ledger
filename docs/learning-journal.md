# Learning Journal

Use this journal to write explanations in your own words. Do not copy the mentor's response verbatim. The mild discomfort is the educational part.

## Entry template

### YYYY-MM-DD: Topic

#### Explain the concept simply

#### Why it matters in a fintech backend

#### What I implemented

#### What failed or confused me

#### How I verified the behavior

#### Production trade-offs

#### Questions for the next session

### 2026-06-04: Phase 0 bootstrap

#### Explain the concept simply

A Go module is the root boundary for building and testing code in this repo. Packages group related `.go` files; a command is a `package main` with `func main()` that compiles to a binary under `cmd/`.

#### Why it matters in a fintech backend

Financial services need reproducible builds and small, reviewable changes. Phase 0 proves the toolchain before any ledger logic.

#### What I implemented

Initialized `github.com/siddharthkundu/fintech-ledger`, added `cmd/ledger-check` and `internal/version` with a banner test, plus root `README.md`.

#### What failed or confused me

Nothing blocked; `internal/` is allowed for small shared helpers even though domain code comes later.

#### How I verified the behavior

`go fmt`, `go vet`, `go test ./...`, and `go run ./cmd/ledger-check` all succeeded.

#### Production trade-offs

Real teams run the same checks in CI on every push; local discipline now avoids fixing hygiene later under time pressure.

#### Questions for the next session

When should value types use pointer receivers vs value receivers? (Phase 1 topic.)
