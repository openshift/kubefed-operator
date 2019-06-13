#!/bin/bash
VERSION="${VERSION:-v0.1.0-rc2}"
OS=$(go env GOOS)
ARCH=$(go env GOARCH)

cd ~/tmp

wget --continue https://github.com/kubernetes-sigs/kubefed/releases/download/${VERSION}/kubefedctl-${VERSION}-${OS}-${ARCH}.tgz &

tar -xvzf kubefedctl-${VERSION}-${OS}-${ARCH}.tgz

chmod u+x kubefedctl

export PATH=$PATH:/usr/local/bin

sudo mv kubefedctl /usr/local/bin/ 
