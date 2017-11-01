# sopstool

## Installation

1. See from "Using these modules in other projects" section [here](https://github.com/Ibotta/go-commons#using-these-modules-in-other-projects)
1. Download from S3 and mark as executable

## Usage

This is a package that builds a single binary (per architecture) for wrapping [sops](https://github.com/mozilla/sops) with multi-file capabilities.

for more details, use the built-in documentation on commands:

```sh
sopstool -h
```

to get the shell completion helpers:

```sh
#!/usr/bin/env bash
sopstool completion
```

```sh
#!/usr/bin/env zsh
sopstool completion --sh zsh
```

## Configuration

1. use a [`.sops.yaml`](https://github.com/mozilla/sops#using-sops-yaml-conf-to-select-kms-pgp-for-new-files) file
    * this will be at the root of your project. this file is used to both configure keys as well as hold the list of files managed.
    * it needs to specify at least one KMS key accessible by your environment

        ```yaml
        creation_rules:
          - kms: arn:aws:kms:REGION:ACCOUNT:key/KEY_ID
        ```

    * it can specify more complex cases of patterns vs keys too (see link)

## Installation and prereqs

See the root README on `Installing all binaries`.  This is probably what you want to do.

OR do it by hand and just install the one binary

1. install sops into your $PATH for your platform

    ```sh
    # OSX (darwin)
    brew install sops
    ```

    ```sh
    # Debian/Ubuntu (linux) sops 3.0
    wget -O /tmp/sops_3.0.0_amd64.deb https://github.com/mozilla/sops/releases/download/3.0.0/sops_3.0.0_amd64.deb && \
    dpkg -i /tmp/sops_3.0.0_amd64.deb && \
    rm /tmp/sops_3.0.0_amd64.deb
    ```

    or [from the github release](https://github.com/mozilla/sops/releases)

1. install the sopstool binary into your $PATH for your platform

    ```sh
    # OSX (darwin)
    aws s3 cp s3://ibotta-source/binaries/go-commons/darwin/amd64/sopstool /tmp/sopstool && \
    sudo install -Sv /tmp/sopstool /usr/local/bin && \
    rm -r /tmp/sopstool
    ```

    ```sh
    # Debian/Ubuntu (linux)
    aws s3 cp  s3://ibotta-source/binaries/go-commons/linux/amd64/sopstool /tmp/sopstool && \
    sudo install -v /tmp/sopstool /usr/local/bin && \
    rm -r /tmp/sopstool
    ```

## Contributing

Bug reports and pull requests are welcome at <https://github.com/Ibotta/go-commons>

### docs

Generate markdown docs for the commands via

```sh
sopstool docs
```
