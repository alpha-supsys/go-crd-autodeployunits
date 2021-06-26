package crd

import (
	"context"
	"encoding/json"
	"fmt"

	autodeploy_v1 "github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/apis/autodeploy/v1"
	"github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type AutoDeployUnitCRDEtcdClient interface {
	Create(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error)
	Update(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error)
	Apply(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error)
	Delete(name string, opts *metav1.DeleteOptions) (*autodeploy_v1.AutoDeployUnit, error)
	List(opts *metav1.ListOptions) (*autodeploy_v1.AutoDeployUnitList, error)
	Get(name string, opts *metav1.GetOptions) (*autodeploy_v1.AutoDeployUnit, error)
	Namespace(namespace string) AutoDeployUnitCRDEtcdClient
}

type AutoDeployUnitCRDEtcdDefaultClient struct {
	AutoDeployUnitCRDEtcdClient
	restClient rest.Interface
	namespace  string
}

func NewAutoDeployUnitCRDEtcdDefaultClient(restClient rest.Interface) AutoDeployUnitCRDEtcdClient {
	return &AutoDeployUnitCRDEtcdDefaultClient{
		restClient: restClient,
	}
}

// Create todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Create(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error) {
	result := &autodeploy_v1.AutoDeployUnit{}
	req := s.restClient.Post().Namespace(s.namespace).Resource("autodeployunits").Body(obj)
	fmt.Println(req.URL())
	res := req.Do(context.TODO())
	fmt.Println(res.Error())
	err := res.Into(result)
	return result, err
}

// Update todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Update(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error) {
	result := &autodeploy_v1.AutoDeployUnit{}
	namespace := s.namespace
	if len(namespace) == 0 {
		namespace = obj.Namespace
	}
	err := s.restClient.Put().Namespace(namespace).Resource("autodeployunits").Body(obj).Do(context.TODO()).Into(result)
	return result, err
}

// Apply todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Apply(obj *autodeploy_v1.AutoDeployUnit) (*autodeploy_v1.AutoDeployUnit, error) {
	result := &autodeploy_v1.AutoDeployUnit{}
	namespace := s.namespace
	if len(namespace) == 0 {
		namespace = obj.Namespace
	}
	bs, _ := json.Marshal(obj)
	err := s.restClient.Patch(types.ApplyPatchType).Namespace(namespace).Resource("autodeployunits").Name(obj.Name).Body(bs).Param("fieldManager", "apply_test").Do(context.TODO()).Into(result)
	return result, err
}

// Delete todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Delete(name string, opts *metav1.DeleteOptions) (*autodeploy_v1.AutoDeployUnit, error) {
	result := &autodeploy_v1.AutoDeployUnit{}
	err := s.restClient.Delete().Namespace(s.namespace).Resource("autodeployunits").Name(name).VersionedParams(opts, scheme.ParameterCodec).Do(context.TODO()).Into(result)
	return result, err
}

// List todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) List(opts *metav1.ListOptions) (*autodeploy_v1.AutoDeployUnitList, error) {
	result := &autodeploy_v1.AutoDeployUnitList{}
	err := s.restClient.Get().Namespace(s.namespace).Resource("autodeployunits").VersionedParams(opts, scheme.ParameterCodec).Do(context.TODO()).Into(result)
	return result, err
}

// Get todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Get(name string, opts *metav1.GetOptions) (*autodeploy_v1.AutoDeployUnit, error) {
	result := &autodeploy_v1.AutoDeployUnit{}
	err := s.restClient.Get().Namespace(s.namespace).Resource("autodeployunits").Name(name).VersionedParams(opts, scheme.ParameterCodec).Do(context.TODO()).Into(result)
	return result, err
}

// Namespace todo
func (s *AutoDeployUnitCRDEtcdDefaultClient) Namespace(namespace string) AutoDeployUnitCRDEtcdClient {
	return &AutoDeployUnitCRDEtcdDefaultClient{
		namespace:  namespace,
		restClient: s.restClient,
	}
}
