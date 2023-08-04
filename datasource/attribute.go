// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

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
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type BoolAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.BoolValidator         `json:"validators,omitempty"`
}

type Float64Attribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Float64Validator      `json:"validators,omitempty"`
}

type Int64Attribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Int64Validator        `json:"validators,omitempty"`
}

type ListAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ListValidator         `json:"validators,omitempty"`
}

type ListNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType     `json:"custom_type,omitempty"`
	DeprecationMessage *string                `json:"deprecation_message,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Sensitive          *bool                  `json:"sensitive,omitempty"`
	Validators         []schema.ListValidator `json:"validators,omitempty"`
}

type MapAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.MapValidator          `json:"validators,omitempty"`
}

type MapNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.MapValidator `json:"validators,omitempty"`
}

type NumberAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.NumberValidator       `json:"validators,omitempty"`
}

type ObjectAttribute struct {
	AttributeTypes           schema.ObjectAttributeTypes     `json:"attribute_types"`
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type SetAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.SetValidator          `json:"validators,omitempty"`
}

type SetNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.SetValidator `json:"validators,omitempty"`
}

type SingleNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	Attributes               Attributes                      `json:"attributes,omitempty"`
	AssociatedExternalType   *schema.AssociatedExternalType  `json:"associated_external_type,omitempty"`
	CustomType               *schema.CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                         `json:"deprecation_message,omitempty"`
	Description              *string                         `json:"description,omitempty"`
	Sensitive                *bool                           `json:"sensitive,omitempty"`
	Validators               []schema.ObjectValidator        `json:"validators,omitempty"`
}

type StringAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.StringValidator       `json:"validators,omitempty"`
}
