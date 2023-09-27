// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// BoolType is a representation of a boolean.
type BoolType struct {
	// CustomType is a customization of the BoolType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
