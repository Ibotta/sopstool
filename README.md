# sopstool

## Installation and prereqs

> TODO if fpm or brew work, use those first

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

    > TODO adjust these to download from the github release url

    ```sh
    # OSX (darwin)
    wget $URL sopstool.tar.gz && \
    sudo install -Sv /tmp/sopstool /usr/local/bin && \
    rm -r /tmp/sopstool
    ```

    ```sh
    # Debian/Ubuntu (linux)
    wget $URL sopstool.tar.gz && \
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

1. Create a KMS key. [KMS Keys](https://aws.amazon.com/kms/) are managed via [Terraform](https://www.terraform.io/docs/providers/aws/r/kms_key.html). You can find an [example here](https://github.com/Ibotta/infrastructure/pull/265).
1. Follow along the [Configuration Steps](https://github.com/Ibotta/go-commons/tree/develop/sopstool#configuration), and place the `.sops.yaml` file at the root directory where your scripts will run.
    * All files added to SOPS are relative, or in child directories to the `.sops.yaml` configuration file.
1. Create a file to encrypt(any extension other than `.yaml` if you wish to do the **ENTIRE** file), or create a yaml file with `key: value` pairs(and make sure it's extension is `.yaml`). Sops will encrypt the keys, but not it's values.
    * You can read more about [SOPS Here](https://github.com/mozilla/sops).
1. At this point, `sopstool` is ready and you can now `sopstool add filename`. You'll notice it will create a `filename.sops.extension`. This is your newly encrypted file.
    * When your files are properly encyrepted, you can run `sopstool clean` to remove the original plain text secret files.
1. Now, you can interact via the command line in various ways.
    * **Editing an encrypted file** - `sopstool edit filename.sops.extension`. You can also use your original filename too! `sopstool edit filename.extension`
    * **Listing all encrypted files** - `sopstool list`
    * **Removing encrypted file** - `sopstool remove filename.extension`
    * **Display the contents of encrypted file** - `sopstool cat filename.extension`

### Walkthrough

In this walkthrough, we will go through the steps required to get a secure yaml configuration file running.

** Configure your `.sops.yaml` **

```yaml
# .sops.yaml
creation_rules:
   - kms: arn:aws:kms:REGION:ACCOUNT:key/KEY_ID
```

** Create a secrets yaml configuration file **

```yaml
# credentials.yaml
database.password: supersecretdb
database.user: supersecretpassword
redshift:
    user: my.user.name
    password: my.password
```

** Encrypt the newly created file **

```sh
sopstool add credentials.yaml
```

** Create a sample script **

```python
# myscript.py
import yaml
with open('credentials.yaml', 'r') as f:
    credentials = yaml.load(f)

print credentials["database.user"]
print credentials["database.password"]
print credentials["redshift"]["user"]
print credentials["redshift"]["password"]
```

Here is what your folder structure would look like to this point(after deleting the unencrypted credentials.yaml file)

```text
my-project/
├── .sops.yaml
├── credentials.sops.yaml
└── myscript.py
```

** Accessing credentials **

The flow should be as follows: unencrypt credentials -> run script -> destroy credentials. You can use the `sopstool entrypoint` to achieve this.

```sh
sopstool entrypoint python myscript.py
```

## Contributing

Bug reports and pull requests are welcome at <https://github.com/Ibotta/go-commons>

### docs

Generate markdown docs for the commands via

```sh
sopstool docs
```
