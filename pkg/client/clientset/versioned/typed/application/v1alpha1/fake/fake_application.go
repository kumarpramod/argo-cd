package fake

import (
	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplications implements ApplicationInterface
type FakeApplications struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var applicationsResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}

var applicationsKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "Application"}

// Get takes name of the application, and returns the corresponding application object, and an error if there is any.
func (c *FakeApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationsResource, c.ns, name), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// List takes label and field selectors, and returns the list of Applications that match those selectors.
func (c *FakeApplications) List(opts v1.ListOptions) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationsResource, applicationsKind, c.ns, opts), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationList{ListMeta: obj.(*v1alpha1.ApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applications.
func (c *FakeApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationsResource, c.ns, opts))

}

// Create takes the representation of a application and creates it.  Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Create(application *v1alpha1.Application) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationsResource, c.ns, application), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// Update takes the representation of a application and updates it. Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Update(application *v1alpha1.Application) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationsResource, c.ns, application), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// Delete takes name of the application and deletes it. Returns an error if one occurs.
func (c *FakeApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationsResource, c.ns, name), &v1alpha1.Application{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationList{})
	return err
}

// Patch applies the patch and returns the patched application.
func (c *FakeApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationsResource, c.ns, name, data, subresources...), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}