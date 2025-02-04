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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscoordinationv1beta1 "k8s.io/client-go/applyconfigurations/coordination/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// LeasesGetter has a method to return a LeaseInterface.
// A group's client should implement this interface.
type LeasesGetter interface {
	Leases(namespace string) LeaseInterface
}

// LeaseInterface has methods to work with Lease resources.
type LeaseInterface interface {
	Create(ctx context.Context, lease *coordinationv1beta1.Lease, opts v1.CreateOptions) (*coordinationv1beta1.Lease, error)
	Update(ctx context.Context, lease *coordinationv1beta1.Lease, opts v1.UpdateOptions) (*coordinationv1beta1.Lease, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*coordinationv1beta1.Lease, error)
	List(ctx context.Context, opts v1.ListOptions) (*coordinationv1beta1.LeaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *coordinationv1beta1.Lease, err error)
	Apply(ctx context.Context, lease *applyconfigurationscoordinationv1beta1.LeaseApplyConfiguration, opts v1.ApplyOptions) (result *coordinationv1beta1.Lease, err error)
	LeaseExpansion
}

// leases implements LeaseInterface
type leases struct {
	*gentype.ClientWithListAndApply[*coordinationv1beta1.Lease, *coordinationv1beta1.LeaseList, *applyconfigurationscoordinationv1beta1.LeaseApplyConfiguration]
}

// newLeases returns a Leases
func newLeases(c *CoordinationV1beta1Client, namespace string) *leases {
	return &leases{
		gentype.NewClientWithListAndApply[*coordinationv1beta1.Lease, *coordinationv1beta1.LeaseList, *applyconfigurationscoordinationv1beta1.LeaseApplyConfiguration](
			"leases",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *coordinationv1beta1.Lease { return &coordinationv1beta1.Lease{} },
			func() *coordinationv1beta1.LeaseList { return &coordinationv1beta1.LeaseList{} },
			gentype.PrefersProtobuf[*coordinationv1beta1.Lease](),
		),
	}
}
