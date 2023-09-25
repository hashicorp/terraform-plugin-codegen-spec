// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
)

// ValidateRequest defines the Path of the provider that is
// being validated.
type ValidateRequest struct {
	Path string
}

// Provider defines an individual provider.
type Provider struct {
	// Name is the string identifier for the provider.
	Name string `json:"name"`

	// Schema defines the Attributes and Blocks for the provider.
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
