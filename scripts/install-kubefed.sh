#!/bin/bash

# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./install.sh <location> <namespace>

set -e
 

NAMESPACE=""
LOCATION=""
NAMESPACE_STR=""

while getopts “n:d:” opt; do
    case $opt in
	n) NAMESPACE=$OPTARG ;;
	d) LOCATION=$OPTARG ;;
    esac
done

echo "NS=$NAMESPACE"
echo "LOC=$LOCATION"

if test X"$NAMESPACE" != X; then
    # create a namespace 
    kubectl create ns ${NAMESPACE}
    NAMESPACE_STR="--namespace=${NAMESPACE}"
fi

# Install crds 
kubectl apply -f ./deploy/crds/operator_v1alpha1_install_crd.yaml

# Install CR
kubectl apply -f ./deploy/crds/operator_v1alpha1_install_cr.yaml $NAMESPACE_STR

# Check if operator-sdk is installed or not and accordinlgy execute the command.
if test X"$LOCATION" = Xlocal; then
    operator-sdk &> /dev/null
    if [ $? == 0 ]; then
    # operator-sdk up local command doesn't install the requried CRD's
    for f in ./deploy/crds/*_crd.yaml ; do     
	kubectl apply -f "${f}" --validate=false 
    done
	operator-sdk up local $NAMESPACE_STR &
    else
	echo "Operator SDK is not installed."
	exit 1
    fi
elif test X"$LOCATION" = Xapply; then
    #TODO: change the location in the container stanza of the operator yaml
    for f in ./deploy/crds/*_crd.yaml ; do     
	kubectl apply -f "${f}" --validate=false 
    done
    echo "Deployed all the operator yamls for kubefed-operator in the cluster"
elif test X"$LOCATION" = Xolm-kube; then

kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.10.0/crds.yaml

kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.10.0/olm.yaml

echo "OLM is deployed in the cluster"
 
./hack/catalog.sh | kubectl apply $NAMESPACE_STR -f -

cat <<-EOF | kubectl apply -f -
---
apiVersion: v1
kind: Namespace
metadata:
  name: ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: kubefed
  namespace: ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: kubefed-operator-sub
  generateName: kubefed-operator-
  namespace: ${NAMESPACE}
spec:
  source: kubefed-operator
  sourceNamespace: ${NAMESPACE}
  name: kubefed-operator
  channel: alpha
EOF
elif test X"$LOCATION" = Xolm-openshift; then

./hack/catalog.sh | oc apply $NAMESPACE_STR -f -

cat <<-EOF | kubectl apply -f -
---
apiVersion: v1
kind: Namespace
metadata:
  name: ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: kubefed
  namespace: ${NAMESPACE}
spec:
 targetNamespaces:
   - ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: kubefed-operator-sub
  generateName: kubefed-operator-
  namespace: ${NAMESPACE}
spec:
  source: kubefed-operator
  sourceNamespace: ${NAMESPACE}
  name: kubefed-operator
  channel: alpha
EOF
   
fi



