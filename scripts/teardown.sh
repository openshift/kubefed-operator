#!/bin/bash

NAMESPACE="${NAMESPACE:-federation-test}"
CONFIGMAP="${CONFIGMAP:-test-configmap}"
CLUSTERNAME="${CLUSTERNAME:-cluster1}"

kubectl delete configmap ${CONFIGMAP} -n ${NAMESPACE}

kubectl delete crd --all

# kill the process id for kubefed-operator process
kill $(ps ax | grep "kubefed-operator" | awk '{print $1}'| head -n 1)

kubectl delete ns ${NAMESPACE}

kind delete cluster --name=${CLUSTERNAME}