// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	keycloakv1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	versioned "github.com/keycloak/keycloak-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/keycloak/keycloak-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/client/listers/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KeycloakUserInformer provides access to a shared informer and lister for
// KeycloakUsers.
type KeycloakUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KeycloakUserLister
}

type keycloakUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKeycloakUserInformer constructs a new informer for KeycloakUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKeycloakUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKeycloakUserInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKeycloakUserInformer constructs a new informer for KeycloakUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKeycloakUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakUsers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakUsers(namespace).Watch(context.TODO(), options)
			},
		},
		&keycloakv1alpha1.KeycloakUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *keycloakUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKeycloakUserInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *keycloakUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&keycloakv1alpha1.KeycloakUser{}, f.defaultInformer)
}

func (f *keycloakUserInformer) Lister() v1alpha1.KeycloakUserLister {
	return v1alpha1.NewKeycloakUserLister(f.Informer().GetIndexer())
}
