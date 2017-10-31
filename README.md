# sopstool

## Installation

1. See from "Using these modules in other projects" section [here](https://github.com/Ibotta/go-commons#using-these-modules-in-other-projects)
1. Download from S3 and mark as executable

## Usage

This is a package that builds a single binary (per architecture) for wrapping [sops](https://github.com/mozilla/sops) with multi-file capabilities.

for more details

```sh
sopstool -h
```

to get the shell completion helpers:

```sh
#bash
sopstool completion
```

```sh
#zsh
sopstool completion --sh zsh
```

## Prereqs

1. install sops into your $PATH for your platform

    ```sh
    brew install sops
    ```

    or [from the github release](https://github.com/mozilla/sops/releases)

## Contributing

Bug reports and pull requests are welcome at <https://github.com/Ibotta/go-commons>

### docs

Generate markdown docs for the commands via

```sh
sopstool docs
```
