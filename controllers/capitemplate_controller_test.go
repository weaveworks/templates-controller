package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	capiv1alpha1 "github.com/weaveworks/templates-controller/apis/capi/v1alpha1"
	capiv1alpha2 "github.com/weaveworks/templates-controller/apis/capi/v1alpha2"
	"github.com/weaveworks/templates-controller/apis/core"
	gitopsv1alpha1 "github.com/weaveworks/templates-controller/apis/gitops/v1alpha1"
	gitopsv1alpha2 "github.com/weaveworks/templates-controller/apis/gitops/v1alpha2"
)

func TestCAPITemplateV1Alpha1Conversion(t *testing.T) {
	ctx := context.TODO()

	ct := &capiv1alpha1.CAPITemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-template",
			Namespace: "default",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       capiv1alpha1.Kind,
			APIVersion: "capi.weave.works/v1alpha1",
		},
		Spec: core.TemplateSpecV1{
			Description: "this is test template 1",
			Params: []core.TemplateParam{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
					Required:    false,
				},
			},
		},
	}

	assertNoError(t, testEnv.Create(ctx, ct))
	defer cleanupResource(t, testEnv, ct)

	updated := &capiv1alpha2.CAPITemplate{}
	assertNoError(t, testEnv.Get(ctx, client.ObjectKeyFromObject(ct), updated))

	want := &capiv1alpha2.CAPITemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-template",
			Namespace: "default",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       capiv1alpha2.Kind,
			APIVersion: "capi.weave.works/v1alpha2",
		},
		Spec: core.TemplateSpec{
			Description: "this is test template 1",
			RenderType:  "envsubst",
			Params: []core.TemplateParam{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
					Required:    false,
				},
			},
			ResourceTemplates: []core.ResourceTemplate{{}},
		},
	}

	if diff := cmp.Diff(want, updated, ignoreObjectMeta); diff != "" {
		t.Fatalf("failed to update the template:\n%s\n", diff)
	}
}

func TestGitOpsTemplateV1Alpha1Conversion(t *testing.T) {
	ctx := context.TODO()

	ct := &gitopsv1alpha1.GitOpsTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-template",
			Namespace: "default",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       capiv1alpha1.Kind,
			APIVersion: "capi.weave.works/v1alpha1",
		},
		Spec: core.TemplateSpecV1{
			Description: "this is test template 1",
			Params: []core.TemplateParam{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
					Required:    false,
				},
			},
		},
	}

	assertNoError(t, testEnv.Create(ctx, ct))
	defer cleanupResource(t, testEnv, ct)

	updated := &gitopsv1alpha2.GitOpsTemplate{}
	assertNoError(t, testEnv.Get(ctx, client.ObjectKeyFromObject(ct), updated))

	want := &gitopsv1alpha2.GitOpsTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-template",
			Namespace: "default",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       gitopsv1alpha2.Kind,
			APIVersion: "templates.weave.works/v1alpha2",
		},
		Spec: core.TemplateSpec{
			Description: "this is test template 1",
			RenderType:  "envsubst",
			Params: []core.TemplateParam{
				{
					Name:        "CLUSTER_NAME",
					Description: "This is used for the cluster naming.",
					Required:    false,
				},
			},
			ResourceTemplates: []core.ResourceTemplate{{}},
		},
	}

	if diff := cmp.Diff(want, updated, ignoreObjectMeta); diff != "" {
		t.Fatalf("failed to update the template:\n%s\n", diff)
	}
}

var ignoreObjectMeta = cmpopts.IgnoreFields(metav1.ObjectMeta{}, "ResourceVersion", "UID", "Generation", "CreationTimestamp", "ManagedFields")

func cleanupResource(t *testing.T, cl client.Client, obj client.Object) {
	t.Helper()
	if err := cl.Delete(context.TODO(), obj); err != nil {
		t.Fatal(err)
	}
}

func rawExtension(s string) runtime.RawExtension {
	return runtime.RawExtension{
		Raw: []byte(s),
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
