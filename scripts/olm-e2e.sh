#!/bin/bash

set -e

./scripts/download-binaries.sh

mv kubectl kubefedctl /go/bin

# Check for Namespaced-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Namespaced

# Check for Cluster-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Cluster
