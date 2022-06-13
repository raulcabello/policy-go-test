// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// FCVolumeSource Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
//
// swagger:model FCVolumeSource
type FCVolumeSource struct {

	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType string `json:"fsType,omitempty"`

	// lun is Optional: FC target lun number
	Lun int32 `json:"lun,omitempty"`

	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// targetWWNs is Optional: FC target worldwide names (WWNs)
	TargetWWNs []string `json:"targetWWNs"`

	// wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids []string `json:"wwids"`
}