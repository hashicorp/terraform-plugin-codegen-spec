// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"errors"
)

// Schema defines the Attributes and Blocks associated with a Provider.
type Schema struct {
	// Attributes defines the Attribute types for the Schema..
	Attributes Attributes `json:"attributes,omitempty"`

	// Blocks defines the Block types for the Schema.
	Blocks Blocks `json:"blocks,omitempty"`

	// Description is used in various tooling, like the language server, to
	// give practitioners more information about what this provider is,
	// what it's for, and how it should be used. It should be written as
	// plain text, with no special formatting.
	Description *string `json:"description,omitempty"`

	// MarkdownDescription is used in various tooling, like the
	// documentation generator, to give practitioners more information
	// about what this provider is, what it's for, and how it should be
	// used. It should be formatted using Markdown.
	MarkdownDescription *string `json:"markdown_description,omitempty"`

	// DeprecationMessage defines warning diagnostic details to display when
	// practitioner configurations use this provider. The warning diagnostic
	// summary is automatically set to "Provider Deprecated" along with
	// configuration source file and line information.
	//
	// Set this field to a practitioner actionable message such as:
	//
	//  - "Use examplenewcloud provider instead."
	//  - "Remove this provider as it no longer is valid."
	//
	DeprecationMessage *string `json:"deprecation_message,omitempty"`
}

// SchemaValidateRequest specifies the provider being validated.
type SchemaValidateRequest struct {
	Path string
}

// Validate delegates to Attributes.Validate and Blocks.Validate.
func (s Schema) Validate(ctx context.Context, req SchemaValidateRequest) error {
	var errs []error

	attributeValidateRequest := AttributeValidateRequest(req)

	err := s.Attributes.Validate(ctx, attributeValidateRequest)

	if err != nil {
		errs = append(errs, err)
	}

	blockValidateRequest := BlockValidateRequest(req)

	err = s.Blocks.Validate(ctx, blockValidateRequest)

	if err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
