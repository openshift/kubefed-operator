#!/bin/bash

# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./install.sh -n <namespace> -d <location>

set -e
 

NAMESPACE=""
LOCATION=""
NAMESPACE_STR=""
OLM_VERSION="0.10.0"
OPERATOR="kubefed-operator"
IMAGE_NAME=""
OPERATOR_YAML_PATH="./deploy/operator.yaml"

while getopts “n:d:i:” opt; do
    case $opt in
	n) NAMESPACE=$OPTARG ;;
	d) LOCATION=$OPTARG ;;
    i) IMAGE_NAME=$OPTARG ;;
    esac
done

echo "NS=$NAMESPACE"
echo "LOC=$LOCATION"
echo "Operator Image Name=$IMAGE_NAME"

if test X"$NAMESPACE" != X; then
    # create a namespace 
    kubectl create ns ${NAMESPACE}
    NAMESPACE_STR="--namespace=${NAMESPACE}"
fi

# Install CRD
kubectl apply -f ./deploy/crds/operator_v1alpha1_kubefed_crd.yaml

# Install CR
kubectl apply -f ./deploy/crds/operator_v1alpha1_kubefed_cr.yaml $NAMESPACE_STR

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
elif test X"$LOCATION" = Xcluster; then
  #TODO: change the location in the container stanza of the operator yaml
  for f in ./deploy/*.yaml ; do
   if test X"$OPERATOR_YAML_PATH" = X"$f" ; then
      echo "Reading the image name and sed it in"
      sed "/image: /s|: .*|: ${IMAGE_NAME}|" $f | kubectl apply $NAMESPACE_STR --validate=false -f -
    else
      kubectl apply -f "${f}" --validate=false $NAMESPACE_STR
    fi
  done
  for f in ./deploy/crds/*_crd.yaml ; do     
	  kubectl apply -f "${f}" --validate=false 
  done
  echo "Deployed all the operator yamls for kubefed-operator in the cluster"

elif test X"$LOCATION" = Xolm-kube; then
./scripts/kubernetes/olm-install.sh ${OLM_VERSION}

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
spec:
 targetNamespaces:
   - ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: ${OPERATOR}-sub
  generateName: ${OPERATOR}-
  namespace: ${NAMESPACE}
spec:
  source: ${OPERATOR}
  sourceNamespace: ${NAMESPACE}
  name: ${OPERATOR}
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
  name: ${OPERATOR}-sub
  generateName: ${OPERATOR}-
  namespace: ${NAMESPACE}
spec:
  source: ${OPERATOR}
  sourceNamespace: ${NAMESPACE}
  name: ${OPERATOR}
  channel: alpha
EOF
   
fi



