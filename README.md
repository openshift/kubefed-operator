[![Go Report Card](https://goreportcard.com/badge/github.com/openshift/kubefed-operator)](https://goreportcard.com/report/github.com/openshift/kubefed-operator)
[![LICENSE](https://img.shields.io/badge/license-apache2.0-green.svg)](https://github.com/openshift/kubefed-operator/blob/master/LICENSE)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [kubefed-operator](#kubefed-operator)
  - [Deploying and testing](#deploying-and-testing)
    - [Prerequisites](#prerequisites)
    - [Using `operator-sdk up local`](#using-operator-sdk-up-local)
    - [Smoke Testing](#smoke-testing)
      - [Usage:](#usage)
      - [Example: for in-cluster deployment](#example-for-in-cluster-deployment)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# kubefed-operator

This repository contains code for an operator for deploying [KubeFed](https://github.com/kubernetes-sigs/kubefed). It is planned to replace the [federation-v2-operator repo](https://github.com/openshift/federation-v2-operator).

## Deploying and testing

This work-in-progress section describes how people _developing_ this operator can deploy it.

### Prerequisites
- [Kind](https://github.com/kubernetes-sigs/kind) tool
-  The latest version of [kubefedctl](https://github.com/kubernetes-sigs/kubefed/blob/master/docs/userguide.md#kubefedctl-cli) binary installed in your $PATH 
-  If you’re planning to test against the openshift cluster then ensure that it’s already up and running.
-  [Operator-SDK](https://github.com/operator-framework/operator-sdk) CLI tool

### Using `operator-sdk up local`

The operator SDK provides a way to run your operator locally outside a cluster. This allows you to easily iterate on changes without having to push an image.

All you need to do is run the following command from the root directory of this project:

```bash
# if you already have a kubernetes cluster
$ kubectl create ns kubefed-test
# or use create-cluster.sh script to create one
$ ./scripts/create-cluster.sh
$ kubectl apply -f deploy/crds/operator_v1alpha1_kubefed_cr.yaml
$ operator-sdk up local --namespace=kubefed-test
```

This will run the operator configured to watch the `kubefed-test` namespace.

```bash
{"level":"info","ts":1560806650.3372273,"logger":"kubebuilder.controller","msg":"Starting workers","controller":"kubefed-controller","worker count":1}
```
Once you get the above message, you can create a `KubeFed` in the `kubefed-test` namespace to drive the installation in that namespace:

```
$ kubectl create -f deploy/crds/operator_v1alpha1_kubefed_cr.yaml -n kubefed-test
```
When the operator is ready, you will see the following message:

```bash
{"level":"info","ts":1560806722.3100324,"logger":"controller_kubefed","msg":"Finished reconciling kubefed","Request.Namespace":"kubefed-test","Request.Name":"kubefed-resource"}
```
Altenatively, you can also install the operator locally by using the following script:
```bash
$ ./scripts/install-kubefed.sh -n kubefed-test -d local -s Namespaced
```
### Smoke Testing
The smoke-test script is designed to peform smoke testing on the kubefed-operator. It's handling the following scenarios:
*  A local deployment (Location: local)
*  in-cluster deployment using Kind (Location: cluster)
* Deployment via OLM on Minikube (Location: olm-kube)
* Deployment via OLM on Openshift Cluster (Location: olm-openshift)

 

#### Usage:
```bash
$ ./scripts/smoke-test.sh  -n ${NAMESPACE} -d ${LOCATION} -i ${IMAGE_NAME} -s ${SCOPE} -o ${OPERATOR_VERSION}
```
<br>

|Flag | Description | Default value |
| -------- | -------- | -------- |
|n| **Namespace** where you will be installing the kubefed-operator|default|
|d| **Location** defines the mode of deployment|local|
|i| Docker **Image** location | quay.io/openshift/kubefed-operator:v0.1.0-rc6|
|s|**Scope** of the kubefed resource| Namespaced|
|o|**Operator Version** the versioned directory from where the CRD definitions will be applied for olm-kube and olm-openshift deployments
 #### Example: for in-cluster deployment 

```bash
$ ./scripts/smoke-test.sh -n kubefed-test -d cluster -s Namespaced
```
