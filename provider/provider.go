package provider

import (
	"context"
	"fmt"
)

type ValidateRequest struct {
	Path string
}

type Provider struct {
	Name string `json:"name"`

	Schema *Schema `json:"schema,omitempty"`
}

// Validate delegates to Schema.Validate.
func (r Provider) Validate(ctx context.Context, req ValidateRequest) error {
	if r.Schema == nil {
		return nil
	}

	schemaValidateRequest := SchemaValidateRequest{
		Path: fmt.Sprintf("provider %q", r.Name),
	}

	return r.Schema.Validate(ctx, schemaValidateRequest)
}
