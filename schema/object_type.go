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

	if !o.CustomType.Equal(other.CustomType) {
		return false
	}

	if o.AttributeTypes == nil && other.AttributeTypes != nil {
		return false
	}

	if o.AttributeTypes != nil && other.AttributeTypes == nil {
		return false
	}

	for k, v := range o.AttributeTypes {
		if v.Name != other.AttributeTypes[k].Name {
			return false
		}

		a := ElementType{
			Bool:    v.Bool,
			Float64: v.Float64,
			Int64:   v.Int64,
			List:    v.List,
			Map:     v.Map,
			Number:  v.Number,
			Object:  v.Object,
			Set:     v.Set,
			String:  v.String,
		}

		b := ElementType{
			Bool:    other.AttributeTypes[k].Bool,
			Float64: other.AttributeTypes[k].Float64,
			Int64:   other.AttributeTypes[k].Int64,
			List:    other.AttributeTypes[k].List,
			Map:     other.AttributeTypes[k].Map,
			Number:  other.AttributeTypes[k].Number,
			Object:  other.AttributeTypes[k].Object,
			Set:     other.AttributeTypes[k].Set,
			String:  other.AttributeTypes[k].String,
		}

		if !a.Equal(b) {
			return false
		}
	}

	return true
}
