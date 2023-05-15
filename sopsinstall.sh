#!/bin/sh
set -e

usage() {
  this=$1
  cat <<EOF

$this: download binaries for mozilla/sops

Usage: $this [-b bindir] [-o OS] [-a ARCH] [-z dir] [-d] [tag]
  -b sets bindir or installation directory, Defaults to ./bin
  -o target OS (windows, linux, darwin) - uses uname by default
  -a target architecture (amd64, arm64) - uses uname by default
  -z save an archive instead of installing
  -d turns on debug logging
  [tag] is a tag from https://github.com/mozilla/sops/releases
        If tag is missing, then the latest release will be used.

EOF
  exit 2
}
parse_args() {
  BINDIR="${BINDIR:-./bin}"
  while getopts "b:o:a:z:dxh?" arg; do
    case "${arg}" in
      b) BINDIR="${OPTARG}" ;;
      o) TARGET_OS="${OPTARG}" ;;
      a) TARGET_ARCH="${OPTARG}" ;;
      z) ARCHIVEDIR="${OPTARG}" ;;
      d) log_set_priority 10 ;;
      x) set -x ;;
      h | \? | *) usage "$0" ;;
    esac
  done
  shift $((OPTIND - 1))
  TAG="$1"
}
tag_to_version() {
  if [ -z "${TAG}" ]; then
    log_info "checking GitHub for latest tag"
  else
    LATEST=""
    log_info "checking GitHub for tag '${TAG}'"
  fi
  REALTAG=$(github_release "${OWNER}/${REPO}" "${TAG}") && true
  if test -z "${REALTAG}"; then
    log_crit "unable to find '${TAG}' - use 'latest' or see https://github.com/${PREFIX}/releases for details"
    exit 1
  fi
  # if version starts with 'v', remove it
  TAG="${REALTAG}"
  VERSION="${TAG#v}"
}
adjust_binary() {
  # cribbed from https://havoc.io/post/shellsemver/
  oldversion=""
  MIN_V_VERSION="3.5.0"
  LOWEST_V_VERSION=$(printf "${VERSION}\n${MIN_V_VERSION}" |
    sort -t "." -n -k1,1 -k2,2 -k3,3 -k4,4 | head -n 1)
  if [ "${LOWEST_V_VERSION}" != "${MIN_V_VERSION}" ]; then
    oldversion="${VERSION}"
  fi

  oldarch=""
  MIN_ARCH_VERSION="3.7.2"
  LOWEST_ARCH_VERSION=$(printf "${VERSION}\n${MIN_ARCH_VERSION}" |
    sort -t "." -n -k1,1 -k2,2 -k3,3 -k4,4 | head -n 1)
  if [ "${LOWEST_ARCH_VERSION}" != "${MIN_ARCH_VERSION}" ]; then
    oldarch="${VERSION}"
  fi

  if [ "${OS}" = "windows" ]; then
    if [ "${ARCH}" != "amd64" ]; then
      log_crit "unsupported ARCH=${ARCH} for ${TAG}"
      return 1
    fi
    if [ -n "${oldversion}" ]; then
      NAME="sops-${VERSION}.exe"
    else
      NAME="sops-v${VERSION}.exe"
    fi
    BINARY="${BINARY}.exe"
  else
    if [ -n "${oldarch}" ] && [ "${ARCH}" != "amd64" ]; then
      log_crit "unsupported ARCH=${ARCH} for ${TAG}"
      return 1
    fi
    if [ -n "${oldversion}" ]; then
      NAME="sops-${VERSION}.${OS}"
    else
      NAME="sops-v${VERSION}.${OS}.${ARCH}"
    fi
  fi
}
install_binary() {
  TMPDIR="$1"
  test ! -d "${BINDIR}" && install -d "${BINDIR}"
  install "${TMPDIR}/${NAME}" "${BINDIR}/${BINARY}"
  log_info "installed ${BINDIR}/${BINARY}"
}
archive_binary() {
  chmod a+x "${TMPDIR}/${NAME}"
  TAGDIR="${ARCHIVEDIR}/${TAG}"
  mkdir -p "${TAGDIR}"

  if [ "${OS}" = "windows" ]; then
    ARCHIVE="sops_${OS}.zip"
  else
    ARCHIVE="sops_${OS}_${ARCH}.tar.gz"
  fi

  ARCHIVE_TARGET="$(get_abs_filename "${TAGDIR}/${ARCHIVE}")"
  compress_binary "${ARCHIVE_TARGET}" "${NAME}"
  log_info "archived to ${TAGDIR}/${ARCHIVE}"
  if [ -n "$LATEST" ]; then
    cp "${TAGDIR}/${ARCHIVE}" "${ARCHIVEDIR}/${ARCHIVE}"
    log_info "marked as latest"
  fi
}
compress_binary() {
  if [ "${OS}" = "windows" ]; then
    (cd "${TMPDIR}" && zip "${ARCHIVE_TARGET}" "${NAME}")
  else
    (cd "${TMPDIR}" && tar -czf "${ARCHIVE_TARGET}" "${NAME}")
  fi
}
get_abs_filename() {
  # $1 : relative filename
  echo "$(cd "$(dirname "$1")" && pwd)/$(basename "$1")"
}
# wrap all destructive operations into a function
# to prevent curl|bash network truncation and disaster
execute() {
  TMPDIR=$(mktemp -d)
  log_info "downloading from ${TARBALL_URL}"
  http_download "${TMPDIR}/${NAME}" "${TARBALL_URL}"
  if [ "$ARCHIVEDIR" != "" ]; then
    archive_binary "$TMPDIR"
  else
    install_binary "$TMPDIR"
  fi
  rm -rf "${TMPDIR}"
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
echoerr() {
  echo "$@" 1>&2
}
log_prefix() {
  echo "$0"
}
_logp=6
log_set_priority() {
  _logp="$1"
}
log_priority() {
  if test -z "$1"; then
    echo "$_logp"
    return
  fi
  [ "$1" -le "$_logp" ]
}
log_tag() {
  case $1 in
    0) echo "emerg" ;;
    1) echo "alert" ;;
    2) echo "crit" ;;
    3) echo "err" ;;
    4) echo "warning" ;;
    5) echo "notice" ;;
    6) echo "info" ;;
    7) echo "debug" ;;
    *) echo "$1" ;;
  esac
}
log_debug() {
  log_priority 7 || return 0
  echoerr "$(log_prefix)" "$(log_tag 7)" "$@"
}
log_info() {
  log_priority 6 || return 0
  echoerr "$(log_prefix)" "$(log_tag 6)" "$@"
}
log_err() {
  log_priority 3 || return 0
  echoerr "$(log_prefix)" "$(log_tag 3)" "$@"
}
log_crit() {
  log_priority 2 || return 0
  echoerr "$(log_prefix)" "$(log_tag 2)" "$@"
}
uname_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    cygwin_nt*) os="windows" ;;
    mingw*) os="windows" ;;
    msys_nt*) os="windows" ;;
  esac
  echo "$os"
}
uname_arch() {
  arch=$(uname -m)
  case $arch in
    x86_64) arch="amd64" ;;
    x86) arch="386" ;;
    i686) arch="386" ;;
    i386) arch="386" ;;
    aarch64) arch="arm64" ;;
    armv5*) arch="armv5" ;;
    armv6*) arch="armv6" ;;
    armv7*) arch="armv7" ;;
  esac
  echo ${arch}
}
uname_os_check() {
  os=$1
  case "$os" in
    darwin) return 0 ;;
    linux) return 0 ;;
    windows) return 0 ;;
  esac
  log_crit "uname_os_check '$os' is not a supported value."
  return 1
}
uname_arch_check() {
  arch=$1
  case "$arch" in
    amd64) return 0 ;;
    arm64) return 0 ;;
  esac
  log_crit "uname_arch_check '$arch' is not a supported value."
  return 1
}
http_download_curl() {
  local_file=$1
  source_url=$2
  header=$3
  if [ -z "$header" ]; then
    code=$(curl -w '%{http_code}' -sL -o "$local_file" "$source_url")
  else
    code=$(curl -w '%{http_code}' -sL -H "$header" -o "$local_file" "$source_url")
  fi
  if [ "$code" != "200" ]; then
    log_debug "http_download_curl received HTTP status $code"
    return 1
  fi
  return 0
}
http_download_wget() {
  local_file=$1
  source_url=$2
  header=$3
  if [ -z "$header" ]; then
    wget -q -O "$local_file" "$source_url"
  else
    wget -q --header "$header" -O "$local_file" "$source_url"
  fi
}
http_download() {
  log_debug "http_download $2"
  if is_command curl; then
    http_download_curl "$@"
    return
  elif is_command wget; then
    http_download_wget "$@"
    return
  fi
  log_crit "http_download unable to find wget or curl"
  return 1
}
http_copy() {
  tmp=$(mktemp)
  http_download "${tmp}" "$1" "$2" || return 1
  body=$(cat "$tmp")
  rm -f "${tmp}"
  echo "$body"
}
github_release() {
  owner_repo=$1
  version=$2
  test -z "$version" && version="latest"
  giturl="https://github.com/${owner_repo}/releases/${version}"
  json=$(http_copy "$giturl" "Accept:application/json")
  test -z "$json" && return 1
  version=$(echo "$json" | tr -s '\n' ' ' | sed 's/.*"tag_name":"//' | sed 's/".*//')
  test -z "$version" && return 1
  echo "$version"
}
cat /dev/null <<EOF
------------------------------------------------------------------------
End of functions from https://github.com/client9/shlib
------------------------------------------------------------------------
EOF

OWNER="mozilla"
REPO="sops"
BINARY="sops"
BINDIR="${BINDIR:-./bin}"
PREFIX="${OWNER}/${REPO}"
ARCHIVEDIR=""
LATEST="latest"

# use in logging routines
log_prefix() {
  echo "${PREFIX}"
}

GITHUB_DOWNLOAD=https://github.com/${OWNER}/${REPO}/releases/download

# parse_args, show usage and exit if necessary
parse_args "$@"

OS="${TARGET_OS:-$(uname_os)}"
ARCH="${TARGET_ARCH:-$(uname_arch)}"

# make sure we are on a platform that makes sense
uname_os_check "${OS}"
uname_arch_check "${ARCH}"

# setup version from tag
tag_to_version

log_info "found version ${VERSION} for ${TAG}/${OS}/${ARCH}"

# adjust binary name based on OS and version
adjust_binary

# compute URL to download
TARBALL_URL="${GITHUB_DOWNLOAD}/${TAG}/${NAME}"

# do it
execute
