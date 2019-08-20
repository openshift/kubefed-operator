#!/usr/bin/env bash

set -e

./scripts/download-binaries.sh

mv kubectl kubefedctl /go/bin

if [ -n "${IMAGE_FORMAT:-}" ] ; then
    IMAGE_NAME=$(sed -e "s,\${component},kubefed-operator," <(echo $IMAGE_FORMAT))
else
    IMAGE_NAME=${IMAGE_NAME:-quay.io/openshift/kubefed-operator:latest}
fi

echo "Running the e2e tests with image: $IMAGE_NAME"

GIT_BRANCH=`git branch | grep "*" | cut -f2 -d " "`
if [ ${GIT_BRANCH} = master ] ; then
    OPERATOR_VERSION="0.1.0"
elif [ X`echo "${GIT_BRANCH}" | awk -F "release-" '{print $2}'` != X ]; then
    RELEASE_BRANCH=`echo "${GIT_BRANCH}" | awk -F "release-" '{print $2}'`
    if [ X`echo "${RELEASE_BRANCH}" | awk -F "0.1.0" '{print $2}'` != X ]; then
	OPERATOR_VERSION="0.1.0"
    else 
	OPERATOR_VERSION="${RELEASE_BRANCH}"
    fi
else
    RELEASE_BRANCH=`echo "${GIT_BRANCH}" | awk -F "release-" '{print $2}'`
    echo "${RELEASE_BRANCH}"
    OPERATOR_VERSION="0.1.0"
fi

echo "Operator version in use is: $OPERATOR_VERSION"

# Check for Namespaced-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Namespaced -i ${IMAGE_NAME} -o ${OPERATOR_VERSION}

# Check for Cluster-scoped deployment

./scripts/smoke-test.sh -n kubefed-test -d olm-openshift -s Cluster -i ${IMAGE_NAME} -o ${OPERATOR_VERSION}

