#!/bin/bash

set -e

./scripts/download-binaries.sh

mv kubectl kubefedctl /go/bin

if [ -n "${IMAGE_FORMAT:-}" ] ; then
    IMAGE_NAME=$(sed -e "s,\${component},kubefed-operator," <(echo $IMAGE_FORMAT))
else
    IMAGE_NAME=${IMAGE_NAME:-quay.io/openshift/kubefed-operator:latest}
fi

echo "Running the e2e tests with image: $IMAGE_NAME"

# Check for Namespaced-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Namespaced -i ${IMAGE_NAME}

# Check for Cluster-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Cluster -i ${IMAGE_NAME}
