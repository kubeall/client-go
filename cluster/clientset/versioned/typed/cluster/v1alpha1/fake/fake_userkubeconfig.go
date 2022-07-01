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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kubeall/api/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserKubeConfigs implements UserKubeConfigInterface
type FakeUserKubeConfigs struct {
	Fake *FakeClusterV1alpha1
}

var userkubeconfigsResource = schema.GroupVersionResource{Group: "cluster.kubeall.com", Version: "v1alpha1", Resource: "userkubeconfigs"}

var userkubeconfigsKind = schema.GroupVersionKind{Group: "cluster.kubeall.com", Version: "v1alpha1", Kind: "UserKubeConfig"}

// Get takes name of the userKubeConfig, and returns the corresponding userKubeConfig object, and an error if there is any.
func (c *FakeUserKubeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserKubeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(userkubeconfigsResource, name), &v1alpha1.UserKubeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserKubeConfig), err
}

// List takes label and field selectors, and returns the list of UserKubeConfigs that match those selectors.
func (c *FakeUserKubeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserKubeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(userkubeconfigsResource, userkubeconfigsKind, opts), &v1alpha1.UserKubeConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserKubeConfigList{ListMeta: obj.(*v1alpha1.UserKubeConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserKubeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userKubeConfigs.
func (c *FakeUserKubeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(userkubeconfigsResource, opts))
}

// Create takes the representation of a userKubeConfig and creates it.  Returns the server's representation of the userKubeConfig, and an error, if there is any.
func (c *FakeUserKubeConfigs) Create(ctx context.Context, userKubeConfig *v1alpha1.UserKubeConfig, opts v1.CreateOptions) (result *v1alpha1.UserKubeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(userkubeconfigsResource, userKubeConfig), &v1alpha1.UserKubeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserKubeConfig), err
}

// Update takes the representation of a userKubeConfig and updates it. Returns the server's representation of the userKubeConfig, and an error, if there is any.
func (c *FakeUserKubeConfigs) Update(ctx context.Context, userKubeConfig *v1alpha1.UserKubeConfig, opts v1.UpdateOptions) (result *v1alpha1.UserKubeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(userkubeconfigsResource, userKubeConfig), &v1alpha1.UserKubeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserKubeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserKubeConfigs) UpdateStatus(ctx context.Context, userKubeConfig *v1alpha1.UserKubeConfig, opts v1.UpdateOptions) (*v1alpha1.UserKubeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(userkubeconfigsResource, "status", userKubeConfig), &v1alpha1.UserKubeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserKubeConfig), err
}

// Delete takes name of the userKubeConfig and deletes it. Returns an error if one occurs.
func (c *FakeUserKubeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(userkubeconfigsResource, name, opts), &v1alpha1.UserKubeConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserKubeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(userkubeconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserKubeConfigList{})
	return err
}

// Patch applies the patch and returns the patched userKubeConfig.
func (c *FakeUserKubeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserKubeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(userkubeconfigsResource, name, pt, data, subresources...), &v1alpha1.UserKubeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserKubeConfig), err
}
