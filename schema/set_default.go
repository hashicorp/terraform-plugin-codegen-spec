// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// SetDefault defines a custom type for a default set value.
type SetDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`
}

// Equal returns true if all fields of the given SetDefault are equal.
func (d *SetDefault) Equal(other *SetDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
