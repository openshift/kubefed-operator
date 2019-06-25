#!/bin/bash

#Default values
NAMESPACE="${NAMESPACE:-default}"
LOCATION="${LOCATION:-local}"
VERSION="${VERSION:-v0.1.0-rc2}"
IMAGE_NAME="${IMAGE_NAME:-quay.io/sohankunkerkar/kubefed-operator:v0.1.0}"
SCOPE="${SCOPE:-Namespaced}"
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

function setup-infrastructure () {

   if [[ "$LOCATION" != "olm-openshift" && "$LOCATION" != "olm-kube" ]]; then
      ./scripts/create-cluster.sh
   elif test X"$LOCATION" = Xolm-kube; then
       minikube start
   fi
  
  ./scripts/install-kubefed.sh -n ${NAMESPACE} -d ${LOCATION} -i ${IMAGE_NAME} -s ${SCOPE} &

  retries=100
  until [[ $retries == 0 || $RESOURCE =~ "kubefed" ]]; do
    RESOURCE=$(kubectl get kubefedconfig -n ${NAMESPACE} -o jsonpath='{.items[*].metadata.name}' 2>/dev/null)
    if [[ $RESOURCE != *"kubefed"* ]]; then
        echo "Waiting for kubefedconfig to appear"
        sleep 1
        retries=$((retries - 1))
    fi
  done

  if [ $retries == 0 ]; then
    echo "Failed to retrieve kubefedconfig resource"
    exit 1
  fi

  #./scripts/download-binaries.sh
  
}

function enable-resources () {

if test X"$LOCATION" = Xolm-openshift || test X"$LOCATION" = Xolm-kube ; then
  # renaming context for openshift cluster to consumable format
  oc config rename-context $(oc config current-context) cluster1
fi

echo "Performing the join operation on cluster1"
kubefedctl join cluster1 --kubefed-namespace=${NAMESPACE} --host-cluster-context=cluster1 --host-cluster-name=cluster1 --cluster-context=cluster1

if test X"$SCOPE" = XNamespaced; then
echo "Enable FederatedTypeconfigs"
kubefedctl enable namespaces --kubefed-namespace=${NAMESPACE}
kubefedctl enable configmaps --kubefed-namespace=${NAMESPACE}
echo "Creating a FederatedConfigMap resource"

cat <<EOF | kubectl --namespace=${NAMESPACE} apply -f -
apiVersion: types.kubefed.k8s.io/v1beta1
kind: FederatedConfigMap
metadata:
  name: test-configmap
  namespace: ${NAMESPACE}
spec:
  template:
    data:
      key: value
  placement:
    clusters:
    - name: cluster1
EOF

# check for a FederatedConfigMap name
retries=100
until [[ $retries == 0 || $CONFIGMAP =~ "test-configmap" ]]; do
  CONFIGMAP=$(kubectl get configmap -n ${NAMESPACE} -o jsonpath='{.items[*].metadata.name}' 2>/dev/null)
  if [[ $CONFIGMAP != *"test-configmap"* ]]; then
      echo "Waiting for test-configmap to appear"
      sleep 1
      retries=$((retries - 1))
  fi
done

 if [ $retries == 0 ]; then
    echo "Failed to retrieve test-configmap resource"
    exit 1
 fi

 echo "The configmap resource is federated successfully"

elif test X"$SCOPE" = XCluster; then
echo "Enable FederatedTypeconfigs"
kubefedctl enable namespaces --kubefed-namespace=${NAMESPACE}
sleep 5
kubefedctl enable storageclass --kubefed-namespace=${NAMESPACE}
echo "Creating a FederatedStorageClass resource"
cat <<EOF | kubectl apply -f -
apiVersion: types.kubefed.k8s.io/v1beta1                                                                                                                                                      
kind: FederatedStorageClass                                                                                                                                                                    
metadata:                                                                                                                                                                                      
 name: test-storageclass                                                                                                                                                                      
spec:                                                                                                                                                                                          
 template:                                                                                                                                                                                    
   provisioner: local                                                                                                                                                                        
   reclaimPolicy: Retain                                                                                                                                                                      
 placement:                                                                                                                                                                                  
   clusters:                                                                                                                                                                                  
   - name: cluster1
EOF
# check for a FederatedStorageClass name
retries=100
until [[ $retries == 0 || $STORAGECLASS =~ "test-storageclass" ]]; do
  STORAGECLASS=$(kubectl get storageclass -o jsonpath='{.items[*].metadata.name}' 2>/dev/null)
  if [[ $STORAGECLASS != *"test-storageclass"* ]]; then
      echo "Waiting for test-storageclass to appear"
      sleep 1
      retries=$((retries - 1))
  fi
done

 if [ $retries == 0 ]; then
    echo "Failed to retrieve test-storageclass resource"
    exit 1
 fi

 echo "The storageclass resource is federated successfully"
else
   echo "Please enter the valid scope"
   exit 1
fi
}


echo "==========Setting up the infrastructure for kubefed-operator============="
setup-infrastructure

echo "==========Enabling resources=============="
enable-resources

echo "==========Teardown the infrastructure======"
./scripts/teardown.sh -n ${NAMESPACE} -d ${LOCATION}

echo "The smoke testing is completed successfully for the ${SCOPE}-scoped deployment"
