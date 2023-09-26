// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// NumberType is a representation of a generic number with
// up to 512 bits of floating point or integer precision.
type NumberType struct {
	// CustomType is a customization of the NumberType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
