// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// BoolDefault defines a value, or a custom type for a default boolean value.
type BoolDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`

	// Static defines a specific boolean value.
	Static *bool `json:"static,omitempty"`
}

// Equal returns true if all fields of the given BoolDefault are equal.
func (d *BoolDefault) Equal(other *BoolDefault) bool {
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
