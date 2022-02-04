#!/usr/bin/env bash
set -eu

cd "$(dirname "${BASH_SOURCE[0]}")"

ROOT_DIR="$(realpath ".")"
export GOPATH="${ROOT_DIR}/workspace"
echo $GOPATH
(go build -o "${ROOT_DIR}/bin/compiler" mit.edu/compilers/compiler)