// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// DynamicType is a representation of dynamic type.
type DynamicType struct {
	// CustomType is a customization of the DynamicType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
