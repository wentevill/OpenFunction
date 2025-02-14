/*
Copyright 2021.

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

package fake

import (
	"context"

	v1alpha1 "github.com/openfunction/apis/events/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterEventBuses implements ClusterEventBusInterface
type FakeClusterEventBuses struct {
	Fake *FakeEventsV1alpha1
	ns   string
}

var clustereventbusesResource = schema.GroupVersionResource{Group: "events.openfunction.io", Version: "v1alpha1", Resource: "clustereventbuses"}

var clustereventbusesKind = schema.GroupVersionKind{Group: "events.openfunction.io", Version: "v1alpha1", Kind: "ClusterEventBus"}

// Get takes name of the clusterEventBus, and returns the corresponding clusterEventBus object, and an error if there is any.
func (c *FakeClusterEventBuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterEventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustereventbusesResource, c.ns, name), &v1alpha1.ClusterEventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterEventBus), err
}

// List takes label and field selectors, and returns the list of ClusterEventBuses that match those selectors.
func (c *FakeClusterEventBuses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterEventBusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustereventbusesResource, clustereventbusesKind, c.ns, opts), &v1alpha1.ClusterEventBusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterEventBusList{ListMeta: obj.(*v1alpha1.ClusterEventBusList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterEventBusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterEventBuses.
func (c *FakeClusterEventBuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustereventbusesResource, c.ns, opts))

}

// Create takes the representation of a clusterEventBus and creates it.  Returns the server's representation of the clusterEventBus, and an error, if there is any.
func (c *FakeClusterEventBuses) Create(ctx context.Context, clusterEventBus *v1alpha1.ClusterEventBus, opts v1.CreateOptions) (result *v1alpha1.ClusterEventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustereventbusesResource, c.ns, clusterEventBus), &v1alpha1.ClusterEventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterEventBus), err
}

// Update takes the representation of a clusterEventBus and updates it. Returns the server's representation of the clusterEventBus, and an error, if there is any.
func (c *FakeClusterEventBuses) Update(ctx context.Context, clusterEventBus *v1alpha1.ClusterEventBus, opts v1.UpdateOptions) (result *v1alpha1.ClusterEventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustereventbusesResource, c.ns, clusterEventBus), &v1alpha1.ClusterEventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterEventBus), err
}

// Delete takes name of the clusterEventBus and deletes it. Returns an error if one occurs.
func (c *FakeClusterEventBuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustereventbusesResource, c.ns, name), &v1alpha1.ClusterEventBus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterEventBuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustereventbusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterEventBusList{})
	return err
}

// Patch applies the patch and returns the patched clusterEventBus.
func (c *FakeClusterEventBuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterEventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustereventbusesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterEventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterEventBus), err
}
