/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedClusterRolesGetter has a method to return a FederatedClusterRoleInterface.
// A group's client should implement this interface.
type FederatedClusterRolesGetter interface {
	FederatedClusterRoles(namespace string) FederatedClusterRoleInterface
}

// FederatedClusterRoleInterface has methods to work with FederatedClusterRole resources.
type FederatedClusterRoleInterface interface {
	Create(*v1beta1.FederatedClusterRole) (*v1beta1.FederatedClusterRole, error)
	Update(*v1beta1.FederatedClusterRole) (*v1beta1.FederatedClusterRole, error)
	UpdateStatus(*v1beta1.FederatedClusterRole) (*v1beta1.FederatedClusterRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedClusterRole, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedClusterRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRole, err error)
	FederatedClusterRoleExpansion
}

// federatedClusterRoles implements FederatedClusterRoleInterface
type federatedClusterRoles struct {
	client rest.Interface
	ns     string
}

// newFederatedClusterRoles returns a FederatedClusterRoles
func newFederatedClusterRoles(c *TypesV1beta1Client, namespace string) *federatedClusterRoles {
	return &federatedClusterRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedClusterRole, and returns the corresponding federatedClusterRole object, and an error if there is any.
func (c *federatedClusterRoles) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedClusterRole, err error) {
	result = &v1beta1.FederatedClusterRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedClusterRoles that match those selectors.
func (c *federatedClusterRoles) List(opts v1.ListOptions) (result *v1beta1.FederatedClusterRoleList, err error) {
	result = &v1beta1.FederatedClusterRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedClusterRoles.
func (c *federatedClusterRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedClusterRole and creates it.  Returns the server's representation of the federatedClusterRole, and an error, if there is any.
func (c *federatedClusterRoles) Create(federatedClusterRole *v1beta1.FederatedClusterRole) (result *v1beta1.FederatedClusterRole, err error) {
	result = &v1beta1.FederatedClusterRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		Body(federatedClusterRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedClusterRole and updates it. Returns the server's representation of the federatedClusterRole, and an error, if there is any.
func (c *federatedClusterRoles) Update(federatedClusterRole *v1beta1.FederatedClusterRole) (result *v1beta1.FederatedClusterRole, err error) {
	result = &v1beta1.FederatedClusterRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		Name(federatedClusterRole.Name).
		Body(federatedClusterRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedClusterRoles) UpdateStatus(federatedClusterRole *v1beta1.FederatedClusterRole) (result *v1beta1.FederatedClusterRole, err error) {
	result = &v1beta1.FederatedClusterRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		Name(federatedClusterRole.Name).
		SubResource("status").
		Body(federatedClusterRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedClusterRole and deletes it. Returns an error if one occurs.
func (c *federatedClusterRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedClusterRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedclusterroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedClusterRole.
func (c *federatedClusterRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRole, err error) {
	result = &v1beta1.FederatedClusterRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedclusterroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
