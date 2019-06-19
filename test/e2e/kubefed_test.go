package e2e

import (
	"testing"
	"time"

	"github.com/openshift/kubefed-operator/pkg/apis"
	kubefedv1alpha1 "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
)

var (
	retryInterval        = time.Second * 25
	timeout              = time.Second * 60
	cleanupRetryInterval = time.Second * 1
	cleanupTimeout       = time.Second * 5
)

func TestKubeFed(t *testing.T) {
	kubefedList := &kubefedv1alpha1.KubeFedList{}

	err := framework.AddToFrameworkScheme(apis.AddToScheme, kubefedList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}

	// run subtests
	t.Run("kubefed-group", func(t *testing.T) {
		t.Run("Cluster", KubeFedCluster)
	})
}

func KubeFedCluster(t *testing.T) {
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()
	err := ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: cleanupTimeout, RetryInterval: cleanupRetryInterval})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	t.Log("Initialized cluster resources")
	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatal(err)
	}
	// get global framework variables
	f := framework.Global
	err = e2eutil.WaitForDeployment(t, f.KubeClient, namespace, "kubefed-controller-manager", 1, retryInterval, timeout)
	if err != nil {
		t.Fatal(err)
	}
}
