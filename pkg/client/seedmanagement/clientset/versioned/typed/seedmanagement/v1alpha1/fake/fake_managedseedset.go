/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	v1alpha1 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedSeedSets implements ManagedSeedSetInterface
type FakeManagedSeedSets struct {
	Fake *FakeSeedmanagementV1alpha1
	ns   string
}

var managedseedsetsResource = schema.GroupVersionResource{Group: "seedmanagement.gardener.cloud", Version: "v1alpha1", Resource: "managedseedsets"}

var managedseedsetsKind = schema.GroupVersionKind{Group: "seedmanagement.gardener.cloud", Version: "v1alpha1", Kind: "ManagedSeedSet"}

// Get takes name of the managedSeedSet, and returns the corresponding managedSeedSet object, and an error if there is any.
func (c *FakeManagedSeedSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedSeedSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedseedsetsResource, c.ns, name), &v1alpha1.ManagedSeedSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSeedSet), err
}

// List takes label and field selectors, and returns the list of ManagedSeedSets that match those selectors.
func (c *FakeManagedSeedSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedSeedSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedseedsetsResource, managedseedsetsKind, c.ns, opts), &v1alpha1.ManagedSeedSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedSeedSetList{ListMeta: obj.(*v1alpha1.ManagedSeedSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedSeedSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedSeedSets.
func (c *FakeManagedSeedSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedseedsetsResource, c.ns, opts))

}

// Create takes the representation of a managedSeedSet and creates it.  Returns the server's representation of the managedSeedSet, and an error, if there is any.
func (c *FakeManagedSeedSets) Create(ctx context.Context, managedSeedSet *v1alpha1.ManagedSeedSet, opts v1.CreateOptions) (result *v1alpha1.ManagedSeedSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedseedsetsResource, c.ns, managedSeedSet), &v1alpha1.ManagedSeedSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSeedSet), err
}

// Update takes the representation of a managedSeedSet and updates it. Returns the server's representation of the managedSeedSet, and an error, if there is any.
func (c *FakeManagedSeedSets) Update(ctx context.Context, managedSeedSet *v1alpha1.ManagedSeedSet, opts v1.UpdateOptions) (result *v1alpha1.ManagedSeedSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedseedsetsResource, c.ns, managedSeedSet), &v1alpha1.ManagedSeedSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSeedSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedSeedSets) UpdateStatus(ctx context.Context, managedSeedSet *v1alpha1.ManagedSeedSet, opts v1.UpdateOptions) (*v1alpha1.ManagedSeedSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedseedsetsResource, "status", c.ns, managedSeedSet), &v1alpha1.ManagedSeedSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSeedSet), err
}

// Delete takes name of the managedSeedSet and deletes it. Returns an error if one occurs.
func (c *FakeManagedSeedSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedseedsetsResource, c.ns, name), &v1alpha1.ManagedSeedSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedSeedSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedseedsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedSeedSetList{})
	return err
}

// Patch applies the patch and returns the patched managedSeedSet.
func (c *FakeManagedSeedSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedSeedSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedseedsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedSeedSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSeedSet), err
}
