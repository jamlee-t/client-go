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

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
)

// PodFailurePolicyOnExitCodesRequirementApplyConfiguration represents a declarative configuration of the PodFailurePolicyOnExitCodesRequirement type for use
// with apply.
type PodFailurePolicyOnExitCodesRequirementApplyConfiguration struct {
	ContainerName *string                                      `json:"containerName,omitempty"`
	Operator      *batchv1.PodFailurePolicyOnExitCodesOperator `json:"operator,omitempty"`
	Values        []int32                                      `json:"values,omitempty"`
}

// PodFailurePolicyOnExitCodesRequirementApplyConfiguration constructs a declarative configuration of the PodFailurePolicyOnExitCodesRequirement type for use with
// apply.
func PodFailurePolicyOnExitCodesRequirement() *PodFailurePolicyOnExitCodesRequirementApplyConfiguration {
	return &PodFailurePolicyOnExitCodesRequirementApplyConfiguration{}
}
func (b PodFailurePolicyOnExitCodesRequirementApplyConfiguration) IsApplyConfiguration() {}

// WithContainerName sets the ContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerName field is set to the value of the last call.
func (b *PodFailurePolicyOnExitCodesRequirementApplyConfiguration) WithContainerName(value string) *PodFailurePolicyOnExitCodesRequirementApplyConfiguration {
	b.ContainerName = &value
	return b
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *PodFailurePolicyOnExitCodesRequirementApplyConfiguration) WithOperator(value batchv1.PodFailurePolicyOnExitCodesOperator) *PodFailurePolicyOnExitCodesRequirementApplyConfiguration {
	b.Operator = &value
	return b
}

// WithValues adds the given value to the Values field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Values field.
func (b *PodFailurePolicyOnExitCodesRequirementApplyConfiguration) WithValues(values ...int32) *PodFailurePolicyOnExitCodesRequirementApplyConfiguration {
	for i := range values {
		b.Values = append(b.Values, values[i])
	}
	return b
}
