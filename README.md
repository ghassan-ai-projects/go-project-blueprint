# Go Project Blueprint

> **Author:** [Ghassan Alhamoud](https://ghassan-alhamoud.com)

Production-ready scaffolding for Go projects that expect humans and coding agents to work in the same repository.

This template gives agents a crisp operating contract, gives humans reproducible quality gates, and keeps the two aligned through `AGENTS.md`, linting, pre-commit hooks, CI, and a small set of agent bridge files.

---

## What's Included

This repository is a **template**, not an application. Fork it, replace the project identity, add your Go code, and keep the quality gates intact.

It includes:

- `AGENTS.md` - the canonical instructions for agents and humans; Codex reads this natively.
- `.agents/context/` - small, task-targeted context files for project, architecture, testing, style, and review.
- `.agents/prompts/` - reusable prompts for planning, implementation, review, feature work, and bugfixes.
- `CLAUDE.md`, `GEMINI.md`, `.github/copilot-instructions.md` - thin bridge files that point popular agents at the canonical instructions.
- `CONTRIBUTING.md`, `SECURITY.md`, PR and issue templates - human governance for reviews, reports, and contribution flow.
- `.golangci.yml` - strict Go linting with security and error-handling checks.
- `.pre-commit-config.yaml` - local hooks for formatting, vetting, imports, linting, and file hygiene.
- `.github/workflows/ci.yml`, `.github/dependabot.yml` - GitHub Actions pipeline and dependency update automation.
- `Makefile` - reproducible local commands for daily development and CI parity.
- `tools.go` - pinned tool dependencies for reproducible installs.
- `.editorconfig`, `.gitignore`, `LICENSE`.

## Why

Most repositories hand coding agents a wall of ambiguity: where do files go, which command is authoritative, what needs tests, which architecture boundaries matter, and what may be changed without review? The result is predictable: the agent guesses, the human corrects, and the same mistakes recur.

This blueprint makes the contract explicit:

- Agents read one canonical instruction file before editing.
- Agents can then load only the smallest relevant `.agents/context/*.md` file instead of re-reading the full policy surface.
- Humans get a short setup checklist and reproducible commands.
- Hooks and CI enforce what prose alone cannot.
- Tool-specific files stay tiny and delegate to `AGENTS.md`, avoiding drift. Codex needs no bridge file because `AGENTS.md` is its native project-guidance format.

## Who

Fork this if you:

- Start a new Go project and want production-grade scaffolding from minute one.
- Maintain an existing Go project and want to give every coding agent a level playing field.
- Run a team of humans plus agents and want one convention document that governs both.

## How

### Quick start (forking this template)

```bash
# 1. Clone (or use the GitHub "Use this template" button)
git clone https://github.com/your-org/your-project.git
cd your-project

# 2. Update the module path
go mod edit -module github.com/your-org/your-project

# 3. Update the template metadata and command defaults
$EDITOR README.md AGENTS.md Makefile .pre-commit-config.yaml

# 4. Install the pre-commit hook
pre-commit install

# 5. Verify the toolchain
make ci-check

# 6. Add your entry point
mkdir -p cmd/myapp
echo 'package main

func main() {}' > cmd/myapp/main.go

# 7. Build
make build
```

Then update the remaining template placeholders called out in `AGENTS.md`, `go.mod`, `Makefile`, and `.pre-commit-config.yaml`.

### For Coding Agents

Read [AGENTS.md](AGENTS.md) before editing. It contains the entrypoint rules and points to the smaller context files under `.agents/context/` for architecture, testing, style, and review.

## Quality

```bash
make ci-check        # tidy + build + vet + lint + test-short + deadcode + vulncheck
make test            # race + shuffle + coverage
make test-coverage   # coverage HTML report
make lint            # golangci-lint
make cross-compile   # linux/amd64 binary
```

CI runs the same core quality gates and also publishes a short-lived linux/amd64 artifact.

## License

MIT - see [LICENSE](LICENSE).
