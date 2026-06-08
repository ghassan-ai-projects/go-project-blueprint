// Package blueprint anchors the project root so that all Go tooling
// (build, vet, test, lint) finds at least one package to operate on
// in a freshly-cloned template. Once you add your own packages
// (e.g. cmd/myapp/main.go, internal/...), this file can be deleted
// or repurposed to host shared types at the module root.
//
// See AGENTS.md for the canonical project definition.
package blueprint
