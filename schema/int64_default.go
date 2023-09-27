// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int64Default defines a value, or a custom type for a default 64-bit integer value.
type Int64Default struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`

	// Static defines a specific 64-bit integer value.
	Static *int64 `json:"static,omitempty"`
}

// Equal returns true if all fields of the given Int64Default are equal.
func (d *Int64Default) Equal(other *Int64Default) bool {
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
