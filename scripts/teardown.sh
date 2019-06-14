#!/bin/bash

NAMESPACE="${NAMESPACE:-default}"
CLUSTERNAME="${CLUSTERNAME:-cluster1}"
LOCATION="${LOCATION:-local}"

while getopts “n:d:” opt; do
    case $opt in
	n) NAMESPACE=$OPTARG ;;
    d) LOCATION=$OPTARG ;;
    esac
done

echo ">> Uninstalling kubefed operator"
  CONFIGMAP=$(kubectl get configmap -n ${NAMESPACE} -o jsonpath='{.items[1].metadata.name}' 2>/dev/null)
  if test X"$CONFIGMAP" = Xtest-configmap; then
  echo ">> Deleting test-configmap resource"
  kubectl delete configmap ${CONFIGMAP} -n ${NAMESPACE}
  fi
 STORAGECLASS=$(kubectl get storageclass -o jsonpath='{.items[1].metadata.name}' 2>/dev/null)
   if test X"$STORAGECLASS" = Xtest-storageclass; then
  echo ">> Deleting test-configmap resource"
  kubectl delete storageclass ${STORAGECLASS}
  fi
 echo ">> Deleting all the CRD's"
 kubectl delete crd --all

# kill the process id for kubefed-operator process
echo ">> Kill the kubefed-operator process"
kill $(ps ax | grep "kubefed-operator" | awk '{print $1}'| head -n 1)

echo ">> Deleting ${NAMESPACE}"
kubectl delete ns ${NAMESPACE}

# For deleting a given cluster
 if test X"$LOCATION" != Xolm-openshift; then
    kind delete cluster --name=${CLUSTERNAME}
 else
  # Assumption: openshift-install binary is already installed 
  # and present in the $PATH
  openshift-install destroy cluster
 fi