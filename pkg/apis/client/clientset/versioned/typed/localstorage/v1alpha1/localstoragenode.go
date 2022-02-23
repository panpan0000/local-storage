// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/hwameistor/local-storage/pkg/apis/client/clientset/versioned/scheme"
	v1alpha1 "github.com/hwameistor/local-storage/pkg/apis/localstorage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocalStorageNodesGetter has a method to return a LocalStorageNodeInterface.
// A group's client should implement this interface.
type LocalStorageNodesGetter interface {
	LocalStorageNodes() LocalStorageNodeInterface
}

// LocalStorageNodeInterface has methods to work with LocalStorageNode resources.
type LocalStorageNodeInterface interface {
	Create(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.CreateOptions) (*v1alpha1.LocalStorageNode, error)
	Update(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.UpdateOptions) (*v1alpha1.LocalStorageNode, error)
	UpdateStatus(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.UpdateOptions) (*v1alpha1.LocalStorageNode, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LocalStorageNode, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LocalStorageNodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalStorageNode, err error)
	LocalStorageNodeExpansion
}

// localStorageNodes implements LocalStorageNodeInterface
type localStorageNodes struct {
	client rest.Interface
}

// newLocalStorageNodes returns a LocalStorageNodes
func newLocalStorageNodes(c *LocalStorageV1alpha1Client) *localStorageNodes {
	return &localStorageNodes{
		client: c.RESTClient(),
	}
}

// Get takes name of the localStorageNode, and returns the corresponding localStorageNode object, and an error if there is any.
func (c *localStorageNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalStorageNode, err error) {
	result = &v1alpha1.LocalStorageNode{}
	err = c.client.Get().
		Resource("localstoragenodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalStorageNodes that match those selectors.
func (c *localStorageNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalStorageNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocalStorageNodeList{}
	err = c.client.Get().
		Resource("localstoragenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localStorageNodes.
func (c *localStorageNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("localstoragenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a localStorageNode and creates it.  Returns the server's representation of the localStorageNode, and an error, if there is any.
func (c *localStorageNodes) Create(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.CreateOptions) (result *v1alpha1.LocalStorageNode, err error) {
	result = &v1alpha1.LocalStorageNode{}
	err = c.client.Post().
		Resource("localstoragenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localStorageNode).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a localStorageNode and updates it. Returns the server's representation of the localStorageNode, and an error, if there is any.
func (c *localStorageNodes) Update(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.UpdateOptions) (result *v1alpha1.LocalStorageNode, err error) {
	result = &v1alpha1.LocalStorageNode{}
	err = c.client.Put().
		Resource("localstoragenodes").
		Name(localStorageNode.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localStorageNode).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *localStorageNodes) UpdateStatus(ctx context.Context, localStorageNode *v1alpha1.LocalStorageNode, opts v1.UpdateOptions) (result *v1alpha1.LocalStorageNode, err error) {
	result = &v1alpha1.LocalStorageNode{}
	err = c.client.Put().
		Resource("localstoragenodes").
		Name(localStorageNode.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localStorageNode).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the localStorageNode and deletes it. Returns an error if one occurs.
func (c *localStorageNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("localstoragenodes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localStorageNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("localstoragenodes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched localStorageNode.
func (c *localStorageNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalStorageNode, err error) {
	result = &v1alpha1.LocalStorageNode{}
	err = c.client.Patch(pt).
		Resource("localstoragenodes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
