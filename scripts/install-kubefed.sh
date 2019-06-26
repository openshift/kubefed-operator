#!/bin/bash

# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./scripts/install-kubefed.sh -n <namespace> -d <location> -i <image_name> -s <scope>
# for a local deployment, you don't need to specify an image_name flag

set -e
 
#default values
NAMESPACE="default"
LOCATION="local"
OLM_VERSION="0.10.0"
OPERATOR_VERSION="0.1.0"
OPERATOR="kubefed-operator"
IMAGE_NAME="quay.io/sohankunkerkar/kubefed-operator:v0.1.0"
OPERATOR_YAML_PATH="./deploy/operator.yaml"
CLUSTER_ROLEBINDING="./deploy/role_binding.yaml"
CSV_PATH="./deploy/olm-catalog/kubefed-operator/${OPERATOR_VERSION}/kubefed-operator.v${OPERATOR_VERSION}.clusterserviceversion.yaml"
SCOPE="Namespaced"

while getopts “n:d:i:s:” opt; do
    case $opt in
	n) NAMESPACE=$OPTARG ;;
	d) LOCATION=$OPTARG ;;
        i) IMAGE_NAME=$OPTARG ;;
        s) SCOPE=$OPTARG;;
    esac
done

echo "NS=$NAMESPACE"
echo "LOC=$LOCATION"
echo "Operator Image Name=$IMAGE_NAME"
echo "Scope=$SCOPE"

if test X"$NAMESPACE" != Xdefault; then
    # create a namespace 
    kubectl create ns ${NAMESPACE}
fi

# Install kubefed CRD
kubectl apply -f ./deploy/crds/operator_v1alpha1_kubefed_crd.yaml

# Install kubefed CR based on the scope
if test X"$SCOPE" = XCluster; then
  sed "s,scope:.*,scope: ${SCOPE}," ./deploy/crds/operator_v1alpha1_kubefed_cr.yaml | kubectl apply -n $NAMESPACE -f -
else
  kubectl apply -f ./deploy/crds/operator_v1alpha1_kubefed_cr.yaml -n $NAMESPACE
fi


# A local deployment
if test X"$LOCATION" = Xlocal; then
  operator-sdk &> /dev/null
  if [ $? == 0 ]; then
  # operator-sdk up local command doesn't install the requried CRD's
  for f in ./deploy/crds/*_crd.yaml ; do     
	  kubectl apply -f "${f}" --validate=false 
  done
	    operator-sdk up local --namespace=$NAMESPACE &
  else
	  echo "Operator SDK is not installed."
	  exit 1
  fi

# in-cluster deployment on kind cluster
elif test X"$LOCATION" = Xcluster; then
  for f in ./deploy/*.yaml ; do
   if test X"$OPERATOR_YAML_PATH" = X"$f" ; then
      echo "Reading the image name and sed it in"
      sed "/image: /s|: .*|: ${IMAGE_NAME}|" $f | kubectl apply -n $NAMESPACE --validate=false -f -
   elif test X"$CLUSTER_ROLEBINDING" = X"$f" ; then
      echo "Reading the namespace in clusterrolebinding and sed it in"
      sed "/namespace: /s|: .*|: ${NAMESPACE}|" $f | kubectl -n $NAMESPACE apply -f -
   else
      kubectl apply -f "${f}" --validate=false -n $NAMESPACE
   fi
  done
  for f in ./deploy/crds/*_crd.yaml ; do     
	  kubectl apply -f "${f}" --validate=false 
  done
  echo "Deployed all the operator yamls for kubefed-operator in the cluster"

# olm-deployment on minikube cluster
elif test X"$LOCATION" = Xolm-kube; then
 ./scripts/kubernetes/olm-install.sh ${OLM_VERSION}
 
 echo "OLM is deployed on kube cluster"
 sed "s,image: quay.*$,image: ${IMAGE_NAME}," -i.bak $CSV_PATH|./hack/catalog.sh | kubectl apply -n $NAMESPACE -f -
 mv $CSV_PATH.bak $CSV_PATH
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

# olm deployment on openshift cluster   
elif test X"$LOCATION" = Xolm-openshift; then
 sed "s,image: quay.*$,image: ${IMAGE_NAME}," -i.bak $CSV_PATH|./hack/catalog.sh | oc apply -n $NAMESPACE -f -
 mv $CSV_PATH.bak $CSV_PATH
 cat <<-EOF | oc apply -f -
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
else
  echo "Please enter the valid location"
  exit 1
fi
