package crd

import (
	"context"
	"encoding/json"
	"fmt"

	autodeploy_v1 "github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/apis/autodeploy/v1"
	autodeployscheme "github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/client/clientset/versioned/scheme"
	api_x_v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
)

var AutoDeployUnitCRD = RegistCRD{
	CustomResourceDefinition: api_x_v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{Kind: "CustomResourceDefinition", APIVersion: "apiextensions.k8s.io/v1"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "autodeployunits." + autodeploy_v1.SchemeGroupVersion.Group,
		},
		Spec: api_x_v1.CustomResourceDefinitionSpec{
			Group: autodeploy_v1.SchemeGroupVersion.Group,
			// Version: autodeploy_v1.SchemeGroupVersion.Version,
			Versions: []api_x_v1.CustomResourceDefinitionVersion{{
				Name:    autodeploy_v1.SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
				Subresources: &api_x_v1.CustomResourceSubresources{
					Status: &api_x_v1.CustomResourceSubresourceStatus{},
				},
				Schema: &api_x_v1.CustomResourceValidation{
					OpenAPIV3Schema: &api_x_v1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]api_x_v1.JSONSchemaProps{
							"spec": {
								Type: "object",
								Properties: map[string]api_x_v1.JSONSchemaProps{
									"image": {
										Type: "string",
										// Nullable: true,
									},
									"cmd": {
										Type: "string",
										// Nullable: true,
									},
									"url": {
										Type: "string",
										// Nullable: true,
									},
									"env": {
										Type: "array",
										Items: &api_x_v1.JSONSchemaPropsOrArray{
											Schema: &api_x_v1.JSONSchemaProps{
												Type: "object",
												Properties: map[string]api_x_v1.JSONSchemaProps{
													"name": {
														Type:        "string",
														Description: "name",
													},
													"value": {
														Type:        "string",
														Description: "value",
													},
												},
											},
										},
									},
									"sync": {
										Type: "boolean",
										// Nullable: true,
									},
								},
							},
						},
					},
				},
			}},
			Names: api_x_v1.CustomResourceDefinitionNames{
				Plural:     "autodeployunits",
				Singular:   "autodeployunit",
				ShortNames: []string{"adunit"},
				Kind:       "AutoDeployUnit",
			},
			Scope: api_x_v1.NamespaceScoped,
		},
		Status: api_x_v1.CustomResourceDefinitionStatus{},
	},
}

type RegistCRD struct {
	api_x_v1.CustomResourceDefinition
}

func (s *RegistCRD) Regist(k8sXClient *apiextensionsclient.Clientset) error {
	crd := &AutoDeployUnitCRD.CustomResourceDefinition

	bs, _ := json.Marshal(crd)

	_, err := k8sXClient.ApiextensionsV1().CustomResourceDefinitions().Patch(context.TODO(), crd.ObjectMeta.Name, types.ApplyPatchType, bs, metav1.PatchOptions{
		FieldManager: "crd-init",
	})
	if err != nil {
		fmt.Println(err)
	}
	return autodeployscheme.AddToScheme(scheme.Scheme)
}
