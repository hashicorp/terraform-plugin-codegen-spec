// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// MapDefault defines a custom type for a default map value.
type MapDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`
}

// Equal returns true if all fields of the given MapDefault are equal.
func (d *MapDefault) Equal(other *MapDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
