/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	applicationv1alpha1 "kubesphere.io/api/application/v1alpha1"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubesphere.io/kubesphere/pkg/client/listers/application/v1alpha1"
)

// OperatorApplicationInformer provides access to a shared informer and lister for
// OperatorApplications.
type OperatorApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OperatorApplicationLister
}

type operatorApplicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOperatorApplicationInformer constructs a new informer for OperatorApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOperatorApplicationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOperatorApplicationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOperatorApplicationInformer constructs a new informer for OperatorApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOperatorApplicationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationV1alpha1().OperatorApplications().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationV1alpha1().OperatorApplications().Watch(context.TODO(), options)
			},
		},
		&applicationv1alpha1.OperatorApplication{},
		resyncPeriod,
		indexers,
	)
}

func (f *operatorApplicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOperatorApplicationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *operatorApplicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&applicationv1alpha1.OperatorApplication{}, f.defaultInformer)
}

func (f *operatorApplicationInformer) Lister() v1alpha1.OperatorApplicationLister {
	return v1alpha1.NewOperatorApplicationLister(f.Informer().GetIndexer())
}
