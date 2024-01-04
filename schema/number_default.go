// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// NumberDefault defines a custom type for a default generic
// number with up to 512 bits of floating point or integer precision.
type NumberDefault struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`
}

// CustomDefault returns *CustomDefault.
func (d *NumberDefault) CustomDefault() *CustomDefault {
	if d == nil {
		return nil
	}

	return d.Custom
}

// Equal returns true if all fields of the given Int64Default are equal.
func (d *NumberDefault) Equal(other *NumberDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
