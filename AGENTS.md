# AGENTS.md - Go Project Blueprint

Canonical instructions for coding agents in this repository. Read this file first, then load only the context files needed for the task.

## Purpose

This repository is a Go project template. It does not contain an application yet. Its job is to provide:

- agent operating rules
- Go quality gates
- CI and pre-commit automation
- reusable scaffold for future Go services

Treat the repository itself as the product. Most changes here affect future projects that adopt this blueprint.

## Engineering Priorities

Use these priorities in order:

1. Correctness and safety.
2. Simplicity of design and implementation.
3. Test evidence for changed behavior.
4. Consistency with existing repo rules and automation.
5. Speed of implementation.

If two options both work, choose the one that is easier to read, easier to test, and easier for the next agent to extend.

## Read Order

Before editing:

1. Read this file.
2. Read [README.md](README.md).
3. Check the worktree with `git status --short`.
4. Read the smallest relevant context files under `.agents/context/`.
5. Make a short plan before editing.

Start with these context files:

- [.agents/context/project.md](.agents/context/project.md) for repository scope and current state.
- [.agents/context/architecture.md](.agents/context/architecture.md) for layout and dependency direction.
- [.agents/context/testing.md](.agents/context/testing.md) for commands and testing bar.
- [.agents/context/go-style.md](.agents/context/go-style.md) for coding conventions.
- [.agents/context/review-checklist.md](.agents/context/review-checklist.md) before handoff.

Use the prompt files under `.agents/prompts/` when the task matches them.

## Current Repository State

- The module path is still a placeholder: `github.com/your-org/your-project`.
- There is no `cmd/` tree yet.
- There is no `internal/` tree yet.
- There are no real application packages yet beyond the root scaffold package in [doc.go](doc.go).
- Many commands intentionally degrade gracefully in a fresh clone.

Do not invent missing application architecture. If a task needs product code, either scaffold it explicitly or update the template docs and rules that govern future code.

## Operating Rules

- Keep changes scoped. Prefer the smallest coherent change that improves the template.
- Do not overwrite user changes.
- Do not add broad always-loaded instructions when a narrower context file is enough.
- Prefer straightforward solutions over clever ones.
- Do not add a new abstraction, helper, or dependency unless it removes repeated or proven complexity.
- Keep repo facts in files, not in tool-specific chat history.
- Update docs when commands, expectations, structure, or workflow change.
- For production-code changes, add or update tests in the same change.
- For production behavior changes, write the test before the implementation when feasible. If you do not, explain why in the handoff.
- Explain new dependencies, generated artifacts, workflow permission changes, and public behavior changes in the final handoff.

## Plan First

Before editing, write a short plan that covers:

- files to create or update
- why each file is needed
- validation commands
- risks or assumptions

The plan must also state:

- why the chosen approach is the simplest correct one
- what test will prove the behavior change
- whether the test can be written before implementation

Proceed without waiting only for small, obvious, low-risk changes. Pause for human review if the change is broad, destructive, security-sensitive, architectural, or dependency-changing.

## Architecture Overview

Expected project layout for repositories that adopt this blueprint:

- `cmd/<binary>/main.go` for entrypoints and wiring
- `internal/config` for config loading
- `internal/models` for data structures and validation
- `internal/service` for business logic
- `internal/store` for persistence and integrations
- `internal/server` for HTTP, gRPC, or MCP transport
- `docs/` for durable reference docs
- `test/` for integration suites

Dependency direction:

- `cmd` -> `server` -> `service` -> `store` -> `models`
- Dependencies flow downward only.
- `service` should depend on store interfaces owned by the consumer package.

See [.agents/context/architecture.md](.agents/context/architecture.md) for the concise version agents should actually load.

## Build, Test, and Lint

Primary commands:

- `make ci-check`
- `make build`
- `go vet ./...`
- `go test ./...`
- `make lint`
- `git diff --check`
- `pre-commit run --all-files` when `pre-commit` is installed

Important template behavior:

- `make build` skips gracefully when `cmd/` does not exist.
- `make test` and `go test ./...` operate on the root scaffold package until real packages exist.
- `make lint` depends on `golangci-lint` and may fail if the environment cannot write to its cache.

See [.agents/context/testing.md](.agents/context/testing.md) for the testing and validation bar.

## Go Standards

- Use `context.Context` as the first parameter for cancellable or I/O work.
- Use `log/slog` for logging.
- Wrap errors with `%w`.
- Keep handlers thin and business logic in `service`.
- Prefer standard library helpers such as `cmp`, `maps`, and `slices`.
- Prefer existing package boundaries and local helpers over new abstractions.
- Document exported symbols.
- Use table-driven tests with `t.Run()` and `t.Parallel()` where safe.
- Use `t.Context()` in tests when appropriate.

See [.agents/context/go-style.md](.agents/context/go-style.md) for the repo-specific style rules.

## Forbidden Changes

- Do not add secrets, credentials, or machine-specific private data.
- Do not add network calls to unit tests.
- Do not introduce `database/sql` as the default persistence abstraction in this blueprint.
- Do not add top-level dependencies without clear justification.
- Do not add abstraction layers "for future flexibility" without a current concrete need.
- Do not replace direct code with indirection unless at least one current pain point is removed.
- Do not create duplicate canonical agent files such as `CODEX.md`.
- Do not broaden repository instructions with vague policy that cannot be enforced or reviewed.
- Do not do unrelated refactors while touching template files.

## Definition Of Done

A task is done when:

- the requested scope is complete
- the change is specific to this repository and useful to future agents
- the chosen implementation is the simplest correct version that fits current needs
- tests/docs are updated when required
- validation commands were run at the appropriate scope
- skipped or failing checks are explained clearly
- instructions remain concise and non-duplicative

## Review Checklist

Before handoff:

1. Re-read changed files for vague or duplicated guidance.
2. Confirm `AGENTS.md` points to smaller context files instead of absorbing everything.
3. Confirm commands in docs match `Makefile` and CI behavior.
4. Confirm tests were added for production-code changes.
5. Confirm complexity was justified and unnecessary abstraction was avoided.
6. Confirm tests prove the changed behavior, not just raise coverage.
7. Run the relevant validation commands and record failures accurately.
8. Summarize changed files, validation, remaining risks, and suggested next improvements.

Use [.agents/context/review-checklist.md](.agents/context/review-checklist.md) as the final self-review pass.
