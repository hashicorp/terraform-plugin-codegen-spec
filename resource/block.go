// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type BlockValidateRequest struct {
	Path string
}

type Blocks []Block

// Validate checks for duplicated block names. Validate is called recursively in
// instances where a block contains nested blocks. Validate delegates to
// Attributes.Validate in instances where the block has attributes.
func (b Blocks) Validate(ctx context.Context, req BlockValidateRequest) error {
	blockNames := make(map[string]struct{}, len(b))

	var errs, nestedErrs []error

	for _, block := range b {
		if _, ok := blockNames[block.Name]; ok {
			errs = append(errs, fmt.Errorf("%s block %q is duplicated", req.Path, block.Name))
		}

		blockNames[block.Name] = struct{}{}

		attributeValidateRequest := AttributeValidateRequest{
			Path: fmt.Sprintf("%s block %q", req.Path, block.Name),
		}

		blockValidateRequest := BlockValidateRequest{
			Path: fmt.Sprintf("%s block %q", req.Path, block.Name),
		}

		var attributeErr, blockErr error

		switch {
		case block.ListNested != nil:
			attributeErr = block.ListNested.NestedObject.Attributes.Validate(ctx, attributeValidateRequest)
			blockErr = block.ListNested.NestedObject.Blocks.Validate(ctx, blockValidateRequest)
		case block.SetNested != nil:
			attributeErr = block.SetNested.NestedObject.Attributes.Validate(ctx, attributeValidateRequest)
			blockErr = block.SetNested.NestedObject.Blocks.Validate(ctx, blockValidateRequest)
		case block.SingleNested != nil:
			attributeErr = block.SingleNested.Attributes.Validate(ctx, attributeValidateRequest)
			blockErr = block.SingleNested.Blocks.Validate(ctx, blockValidateRequest)
		}

		if attributeErr != nil {
			nestedErrs = append(nestedErrs, attributeErr)
		}

		if blockErr != nil {
			nestedErrs = append(nestedErrs, blockErr)
		}
	}

	e := append(errs, nestedErrs...)

	return errors.Join(e...)
}

type Block struct {
	Name string `json:"name"`

	ListNested   *ListNestedBlock   `json:"list_nested,omitempty"`
	SetNested    *SetNestedBlock    `json:"set_nested,omitempty"`
	SingleNested *SingleNestedBlock `json:"single_nested,omitempty"`
}

type NestedBlockObject struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Blocks     Blocks     `json:"blocks,omitempty"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	PlanModifiers          []schema.ObjectPlanModifier    `json:"plan_modifiers,omitempty"`
	Validators             schema.ObjectValidators        `json:"validators,omitempty"`
}

type ListNestedBlock struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedBlockObject               `json:"nested_object"`

	CustomType         *schema.CustomType        `json:"custom_type,omitempty"`
	Default            *schema.ListDefault       `json:"default,omitempty"`
	DeprecationMessage *string                   `json:"deprecation_message,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	PlanModifiers      []schema.ListPlanModifier `json:"plan_modifiers,omitempty"`
	Sensitive          *bool                     `json:"sensitive,omitempty"`
	Validators         schema.ListValidators     `json:"validators,omitempty"`
}

type SetNestedBlock struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedBlockObject               `json:"nested_object"`

	CustomType         *schema.CustomType       `json:"custom_type,omitempty"`
	Default            *schema.SetDefault       `json:"default,omitempty"`
	DeprecationMessage *string                  `json:"deprecation_message,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	PlanModifiers      []schema.SetPlanModifier `json:"plan_modifiers,omitempty"`
	Sensitive          *bool                    `json:"sensitive,omitempty"`
	Validators         schema.SetValidators     `json:"validators,omitempty"`
}

type SingleNestedBlock struct {
	Attributes               Attributes                      `json:"attributes,omitempty"`
	Blocks                   Blocks                          `json:"blocks,omitempty"`
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.ObjectDefault          `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.ObjectPlanModifier    `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             schema.ObjectValidators        `json:"validators,omitempty"`
}
