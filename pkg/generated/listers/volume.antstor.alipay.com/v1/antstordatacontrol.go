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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "code.alipay.com/dbplatform/node-disk-controller/pkg/api/volume.antstor.alipay.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AntstorDataControlLister helps list AntstorDataControls.
// All objects returned here must be treated as read-only.
type AntstorDataControlLister interface {
	// List lists all AntstorDataControls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AntstorDataControl, err error)
	// AntstorDataControls returns an object that can list and get AntstorDataControls.
	AntstorDataControls(namespace string) AntstorDataControlNamespaceLister
	AntstorDataControlListerExpansion
}

// antstorDataControlLister implements the AntstorDataControlLister interface.
type antstorDataControlLister struct {
	indexer cache.Indexer
}

// NewAntstorDataControlLister returns a new AntstorDataControlLister.
func NewAntstorDataControlLister(indexer cache.Indexer) AntstorDataControlLister {
	return &antstorDataControlLister{indexer: indexer}
}

// List lists all AntstorDataControls in the indexer.
func (s *antstorDataControlLister) List(selector labels.Selector) (ret []*v1.AntstorDataControl, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AntstorDataControl))
	})
	return ret, err
}

// AntstorDataControls returns an object that can list and get AntstorDataControls.
func (s *antstorDataControlLister) AntstorDataControls(namespace string) AntstorDataControlNamespaceLister {
	return antstorDataControlNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AntstorDataControlNamespaceLister helps list and get AntstorDataControls.
// All objects returned here must be treated as read-only.
type AntstorDataControlNamespaceLister interface {
	// List lists all AntstorDataControls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AntstorDataControl, err error)
	// Get retrieves the AntstorDataControl from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AntstorDataControl, error)
	AntstorDataControlNamespaceListerExpansion
}

// antstorDataControlNamespaceLister implements the AntstorDataControlNamespaceLister
// interface.
type antstorDataControlNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AntstorDataControls in the indexer for a given namespace.
func (s antstorDataControlNamespaceLister) List(selector labels.Selector) (ret []*v1.AntstorDataControl, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AntstorDataControl))
	})
	return ret, err
}

// Get retrieves the AntstorDataControl from the indexer for a given namespace and name.
func (s antstorDataControlNamespaceLister) Get(name string) (*v1.AntstorDataControl, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("antstordatacontrol"), name)
	}
	return obj.(*v1.AntstorDataControl), nil
}