# Contributing

Thanks for improving this blueprint. The goal is simple: make Go repositories easier for humans and coding agents to work in safely.

## Before You Start

1. Read [AGENTS.md](AGENTS.md).
2. Check existing issues and pull requests for related work.
3. Keep changes focused. Avoid mixing documentation, tooling, and behavior changes unless they must ship together.
4. Run `git status --short` before editing so you do not overwrite someone else's work.

## Development Workflow

Use this loop for non-trivial changes:

1. Think: clarify the goal, constraints, affected files, and risks.
2. Plan: choose the smallest safe path and the checks that will validate it.
3. Review plan: get human review before broad, ambiguous, security-sensitive, destructive, architectural, or dependency-changing work.
4. Write tests: add or update tests before implementation for production-code behavior changes.
5. Implement: keep the diff cohesive and reviewable.
6. Validate: prove the change meets the Definition of Done.
7. Handoff: summarize changes, checks, skipped checks, and remaining risk.

## Definition of Done

A change is done when:

- The requested scope is complete and the diff avoids unrelated refactors.
- Production-code behavior changes include meaningful tests, and modified packages do not show 0% coverage.
- `make ci-check` passes, unless the change is documentation-only and a narrower check is clearly sufficient.
- `git diff --check` passes, and `pre-commit run --all-files` passes when `pre-commit` is installed.
- Documentation is updated when behavior, commands, setup, or agent expectations change.
- Secrets are not added, security-sensitive changes are called out, and dependency or workflow permission changes receive extra review.
- Makefile targets, CI, hooks, and documented commands agree.
- The PR or handoff lists checks run, skipped checks with reasons, and any residual risk.

```bash
make help
make ci-check
```

For code changes, include tests in every modified package. For documentation-only changes, run the narrowest useful checks and explain anything skipped.

## Pull Request Expectations

Every PR should include:

- A short summary of what changed and why.
- Tests or checks run.
- Any skipped checks and the reason.
- Any new dependency, generated file, security implication, or migration step.
- Updates to [AGENTS.md](AGENTS.md) when the change affects future agent behavior.

## Commit Style

Use Conventional Commits:

- `feat:` for user-visible features.
- `fix:` for bug fixes.
- `docs:` for documentation.
- `test:` for tests.
- `ci:` for CI changes.
- `build:` for build tooling.
- `chore:` for maintenance.

## Agent Contributions

Agent-authored changes are welcome, but the handoff must be reviewable:

- Keep diffs small enough for a human to audit.
- Prefer existing patterns over new abstractions.
- Do not commit secrets, machine-local paths, or private customer data.
- Include command output summaries rather than pasting noisy logs.
