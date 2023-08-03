// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

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
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type ListNestedBlock struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedBlockObject       `json:"nested_object"`

	CustomType         *schema.CustomType     `json:"custom_type,omitempty"`
	DeprecationMessage *string                `json:"deprecation_message,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Sensitive          *bool                  `json:"sensitive,omitempty"`
	Validators         []schema.ListValidator `json:"validators,omitempty"`
}

type SetNestedBlock struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedBlockObject       `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.SetValidator `json:"validators,omitempty"`
}

type SingleNestedBlock struct {
	Attributes       Attributes              `json:"attributes,omitempty"`
	Blocks           Blocks                  `json:"blocks,omitempty"`
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}
