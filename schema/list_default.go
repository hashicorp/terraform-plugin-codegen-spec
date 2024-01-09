// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ListDefault defines a custom type for a default list value.
type ListDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`
}

// CustomDefault returns *CustomDefault.
func (d *ListDefault) CustomDefault() *CustomDefault {
	if d == nil {
		return nil
	}

	return d.Custom
}

// Equal returns true if all fields of the given ListDefault are equal.
func (d *ListDefault) Equal(other *ListDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
