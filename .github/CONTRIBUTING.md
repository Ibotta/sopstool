# Contributing

## Getting Started

## Layout

This is a single top-level namespace filled with packages.  Each directory is potentially a package. Binary builds are done on packages with a main subpackage.

## Building Locally

### Go Version

Requires Go `>= 1.12.6`.


1. Install go (currently 1.12)

    You may want to use a version manager.  

    [asdf](https://github.com/kennyp/asdf-golang)

    ```sh
    asdf install golang 1.12.6
    # use global to update default go version. local set just for current directory
    asdf local golang 1.12.6
    ```

    [goenv](https://github.com/syndbg/goenv),   are all options.

    ```sh
    goenv install 1.12.6
    go version
    # go version go1.12.6 darwin/amd64
    ```
        
    * [gimme](https://github.com/travis-ci/gimme)


1. Install gomock

    ```sh
    go get -u github.com/golang/mock/gomock && go install github.com/golang/mock/mockgen
    ```
    
### Build


With go 1.11 and addition of [Modules](https://github.com/golang/go/wiki/Modules),
go projects can be located outside the GOPATH.

If you are having issues review
[faq](https://github.com/golang/go/wiki/Modules#faqs--most-common)


If generate is ran above, then it does not need to run again.

```bash
go build
go fmt ./...
golint ./...
```

### Unit Test

```bash
go test ./...
```

## Releasing

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

    This **MUST** be a fast-forward merge, otherwise the release tag(s) will be lost.

    ```sh
    git checkout master && git pull
    git merge develop --ff-only
    ```

    If ff-only fails, you must resolve the merge so that it is linear, which could involve rebuilding the release tag(s).

1. Watch for the release to pass CI

> TODO the tag+branch is a little tricky and one misstep can break the process. Look into a more foolproof or just simpler way to detect and release tags

### Release Gotchas

> TODO This process under review

## Versioning

We use [Semantic Versioning](http://semver.org/spec/v2.0.0.html) (2.0) for numbering our releases.

Summary: Given a version number **MAJOR**.**MINOR**.**PATCH**, increment the:

1. **MAJOR** (first) version when you make incompatible API changes
1. **MINOR** (second) version when you add functionality in a backwards-compatible manner
1. **PATCH** (third) version when you make backwards-compatible bug fixes or internal changes
1. In general, don't manually increment prerelease versions - we typically use that for automatic per-build numbering

## Patterns

### Godownloader

We use [godownloader](https://github.com/goreleaser/godownloader) to generate the installer scripts.

* for sops, it uses the 'raw repo' method

    ```sh
    godownloader -source raw -repo mozilla/sops -exe sops -nametpl 'sops-{{ .Version }}.{{ .Os }}' > sopsdownload.sh
    ```

* for sopstool, we can use the goreleaser file

    ```sh
    godownloader -repo Ibotta/sopstool .goreleaser.yml > sopstoolinstall.sh
    ```

### Common third-party modules in use

* cobra
* yaml
* mock

### Errors

> TODO

### Tests

Write tests for all your public APIs.  It can also be useful to write unit tests for private APIs if the methods are complex. Use mocks intelligently. Remember to properly handle and return promises and other async code to avoid those tests getting missed. Don't check in tests with the 'only' flag.

### Documentation

Document all public APIs to help users understand the module at a glance. Also consider documenting private APIs, especially as the methods become less obvious and the documentation can help future maintainers.

### Style

Clean up style warnings thrown by gofmt/golint (configured at the base of this repository).  These will be marked as build failures in CI.  Also consider using 'gofmt' to automatically clean up your code style while conforming to the configuration.
