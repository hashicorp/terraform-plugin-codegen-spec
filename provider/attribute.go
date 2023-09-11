// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type AttributeValidateRequest struct {
	Path string
}

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

type Attribute struct {
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

type NestedAttributeObject struct {
	Attributes Attributes `json:"attributes,omitempty"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Validators             schema.ObjectValidators        `json:"validators,omitempty"`
}

type BoolAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.BoolValidators          `json:"validators,omitempty"`
}

type Float64Attribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.Float64Validators       `json:"validators,omitempty"`
}

type Int64Attribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.Int64Validators         `json:"validators,omitempty"`
}

type ListAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.ListValidators          `json:"validators,omitempty"`
}

type ListNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         schema.ListValidators `json:"validators,omitempty"`
}

type MapAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.MapValidators           `json:"validators,omitempty"`
}

type MapNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType   `json:"custom_type,omitempty"`
	DeprecationMessage *string              `json:"deprecation_message,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Sensitive          *bool                `json:"sensitive,omitempty"`
	Validators         schema.MapValidators `json:"validators,omitempty"`
}

type NumberAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.NumberValidators        `json:"validators,omitempty"`
}

type ObjectAttribute struct {
	AttributeTypes   schema.ObjectAttributeTypes `json:"attribute_types"`
	OptionalRequired schema.OptionalRequired     `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.ObjectValidators        `json:"validators,omitempty"`
}

type SetAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.SetValidators           `json:"validators,omitempty"`
}

type SetNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType   `json:"custom_type,omitempty"`
	DeprecationMessage *string              `json:"deprecation_message,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Sensitive          *bool                `json:"sensitive,omitempty"`
	Validators         schema.SetValidators `json:"validators,omitempty"`
}

type SingleNestedAttribute struct {
	OptionalRequired       schema.OptionalRequired        `json:"optional_required"`
	Attributes             Attributes                     `json:"attributes,omitempty"`
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.ObjectValidators        `json:"validators,omitempty"`
}

type StringAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.StringValidators        `json:"validators,omitempty"`
}
