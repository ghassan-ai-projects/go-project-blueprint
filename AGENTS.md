# AGENTS.md - Go Project Blueprint

This is the canonical operating guide for coding agents working in this repository. Keep it specific, enforceable, and short enough to stay useful in every session.

---

## What This Is

This repository is a **template**: agent instructions, lint rules, CI, pre-commit hooks, and Make targets that Go projects can adopt. The included files are scaffold, not the application.

---

## For Coding Agents

Before editing:

1. Read this file and the "Project Identity" section.
2. Read [README.md](README.md) for repository-specific context.
3. Check the worktree with `git status --short`; never overwrite user changes.
4. Run `make ci-check` before changes when feasible. If it fails because dependencies or tools are missing, report that and continue with the narrowest useful checks.
5. Follow the Test Mandate for every production-code change.

During work:

- Prefer small, reviewable changes over broad rewrites.
- Use existing package boundaries and local helpers before adding new abstractions.
- Explain any new dependency, generated artifact, skipped test, or changed public behavior in the handoff.
- If human feedback conflicts with this file, follow the human and propose a doc update.

Default workflow:

1. Think: restate the goal, constraints, relevant files, and risks.
2. Plan: outline the smallest safe change and the checks that will prove it works.
3. Review plan: pause for human review when the change is broad, ambiguous, security-sensitive, destructive, architectural, or dependency-changing. For small obvious fixes, proceed and summarize the plan in the handoff.
4. Write tests: add or update failing tests first for production-code behavior changes.
5. Implement: make the smallest cohesive change that satisfies the plan.
6. Validate: prove the change meets the Definition of Done below.
7. Handoff: summarize changed files, behavior, checks, skipped checks, and residual risk.

Definition of Done:

- Scope is satisfied: the requested behavior or documentation change is complete, with no unrelated refactors.
- Tests are meaningful: production-code changes include tests for the new or changed behavior, and modified packages do not show 0% coverage.
- Quality gates pass: run `make ci-check` unless the change is documentation-only and a narrower check is clearly sufficient.
- Formatting and hygiene pass: run `git diff --check`; run `pre-commit run --all-files` when `pre-commit` is installed.
- Documentation is current: update `README.md`, `AGENTS.md`, config examples, or reference docs when behavior, commands, setup, or agent expectations change.
- Security is considered: no secrets are added, new inputs are validated, dependency or workflow permission changes are called out, and security-sensitive changes receive plan review.
- Automation matches prose: Makefile targets, CI, hooks, and documented commands agree.
- Handoff is complete: final response or PR notes list changed files, validation performed, skipped checks with reasons, and residual risks.

Handoff checklist:

- State what changed and why.
- List tests/checks run, including failures and skipped checks.
- Call out migrations, configuration changes, security considerations, and follow-up work.

---

## For Humans Forking This Template

Setup checklist after cloning or using this template:

1. Update `module` in `go.mod` to your module path (e.g. `github.com/your-org/your-project`).
2. Update `BINARY` in the `Makefile` to your binary name (e.g. `BINARY := bin/myapp`).
3. Update `MODULE` in the `Makefile` to your module path (same as `go.mod`).
4. Update the `-local` flag in `.pre-commit-config.yaml` to your module path.
5. Replace the placeholder module path in `.golangci.yml` exclusions (if any) and `run` config.
6. Update the **Project Identity** section below with 1–2 paragraphs about your project.
7. Add your entry point at `cmd/<binary>/main.go`.
8. Update copyright in `LICENSE` if not the template author.
9. Install the pre-commit hook: `pre-commit install`.
10. Update `CLAUDE.md`, `GEMINI.md`, and `.github/copilot-instructions.md` only if your team needs tool-specific behavior beyond importing `AGENTS.md`.
11. Run `make ci-check` once to confirm everything is green.

Keep `AGENTS.md` canonical. Tool-specific files should bridge to it instead of duplicating rules.

---

## Agent Configuration Strategy

- `AGENTS.md` is the cross-agent source of truth.
- Codex reads `AGENTS.md` natively. Do not add a `CODEX.md` duplicate.
- `CLAUDE.md` imports `AGENTS.md` for Claude Code.
- `GEMINI.md` points Gemini CLI at `AGENTS.md`; teams can also configure Gemini's context file name to `AGENTS.md`.
- `.github/copilot-instructions.md` points GitHub Copilot at the same rules.
- Use `.codex/config.toml` for Codex settings such as sandbox, MCP servers, hooks, models, or approval defaults; keep repository conventions in `AGENTS.md`.
- Add path-specific or workflow-specific agent files only when a rule is too narrow to load into every session.
- Keep durable facts in files, not chat history: build commands, project identity, architecture, test commands, security rules, and review expectations.
- Do not put secrets, credentials, private customer data, or environment-specific machine paths in committed agent instructions.

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
| Files | `lowercase.go` | `agentstore.go` |
| Test files | `<file>_test.go` | `agentstore_test.go` |
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

Default layout for projects adopting this blueprint:

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
├── .github/ISSUE_TEMPLATE/  ← issue templates
├── AGENTS.md                ← this file
├── CONTRIBUTING.md          ← contribution workflow
├── SECURITY.md              ← vulnerability reporting policy
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

Layer rules:

- `cmd/` → `server/` → `service/` → `store/` → `models/`
- Imports flow **downward only.** Never circular.
- `service/` depends on `store/` **interfaces**, not concrete types. Mocks live in `service/storemock/`.
- `models/` has no dependencies on other `internal/` packages.
- `cmd/<binary>/main.go` wires concrete implementations together.
- Shared code belongs under `internal/` unless it is intentionally importable by other modules.
- Keep generated files clearly marked and document the generator command.

---

## Quality Gates

All gates must pass before merge. `make ci-check` runs the local core pipeline: `tidy -> build -> vet -> lint-ci -> test-short -> deadcode -> vulncheck`. GitHub Actions runs the same core gates and then cross-compiles a linux/amd64 artifact.

| Gate | Command | Required |
|---|---|---|
| Format | `gofmt -l .` | clean |
| Imports | `goimports -l -local <module-path> .` | clean |
| Vet | `go vet ./...` | clean |
| Lint | `make lint` (golangci-lint, 14 linters) | clean |
| Build | `make build` | success |
| Test | `make test` (race + shuffle + count=1 + coverage) | passes |
| Coverage | `make test-coverage` | meets layer targets (review-enforced) |
| Deadcode | `make deadcode` | no unused funcs |
| Vulns | `make vulncheck` | no known vulns |
| CI | `make ci-check` | green |

---

## Test Mandate

**Every code-writing task MUST produce `*_test.go` files alongside production code.** This is non-negotiable. Hooks and CI enforce the runnable checks; reviewers enforce package-level coverage expectations and the presence of meaningful tests.

### Requirements

- Every **package modified** must have a `*_test.go` file committed in the same push.
- Table-driven tests with `t.Run()` for every new exported function and every meaningful branch in unexported logic.
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

If `make test` fails or any modified package shows 0% coverage, the work is rejected until fixed. There is no "we'll add tests later."

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
- Keep handlers thin: parse/validate input, call service, translate response.
- Keep business logic in `service/`, persistence details in `store/`, transport details in `server/`.
- Prefer explicit validation at boundaries over relying on downstream failures.

### DO NOT

- Use the `database/sql` interface — use a native driver directly (`pgx`, `modernc.org/sqlite`, etc.).
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
- Log secrets, tokens, passwords, API keys, session cookies, or raw credentials.
- Introduce network calls in unit tests; use integration tests under `test/` or build-tagged suites.

### Security Defaults

- Treat all external input as untrusted.
- Validate request payloads, config files, environment variables, and CLI flags before use.
- Set server timeouts for HTTP servers and clients.
- Use least-privilege file permissions for generated credentials, sockets, and state files.
- Prefer allowlists over denylists for commands, paths, enum values, and external integrations.
- Keep secrets out of source, test fixtures, logs, traces, screenshots, and agent memory files.

---

## Build & Run

### Prerequisites

| Tool | Version | Install |
|---|---|---|
| Go | 1.25+ | https://go.dev/dl/ |
| `golangci-lint` | v2.x | `brew install golangci-lint` |
| `pre-commit` | latest | `brew install pre-commit` |
| `deadcode`, `govulncheck` | latest | `go install` from `tools.go` deps |

CI pins tool installation versions in `.github/workflows/ci.yml`. When updating Go tool dependencies in `go.mod`, update the matching CI environment variables in the same change.

### Common Commands

```bash
make help             # show all available targets
make fmt              # format code with goimports
make tidy             # go mod tidy
make build            # go build ./...
make run              # go run ./cmd/<binary>/
make test             # go test -race -count=1 -shuffle=on
make test-short       # go test -short -race -count=1
make test-coverage    # test with coverage report
make test-race        # test with race detector
make lint             # golangci-lint run
make lint-ci          # golangci-lint with --timeout=5m
make ci-check         # tidy + build + vet + lint-ci + test-short + deadcode + vulncheck
make deadcode         # detect unused exported functions
make vulncheck        # govulncheck
make clean            # rm coverage.out, rm -rf bin/
```

Use `make help` as the command index. If a Make target and this document disagree, fix the document or the target in the same change.

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
| `check-added-large-files` | Rejects files over 1 MB |
| `check-case-conflict` | Blocks case-insensitive filename collisions |
| `mixed-line-ending` | Ensures LF line endings |
| `go-fmt` | Runs `gofmt` |
| `go-vet` | Runs `go vet` |
| `go-imports` | Runs `goimports -local <module-path>` |
| `golangci-lint` | Runs the linter (golangci-lint, timeout 1m) |

If any of these fail, the commit is **blocked**. Fix the issues and re-commit.

Run all hooks manually with:

```bash
pre-commit run --all-files
```

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
