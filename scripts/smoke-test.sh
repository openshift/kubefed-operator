#!/bin/bash

NAMESPACE="${NAMESPACE:-kubefed-test}"
LOCATION="${LOCATION:-local}"
VERSION="${VERSION:-v0.1.0-rc2}"

function setup-infrastructure () {

  ./scripts/create-cluster.sh
  
  ./scripts/install-kubefed.sh -n ${NAMESPACE} -d ${LOCATION} -i ${IMAGE_NAME} &

  retries=70
  until [[ $retries == 0 || $name == "kubefed" ]]; do
    name=$(kubectl get kubefedconfig -n ${NAMESPACE} -o jsonpath='{.items[0].metadata.name}' 2>/dev/null)
    if [[ $name != "kubefed" ]]; then
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

echo "Performing join operation using kubefedctl"
kubefedctl join cluster1 --kubefed-namespace=${NAMESPACE} --host-cluster-context=cluster1 --host-cluster-name=cluster1 --cluster-context=cluster1

echo "Enable FederatedTypeconfigs"
kubefedctl enable namespaces --kubefed-namespace=${NAMESPACE}

kubefedctl enable configmaps --kubefed-namespace=${NAMESPACE}

echo "Creating test-configmap resource"

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

# check for test-configmap name
retries=70
until [[ $retries == 0 || $name == "test-configmap" ]]; do
  name=$(kubectl get configmap -n ${NAMESPACE} -o jsonpath='{.items[1].metadata.name}' 2>/dev/null)
  if [[ $name != "test-configmap" ]]; then
      echo "Waiting for test-configmap to appear"
      sleep 1
      retries=$((retries - 1))
  fi
done

 if [ $retries == 0 ]; then
    echo "Failed to retrieve test-configmap resource"
    exit 1
 fi

 echo "Configmap resource is federated successfully"

}


echo "==========Setting up the infrastructure for kubefed operator============="
setup-infrastructure

echo "==========Enabling resources=============="
enable-resources

echo "==========Teardown the infrastructure======"
./scripts/teardown.sh

echo "Smoke testing is completed successfully"
