#!/bin/bash

#VERSION="${VERSION:-v0.1.0-rc2}"
VERSION="${VERSION:-0.0.10}"

#OS="${OS:-linux}"

cd ~/tmp

# wget https://github.com/kubernetes-sigs/kubefed/releases/download/${VERSION}/kubefedctl-${VERSION}-${OS}-amd64.tgz

curl -LO https://github.com/kubernetes-sigs/federation-v2/releases/download/${VERSION}/kubefedctl.tgz

#sudo tar -xvzf kubefedctl-${VERSION}-${OS}-amd64.tgz

sudo tar xzfP kubefedctl.tgz

sudo chmod u+x kubefedctl

export PATH=$PATH:/usr/local/bin

sudo mv kubefedctl /usr/local/bin/ 
