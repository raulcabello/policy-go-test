// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ConfigMapNodeConfigSource ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node. This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
//
// swagger:model ConfigMapNodeConfigSource
type ConfigMapNodeConfigSource struct {

	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases.
	// Required: true
	KubeletConfigKey *string `json:"kubeletConfigKey"`

	// Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
	// Required: true
	Name *string `json:"name"`

	// Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
	// Required: true
	Namespace *string `json:"namespace"`

	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	UID string `json:"uid,omitempty"`
}