// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ObjectType is a representation of an object.
type ObjectType struct {
	// AttributeTypes defines the types of the attributes
	// within the object.
	AttributeTypes ObjectAttributeTypes `json:"attribute_types"`

	// CustomType is a customization of the ObjectType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}

// Equal returns true if the fields of the given ObjectType are equal.
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
