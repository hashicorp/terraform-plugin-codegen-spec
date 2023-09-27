// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// StringDefault defines a value, or a custom type for a default string value.
type StringDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`

	// Static defines a specific string value.
	Static *string `json:"static,omitempty"`
}

// Equal returns true if all fields of the given StringDefault are equal.
func (d *StringDefault) Equal(other *StringDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	if !d.Custom.Equal(other.Custom) {
		return false
	}

	if d.Static == nil && other.Static != nil {
		return false
	}

	if d.Static != nil && other.Static == nil {
		return false
	}

	return *d.Static == *other.Static
}
