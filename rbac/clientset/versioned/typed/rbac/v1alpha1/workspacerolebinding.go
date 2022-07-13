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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubeall/api/rbac/v1alpha1"
	scheme "github.com/kubeall/client-go/rbac/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkspaceRoleBindingsGetter has a method to return a WorkspaceRoleBindingInterface.
// A group's client should implement this interface.
type WorkspaceRoleBindingsGetter interface {
	WorkspaceRoleBindings() WorkspaceRoleBindingInterface
}

// WorkspaceRoleBindingInterface has methods to work with WorkspaceRoleBinding resources.
type WorkspaceRoleBindingInterface interface {
	Create(ctx context.Context, workspaceRoleBinding *v1alpha1.WorkspaceRoleBinding, opts v1.CreateOptions) (*v1alpha1.WorkspaceRoleBinding, error)
	Update(ctx context.Context, workspaceRoleBinding *v1alpha1.WorkspaceRoleBinding, opts v1.UpdateOptions) (*v1alpha1.WorkspaceRoleBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkspaceRoleBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkspaceRoleBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceRoleBinding, err error)
	WorkspaceRoleBindingExpansion
}

// workspaceRoleBindings implements WorkspaceRoleBindingInterface
type workspaceRoleBindings struct {
	client rest.Interface
}

// newWorkspaceRoleBindings returns a WorkspaceRoleBindings
func newWorkspaceRoleBindings(c *RbacV1alpha1Client) *workspaceRoleBindings {
	return &workspaceRoleBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceRoleBinding, and returns the corresponding workspaceRoleBinding object, and an error if there is any.
func (c *workspaceRoleBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceRoleBinding, err error) {
	result = &v1alpha1.WorkspaceRoleBinding{}
	err = c.client.Get().
		Resource("workspacerolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceRoleBindings that match those selectors.
func (c *workspaceRoleBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkspaceRoleBindingList{}
	err = c.client.Get().
		Resource("workspacerolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceRoleBindings.
func (c *workspaceRoleBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspacerolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workspaceRoleBinding and creates it.  Returns the server's representation of the workspaceRoleBinding, and an error, if there is any.
func (c *workspaceRoleBindings) Create(ctx context.Context, workspaceRoleBinding *v1alpha1.WorkspaceRoleBinding, opts v1.CreateOptions) (result *v1alpha1.WorkspaceRoleBinding, err error) {
	result = &v1alpha1.WorkspaceRoleBinding{}
	err = c.client.Post().
		Resource("workspacerolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workspaceRoleBinding and updates it. Returns the server's representation of the workspaceRoleBinding, and an error, if there is any.
func (c *workspaceRoleBindings) Update(ctx context.Context, workspaceRoleBinding *v1alpha1.WorkspaceRoleBinding, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceRoleBinding, err error) {
	result = &v1alpha1.WorkspaceRoleBinding{}
	err = c.client.Put().
		Resource("workspacerolebindings").
		Name(workspaceRoleBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workspaceRoleBinding and deletes it. Returns an error if one occurs.
func (c *workspaceRoleBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspacerolebindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceRoleBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspacerolebindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workspaceRoleBinding.
func (c *workspaceRoleBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceRoleBinding, err error) {
	result = &v1alpha1.WorkspaceRoleBinding{}
	err = c.client.Patch(pt).
		Resource("workspacerolebindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
