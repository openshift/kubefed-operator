package common

import (
	"os"
	"strings"

	mf "github.com/jcrossley3/manifestival"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ResourceImageReplace is a transform method that replace the deployment resource's image location
func ResourceImageReplace(ns, name string) mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		reqLogger := log.WithValues("Instance.Namespace", ns, "Instance.Name", name)
		image := os.Getenv("IMAGE")
		if len(image) == 0 {
			reqLogger.Info("Failed to get the image")
		} else {
			switch strings.ToLower(u.GetKind()) {
			case "deployment":
				var err error
				if containers, ok, err := unstructured.NestedSlice(u.Object,
					"spec", "template", "spec", "containers"); ok {
					for _, container := range containers {
						imageErr := unstructured.SetNestedField(container.(map[string]interface{}), image, "image")
						if imageErr != nil {
							reqLogger.Info("Failed to update the image")
						}
					}
					err = unstructured.SetNestedSlice(u.Object, containers, "spec", "template", "spec", "containers")
					if err != nil {
						reqLogger.Info("Failed to update the container array")
					}
				} else {
					reqLogger.Info("Failed to get the containers slice")
				}
				return err
			}
		}
		return nil
	}
}
