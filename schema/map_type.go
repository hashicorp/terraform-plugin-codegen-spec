// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type MapType struct {
	ElementType

	// CustomType is a customization of the MapType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
