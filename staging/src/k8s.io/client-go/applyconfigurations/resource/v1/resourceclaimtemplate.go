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
	resourcev1 "k8s.io/api/resource/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ResourceClaimTemplateApplyConfiguration represents a declarative configuration of the ResourceClaimTemplate type for use
// with apply.
type ResourceClaimTemplateApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                                 *ResourceClaimTemplateSpecApplyConfiguration `json:"spec,omitempty"`
}

// ResourceClaimTemplate constructs a declarative configuration of the ResourceClaimTemplate type for use with
// apply.
func ResourceClaimTemplate(name, namespace string) *ResourceClaimTemplateApplyConfiguration {
	b := &ResourceClaimTemplateApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ResourceClaimTemplate")
	b.WithAPIVersion("resource.k8s.io/v1")
	return b
}

// ExtractResourceClaimTemplate extracts the applied configuration owned by fieldManager from
// resourceClaimTemplate. If no managedFields are found in resourceClaimTemplate for fieldManager, a
// ResourceClaimTemplateApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// resourceClaimTemplate must be a unmodified ResourceClaimTemplate API object that was retrieved from the Kubernetes API.
// ExtractResourceClaimTemplate provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractResourceClaimTemplate(resourceClaimTemplate *resourcev1.ResourceClaimTemplate, fieldManager string) (*ResourceClaimTemplateApplyConfiguration, error) {
	return extractResourceClaimTemplate(resourceClaimTemplate, fieldManager, "")
}

// ExtractResourceClaimTemplateStatus is the same as ExtractResourceClaimTemplate except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractResourceClaimTemplateStatus(resourceClaimTemplate *resourcev1.ResourceClaimTemplate, fieldManager string) (*ResourceClaimTemplateApplyConfiguration, error) {
	return extractResourceClaimTemplate(resourceClaimTemplate, fieldManager, "status")
}

func extractResourceClaimTemplate(resourceClaimTemplate *resourcev1.ResourceClaimTemplate, fieldManager string, subresource string) (*ResourceClaimTemplateApplyConfiguration, error) {
	b := &ResourceClaimTemplateApplyConfiguration{}
	err := managedfields.ExtractInto(resourceClaimTemplate, internal.Parser().Type("io.k8s.api.resource.v1.ResourceClaimTemplate"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(resourceClaimTemplate.Name)
	b.WithNamespace(resourceClaimTemplate.Namespace)

	b.WithKind("ResourceClaimTemplate")
	b.WithAPIVersion("resource.k8s.io/v1")
	return b, nil
}
func (b ResourceClaimTemplateApplyConfiguration) IsApplyConfiguration() {}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithKind(value string) *ResourceClaimTemplateApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithAPIVersion(value string) *ResourceClaimTemplateApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithName(value string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithGenerateName(value string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithNamespace(value string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithUID(value types.UID) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithResourceVersion(value string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithGeneration(value int64) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ResourceClaimTemplateApplyConfiguration) WithLabels(entries map[string]string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ResourceClaimTemplateApplyConfiguration) WithAnnotations(entries map[string]string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ResourceClaimTemplateApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ResourceClaimTemplateApplyConfiguration) WithFinalizers(values ...string) *ResourceClaimTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *ResourceClaimTemplateApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ResourceClaimTemplateApplyConfiguration) WithSpec(value *ResourceClaimTemplateSpecApplyConfiguration) *ResourceClaimTemplateApplyConfiguration {
	b.Spec = value
	return b
}

// GetKind retrieves the value of the Kind field in the declarative configuration.
func (b *ResourceClaimTemplateApplyConfiguration) GetKind() *string {
	return b.TypeMetaApplyConfiguration.Kind
}

// GetAPIVersion retrieves the value of the APIVersion field in the declarative configuration.
func (b *ResourceClaimTemplateApplyConfiguration) GetAPIVersion() *string {
	return b.TypeMetaApplyConfiguration.APIVersion
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ResourceClaimTemplateApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}

// GetNamespace retrieves the value of the Namespace field in the declarative configuration.
func (b *ResourceClaimTemplateApplyConfiguration) GetNamespace() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Namespace
}
