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

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	bkbcs_v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/informers/internalinterfaces"
	internalclientset "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/internalclientset"
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/lister/bkbcs/v2"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AdmissionWebhookConfigurationInformer provides access to a shared informer and lister for
// AdmissionWebhookConfigurations.
type AdmissionWebhookConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.AdmissionWebhookConfigurationLister
}

type admissionWebhookConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAdmissionWebhookConfigurationInformer constructs a new informer for AdmissionWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAdmissionWebhookConfigurationInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAdmissionWebhookConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAdmissionWebhookConfigurationInformer constructs a new informer for AdmissionWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAdmissionWebhookConfigurationInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV2().AdmissionWebhookConfigurations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV2().AdmissionWebhookConfigurations(namespace).Watch(options)
			},
		},
		&bkbcs_v2.AdmissionWebhookConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *admissionWebhookConfigurationInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAdmissionWebhookConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *admissionWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bkbcs_v2.AdmissionWebhookConfiguration{}, f.defaultInformer)
}

func (f *admissionWebhookConfigurationInformer) Lister() v2.AdmissionWebhookConfigurationLister {
	return v2.NewAdmissionWebhookConfigurationLister(f.Informer().GetIndexer())
}
