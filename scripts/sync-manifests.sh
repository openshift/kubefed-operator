#!/usr/bin/env bash

# This script synchronizes the CRDs for the upstream kubefed version requested,
# into this repository.
#
# Example usage:
#   ./scripts/sync-manifests.sh v0.1.0-rc4

set -o errexit
set -o nounset
set -o pipefail

KUBEFED_VERSION=${1:-}
if [[ -z "${KUBEFED_VERSION}" ]]; then
  >&2 echo "Usage: $0 <kubefed_version|master>"
  exit 1
fi

ROOT_DIR="$(cd "$(dirname "$0")/.." ; pwd)"
TMP_GENDIR="$(mktemp -d /tmp/kubefed-gen.XXXXXXXX)"

# Called on EXIT after the temporary directories have been created.
function clean-up() {
  if [[ "${TMP_GENDIR}" == "/tmp/kubefed-gen."* ]]; then
    rm -rf "${TMP_GENDIR}"
  fi
}

trap clean-up EXIT

# Set up temporary workspace
mkdir -p "${TMP_GENDIR}/src"
pushd "${TMP_GENDIR}/src"

# Clone repo at specific KUBEFED_VERSION
git clone --branch "${KUBEFED_VERSION}" https://github.com/kubernetes-sigs/kubefed.git sigs.k8s.io/kubefed
cd sigs.k8s.io/kubefed/

# Generate kubefed CRDs
export GOPATH="${TMP_GENDIR}"
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd

# Rename CRDs to convention used by CRD generation code:
# <group>_<version>_<kind>_crd.yaml
for file in config/crds/*.yaml; do
  mv "${file}" "${file%.yaml}_crd.yaml"
done

# Strip off 'v' prefix and anything after the last '-' e.g. v0.1.0-rc4 -> 0.1.0
kubefed_version_crd_dir="${KUBEFED_VERSION#v}"
kubefed_version_crd_dir="${kubefed_version_crd_dir%-*}"

# Copy CRDs to respective directories
crd_dirs="${ROOT_DIR}/deploy/crds"
crd_dirs+=" ${ROOT_DIR}/deploy/olm-catalog/kubefed-operator/${kubefed_version_crd_dir}"
for dir in ${crd_dirs}; do
  if [[ ! -d "${dir}" ]]; then
    echo "No such directory ${dir} exists. Creating it..."
    mkdir -p "${dir}"
  fi
  cp -f config/crds/*.yaml "${dir}/"
done

popd
