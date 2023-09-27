// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ObjectDefault defines a custom type for a default object value.
type ObjectDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`
}

// Equal returns true if all fields of the given ObjectDefault are equal.
func (d *ObjectDefault) Equal(other *ObjectDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
