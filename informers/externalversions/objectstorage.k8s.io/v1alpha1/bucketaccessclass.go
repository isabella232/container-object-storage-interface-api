/*
Copyright The Kubernetes Authors.

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

	objectstoragek8siov1alpha1 "github.com/kubernetes-sigs/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
	clientset "github.com/kubernetes-sigs/container-object-storage-interface-api/clientset"
	internalinterfaces "github.com/kubernetes-sigs/container-object-storage-interface-api/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-sigs/container-object-storage-interface-api/listers/objectstorage.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BucketAccessClassInformer provides access to a shared informer and lister for
// BucketAccessClasses.
type BucketAccessClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BucketAccessClassLister
}

type bucketAccessClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBucketAccessClassInformer constructs a new informer for BucketAccessClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBucketAccessClassInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBucketAccessClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBucketAccessClassInformer constructs a new informer for BucketAccessClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBucketAccessClassInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectstorageV1alpha1().BucketAccessClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectstorageV1alpha1().BucketAccessClasses().Watch(context.TODO(), options)
			},
		},
		&objectstoragek8siov1alpha1.BucketAccessClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *bucketAccessClassInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBucketAccessClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bucketAccessClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&objectstoragek8siov1alpha1.BucketAccessClass{}, f.defaultInformer)
}

func (f *bucketAccessClassInformer) Lister() v1alpha1.BucketAccessClassLister {
	return v1alpha1.NewBucketAccessClassLister(f.Informer().GetIndexer())
}
