package datasource

import (
	"context"
	"errors"
)

type Schema struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Blocks     Blocks     `json:"blocks,omitempty"`
}

type SchemaValidateRequest struct {
	Path string
}

// Validate delegates to Attributes.Validate and Blocks.Validate.
func (s Schema) Validate(ctx context.Context, req SchemaValidateRequest) error {
	var errs []error

	attributeValidateRequest := AttributeValidateRequest{
		Path: req.Path,
	}

	err := s.Attributes.Validate(ctx, attributeValidateRequest)

	if err != nil {
		errs = append(errs, err)
	}

	blockValidateRequest := BlockValidateRequest{
		Path: req.Path,
	}

	err = s.Blocks.Validate(ctx, blockValidateRequest)

	if err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
