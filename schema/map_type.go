// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// MapType is a representation of a map.
type MapType struct {
	// ElementType defines the type of the elements within
	// the map.
	ElementType ElementType `json:"element_type"`

	// CustomType is a customization of the MapType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
