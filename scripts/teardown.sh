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
  CONFIGMAP=$(kubectl get configmap -n ${NAMESPACE} -o jsonpath='{.items[*].metadata.name}' 2>/dev/null)
  if [[ $CONFIGMAP == *"test-configmap"* ]]; then
  echo ">> Deleting test-configmap resource"
  kubectl delete configmap ${CONFIGMAP} -n ${NAMESPACE}
  kubectl delete  federatedconfigmap  test-configmap -n ${NAMESPACE}
  fi
 STORAGECLASS=$(kubectl get storageclass -o jsonpath='{.items[*].metadata.name}' 2>/dev/null)
   if [[ $STORAGECLASS == *"test-storageclass"* ]]; then
  echo ">> Deleting test-configmap resource"
  kubectl delete storageclass ${STORAGECLASS}
  kubectl delete  federatedstorageclass  test-storageclass
  fi
 echo ">> Deleting all the CRDs related to kubefed"
 kubectl get crd | awk '/kubefed/{print $1}' | xargs kubectl delete crd

# kill the process id for kubefed-operator process
echo ">> Kill the kubefed-operator process"
kill $(ps ax | grep "kubefed-operator" | awk '{print $1}'| head -n 1)

echo ">> Deleting Namespaces"
if [[ "$LOCATION" == "olm-openshift" && "$NAMESPACE" == "default" ]]; then
   kubectl delete ns olm
   kubectl delete ns operators 
elif test X"$NAMESPACE" != Xdefault ; then
     kubectl delete ns ${NAMESPACE}
else
   echo "Skipping this step as ${NAMESPACE} namespace may not be deleted "
fi
# For deleting a given cluster
 if [[ "$LOCATION" != "olm-openshift" ]]; then
    kind delete cluster --name=${CLUSTERNAME}
 else
   # Assumption: openshift-install binary is already installed 
   # and present in the $PATH
   #openshift-install destroy cluster
    echo "Please delete the openshift cluster on your own!"
 fi
 
