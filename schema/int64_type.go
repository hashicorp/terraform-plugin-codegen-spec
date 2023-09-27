// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int64Type is a representation of a 64-bit integer.
type Int64Type struct {
	// CustomType is a customization of the Int64Type.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
