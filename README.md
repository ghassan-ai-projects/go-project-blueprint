# Go Project Blueprint

> **The starter kit every Go coding agent wishes you had.**

Industry-best agent configuration files for Go projects — `AGENTS.md`, linters, CI, pre-commit hooks, `Makefile`s — the **ultimate coding-agent onboarding kit**.

---

## What

This repository is a **template** — the configuration files (not the application code) that any Go project can adopt to give coding agents (Claude, Cursor, Qwen, Aider, etc.) the best possible context to work effectively.

It includes:

- **`AGENTS.md`** — the canonical reference for any coding agent, covering naming, structure, quality gates, the test mandate, and coding rules.
- **`.golangci.yml`** — 14 linters, opinionated settings for `gosec` and `wrapcheck`.
- **`.pre-commit-config.yaml`** — automated quality gates on every commit.
- **`Makefile`** — `tidy`, `build`, `vet`, `lint`, `test`, `test-coverage`, `ci-check`, `deadcode`, `vulncheck`, `cross-compile`.
- **`.github/workflows/ci.yml`** — full CI pipeline: tidy → build → vet → lint → test → deadcode → vulncheck → cross-compile.
- **`tools.go`** — pins `deadcode` and `govulncheck` in `go.mod` for reproducible installs.
- **`.editorconfig`**, **`.gitignore`**, **`LICENSE`**.

## Why

Most Go projects hand coding agents a wall of ambiguity: where do files go? What naming? Is there a test mandate? What about logging? CI? The agent guesses, the human reviews, the cycle repeats. **This blueprint removes the guessing.**

`AGENTS.md` is the single source of truth. The pre-commit hook enforces the rules before push. CI enforces them again. **The agent, the human, and CI all read the same document.**

## Who

Fork this if you:

- Start a new Go project and want production-grade scaffolding from minute one.
- Maintain an existing Go project and want to give every coding agent a level playing field.
- Run a team of humans + agents and want a single convention document that governs both.

## How

### Quick start (forking this template)

```bash
# 1. Clone (or use the GitHub "Use this template" button)
git clone https://github.com/your-org/your-project.git
cd your-project

# 2. Update the module path
go mod edit -module github.com/your-org/your-project

# 3. Install the pre-commit hook
pre-commit install

# 4. Verify the toolchain
make ci-check

# 5. Add your entry point
mkdir -p cmd/myapp
echo 'package main

func main() {}' > cmd/myapp/main.go

# 6. Build
make build
```

For the full setup checklist, see [**`AGENTS.md` → For Humans Forking This Template**](AGENTS.md#for-humans-forking-this-template).

### Coding agents

Read [**`AGENTS.md`**](AGENTS.md) end-to-end before you write a single line of code. Every section is load-bearing.

## Quality

```bash
make ci-check     # tidy + build + vet + lint + test (what CI runs)
make test         # go test -race -count=1 -shuffle=on
make lint         # golangci-lint (14 linters)
make deadcode     # detect unused functions
make vulncheck    # govulncheck
```

## License

MIT — see [LICENSE](LICENSE).
