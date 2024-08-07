#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
ROOT_PKG=github.com/41tair/rules-controller

GET_PKG_LOCATION() {
  pkg_name="${1:-}"

  pkg_location="$(go list -m -f '{{.Dir}}' "${pkg_name}" 2>/dev/null)"
  if [ "${pkg_location}" = "" ]; then
    echo "${pkg_name} is missing. Running 'go mod download'."

    go mod download
    pkg_location=$(go list -m -f '{{.Dir}}' "${pkg_name}")
  fi
  echo "${pkg_location}"
}

# Grab code-generator version from go.sum
CODEGEN_PKG="$(GET_PKG_LOCATION "k8s.io/code-generator")"
echo ">> Using ${CODEGEN_PKG}"

source "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/pkg/apis"


kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/pkg/generated" \
    --output-pkg "${ROOT_PKG}/pkg/generated" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/pkg/apis"



