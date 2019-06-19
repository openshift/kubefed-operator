package controller

import "github.com/openshift/kubefed-operator/pkg/controller/kubefed"

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kubefed.Add)
}
