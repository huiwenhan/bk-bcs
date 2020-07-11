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

package fake

import (
	clb_v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/apis/clb/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClbIngresses implements ClbIngressInterface
type FakeClbIngresses struct {
	Fake *FakeClbV1
	ns   string
}

var clbingressesResource = schema.GroupVersionResource{Group: "clb.bmsf.tencent.com", Version: "v1", Resource: "clbingresses"}

var clbingressesKind = schema.GroupVersionKind{Group: "clb.bmsf.tencent.com", Version: "v1", Kind: "ClbIngress"}

// Get takes name of the clbIngress, and returns the corresponding clbIngress object, and an error if there is any.
func (c *FakeClbIngresses) Get(name string, options v1.GetOptions) (result *clb_v1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clbingressesResource, c.ns, name), &clb_v1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clb_v1.ClbIngress), err
}

// List takes label and field selectors, and returns the list of ClbIngresses that match those selectors.
func (c *FakeClbIngresses) List(opts v1.ListOptions) (result *clb_v1.ClbIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clbingressesResource, clbingressesKind, c.ns, opts), &clb_v1.ClbIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clb_v1.ClbIngressList{ListMeta: obj.(*clb_v1.ClbIngressList).ListMeta}
	for _, item := range obj.(*clb_v1.ClbIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clbIngresses.
func (c *FakeClbIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clbingressesResource, c.ns, opts))

}

// Create takes the representation of a clbIngress and creates it.  Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *FakeClbIngresses) Create(clbIngress *clb_v1.ClbIngress) (result *clb_v1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clbingressesResource, c.ns, clbIngress), &clb_v1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clb_v1.ClbIngress), err
}

// Update takes the representation of a clbIngress and updates it. Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *FakeClbIngresses) Update(clbIngress *clb_v1.ClbIngress) (result *clb_v1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clbingressesResource, c.ns, clbIngress), &clb_v1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clb_v1.ClbIngress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClbIngresses) UpdateStatus(clbIngress *clb_v1.ClbIngress) (*clb_v1.ClbIngress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clbingressesResource, "status", c.ns, clbIngress), &clb_v1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clb_v1.ClbIngress), err
}

// Delete takes name of the clbIngress and deletes it. Returns an error if one occurs.
func (c *FakeClbIngresses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clbingressesResource, c.ns, name), &clb_v1.ClbIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClbIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clbingressesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &clb_v1.ClbIngressList{})
	return err
}

// Patch applies the patch and returns the patched clbIngress.
func (c *FakeClbIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *clb_v1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clbingressesResource, c.ns, name, data, subresources...), &clb_v1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clb_v1.ClbIngress), err
}
