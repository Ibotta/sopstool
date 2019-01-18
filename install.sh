#!/bin/sh
set -e
set -o errexit

usage() {
  this=$1
  cat <<EOF

$this: download binaries for sopstool

Usage: $this [bindir]
  [bindir] sets bindir or installation directory, Defaults to ./bin

  SOPS_VERSION overrides the sops version tag downloaded
  SOPSTOOL_VERSION overrides the sopstool version tag downloaded

  Consider setting GITHUB_TOKEN to avoid triggering GitHub rate limits.
  See the following for more details:
  https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/

Inspired by godownloader
 https://github.com/goreleaser/godownloader

EOF
  exit 2
}

parse_args() {
  BINDIR=${1:-"./bin"}
}

execute() {
  if ! is_command sops; then
    http_exec https://raw.githubusercontent.com/Ibotta/sopstool/master/sopsinstall.sh -b "$BINDIR" "$SOPS_VERSION"
  fi

  http_exec https://raw.githubusercontent.com/Ibotta/sopstool/master/sopstoolinstall.sh -b "$BINDIR" "$SOPSTOOL_VERSION"

  echo "Both sops and sopstool installed"
}

http_exec() {
  url=$1
  shift
  if is_command curl; then
    cmd='curl --fail -sSL'
  elif is_command wget; then
    cmd='wget -qO-'
  else
    echo "http_exec: unable to find wget or curl"
    return 1
  fi
  $cmd "$url" | $SHELL -s -- "$@"
}

cat /dev/null <<EOF
------------------------------------------------------------------------
https://github.com/client9/shlib - portable posix shell functions
Public domain - http://unlicense.org
https://github.com/client9/shlib/blob/master/LICENSE.md
but credit (and pull requests) appreciated.
------------------------------------------------------------------------
EOF
is_command() {
  command -v "$1" >/dev/null
}
cat /dev/null <<EOF
------------------------------------------------------------------------
End of functions from https://github.com/client9/shlib
------------------------------------------------------------------------
EOF

# parse_args, show usage and exit if necessary
parse_args "$@"

# do it
execute
