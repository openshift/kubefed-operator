# kubefed-operator

Prototype operator for [KubeFed](https://github.com/kubernetes-sigs/federation-v2). Planned to eventually replace the [federation-v2-operator repo](https://github.com/openshift/federation-v2-operator).

Note, currently, this operator only supports deploying KubeFed in a namespace-scoped fashion.

## Deploying and testing

This work-in-progress section describes how people _developing_ this operator can deploy it.

### Using `operator-sdk up local`

The operator SDK provides a way to run your operator locally outside a cluster. This allows you to easily iterate on changes without having to push an image.

All you need to do is run the following command from the root directory of this project:

```
$ operator-sdk up local --namespace=kubefed-test
```

This will run the operator configured to watch the `kubefed-test` namespace.

After that step, you can create a `KubeFed` in the `kubefed-test` namespace to drive the installation in that namespace:

```
$ kubectl create -f deploy/crds/operator_v1alpha1_kubefed_cr.yaml -n kubefed-test
```
