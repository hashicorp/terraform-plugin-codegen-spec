// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type StringType struct {
	// CustomType is a customization of the StringType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
