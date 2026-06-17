# Project Context

## What This Repo Is

`go-project-blueprint` is a template repository for future Go projects. The repository itself is mostly policy, automation, and scaffold rather than application logic.

The primary output of this repo is a reusable starting point with aligned agent instructions, human contribution guidance, and reproducible quality gates.

## Current State

- `go.mod` still uses the placeholder module path `github.com/your-org/your-project`.
- There is no `cmd/` directory yet.
- There are no `internal/*` packages yet.
- The root [doc.go](../../doc.go) package exists so Go tooling has something to operate on in a fresh clone.
- `README.md`, `AGENTS.md`, `Makefile`, CI, lint config, and pre-commit config are the main assets.

## What Agents Should Optimize For

- Improve the template, not an imaginary application.
- Keep instructions concrete, short, and reviewable.
- Prefer durable repo files over long always-loaded guidance.
- Preserve parity between prose, `Makefile`, and CI.

## Main Risks

- Writing generic guidance that could apply to any repo.
- Duplicating the same rule across `AGENTS.md`, bridge files, README, and prompt files.
- Documenting commands that do not match actual template behavior.
- Assuming packages or binaries exist when this template may still be empty.
