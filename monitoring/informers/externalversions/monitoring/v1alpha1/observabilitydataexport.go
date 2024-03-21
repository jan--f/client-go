// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	monitoringv1alpha1 "github.com/openshift/api/monitoring/v1alpha1"
	versioned "github.com/openshift/client-go/monitoring/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/monitoring/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/client-go/monitoring/listers/monitoring/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ObservabilityDataExportInformer provides access to a shared informer and lister for
// ObservabilityDataExports.
type ObservabilityDataExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ObservabilityDataExportLister
}

type observabilityDataExportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewObservabilityDataExportInformer constructs a new informer for ObservabilityDataExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewObservabilityDataExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredObservabilityDataExportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredObservabilityDataExportInformer constructs a new informer for ObservabilityDataExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredObservabilityDataExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().ObservabilityDataExports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().ObservabilityDataExports(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1alpha1.ObservabilityDataExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *observabilityDataExportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredObservabilityDataExportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *observabilityDataExportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1alpha1.ObservabilityDataExport{}, f.defaultInformer)
}

func (f *observabilityDataExportInformer) Lister() v1alpha1.ObservabilityDataExportLister {
	return v1alpha1.NewObservabilityDataExportLister(f.Informer().GetIndexer())
}
