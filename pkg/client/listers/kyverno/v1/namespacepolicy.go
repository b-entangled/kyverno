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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nirmata/kyverno/pkg/api/kyverno/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespacePolicyLister helps list NamespacePolicies.
type NamespacePolicyLister interface {
	// List lists all NamespacePolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.NamespacePolicy, err error)
	// NamespacePolicies returns an object that can list and get NamespacePolicies.
	NamespacePolicies(namespace string) NamespacePolicyNamespaceLister
	NamespacePolicyListerExpansion
}

// namespacePolicyLister implements the NamespacePolicyLister interface.
type namespacePolicyLister struct {
	indexer cache.Indexer
}

// NewNamespacePolicyLister returns a new NamespacePolicyLister.
func NewNamespacePolicyLister(indexer cache.Indexer) NamespacePolicyLister {
	return &namespacePolicyLister{indexer: indexer}
}

// List lists all NamespacePolicies in the indexer.
func (s *namespacePolicyLister) List(selector labels.Selector) (ret []*v1.NamespacePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NamespacePolicy))
	})
	return ret, err
}

// NamespacePolicies returns an object that can list and get NamespacePolicies.
func (s *namespacePolicyLister) NamespacePolicies(namespace string) NamespacePolicyNamespaceLister {
	return namespacePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespacePolicyNamespaceLister helps list and get NamespacePolicies.
type NamespacePolicyNamespaceLister interface {
	// List lists all NamespacePolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.NamespacePolicy, err error)
	// Get retrieves the NamespacePolicy from the indexer for a given namespace and name.
	Get(name string) (*v1.NamespacePolicy, error)
}

// namespacePolicyNamespaceLister implements the NamespacePolicyNamespaceLister
// interface.
type namespacePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NamespacePolicies in the indexer for a given namespace.
func (s namespacePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.NamespacePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NamespacePolicy))
	})
	return ret, err
}

// Get retrieves the NamespacePolicy from the indexer for a given namespace and name.
func (s namespacePolicyNamespaceLister) Get(name string) (*v1.NamespacePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("namespacepolicy"), name)
	}
	return obj.(*v1.NamespacePolicy), nil
}
