// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"errors"
)

// Schema defines the Attributes and Blocks associated with a DataSource.
type Schema struct {
	// Attributes defines the Attribute types for the Schema..
	Attributes Attributes `json:"attributes,omitempty"`

	// Blocks defines the Block types for the Schema.
	Blocks Blocks `json:"blocks,omitempty"`
}

// SchemaValidateRequest specifies the data source being validated.
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
