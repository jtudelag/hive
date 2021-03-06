// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset-generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSZonesGetter has a method to return a DNSZoneInterface.
// A group's client should implement this interface.
type DNSZonesGetter interface {
	DNSZones(namespace string) DNSZoneInterface
}

// DNSZoneInterface has methods to work with DNSZone resources.
type DNSZoneInterface interface {
	Create(*v1.DNSZone) (*v1.DNSZone, error)
	Update(*v1.DNSZone) (*v1.DNSZone, error)
	UpdateStatus(*v1.DNSZone) (*v1.DNSZone, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.DNSZone, error)
	List(opts metav1.ListOptions) (*v1.DNSZoneList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DNSZone, err error)
	DNSZoneExpansion
}

// dNSZones implements DNSZoneInterface
type dNSZones struct {
	client rest.Interface
	ns     string
}

// newDNSZones returns a DNSZones
func newDNSZones(c *HiveV1Client, namespace string) *dNSZones {
	return &dNSZones{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSZone, and returns the corresponding dNSZone object, and an error if there is any.
func (c *dNSZones) Get(name string, options metav1.GetOptions) (result *v1.DNSZone, err error) {
	result = &v1.DNSZone{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnszones").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSZones that match those selectors.
func (c *dNSZones) List(opts metav1.ListOptions) (result *v1.DNSZoneList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DNSZoneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnszones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSZones.
func (c *dNSZones) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnszones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dNSZone and creates it.  Returns the server's representation of the dNSZone, and an error, if there is any.
func (c *dNSZones) Create(dNSZone *v1.DNSZone) (result *v1.DNSZone, err error) {
	result = &v1.DNSZone{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnszones").
		Body(dNSZone).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dNSZone and updates it. Returns the server's representation of the dNSZone, and an error, if there is any.
func (c *dNSZones) Update(dNSZone *v1.DNSZone) (result *v1.DNSZone, err error) {
	result = &v1.DNSZone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnszones").
		Name(dNSZone.Name).
		Body(dNSZone).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dNSZones) UpdateStatus(dNSZone *v1.DNSZone) (result *v1.DNSZone, err error) {
	result = &v1.DNSZone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnszones").
		Name(dNSZone.Name).
		SubResource("status").
		Body(dNSZone).
		Do().
		Into(result)
	return
}

// Delete takes name of the dNSZone and deletes it. Returns an error if one occurs.
func (c *dNSZones) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnszones").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSZones) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnszones").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dNSZone.
func (c *dNSZones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DNSZone, err error) {
	result = &v1.DNSZone{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnszones").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
