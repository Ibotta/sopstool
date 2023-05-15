#!/bin/sh
set -e
set -o errexit

usage() {
  this=$1
  cat <<EOF

$this: download binaries for sopstool

Usage: $this [-b bindir] [-o OS] [-a ARCH] [-s SOPS_VERSION] [-t SOPSTOOL_VERSION] [-d] [bindir]
  -b sets bindir or installation directory, Defaults to ./bin
  -o target OS (windows, linux, darwin) - uses uname by default
  -a target architecture (amd64, arm64) - uses uname by default
  -s SOPS_VERSION tag to download
  -t SOPSTOOL_VERSION tag to download
	-f force download to the binary directory even if command exists
  -d turns on debug logging

  SOPS_VERSION overrides the sops version tag downloaded
  SOPSTOOL_VERSION overrides the sopstool version tag downloaded
  [bindir] arg sets the installation directory

  Flags are passed to the installers (-a, -o, -d)

EOF
  exit 2
}

parse_args() {
  BINDIR="${BINDIR:-./bin}"
  while getopts "b:o:a:s:t:dfxh?" arg; do
    case "${arg}" in
      x) TARGET_X=1 && set -x ;;
      b) BINDIR="${OPTARG}" ;;
      o) TARGET_OS="${OPTARG}" ;;
      a) TARGET_ARCH="${OPTARG}" ;;
      s) SOPS_VERSION="${OPTARG}" ;;
      t) SOPSTOOL_VERSION="${OPTARG}" ;;
			f) FORCE=1 ;;
      d) TARGET_DEBUG=1 ;;
      h | \? | *) usage "$0" ;;
    esac
  done
  shift $((OPTIND - 1))
  BINDIR=${1:-${BINDIR}}
}

execute() {
  if [ -n "${TARGET_ARCH}" ]; then
    set -- "$@" "-a" "${TARGET_ARCH}"
  fi
  if [ -n "${TARGET_OS}" ]; then
    set -- "$@" "-o" "${TARGET_OS}"
  fi
  if [ -n "${TARGET_DEBUG}" ]; then
    set -- "$@" "-d"
  fi
  if [ -n "${TARGET_X}" ]; then
    set -- "$@" "-x"
  fi

  if [ -n "${FORCE}" ] || ! is_command sops; then
    http_exec https://raw.githubusercontent.com/Ibotta/sopstool/master/sopsinstall.sh -b "${BINDIR}" "$@" "${SOPS_VERSION}"
  fi

  http_exec https://raw.githubusercontent.com/Ibotta/sopstool/master/sopstoolinstall.sh -b "${BINDIR}" "$@" "${SOPSTOOL_VERSION}"

  echo "Both sops and sopstool installed"
}

http_exec() {
  url="$1"
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
