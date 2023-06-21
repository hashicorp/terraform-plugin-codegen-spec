// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ObjectType struct {
	AttributeTypes ObjectAttributeTypes `json:"attribute_types"`

	// CustomType is a customization of the ObjectType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
