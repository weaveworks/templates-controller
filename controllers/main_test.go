package controllers

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/fluxcd/pkg/runtime/testenv"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	capiv1alpha1 "github.com/weaveworks/templates-controller/apis/capi/v1alpha1"
	capiv1alpha2 "github.com/weaveworks/templates-controller/apis/capi/v1alpha2"
)

var (
	testEnv *testenv.Environment
	ctx     = ctrl.SetupSignalHandler()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMain(m *testing.M) {
	utilruntime.Must(capiv1alpha1.AddToScheme(scheme.Scheme))
	utilruntime.Must(capiv1alpha2.AddToScheme(scheme.Scheme))
	testEnv = testenv.New(testenv.WithCRDPath(filepath.Join("..", "config", "crd", "bases")))

	if err := (&CAPITemplateReconciler{
		Client: testEnv,
		Scheme: testEnv.GetScheme(),
	}).SetupWithManager(testEnv); err != nil {
		panic(fmt.Sprintf("Failed to start CAPITemplateReconciler: %v", err))
	}

	// utilruntime.Must((&capiv1alpha1.CAPITemplate{}).SetupWebhookWithManager(testEnv))
	utilruntime.Must((&capiv1alpha2.CAPITemplate{}).SetupWebhookWithManager(testEnv))

	go func() {
		fmt.Println("Starting the test environment")
		if err := testEnv.Start(ctx); err != nil {
			panic(fmt.Sprintf("Failed to start the test environment manager: %v", err))
		}
	}()
	<-testEnv.Manager.Elected()

	code := m.Run()

	fmt.Println("Stopping the test environment")
	if err := testEnv.Stop(); err != nil {
		panic(fmt.Sprintf("Failed to stop the test environment: %v", err))
	}

	os.Exit(code)
}
