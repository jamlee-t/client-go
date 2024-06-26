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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// AllocationResultApplyConfiguration represents a declarative configuration of the AllocationResult type for use
// with apply.
type AllocationResultApplyConfiguration struct {
	ResourceHandles  []ResourceHandleApplyConfiguration `json:"resourceHandles,omitempty"`
	AvailableOnNodes *v1.NodeSelectorApplyConfiguration `json:"availableOnNodes,omitempty"`
	Shareable        *bool                              `json:"shareable,omitempty"`
}

// AllocationResultApplyConfiguration constructs a declarative configuration of the AllocationResult type for use with
// apply.
func AllocationResult() *AllocationResultApplyConfiguration {
	return &AllocationResultApplyConfiguration{}
}

// WithResourceHandles adds the given value to the ResourceHandles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceHandles field.
func (b *AllocationResultApplyConfiguration) WithResourceHandles(values ...*ResourceHandleApplyConfiguration) *AllocationResultApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceHandles")
		}
		b.ResourceHandles = append(b.ResourceHandles, *values[i])
	}
	return b
}

// WithAvailableOnNodes sets the AvailableOnNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableOnNodes field is set to the value of the last call.
func (b *AllocationResultApplyConfiguration) WithAvailableOnNodes(value *v1.NodeSelectorApplyConfiguration) *AllocationResultApplyConfiguration {
	b.AvailableOnNodes = value
	return b
}

// WithShareable sets the Shareable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shareable field is set to the value of the last call.
func (b *AllocationResultApplyConfiguration) WithShareable(value bool) *AllocationResultApplyConfiguration {
	b.Shareable = &value
	return b
}
