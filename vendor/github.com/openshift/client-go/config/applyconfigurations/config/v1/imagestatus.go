// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ImageStatusApplyConfiguration represents an declarative configuration of the ImageStatus type for use
// with apply.
type ImageStatusApplyConfiguration struct {
	InternalRegistryHostname  *string  `json:"internalRegistryHostname,omitempty"`
	ExternalRegistryHostnames []string `json:"externalRegistryHostnames,omitempty"`
}

// ImageStatusApplyConfiguration constructs an declarative configuration of the ImageStatus type for use with
// apply.
func ImageStatus() *ImageStatusApplyConfiguration {
	return &ImageStatusApplyConfiguration{}
}

// WithInternalRegistryHostname sets the InternalRegistryHostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalRegistryHostname field is set to the value of the last call.
func (b *ImageStatusApplyConfiguration) WithInternalRegistryHostname(value string) *ImageStatusApplyConfiguration {
	b.InternalRegistryHostname = &value
	return b
}

// WithExternalRegistryHostnames adds the given value to the ExternalRegistryHostnames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalRegistryHostnames field.
func (b *ImageStatusApplyConfiguration) WithExternalRegistryHostnames(values ...string) *ImageStatusApplyConfiguration {
	for i := range values {
		b.ExternalRegistryHostnames = append(b.ExternalRegistryHostnames, values[i])
	}
	return b
}