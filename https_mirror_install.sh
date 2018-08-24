#!/bin/sh
set -e
set -o errexit

usage() {
  this=$1
  cat <<EOF

$this: download binaries for sopstool and sops from the https mirror

Usage: $this [bindir]
  [bindir] sets bindir or installation directory, Defaults to ./bin

  SOPS_VERSION overrides the sops version tag downloaded
  SOPSTOOL_VERSION overrides the sopstool version tag downloaded

Inspired by godownloader
 https://github.com/goreleaser/godownloader

EOF
  exit 2
}

cd $(mktemp -d)
curl -o sopstool.tar.gz https://oss-pkg.ibotta.com/sopstool/sopstool_linux.tar.gz
tar -xvzf sopstool.tar.gz
curl -o sops.tar.gz https://oss-pkg.ibotta.com/sops/sops_linux.tar.gz
tar -xvzf sops.tar.gz
mv sops /bin
mv sopstool /bin
rm -r $(pwd)
