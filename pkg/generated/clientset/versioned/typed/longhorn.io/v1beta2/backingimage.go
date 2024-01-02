/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta2

import (
	"context"
	"time"

	scheme "github.com/harvester/node-disk-manager/pkg/generated/clientset/versioned/scheme"
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackingImagesGetter has a method to return a BackingImageInterface.
// A group's client should implement this interface.
type BackingImagesGetter interface {
	BackingImages(namespace string) BackingImageInterface
}

// BackingImageInterface has methods to work with BackingImage resources.
type BackingImageInterface interface {
	Create(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.CreateOptions) (*v1beta2.BackingImage, error)
	Update(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.UpdateOptions) (*v1beta2.BackingImage, error)
	UpdateStatus(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.UpdateOptions) (*v1beta2.BackingImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.BackingImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.BackingImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.BackingImage, err error)
	BackingImageExpansion
}

// backingImages implements BackingImageInterface
type backingImages struct {
	client rest.Interface
	ns     string
}

// newBackingImages returns a BackingImages
func newBackingImages(c *LonghornV1beta2Client, namespace string) *backingImages {
	return &backingImages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backingImage, and returns the corresponding backingImage object, and an error if there is any.
func (c *backingImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.BackingImage, err error) {
	result = &v1beta2.BackingImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackingImages that match those selectors.
func (c *backingImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.BackingImageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.BackingImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backingImages.
func (c *backingImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backingImage and creates it.  Returns the server's representation of the backingImage, and an error, if there is any.
func (c *backingImages) Create(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.CreateOptions) (result *v1beta2.BackingImage, err error) {
	result = &v1beta2.BackingImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backingImage and updates it. Returns the server's representation of the backingImage, and an error, if there is any.
func (c *backingImages) Update(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.UpdateOptions) (result *v1beta2.BackingImage, err error) {
	result = &v1beta2.BackingImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimages").
		Name(backingImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backingImages) UpdateStatus(ctx context.Context, backingImage *v1beta2.BackingImage, opts v1.UpdateOptions) (result *v1beta2.BackingImage, err error) {
	result = &v1beta2.BackingImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimages").
		Name(backingImage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backingImage and deletes it. Returns an error if one occurs.
func (c *backingImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backingImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backingImage.
func (c *backingImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.BackingImage, err error) {
	result = &v1beta2.BackingImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
