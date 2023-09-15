// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ObjectType struct {
	AttributeTypes ObjectAttributeTypes `json:"attribute_types"`

	// CustomType is a customization of the ObjectType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}

func (o *ObjectType) Equal(other *ObjectType) bool {
	if o == nil && other == nil {
		return true
	}

	if o == nil || other == nil {
		return false
	}

	if !o.AttributeTypes.Equal(other.AttributeTypes) {
		return false
	}

	return o.CustomType.Equal(other.CustomType)
}
