# sopstool

## Installation and prereqs

See the [root README](../README.md) on `Installing all binaries`.  This is probably what you want to do.

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

## How-To

1. [Create a KMS key](https://github.com/Ibotta/infrastructure/pull/265/files#diff-3c4152d505a5e581de30df76f03f3b3a). Some are in Terraform and others are not, but it's pretty easy to create them via Terraform.
1. Follow along the [Configuration Steps](https://github.com/Ibotta/go-commons/tree/develop/sopstool#configuration), and place the `.sops.yaml` file at the root directory where your scripts will run.
    * This is important because sops uses the same file name that is in it's list, that you specify. So `../myfile.txt` != `myfile.txt`, and the sopstool/sops may not recognize the file as being under it's control.
1. Create a file to encrypt(any extension other than `.yaml` if you wish to do the **ENTIRE** file), or create a yaml file with `key: value` pairs(and make sure it's extension is `.yaml`). Sops will encrypt the keys, but not it's values.
1. At this point, `sopstool` is ready and you can now `sopstool add filename`. You'll notice it will create a `filename.sops.extension`. This is your newly encrypted file.
    * Remember to keep the `*.sops.*` file, and delete your **original** file as we do _NOT_ want to check it into the repository!
1. Now, you can interact via the command line in various ways.
    * **Editing an encrypted file** - `sopstool edit filename.sops.extension`. You can also use your original filename too! `sopstool edit filename.extension`
    * **Listing all encrypted files** - `sopstool list`
    * **Removing encrypted file** - `sopstool remove filename.extension`
    * **Display the contents of encrypted file** - `sopstool cat filename.extension`

### Walkthrough

In this walkthrough, we will go through the steps required to get a secure script running. In this case, we are just setting some env variables to be used only by our script. Once the script exists, the env variables are no longer available to other shells.

** Configure your `.sops.yaml` **
```yaml
# .sops.yaml
creation_rules:
   - kms: arn:aws:kms:REGION:ACCOUNT:key/KEY_ID
```

** Create a secrets file**
```sh
#secrets.sh
export username=sopstoolrocksmysocksoff
export password=yutRy2Hakm3
```

** Encrypt the newly created file **
```sh
$ sopstool add secrets.sh
```

** Create a secure workspace **
```sh
# secure.workspace.sh
#!/usr/bin/env bash

source <(sopstool cat secrets.sh)
eval "$@"
```

Here is what your folder structure would look like to this point:

```
my-project/
├── secrets.sh
├── secrets.sops.sh
└── secure.workspace.sh
```

To use your protected env variables:

```sh
# pass whatever command you'd like to ./secure.workspace.sh
$ ./secure.workspace.sh python my.python.script.py
$ ./secure.workspace.sh ruby my.ruby.script.rb
$ ./secure.workspace.sh ./my.shell.script.sh
```

***

## Contributing

Bug reports and pull requests are welcome at <https://github.com/Ibotta/go-commons>

### docs

Generate markdown docs for the commands via

```sh
sopstool docs
```
