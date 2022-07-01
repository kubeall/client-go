/*
Copyright 2022 The kubeall.com Authors.

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

	clusterv1alpha1 "github.com/kubeall/api/cluster/v1alpha1"
	versioned "github.com/kubeall/client-go/cluster/clientset/versioned"
	internalinterfaces "github.com/kubeall/client-go/cluster/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeall/client-go/cluster/listers/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeUserInformer provides access to a shared informer and lister for
// KubeUsers.
type KubeUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubeUserLister
}

type kubeUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeUserInformer constructs a new informer for KubeUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeUserInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeUserInformer constructs a new informer for KubeUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().KubeUsers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().KubeUsers().Watch(context.TODO(), options)
			},
		},
		&clusterv1alpha1.KubeUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeUserInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterv1alpha1.KubeUser{}, f.defaultInformer)
}

func (f *kubeUserInformer) Lister() v1alpha1.KubeUserLister {
	return v1alpha1.NewKubeUserLister(f.Informer().GetIndexer())
}
