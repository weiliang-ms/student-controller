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

package fake

import (
	wlcontrollerv1 "student-controller/pkg/apis/wlcontroller/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStudents implements StudentInterface
type FakeStudents struct {
	Fake *FakeWlcontrollerV1
	ns   string
}

var studentsResource = schema.GroupVersionResource{Group: "wlcontroller", Version: "v1", Resource: "students"}

var studentsKind = schema.GroupVersionKind{Group: "wlcontroller", Version: "v1", Kind: "Student"}

// Get takes name of the student, and returns the corresponding student object, and an error if there is any.
func (c *FakeStudents) Get(name string, options v1.GetOptions) (result *wlcontrollerv1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(studentsResource, c.ns, name), &wlcontrollerv1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wlcontrollerv1.Student), err
}

// List takes label and field selectors, and returns the list of Students that match those selectors.
func (c *FakeStudents) List(opts v1.ListOptions) (result *wlcontrollerv1.StudentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(studentsResource, studentsKind, c.ns, opts), &wlcontrollerv1.StudentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &wlcontrollerv1.StudentList{ListMeta: obj.(*wlcontrollerv1.StudentList).ListMeta}
	for _, item := range obj.(*wlcontrollerv1.StudentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested students.
func (c *FakeStudents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(studentsResource, c.ns, opts))

}

// Create takes the representation of a student and creates it.  Returns the server's representation of the student, and an error, if there is any.
func (c *FakeStudents) Create(student *wlcontrollerv1.Student) (result *wlcontrollerv1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(studentsResource, c.ns, student), &wlcontrollerv1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wlcontrollerv1.Student), err
}

// Update takes the representation of a student and updates it. Returns the server's representation of the student, and an error, if there is any.
func (c *FakeStudents) Update(student *wlcontrollerv1.Student) (result *wlcontrollerv1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(studentsResource, c.ns, student), &wlcontrollerv1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wlcontrollerv1.Student), err
}

// Delete takes name of the student and deletes it. Returns an error if one occurs.
func (c *FakeStudents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(studentsResource, c.ns, name), &wlcontrollerv1.Student{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStudents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(studentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &wlcontrollerv1.StudentList{})
	return err
}

// Patch applies the patch and returns the patched student.
func (c *FakeStudents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *wlcontrollerv1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(studentsResource, c.ns, name, pt, data, subresources...), &wlcontrollerv1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wlcontrollerv1.Student), err
}
