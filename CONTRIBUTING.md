# Contributing to blogger-go

Thanks for your interest in improving the Blogger API v3 Go SDK.

## Reporting bugs

Open an issue with:

- the Go version you are on (`go version`)
- a minimal reproduction: the call you made, what happened, and what you expected
- the SDK version or commit you are using

## Suggesting changes

For anything beyond a small fix, open an issue first so the design can be
discussed before you invest in a pull request. Typos, documentation fixes, and
clear bugs are welcome as direct pull requests.

## Development setup

Requires the Go version declared in `go.mod` (currently 1.26 or newer).

    git clone https://github.com/itokun99/blogger-go
    cd blogger-go
    go build ./...
    go test ./...

Run the full gate before opening a pull request; all four commands must be
clean:

    gofmt -l .          # must print nothing
    go build ./...
    go vet ./...
    go test ./... -count=1

## Pull requests

- Keep the change focused: one logical change per pull request.
- Add or update tests for any behavior change.
- Match the existing style; `gofmt` is the source of truth.
- Use conventional commit messages, for example `feat(services): ...`,
  `fix(client): ...`, or `docs: ...`.
- Much of `gen/` is generated code (look for the `DO NOT EDIT` header);
  hand-written code lives alongside it in `gen/services/constants.go` and the
  root package. Check the file header before editing.
- The README Quick Start call is compile-checked by `Example_postsListOptions`
  in `quickstart_test.go`. If you touch the public API, keep that example
  compiling.

## License

By contributing, you agree that your contributions are licensed under the
MIT License; see `LICENSE`.
