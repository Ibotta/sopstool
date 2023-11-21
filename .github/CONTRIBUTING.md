# Contributing

## Getting Started

This is a single top-level namespace filled with packages. Each directory is potentially a package. Binary builds are done on packages with a main subpackage.

## Install Golang

### Using asdf or rtx

We utilize a `.tool-versions` file that can be used by [asdf-vm](https://github.com/kennyp/asdf-golang) or [rtx](https://github.com/jdxcode/rtx) like so:

```bash
cd /path/to/sopstool/repository/
asdf plugin add golang
asdf install golang
```

### Using goenv

[goenv](https://github.com/syndbg/goenv) will prefer the `GOENV_VERSION` environment variable first before looking for the `.go-version` file when [determining which Golang version](https://github.com/syndbg/goenv/blob/master/HOW_IT_WORKS.md#choosing-the-go-version) to install.  If you do not have this set (`echo $GOENV_VERSION` is empty), install like so:

```bash
cd /path/to/sopstool/repository/
goenv install
```

### From the developers

You can download and install the Golang [directly from the website](https://go.dev/dl/).

### Additional Go Libraries

Install [gomock](https://github.com/golang/mock)

```sh
go get -u github.com/golang/mock/gomock && go install github.com/golang/mock/mockgen
```

Install [golangci-lint](https://golangci-lint.run/)

```sh
brew tap golangci/tap
brew install golangci/tap/golangci-lint
```

## Build

With Go 1.11+ and addition of [Modules](https://github.com/golang/go/wiki/Modules), Go projects can be located outside the `$GOPATH`.

If you are having issues, review the [FAQ](https://github.com/golang/go/wiki/Modules#faqs--most-common).

If `generate` has already run, then it does not need to run again.

```sh
go build
go fmt ./...
```

## Unit Tests

Each module is unit tested, and passes all tests.

```sh
go test ./...
```

## Linting

`golangci-lint` runs several popular Go linters quickly:

```sh
golangci-lint run
```

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) for builds and releases. Doing the tag/release below triggers the appropriate actions.

1. Preview the release (optional)

   You can preview the package changes by running `scripts/release-preview`. This will show a summary of changes since the last release.

1. Prepare the release:

   ```sh
   git checkout main && git pull
   ```

   Commit and tag with the intended version bump:

   ```sh
   git commit -am "Tagging release $VERSION" && git tag v$VERSION
   ```

   For example:

   ```sh
   git commit -am "Tagging release 0.1.1" && git tag v0.1.1
   ```

   Then push the tag and commit to Github:

   ```sh
   git push && git push --tags # or git push --follow-tags but YMMV
   ```

1. Watch for the release to pass CI. CI publishes the release and pushes the artifacts.

## Versioning

We use [Semantic Versioning](http://semver.org/spec/v2.0.0.html) (2.0) for numbering our releases.

Summary: Given a version number **MAJOR**.**MINOR**.**PATCH**, increment the:

1. **MAJOR** (first) version when you make incompatible API changes
1. **MINOR** (second) version when you add functionality in a backwards-compatible manner
1. **PATCH** (third) version when you make backwards-compatible bug fixes or internal changes
1. In general, don't manually increment prerelease versions - we typically use that for automatic per-build numbering

## Patterns

### Common third-party modules in use

- cobra
- yaml
- mock

### Errors

> TODO

### Tests

Write tests for all your public APIs. It can also be useful to write unit tests for private APIs if the methods are complex. Use mocks intelligently. Remember to properly handle and return promises and other async code to avoid those tests getting missed. Don't check in tests with the 'only' flag.

### Documentation

Document all public APIs to help users understand the module at a glance. Also consider documenting private APIs, especially as the methods become less obvious and the documentation can help future maintainers.

### Style

Clean up style warnings thrown by gofmt/golangci-lint (configured at the base of this repository in `.golangci.yml`). These will be marked as build failures in CI. Also consider using these tools to automatically clean up your code style while conforming to the configuration.
