// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// StringType is a representation of a string.
type StringType struct {
	// CustomType is a customization of the StringType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
