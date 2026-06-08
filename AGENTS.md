# AGENTS.md — Go Project Blueprint

> **The canonical reference for every coding agent (Claude, Cursor, Qwen, …) working in a Go project that adopted this blueprint.**
> Read this file end-to-end before you write a single line of code. Every section is load-bearing.

---

## What This Is

This repository is a **template** — the agent configuration files (`AGENTS.md`, linters, CI, pre-commit hooks, `Makefile`) that any Go project can adopt to give coding agents the best possible context. **Fork it, customize it, ship it.** The included files are not the application — they are the scaffold the application grows inside.

---

## For Coding Agents

You are a coding agent working in a Go project that adopted this blueprint. Before you start:

1. **Read this entire file.** Every section is load-bearing.
2. **Run `make ci-check`** to confirm the toolchain is healthy.
3. **Read [`README.md`](README.md)** for project-specific context, domain, and stack.
4. **Read the "Project Identity" section below** for the *specific* project you're in.
5. **Follow the Test Mandate.** Tests are not optional.

If a human reviewer's feedback conflicts with this file, **defer to the human** — but propose a change to this file so the rule becomes explicit.

---

## For Humans Forking This Template

Setup checklist (do these once, right after `git clone`):

1. Update `module` in `go.mod` to your module path (e.g. `github.com/your-org/your-project`).
2. Update `BINARY` in the `Makefile` to your binary name (e.g. `BINARY := bin/myapp`).
3. Update the `-local` flag in `.pre-commit-config.yaml` to your module path.
4. Replace the placeholder module path in `.golangci.yml` exclusions (if any) and `run` config.
5. Update the **Project Identity** section below with 1–2 paragraphs about your project.
6. Add your entry point at `cmd/<binary>/main.go`.
7. Update copyright in `LICENSE` if not the template author.
8. Install the pre-commit hook: `pre-commit install`.
9. Run `make ci-check` once to confirm everything is green.

---

## Project Identity

> *Replace this section with 1–2 paragraphs describing your specific project: what it does, who it serves, what stack it uses, what its boundaries are. This is the first thing an agent reads to ground itself.*

| Field | Value |
|-------|-------|
| **Project name** | `go-project-blueprint` (rename me) |
| **Purpose** | Industry-best agent configuration files for Go projects |
| **Language** | Go 1.25+ |
| **Module** | `github.com/your-org/your-project` (replace me) |
| **Runtime** | Single static binary, system-native dependencies |
| **License** | MIT |

---

## Naming Conventions

Strict. Non-negotiable for new code. Reviewers will flag violations.

| Element | Convention | Example |
|---|---|---|
| Files | `snake_case.go` | `agent_store.go` |
| Test files | `<file>_test.go` | `agent_store_test.go` |
| Types (exported) | `PascalCase` | `AgentStore` |
| Types (unexported) | `camelCase` | `agentCache` |
| Interfaces | `-er` suffix or descriptive noun | `Reader`, `AgentStore` |
| Methods (exported) | `PascalCase` | `GetAgent` |
| Methods (unexported) | `camelCase` | `fetchFromCache` |
| Variables (exported) | `PascalCase` | `DefaultTimeout` |
| Variables (unexported) | `camelCase` | `maxRetries` |
| Constants | `PascalCase` (preferred) or `SCREAMING_SNAKE` | `MaxRetries` |
| Error vars | `ErrXxx` | `ErrNotFound`, `ErrTimeout` |
| Error types | `XxxError` | `ValidationError`, `NotFoundError` |
| Packages | short, lowercase, singular, no underscores | `agent`, `store`, `http` |

**Acronyms: ALL CAPS, consistent.** `HTTP`, `ID`, `URL`, `JSON`, `API`, `MCP`, `SQL`, `YAML`, `TLS`, `DNS`, `URI`, `UUID`, `gRPC`, `SSH`, `TCP`, `UDP`. Do **not** mix `Http` and `HTTP` in the same project. Do **not** write `userId` — write `userID`.

---

## Project Structure

The default Go project layout. Adopt it unless you have a strong reason not to.

```
.
├── cmd/<binary>/main.go     ← entry point: flag parse, DI, graceful shutdown
├── internal/                ← private packages (not importable externally)
│   ├── config/              ← configuration loading (YAML, env, flags)
│   ├── models/              ← data structs, typed constants, validation
│   ├── service/             ← business logic (depends on store interfaces)
│   │   └── storemock/       ← manual mocks for unit tests of service
│   ├── store/               ← data access (DB, file, network)
│   │   └── migrations/      ← SQL migrations (if applicable)
│   └── server/              ← transport layer (HTTP, gRPC, MCP)
├── deploy/                  ← systemd units, deploy scripts, configs
├── docs/                    ← reference documentation
├── test/                    ← integration & load tests
├── migrations/              ← top-level SQL migrations (alt to internal/store/migrations)
├── .github/workflows/       ← CI pipelines
├── AGENTS.md                ← this file
├── Makefile                 ← development commands
├── go.mod
├── go.sum
├── tools.go                 ← pinned tool dependencies
├── .golangci.yml
├── .pre-commit-config.yaml
├── .editorconfig
├── .gitignore
├── LICENSE
└── README.md
```

**Layer rules (no exceptions):**

- `cmd/` → `server/` → `service/` → `store/` → `models/`
- Imports flow **downward only.** Never circular.
- `service/` depends on `store/` **interfaces**, not concrete types. Mocks live in `service/storemock/`.
- `models/` has no dependencies on other `internal/` packages.
- `cmd/<binary>/main.go` is the **only** place that wires concrete implementations together.

---

## Quality Gates

All gates must pass before any code is merged. The full pipeline `make ci-check` runs `tidy → build → vet → lint → test`. **This is what CI runs and what you run locally before committing.**

| Gate | Command | Required |
|---|---|---|
| Format | `gofmt -l .` | clean |
| Imports | `goimports -l -local <module-path> .` | clean |
| Vet | `go vet ./...` | clean |
| Lint | `make lint` (golangci-lint, 14 linters) | clean |
| Build | `make build` | success |
| Test | `make test` (race + shuffle + count=1) | passes |
| Coverage | `make test-coverage` | meets layer targets |
| Deadcode | `make deadcode` | no unused funcs |
| Vulns | `make vulncheck` | no known vulns |
| CI | `make ci-check` | green |

---

## 🚨 Test Mandate (Hard Requirement)

**Every code-writing task MUST produce `*_test.go` files alongside production code.** This is non-negotiable. The pre-commit hook and CI both enforce it.

### Requirements

- Every **package modified** must have a `*_test.go` file committed in the same push.
- **Table-driven tests with `t.Run()` for every new function** — no exceptions.
- `t.Helper()` in every test helper function.
- Use `t.Context()` (Go 1.24+) or `context.WithCancel` for test contexts. Do **not** call `context.Background()` in business-logic tests.
- Coverage targets by layer (enforced by review, not CI):
  - `internal/models/` — **90%+**
  - `internal/service/` — **80%+**
  - `internal/store/` — **60%+**
  - `internal/server/` — **40%+**
  - `cmd/` — best effort
- `make test` must pass before push.
- `go test -race -count=1 -coverprofile=coverage.out ./...` must show >0% coverage for every modified package.

### Rejection Process

If `make test` fails **or** any modified package shows 0% coverage, the work is **rejected and must be fixed before merging.** There is no "we'll add tests later."

### Table-Driven Test Pattern

Use this template for every new function. Adapt field names, but keep the shape.

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        {name: "happy path", input: ..., want: ...},
        {name: "error case", input: ..., wantErr: true},
        {name: "edge case", input: ..., want: ...},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Fatalf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("FunctionName() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## Coding Rules

Strict. Reviewers will flag violations.

### DO

- Use `context.Context` as the **first parameter** of every function that does I/O, blocks, or may be cancelled.
- Use `log/slog` for all logging (Go 1.21+) — structured, leveled.
- Wrap errors with `%w` and add context: `fmt.Errorf("fetch user %q: %w", id, err)`.
- Use **functional options** or **config structs** for complex constructors (no mega-constructors with 5+ positional args).
- Define `store`/`service` as **interfaces in the consumer package**, not the producer.
- Return `error` as the **last return value**.
- Use `t.Helper()`, `t.Run()`, `t.Parallel()` where safe.
- Use `t.Context()` (Go 1.24+) for test contexts.
- Prefer `slices`, `maps`, `cmp` (Go 1.21+ stdlib) over third-party helpers.
- Use `errors.Is` and `errors.As` for error checks, not `==`.
- Document every exported symbol with a doc comment starting with the symbol name.
- Use `slices.Contains` instead of hand-rolled loops.
- Use `slices.SortFunc` with `cmp.Compare` instead of `sort.Slice`.

### DO NOT

- Use `database/sql` — use a native driver (`pgx`, `modernc.org/sqlite`, etc.).
- Use `init()` outside the `config/` package.
- Use `context.Background()` in business logic — only in `main.go` and tests (use `t.Context()` in tests).
- Use `fmt.Println` for logging — use `slog`.
- Create a package-per-feature — one package per **layer** (`service/`, `store/`, `server/`).
- Create circular dependencies between layers.
- Use global mutable state.
- Suppress errors with `_` (use `_ = err` only when you truly mean it, and add a comment).
- Skip writing tests "for now" — there is no later.
- Mix acronym styles (`Http` vs `HTTP`) — pick one and be consistent.
- Use `interface{}` — use `any`.
- Add new top-level dependencies without justification in the PR description.

---

## Build & Run

### Prerequisites

| Tool | Version | Install |
|---|---|---|
| Go | 1.25+ | https://go.dev/dl/ |
| `golangci-lint` | v2.x | `brew install golangci-lint` |
| `pre-commit` | latest | `brew install pre-commit` |
| `deadcode`, `govulncheck` | latest | `go install` from `tools.go` deps |

### Common Commands

```bash
make tidy             # go mod tidy
make build            # go build ./...
make run              # go run ./cmd/<binary>/
make test             # go test -race -count=1 -shuffle=on
make test-short       # go test -short -race -count=1
make test-coverage    # test with coverage report
make test-race        # test with race detector
make lint             # golangci-lint run
make lint-ci          # golangci-lint with --timeout=5m
make ci-check         # tidy + build + vet + lint + test (what CI runs)
make deadcode         # detect unused exported functions
make vulncheck        # govulncheck
make clean            # rm coverage.out, rm -rf bin/
```

---

## Pre-commit Hook

The pre-commit hook runs on every commit and is the **first line of defense** before CI.

### Install

```bash
pre-commit install
```

### What It Does

| Hook | Purpose |
|---|---|
| `trailing-whitespace` | Strips trailing whitespace |
| `end-of-file-fixer` | Ensures files end with a newline |
| `check-yaml` | Validates YAML syntax |
| `check-merge-conflict` | Blocks commits with conflict markers |
| `go-fmt` | Runs `gofmt` |
| `go-vet` | Runs `go vet` |
| `go-imports` | Runs `goimports -local <module-path>` |
| `golangci-lint --fast` | Runs the linter (fast mode, no type-checking) |

If any of these fail, the commit is **blocked**. Fix the issues and re-commit.

---

## Style References

The definitive Go style guides (in order of authority for this project):

1. [**Effective Go**](https://go.dev/doc/effective_go) — the official guide.
2. [**Go Code Review Comments**](https://github.com/golang/go/wiki/CodeReviewComments) — community-reviewed nits.
3. [**Uber Go Style Guide**](https://github.com/uber-go/guide/blob/master/style.md) — opinionated, battle-tested.
4. [**Google Go Style Guide**](https://google.github.io/styleguide/go/) — Google's internal style.
5. [**Standard Go Project Layout**](https://github.com/golang-standards/project-layout) — directory structure conventions.

---

## Commit Style (Conventional Commits)

Use [Conventional Commits](https://www.conventionalcommits.org/). Examples:

- `feat:` — new feature for user/agent
- `fix:` — bug fix
- `chore:` — dependency bumps, tooling, CI
- `docs:` — documentation only
- `refactor:` — code change with no behavior change
- `test:` — adding or correcting tests
- `perf:` — performance improvement
- `ci:` — CI/CD changes
- `build:` — build system or external dependency changes

---

## License

MIT — see [`LICENSE`](LICENSE).
