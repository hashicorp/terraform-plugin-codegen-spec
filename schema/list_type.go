// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ListType is a representation of a list.
type ListType struct {
	// ElementType defines the type of the elements within
	// the list.
	ElementType ElementType `json:"element_type"`

	// CustomType is a customization of the ListType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
