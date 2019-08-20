#!/usr/bin/env bash
# This script handles the creation of a cluster using kind and the
# ability to create and configure an insecure container registry.

set -e

source "$(dirname "${BASH_SOURCE}")/util.sh"
CREATE_INSECURE_REGISTRY="${CREATE_INSECURE_REGISTRY:-}"
CONFIGURE_INSECURE_REGISTRY="${CONFIGURE_INSECURE_REGISTRY:-}"
CONTAINER_REGISTRY_HOST="${CONTAINER_REGISTRY_HOST:-172.17.0.1:5000}"
KIND_TAG="${KIND_TAG:-}"
kubeconfig="${HOME}/.kube/config"

function create-insecure-registry() {
  # Run insecure registry as container
  docker run -d -p 5000:5000 --restart=always --name registry registry:2
}


function configure-insecure-registry-and-reload() {
  local cmd_context="${1}" # context to run command e.g. sudo, docker exec
  local docker_pid="${2}"
  ${cmd_context} "$(insecure-registry-config-cmd)"
  ${cmd_context} "$(reload-docker-daemon-cmd "${docker_pid}")"
}


function reload-docker-daemon-cmd() {
  echo "kill -s SIGHUP ${1}"
}

function create-cluster() {
  local image_arg=""
  if [[ "${KIND_TAG}" ]]; then
    image_arg="--image=kindest/node:${KIND_TAG}"
  fi
  
  kind create cluster --name "cluster1" ${image_arg}
  fixup-cluster

  echo "Waiting for clusters to be ready"
  check-clusters-ready

  kubectl config view --flatten > ${kubeconfig}
  unset KUBECONFIG
}

function fixup-cluster() {
  local kubeconfig_path="$(kind get kubeconfig-path --name cluster1)"
  export KUBECONFIG="${KUBECONFIG:-}:${kubeconfig_path}"

  # Simplify context name
  kubectl config rename-context "kubernetes-admin@cluster1" "cluster1"

  local container_ip_addr=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cluster1-control-plane)
  # Using the container ip allows the use of port 6443 instead of the
  # random port intended to be exposed on localhost.
  sed -i "s/localhost.*$/${container_ip_addr}:6443/" ${kubeconfig_path}

  sed -i "s/kubernetes-admin/kubernetes-cluster1-admin/" ${kubeconfig_path}
}

function check-clusters-ready() {
  
  util::wait-for-condition 'ok' "kubectl --context cluster1 get --raw=/healthz &> /dev/null" 120
}

echo "Creating a cluster"
create-cluster

echo "Complete"
