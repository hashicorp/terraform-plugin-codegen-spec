// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// SetType is a representation of a set.
type SetType struct {
	// ElementType defines the type of the elements within
	// the set.
	ElementType ElementType `json:"element_type"`

	// CustomType is a customization of the SetType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
