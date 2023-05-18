# sopstool

[![Maintainability](https://api.codeclimate.com/v1/badges/addf39da73692548e1e3/maintainability)](https://codeclimate.com/github/Ibotta/sopstool/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/addf39da73692548e1e3/test_coverage)](https://codeclimate.com/github/Ibotta/sopstool/test_coverage)

sopstool is a multi-file wrapper around [sops](https://github.com/mozilla/sops). It uses the sops binary to encrypt and decrypt files, and piggybacks off the .sops.yaml configuration file.

sopstool provides functionality to manage multiple secret files at once, and even use as an entrypoint to decrypt at startup, for container images. Much of this behavior is inspired by the great [blackbox project](https://github.com/StackExchange/blackbox).

- [sopstool](#sopstool)
	- [1.0.0 Release and Breaking Changes](#100-release-and-breaking-changes)
	- [Installation](#installation)
		- [Package Repositories](#package-repositories)
		- [Container Image](#container-image)
		- [Packages or binaries from Releases](#packages-or-binaries-from-releases)
		- [Shell installer](#shell-installer)
		- [Installing sops manually](#installing-sops-manually)
			- [Installing the sops binary with our script installer](#installing-the-sops-binary-with-our-script-installer)
			- [Download sops from our https mirror](#download-sops-from-our-https-mirror)
		- [Installing sopstool manually](#installing-sopstool-manually)
			- [Installing the sopstool binary using our script installer](#installing-the-sopstool-binary-using-our-script-installer)
			- [Download sopstool from our https mirror](#download-sopstool-from-our-https-mirror)
	- [Usage](#usage)
	- [Configuration](#configuration)
	- [How-To](#how-to)
		- [Walkthrough](#walkthrough)
	- [Contributing](#contributing)
		- [docs](#docs)

## 1.0.0 Release and Breaking Changes

1.0.0 release of `sopstool` introduces M1 / darwin-arm64 support. We also want to match build artifacts produced by GoReleaser to what `sops` produces. Therefore, this version introduces a breaking change where we no longer produce artifacts like `sopstool_linux.(deb|rpm|tar.gz)` and `sopstool_darwin.tar.gz`. Instead, you'll see artifacts like `sopstool_darwin_(arm64|amd64)_(deb|rpm|tar.gz)` and `sopstool_linux_(arm64|amd64)_(deb|rpm|tar.gz)` in future releases.

## Installation

### Package Repositories

sopstool is available in the following repositories

- homebrew via the `Ibotta/public` tap: `brew install Ibotta/public/sopstool`
- asdf (and rtx) via the `sopstool` plugin: `asdf plugin add sopstool`

### Container Image

Images are tagged with the same version numbering as the releases, and `latest` always gets the latest release. Note that your image will need root CA certificates (typically installed with curl, or a `ca-certificates` package).

To use sopstool from container (avoiding doing binary installs):

```sh
docker run --rm -v $(pwd):/work -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -e AWS_REGION -e AWS_SECURITY_TOKEN -e AWS_SESSION_TOKEN ghcr.io/ibotta/sopstool:latest $COMMAND
```

- `sopstool` is the entrypoint, so any sopstool subcommand can be run.
- `/work` is the default WORKDIR - this should be mounted to the root where `.sops.yml` is stored.
- the commands need access to your AWS credentials session to authenticate KMS.

Or, use as a install source in your Dockerfile. `sops` and `sopstool` are in `/usr/local/bin/`:

```docker
COPY --from=ghcr.io/ibotta/sopstool:latest usr/local/bin/sops usr/local/bin/sopstool /usr/local/bin/
```

### Packages or binaries from Releases

Check the [Releases](https://github.com/Ibotta/sopstool/releases) for the latest artifacts

- Binaries (compressed as .tar.gz or .zip) (note, you will need `sops` installed manually)
- RPM, Debian and APK packages

All artifacts have their sha256 checksums recorded in `sopstool_checksums.txt`, and SPDX SBOM artifacts are available.

### Shell installer

The most direct install uses a shell script hosted in this repository. This script will install the latest sops (if the command does not exist) and sopstool to `./bin` by default.

```sh
curl https://raw.githubusercontent.com/Ibotta/sopstool/main/install.sh | bash
```

- Override the sops version with the `-s` argument
- Override the sopstool version with the `-t` argument
- Override the binary install location with the `-b` argument
  - remember, you may need `sudo` or root access if you are installing to `/usr/*`

Example with overrides:

```sh
curl https://raw.githubusercontent.com/Ibotta/sopstool/main/install.sh | bash -s -- -b /usr/local/bin -s 3.0.0 -t 0.3.0
```

### Installing sops manually

sopstool requires [sops](https://github.com/mozilla/sops). You can use one of the following methods:

- From one of the public repositories (it is available in most)
- From the [official releases](https://github.com/mozilla/sops/releases)

#### Installing the sops binary with our script installer

The install script above uses a separate script to download sops

```sh
curl https://raw.githubusercontent.com/Ibotta/sopstool/main/sopsinstall.sh | bash
```

- Override the tag with the first shell argument (defaults to latest)
- Override the binary install location with the -b flag (defaults to `/.bin`)

#### Download sops from our https mirror

To avoid needing to find the 'latest' binary by hand or by script, use our https server to download the binary. The latest binary is uploaded automatically whenever sopstool is deployed. The file has the pattern `sops_$OS_$ARCH`, except for `windows`

- OS: `linux`, `darwin`
  - ARCH: `amd64`, `arm64`
  - filenames: `sops_$OS_$ARCH.tar.gz`
- OS: `windows`
  - ARCH `amd64` only
  - filename: `sops_windows.zip`
- Versions
  - latest: `https://oss-pkg.ibotta.com/sops/$filename`
  - specific tags: `https://oss-pkg.ibotta.com/sops/$TAG/$filename`

### Installing sopstool manually

Following the lead of [sops](https://github.com/mozilla/sops), we only build 64bit binaries.

#### Installing the sopstool binary using our script installer

The install script above uses a separate script to download sopstool

```sh
curl https://raw.githubusercontent.com/Ibotta/sopstool/main/sopstoolinstall.sh | bash
```

- Override the tag with the first shell argument (defaults to latest)
- Override the binary install location with the -b flag (defaults to `/.bin`)

#### Download sopstool from our https mirror

To avoid needing to find the 'latest' binary by hand or by script, use our https server to download the binary. The latest binary is uploaded automatically whenever sopstool is deployed.

- OS: `linux`, `darwin`
  - ARCH: `amd64`, `arm64`
  - filenames: `sopstool_$OS_$ARCH.tar.gz`
- OS: `windows`
  - ARCH: `amd64`, `arm64`
  - filename: `sopstool_windows_$ARCH.zip`
- Versions
  - latest: `https://oss-pkg.ibotta.com/sopstool/$filename`
  - specific tags: `https://oss-pkg.ibotta.com/sopstool/$TAG/$filename`

Additionally, all other release assets are also within this folder. This includes the checksums, packages, sboms, as well as installers:

- `https://oss-pkg.ibotta.com/sopstool/install.sh` for the combined installer
- `https://oss-pkg.ibotta.com/sopstool/sopsinstall.sh` for the sops installer
- `https://oss-pkg.ibotta.com/sopstool/sopstoolinstall.sh` for the sopstool installer

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

   - this will be at the root of your project. this file is used to both configure keys as well as hold the list of files managed.
   - it needs to specify at least one KMS key accessible by your environment

     ```yaml
     creation_rules:
       - kms: arn:aws:kms:REGION:ACCOUNT:key/KEY_ID
     ```

   - it can specify more complex cases of patterns vs keys too (see link)

## How-To

1. Create a [KMS Key](https://aws.amazon.com/kms/).
1. Follow along the [Configuration Steps](https://github.com/Ibotta/sopstool/tree/main/#configuration), and place the `.sops.yaml` file at the root directory where your scripts will run.
   - All files added to SOPS are relative, or in child directories to the `.sops.yaml` configuration file.
1. Create a file to encrypt(any extension other than `.yaml` if you wish to do the **ENTIRE** file), or create a yaml file with `key: value` pairs(and make sure it's extension is `.yaml`). Sops will encrypt the values, but not it's keys.
   - You can read more about [SOPS Here](https://github.com/mozilla/sops).
1. At this point, `sopstool` is ready and you can now `sopstool add filename`. You'll notice it will create a `filename.sops.extension`. This is your newly encrypted file.
   - When your files are properly encyrepted, you can run `sopstool clean` to remove the original plain text secret files.
1. Now, you can interact via the command line in various ways.
   - **Editing an encrypted file** - `sopstool edit filename.sops.extension`. You can also use your original filename too! `sopstool edit filename.extension`
   - **Listing all encrypted files** - `sopstool list`
   - **Removing encrypted file** - `sopstool remove filename.extension`
   - **Display the contents of encrypted file** - `sopstool cat filename.extension`

### Walkthrough

In this walkthrough, we will go through the steps required to get a secure yaml configuration file running.

1. Configure your `.sops.yaml`

   ```yaml
   # .sops.yaml
   creation_rules:
     - kms: arn:aws:kms:REGION:ACCOUNT:key/KEY_ID
   ```

1. Create a secrets yaml configuration file

   ```yaml
   # credentials.yaml
   database.password: supersecretdb
   database.user: supersecretpassword
   redshift:
     user: my.user.name
     password: my.password
   ```

1. Encrypt the newly created file

   ```sh
   sopstool add credentials.yaml
   ```

1. Create a sample script

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

1. Here is what your folder structure would look like to this point(after deleting the unencrypted credentials.yaml file)

   ```text
   my-project/
   ├── .sops.yaml
   ├── credentials.sops.yaml
   └── myscript.py
   ```

1. Accessing credentials

   The flow should be as follows: unencrypt credentials -> run script -> destroy credentials. You can use the `sopstool entrypoint` to achieve this.

   ```sh
   sopstool entrypoint python myscript.py
   ```

## Contributing

Bug reports and pull requests are welcome at <https://github.com/Ibotta/sopstool>

### docs

Generate markdown docs for the commands via

```sh
sopstool docs
```
