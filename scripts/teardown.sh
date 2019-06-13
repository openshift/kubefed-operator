#!/bin/bash

echo ">> Uninstalling kubefed operator"
NAMESPACE="${NAMESPACE:-kubefed-test}"
CONFIGMAP="${CONFIGMAP:-test-configmap}"
CLUSTERNAME="${CLUSTERNAME:-cluster1}"

echo ">> Deleting test-configmap resource"
kubectl delete configmap ${CONFIGMAP} -n ${NAMESPACE}

echo ">> Deleting all the CRD's"
kubectl delete crd --all

# kill the process id for kubefed-operator process
echo ">> Kill the kubefed-operator process"
kill $(ps ax | grep "kubefed-operator" | awk '{print $1}'| head -n 1)

echo ">> Deleting ${NAMESPACE}"
kubectl delete ns ${NAMESPACE}

kind delete cluster --name=${CLUSTERNAME}
