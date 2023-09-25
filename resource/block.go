// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// BlockValidateRequest defines the Path of the block that is
// being validated.
type BlockValidateRequest struct {
	Path string
}

// Blocks type defines Block types.
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

// Block defines a structural field inside a Schema. The block types
// (e.g., ListNested, SetNested) are mutually exclusive, one and
// only one must be specified.
type Block struct {
	Name string `json:"name"`

	ListNested   *ListNestedBlock   `json:"list_nested,omitempty"`
	SetNested    *SetNestedBlock    `json:"set_nested,omitempty"`
	SingleNested *SingleNestedBlock `json:"single_nested,omitempty"`
}

// NestedBlockObject is the underlying object defining the Attributes
// for a ListNestedBlock, or SetNestedBlock.
type NestedBlockObject struct {
	// Attributes defines the Attribute types associated with a NestedBlockObject.
	Attributes Attributes `json:"attributes,omitempty"`

	// Blocks defines the Block types associated with a NestedBlockObject.
	Blocks Blocks `json:"blocks,omitempty"`

	// AssociatedExternalType defines a type that can be used as a NestedBlockObject.
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`

	// CustomType defines a custom type and value for the NestedBlockObject.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the NestedBlockObject.
	PlanModifiers schema.ObjectPlanModifiers `json:"plan_modifiers,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the NestedBlockObject.
	Validators schema.ObjectValidators `json:"validators,omitempty"`
}

// ListNestedBlock represents a block that is a list of objects where
// the object attributes can be fully defined
type ListNestedBlock struct {
	// ComputedOptionalRequired indicates whether the block is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// NestedObject defines the underlying object attributes and blocks.
	NestedObject NestedBlockObject `json:"nested_object"`

	// CustomType defines a custom type and value for the block.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the block.
	Default *schema.ListDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the attribute
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the block.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the block.
	PlanModifiers schema.ListPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the block should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ListValidators `json:"validators,omitempty"`
}

// SetNestedBlock represents a block that is a set of objects where
// the object attributes can be fully defined
type SetNestedBlock struct {
	// ComputedOptionalRequired indicates whether the block is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// NestedObject defines the underlying object attributes and blocks.
	NestedObject NestedBlockObject `json:"nested_object"`

	// CustomType defines a custom type and value for the block.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the block.
	Default *schema.SetDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the block
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the block.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the block.
	PlanModifiers schema.SetPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the block should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.SetValidators `json:"validators,omitempty"`
}

// SingleNestedBlock represents a block that is a single object where
// the object attributes can be fully defined.
type SingleNestedBlock struct {
	// Attributes defines the Attribute types associated with the SingleNestedBlock.
	Attributes Attributes `json:"attributes,omitempty"`

	// Blocks defines the Block types associated with the SingleNestedBlock.
	Blocks Blocks `json:"blocks,omitempty"`

	// ComputedOptionalRequired indicates whether the block is required
	// (`required`), optional (`optional`), computed (`computed`), or
	// computed and optional (`computed_optional`).
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	// AssociatedExternalType defines a type that can be used as a NestedAttributeObject.
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`

	// CustomType defines a custom type and value for the block.
	CustomType *schema.CustomType `json:"custom_type,omitempty"`

	// Default defines a default value for the block.
	Default *schema.ObjectDefault `json:"default,omitempty"`

	// DeprecationMessage defines a message describing that the block
	// is deprecated.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`

	// Description defines the purpose and usage of the block.
	Description *string `json:"description,omitempty"`

	// PlanModifiers define types and functions that provide plan modification
	// functionality for the block.
	PlanModifiers schema.ObjectPlanModifiers `json:"plan_modifiers,omitempty"`

	// Sensitive indicates whether the value of the block should
	// be considered sensitive data.
	Sensitive *bool `json:"sensitive,omitempty"`

	// Validators define types and functions that provide validation
	// functionality for the block.
	Validators schema.ObjectValidators `json:"validators,omitempty"`
}
