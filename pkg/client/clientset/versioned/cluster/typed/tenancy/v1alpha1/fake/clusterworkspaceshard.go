//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	tenancyv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1alpha1"
)

var clusterWorkspaceShardsResource = schema.GroupVersionResource{Group: "tenancy.kcp.dev", Version: "v1alpha1", Resource: "clusterworkspaceshards"}
var clusterWorkspaceShardsKind = schema.GroupVersionKind{Group: "tenancy.kcp.dev", Version: "v1alpha1", Kind: "ClusterWorkspaceShard"}

type clusterWorkspaceShardsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterWorkspaceShardsClusterClient) Cluster(cluster logicalcluster.Name) tenancyv1alpha1client.ClusterWorkspaceShardInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterWorkspaceShardsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ClusterWorkspaceShards that match those selectors across all clusters.
func (c *clusterWorkspaceShardsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.ClusterWorkspaceShardList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterWorkspaceShardsResource, clusterWorkspaceShardsKind, logicalcluster.Wildcard, opts), &tenancyv1alpha1.ClusterWorkspaceShardList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.ClusterWorkspaceShardList{ListMeta: obj.(*tenancyv1alpha1.ClusterWorkspaceShardList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.ClusterWorkspaceShardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterWorkspaceShards across all clusters.
func (c *clusterWorkspaceShardsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterWorkspaceShardsResource, logicalcluster.Wildcard, opts))
}

type clusterWorkspaceShardsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *clusterWorkspaceShardsClient) Create(ctx context.Context, clusterWorkspaceShard *tenancyv1alpha1.ClusterWorkspaceShard, opts metav1.CreateOptions) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterWorkspaceShardsResource, c.Cluster, clusterWorkspaceShard), &tenancyv1alpha1.ClusterWorkspaceShard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), err
}

func (c *clusterWorkspaceShardsClient) Update(ctx context.Context, clusterWorkspaceShard *tenancyv1alpha1.ClusterWorkspaceShard, opts metav1.UpdateOptions) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterWorkspaceShardsResource, c.Cluster, clusterWorkspaceShard), &tenancyv1alpha1.ClusterWorkspaceShard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), err
}

func (c *clusterWorkspaceShardsClient) UpdateStatus(ctx context.Context, clusterWorkspaceShard *tenancyv1alpha1.ClusterWorkspaceShard, opts metav1.UpdateOptions) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterWorkspaceShardsResource, c.Cluster, "status", clusterWorkspaceShard), &tenancyv1alpha1.ClusterWorkspaceShard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), err
}

func (c *clusterWorkspaceShardsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterWorkspaceShardsResource, c.Cluster, name, opts), &tenancyv1alpha1.ClusterWorkspaceShard{})
	return err
}

func (c *clusterWorkspaceShardsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterWorkspaceShardsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &tenancyv1alpha1.ClusterWorkspaceShardList{})
	return err
}

func (c *clusterWorkspaceShardsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterWorkspaceShardsResource, c.Cluster, name), &tenancyv1alpha1.ClusterWorkspaceShard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), err
}

// List takes label and field selectors, and returns the list of ClusterWorkspaceShards that match those selectors.
func (c *clusterWorkspaceShardsClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.ClusterWorkspaceShardList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterWorkspaceShardsResource, clusterWorkspaceShardsKind, c.Cluster, opts), &tenancyv1alpha1.ClusterWorkspaceShardList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.ClusterWorkspaceShardList{ListMeta: obj.(*tenancyv1alpha1.ClusterWorkspaceShardList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.ClusterWorkspaceShardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterWorkspaceShardsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterWorkspaceShardsResource, c.Cluster, opts))
}

func (c *clusterWorkspaceShardsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*tenancyv1alpha1.ClusterWorkspaceShard, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterWorkspaceShardsResource, c.Cluster, name, pt, data, subresources...), &tenancyv1alpha1.ClusterWorkspaceShard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.ClusterWorkspaceShard), err
}