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

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BcsClusterAgentSettingLister helps list BcsClusterAgentSettings.
type BcsClusterAgentSettingLister interface {
	// List lists all BcsClusterAgentSettings in the indexer.
	List(selector labels.Selector) (ret []*v2.BcsClusterAgentSetting, err error)
	// BcsClusterAgentSettings returns an object that can list and get BcsClusterAgentSettings.
	BcsClusterAgentSettings(namespace string) BcsClusterAgentSettingNamespaceLister
	BcsClusterAgentSettingListerExpansion
}

// bcsClusterAgentSettingLister implements the BcsClusterAgentSettingLister interface.
type bcsClusterAgentSettingLister struct {
	indexer cache.Indexer
}

// NewBcsClusterAgentSettingLister returns a new BcsClusterAgentSettingLister.
func NewBcsClusterAgentSettingLister(indexer cache.Indexer) BcsClusterAgentSettingLister {
	return &bcsClusterAgentSettingLister{indexer: indexer}
}

// List lists all BcsClusterAgentSettings in the indexer.
func (s *bcsClusterAgentSettingLister) List(selector labels.Selector) (ret []*v2.BcsClusterAgentSetting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsClusterAgentSetting))
	})
	return ret, err
}

// BcsClusterAgentSettings returns an object that can list and get BcsClusterAgentSettings.
func (s *bcsClusterAgentSettingLister) BcsClusterAgentSettings(namespace string) BcsClusterAgentSettingNamespaceLister {
	return bcsClusterAgentSettingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BcsClusterAgentSettingNamespaceLister helps list and get BcsClusterAgentSettings.
type BcsClusterAgentSettingNamespaceLister interface {
	// List lists all BcsClusterAgentSettings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.BcsClusterAgentSetting, err error)
	// Get retrieves the BcsClusterAgentSetting from the indexer for a given namespace and name.
	Get(name string) (*v2.BcsClusterAgentSetting, error)
	BcsClusterAgentSettingNamespaceListerExpansion
}

// bcsClusterAgentSettingNamespaceLister implements the BcsClusterAgentSettingNamespaceLister
// interface.
type bcsClusterAgentSettingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BcsClusterAgentSettings in the indexer for a given namespace.
func (s bcsClusterAgentSettingNamespaceLister) List(selector labels.Selector) (ret []*v2.BcsClusterAgentSetting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsClusterAgentSetting))
	})
	return ret, err
}

// Get retrieves the BcsClusterAgentSetting from the indexer for a given namespace and name.
func (s bcsClusterAgentSettingNamespaceLister) Get(name string) (*v2.BcsClusterAgentSetting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("bcsclusteragentsetting"), name)
	}
	return obj.(*v2.BcsClusterAgentSetting), nil
}
