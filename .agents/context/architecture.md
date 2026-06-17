# Architecture Context

## Repository Architecture

This repository has two layers of architecture:

1. The current template repository.
2. The recommended architecture for projects that adopt the template.

Do not confuse them.

## Current Template Structure

- `AGENTS.md` is the canonical agent entrypoint.
- `README.md` explains what the template includes and how to adopt it.
- `Makefile` defines the authoritative local commands.
- `.github/workflows/ci.yml` defines CI parity for core checks.
- `.golangci.yml` and `.pre-commit-config.yaml` enforce code quality and hygiene.
- `docs/` holds supporting reference material.

## Recommended Adopted-Project Structure

- `cmd/<binary>/main.go`: startup, flags, dependency wiring, shutdown
- `internal/config`: configuration loading and validation
- `internal/models`: data types and validation rules
- `internal/service`: business logic
- `internal/store`: persistence and external integrations
- `internal/server`: HTTP, gRPC, or MCP transport
- `test/`: integration or end-to-end suites

## Dependency Direction

- `cmd` depends on lower layers.
- `server` depends on `service`.
- `service` depends on store interfaces it consumes.
- `store` depends on lower-level drivers or clients.
- `models` should stay independent from other internal packages.

Avoid:

- circular imports
- business logic in handlers
- persistence concerns in `service`
- transport concerns leaking into `models`
