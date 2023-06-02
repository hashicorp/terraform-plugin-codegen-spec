// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type Float64Type struct {
	// CustomType is a customization of the Float64Type.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
