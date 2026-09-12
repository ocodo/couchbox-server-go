# Couchbox Server Go — Project Notes

## Port

- This is a functional/behavioral port of `codefodder/couchbox-server` from Python to Go.
- The Go repository is `ocodo/couchbox-server-go`.
- Compatibility with the Python server is the goal; internal implementation does not need to be line-for-line equivalent.
- The Python repository is the behavioral reference.

## Development Method

- Development is TDD.
- Add the test first.
- Run the focused package test and report the red failure.
- Implement the minimum required behavior.
- Run the focused test until green.
- Run `go test ./...` to check regressions.
- Only then proceed to the next behavior.
- Do not invent APIs or behavior ahead of the tests/source.

## Configuration

- The canonical configuration is the actual current Couchbox configuration, not necessarily the upstream GitHub example.
- `video_exts` is part of the canonical configuration.
- Current canonical video extensions:
  - `.avi`
  - `.mov`
  - `.mkv`
  - `.mp4`
  - `.webm`
- Hard-coded video extension lists found in the Python implementation are technical debt; Go consumers should use configured `VideoExts`.

## Filesystem

- `__thumbs` functionality is removed completely.
- Do not add `__thumbs` references to tests or implementation.
- Filesystem behavior should be implemented only as required by the source/client contract.
- Filesystem is downstream of configuration.

## Media

- The server never handles actual video content.
- Poster/image assets are fixtures for testing.
- SVG/PNG generation is fixture preparation only, never a runtime server concern.

## Repository Structure

- `cmd/couchbox-server`
- `internal/config`
- `internal/filesystem`
- `internal/listings`
- `internal/locations`
- `internal/posters`
- `internal/bundles`
- `internal/server`
- `testdata`

## Testing

- Tests use the idiomatic adjacent `*_test.go` pattern.
- Fixtures belong under `testdata`.
- Test paths should be local and portable; do not copy `/mnt/...` production paths into tests.


## Source of Truth

When there is a discrepancy:
1. Current canonical Couchbox configuration / behavior wins for configuration.
2. The Python implementation wins for server behavior.
3. Tests document the Go port's resulting contract.
4. Do not silently resolve discrepancies by guessing.
