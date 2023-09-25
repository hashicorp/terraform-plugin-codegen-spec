// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// AttributeValidateRequest defines the Path of the attribute that is
// being validated.
type AttributeValidateRequest struct {
	Path string
}

// Attributes type defines Attribute types.
type Attributes []Attribute

// Validate checks for duplicated attribute names. Validate is called recursively in
// instances where an attribute contains nested attributes. Validate delegates to
// ObjectAttributeTypes.Validate when the attribute is an ObjectAttribute.
func (a Attributes) Validate(ctx context.Context, req AttributeValidateRequest) error {
	attributeNames := make(map[string]struct{}, len(a))

	var errs, nestedErrs []error

	for _, attribute := range a {
		if _, ok := attributeNames[attribute.Name]; ok {
			errs = append(errs, fmt.Errorf("%s attribute %q is duplicated", req.Path, attribute.Name))
		}

		attributeNames[attribute.Name] = struct{}{}

		var err error

		attributeValidateRequest := AttributeValidateRequest{
			Path: fmt.Sprintf("%s attribute %q", req.Path, attribute.Name),
		}

		objectValidateRequest := schema.ObjectValidateRequest{
			Path: fmt.Sprintf("%s attribute %q", req.Path, attribute.Name),
		}

		switch {
		case attribute.ListNested != nil:
			err = attribute.ListNested.NestedObject.Attributes.Validate(ctx, attributeValidateRequest)
		case attribute.MapNested != nil:
			err = attribute.MapNested.NestedObject.Attributes.Validate(ctx, attributeValidateRequest)
		case attribute.Object != nil:
			err = attribute.Object.AttributeTypes.Validate(ctx, objectValidateRequest)
		case attribute.SetNested != nil:
			err = attribute.SetNested.NestedObject.Attributes.Validate(ctx, attributeValidateRequest)
		case attribute.SingleNested != nil:
			err = attribute.SingleNested.Attributes.Validate(ctx, attributeValidateRequest)
		}

		if err != nil {
			nestedErrs = append(nestedErrs, err)
		}
	}

	e := append(errs, nestedErrs...)

	return errors.Join(e...)
}

// Attribute defines a value field inside a Schema. The attribute types
// (e.g., Bool, Float64) are mutually exclusive, one and only one must
// be specified.
type Attribute struct {
	// Name defines the attribute name.
	Name string `json:"name"`

	Bool         *BoolAttribute         `json:"bool,omitempty"`
	Float64      *Float64Attribute      `json:"float64,omitempty"`
	Int64        *Int64Attribute        `json:"int64,omitempty"`
	List         *ListAttribute         `json:"list,omitempty"`
	ListNested   *ListNestedAttribute   `json:"list_nested,omitempty"`
	Map          *MapAttribute          `json:"map,omitempty"`
	MapNested    *MapNestedAttribute    `json:"map_nested,omitempty"`
	Number       *NumberAttribute       `json:"number,omitempty"`
	Object       *ObjectAttribute       `json:"object,omitempty"`
	Set          *SetAttribute          `json:"set,omitempty"`
	SetNested    *SetNestedAttribute    `json:"set_nested,omitempty"`
	SingleNested *SingleNestedAttribute `json:"single_nested,omitempty"`
	String       *StringAttribute       `json:"string,omitempty"`
}

// NestedAttributeObject is the underlying object defining the Attributes
// for a ListNestedAttribute, MapNestedAttribute, or SetNestedAttribute.
type NestedAttributeObject struct {
	// Attributes defines the Attribute types associated with a NestedAttributeObject.
	Attributes Attributes `json:"attributes,omitempty"`

	// AssociatedExternalType defines a type that can be used as a NestedAttributeObject.
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`

	// CustomType defines a custom type and value for the NestedAttributeObject.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the NestedAttributeObject.
	PlanModifiers schema.ObjectPlanModifiers `json:"plan_modifiers,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the NestedAttributeObject.
	Validators schema.ObjectValidators `json:"validators,omitempty"`
}

// BoolAttribute represents a Schema attribute that is a boolean.
type BoolAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.BoolDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.BoolPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.BoolValidators `json:"validators,omitempty"`
}

// Float64Attribute represents a Schema attribute that is a 64-bit
// floating point number.
//
// Use Int64Attribute for a 64-bit integer attribute, or NumberAttribute
// for a 512-bit generic number attribute.
type Float64Attribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.Float64Default `json:"default,omitempty"`
	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.Float64PlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.Float64Validators `json:"validators,omitempty"`
}

// Int64Attribute represents a schema attribute that is a 64-bit
// integer.
//
// Use Float64Attribute for a 64-bit floating point number, or
// NumberAttribute for a 512-bit generic number attribute.
type Int64Attribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.Int64Default `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.Int64PlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.Int64Validators `json:"validators,omitempty"`
}

// ListAttribute represents a Schema attribute that is a list with a single
// element type.
type ListAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// ElementType is the type for all elements of the list.
	ElementType schema.ElementType `json:"element_type"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.ListDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.ListPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ListValidators `json:"validators,omitempty"`
}

// ListNestedAttribute represents a Schema attribute that is a list of
// objects, where the object attributes can be fully defined.
type ListNestedAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// NestedObject defines the underlying object attributes.
	NestedObject NestedAttributeObject `json:"nested_object"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.ListDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.ListPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ListValidators `json:"validators,omitempty"`
}

// MapAttribute represents a Schema attribute that is a map with a single
// element type.
type MapAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// ElementType is the type for all elements of the map.
	ElementType schema.ElementType `json:"element_type"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.MapDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.MapPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.MapValidators `json:"validators,omitempty"`
}

// MapNestedAttribute represents a Schema attribute that is a map of
// name to objects, where the object attributes can be fully defined.
type MapNestedAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// NestedObject defines the underlying object attributes.
	NestedObject NestedAttributeObject `json:"nested_object"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.MapDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.MapPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.MapValidators `json:"validators,omitempty"`
}

// NumberAttribute represents a schema attribute that is a generic
// number with up to 512 bits of floating point or integer precision.
//
// Use Float64Attribute for a 64-bit floating point number attribute,
// or Int64Attribute for a 64-bit integer number attribute.
type NumberAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.NumberDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.NumberPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.NumberValidators `json:"validators,omitempty"`
}

// ObjectAttribute represents a Schema attribute that is an object with only
// type information for underlying attributes.
type ObjectAttribute struct {
	// AttributeTypes provides the mapping of underlying names to types.
	AttributeTypes schema.ObjectAttributeTypes `json:"attribute_types"`

	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.ObjectDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.ObjectPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ObjectValidators `json:"validators,omitempty"`
}

// SetAttribute represents a Schema attribute that is a set with a single
// element type.
type SetAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// ElementType is the type for all elements of the set.
	ElementType schema.ElementType `json:"element_type"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.SetDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.SetPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.SetValidators `json:"validators,omitempty"`
}

// SetNestedAttribute represents a Schema attribute that is a list of
// objects, where the object attributes can be fully defined.
type SetNestedAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// NestedObject defines the underlying object attributes.
	NestedObject NestedAttributeObject `json:"nested_object"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.SetDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.SetPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.SetValidators `json:"validators,omitempty"`
}

// SingleNestedAttribute represents a Schema attribute that is a single object where
// the object attributes can be fully defined
type SingleNestedAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// Attributes defines the Attribute types associated with a SingleNestedAttribute.
	Attributes Attributes `json:"attributes,omitempty"`

	// AssociatedExternalType defines a type that can be used as a NestedAttributeObject.
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.ObjectDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.ObjectPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ObjectValidators `json:"validators,omitempty"`
}

// StringAttribute represents a Schema attribute that is a string.
type StringAttribute struct {
	// ComputedOptionalRequired indicates whether the attribute is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// CustomType defines a custom type and value for the attribute.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the attribute.
	Default *schema.StringDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the attribute.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the attribute.
	PlanModifiers schema.StringPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the attribute should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.StringValidators `json:"validators,omitempty"`
}
