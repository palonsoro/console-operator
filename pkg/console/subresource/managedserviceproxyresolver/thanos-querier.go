package managedproxyserviceresolver

import (
	"github.com/openshift/console-operator/pkg/console/assets"
	"github.com/openshift/console-operator/pkg/console/subresource/util"
	proxyv1alpha1 "open-cluster-management.io/cluster-proxy/pkg/apis/proxy/v1alpha1"
)

func DefaultThanosQuerierProxyServiceResolver() *proxyv1alpha1.ManagedProxyServiceResolver {
	return util.ReadManagedProxyServiceResolverOrDie(assets.MustAsset("managedserviceproxyresolvers/thanos-querier.yaml"))
}
