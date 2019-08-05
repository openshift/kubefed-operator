#!/bin/bash
# This script automates the downloading of binaries 
# used for smoke testing
CURL_ARGS="-Ls"
[[ -z "${DEBUG:-""}" ]] || {
  set -x
  CURL_ARGS="-L"
}

logEnd() {
  local msg='done.'
  [ "$1" -eq 0 ] || msg='Error downloading assets'
  echo "$msg"
}
trap 'logEnd $?' EXIT

echo "About to download some binaries. This might take a while..."

KUBEFED_VERSION="${KUBEFED_VERSION:-0.1.0-rc5}"
KUBECTL_VERSION="${KUBECTL_VERSION:-1.15.0}"
OS=$(go env GOOS)
ARCH=$(go env GOARCH)

# Install kubefedctl
curl "${CURL_ARGS}"O https://github.com/kubernetes-sigs/kubefed/releases/download/v${KUBEFED_VERSION}/kubefedctl-${KUBEFED_VERSION}-${OS}-${ARCH}.tgz 

tar -xvzf kubefedctl-${KUBEFED_VERSION}-${OS}-${ARCH}.tgz

chmod u+x kubefedctl

rm kubefedctl-${KUBEFED_VERSION}-${OS}-${ARCH}.tgz

# Install kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl

chmod +x ./kubectl
