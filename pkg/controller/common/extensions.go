/*
Copyright 2019 The kubefed-operator Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package common

import (
	mf "github.com/jcrossley3/manifestival"
	kubefedv1alpha1 "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("common")

type Platforms []func(client.Client, *runtime.Scheme) (*Extension, error)
type Extender func(*kubefedv1alpha1.KubeFedWebHook) error
type Transclosure func(*kubefedv1alpha1.KubeFedWebHook) mf.Transformer
type PreCondition func(*kubefedv1alpha1.KubeFedWebHook) error
type Extensions []Extension
type Extension struct {
	Transformers  []mf.Transformer
	PreInstalls   []Extender
	PostInstalls  []Extender
	Transclosures []Transclosure
	PreConditions []PreCondition
}

func (platforms Platforms) Extend(c client.Client, scheme *runtime.Scheme) (result Extensions, err error) {
	for _, fn := range platforms {
		ext, err := fn(c, scheme)
		if err != nil {
			return result, err
		}
		if ext != nil {
			result = append(result, *ext)
		}
	}
	return
}

func (exts Extensions) Transform(instance *kubefedv1alpha1.KubeFedWebHook) []mf.Transformer {
	result := []mf.Transformer{
		mf.InjectOwner(instance),
		mf.InjectNamespace(instance.GetNamespace()),
	}
	for _, extension := range exts {
		for _, transclosure := range extension.Transclosures {
			result = append(result, transclosure(instance))
		}
		result = append(result, extension.Transformers...)
	}
	return result
}

func (exts Extensions) PreInstall(instance *kubefedv1alpha1.KubeFedWebHook) error {
	for _, extension := range exts {
		for _, f := range extension.PreInstalls {
			if err := f(instance); err != nil {
				return err
			}
		}
	}
	return nil
}

func (exts Extensions) PostInstall(instance *kubefedv1alpha1.KubeFedWebHook) error {
	for _, extension := range exts {
		for _, f := range extension.PostInstalls {
			if err := f(instance); err != nil {
				return err
			}
		}
	}
	return nil
}

func (exts Extensions) PreConditions(instance *kubefedv1alpha1.KubeFedWebHook) error {
	for _, extension := range exts {
		for _, f := range extension.PreConditions {
			if err := f(instance); err != nil {
				return err
			}
		}
	}
	return nil
}
