# Contributing

## Getting Started

> TODO

1. Run the Brewfile bundle

    ```sh
    brew bundle
    ```

1. Set your GOPATH in your startup scripts

    ```sh
    export GOPATH="$HOME/go"
    ```

1. Clone the repo

    ```sh
    git clone git clone git@github.com:Ibotta/sopstool.git $GOPATH/src/github.com/Ibotta/sopstool
    ```

    you may want to symlink this to somewhere nicer

    ```sh
    ln -s $GOPATH/go/src/github.com/Ibotta/sopstool sopstool
    ```

1. Install other go prerequisites (many [recommended by VSCode](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on) and used during development)

    ```sh
    go get -u -v \
    github.com/acroca/go-symbols \
    github.com/cweill/gotests/... \
    github.com/fatih/gomodifytags \
    github.com/golang/lint/golint \
    github.com/golang/mock/gomock \
    github.com/golang/mock/mockgen \
    github.com/josharian/impl \
    github.com/k0kubun/pp \
    github.com/mitchellh/gox \
    github.com/motemen/gore \
    github.com/nsf/gocode \
    github.com/ramya-rao-a/go-outline \
    github.com/rogpeppe/godef \
    github.com/spf13/cobra/cobra \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    github.com/zmb3/gogetdoc \
    golang.org/x/tools/cmd/godoc \
    golang.org/x/tools/cmd/gorename \
    golang.org/x/tools/cmd/guru \
    sourcegraph.com/sqs/goreturns
    ```

1. Install packages required by packages in the repo

    ```sh
    go get -t ./...
    ```

## Layout

This is a single top-level namespace filled with packages.  Each directory is potentially a package. Binary builds are done on packages with a main subpackage.

## Releasing

> Still WIP

1. Preview the release (optional)

    You can preview the package changes by running `scripts/release-preview`. This will show a summary of changes since the last release.

1. Prepare the release

    ```sh
    git checkout develop && git pull
    ```

    Commit and tag

    ```sh
    git commit -am "Tagging release $VERSION" && git tag v$VERSION
    ```

    example:

    ```sh
    git commit -am "Tagging release 0.1.1" && git tag v0.1.1
    ```

    Then push to github

    ```sh
    git push && git push --tags # or git push --follow-tags but YMMV
    ```

    1. Tags the commit with the release versions
    1. pushes the tag to github `develop` branch

1. Review the release

    The easiest way to do this is to create a PR master <- develop. This should show all the changes about to be published, and give a good last review opportunity

1. Merge the release to master

    This **MUST** be a fastforward merge, otherwise the release tag(s) will be lost.

    ```sh
    git checkout master && git pull
    git merge develop --ff-only
    ```

    If ff-only fails, you must resolve the merge so that it is linear, which could involve rebuilding the release tag(s).

1. Watch for the release to pass CI

> TODO the tag+branch is a little tricky and one misstep can break the process. Look into a more foolproof or just simpler way to detect and release tags

### Release Gotchas

> TODO This process is just a little under review

## Versioning

We use [Semantic Versioning](http://semver.org/spec/v2.0.0.html) (2.0) for numbering our releases.

Summary: Given a version number **MAJOR**.**MINOR**.**PATCH**, increment the:

1. **MAJOR** (first) version when you make incompatible API changes
1. **MINOR** (second) version when you add functionality in a backwards-compatible manner
1. **PATCH** (third) version when you make backwards-compatible bug fixes or internal changes
1. In general, don't manually increment prerelease versions - we typically use that for automatic per-build numbering

## Patterns

### Common thirdparty modules to use

> TODO

### Module Configuration

> TODO

### Logging

> TODO

### Metrics

> TODO

### APM

> TODO

### Errors

> TODO

### Tests

Write tests for all your public APIs.  It can also be useful to write unit tests for private APIs if the methods are complex. Use mocks intelligently. Remember to properly handle and return promises and other async code to avoid those tests getting missed. Don't check in tests with the 'only' flag.

### Documentation

Document all public APIs to help users understand the module at a glance. Also consider documenting private APIs, especially as the methods become less obvious and the documentation can help future maintainers.

### Style

Clean up style warnings thrown by gofmt/golint (configured at the base of this repository).  These will be marked as build failures in CI.  Also consider using 'gofmt' to automatically clean up your code style while conforming to the configuration.
