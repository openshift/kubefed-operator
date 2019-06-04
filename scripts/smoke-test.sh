#!/bin/bash

NAMESPACE="${NAMESPACE:-federation-test}"
LOCATION="${LOCATION:-local}"
VERSION="${VERSION:-0.0.10}"

function setup-infrastructure () {

  ./scripts/create-cluster.sh
  
  ./scripts/install-kubefed.sh -n ${NAMESPACE} -d ${LOCATION} &

  sleep 45 

  # ./scripts/download-binaries.sh
  
}

function enable-resources () {

echo "Performing join operation using kubefedctl"
kubefedctl join cluster1 --federation-namespace=${NAMESPACE} --host-cluster-context=cluster1 --host-cluster-name=cluster1 --cluster-context=cluster1 --add-to-registry

echo "Enable FederatedTypeconfigs"
kubefedctl enable namespaces --federation-namespace=${NAMESPACE}

kubefedctl enable configmaps --federation-namespace=${NAMESPACE}


cat <<EOF | kubectl --namespace=federation-test apply -f -
apiVersion: types.federation.k8s.io/v1alpha1
kind: FederatedConfigMap
metadata:
  name: test-configmap
  namespace: federation-test
spec:
  template:
    data:
      key: value
  placement:
    clusterNames:
    - cluster1
EOF

sleep 40

# check for test-configmap name
if kubectl get configmap -n federation-test -o jsonpath="{.items[1].metadata.name}" | grep -q "test-configmap" ; then
   echo "Configmap resource is federated successfully"
else
   exit 1
fi

}


echo "==========Setting up the infrastructure for kubefed operator============="
setup-infrastructure

echo "==========Enabling resources=============="
enable-resources

echo "==========Teardown the infrastructure======"
./scripts/teardown.sh

echo "Smoke testing is completed successfully"
