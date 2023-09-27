// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"errors"
	"fmt"
)

// ValidateRequest defines the Path of the resource that is
// being validated.
type ValidateRequest struct {
	Path string
}

// Resource defines an individual resource.
type Resource struct {
	// Name is the string identifier for the resource.
	Name string `json:"name"`

	// Schema defines the Attributes and Blocks for the data source.
	Schema *Schema `json:"schema,omitempty"`
}

// Validate delegates to Schema.Validate.
func (r Resource) Validate(ctx context.Context, req ValidateRequest) error {
	if r.Schema == nil {
		return nil
	}

	schemaValidateRequest := SchemaValidateRequest(req)

	return r.Schema.Validate(ctx, schemaValidateRequest)
}

// ResourcesValidateRequest defines the request sent during validation of Resources.
type ResourcesValidateRequest struct{}

// Resources type defines Resource types.
type Resources []Resource

// Validate checks for duplicated data source names and delegates to Resource.Validate
// for each data source.
func (rs Resources) Validate(ctx context.Context, req ResourcesValidateRequest) error {
	resourceNames := make(map[string]struct{}, len(rs))

	var errs, nestedErrs []error

	for _, r := range rs {
		if _, ok := resourceNames[r.Name]; ok {
			errs = append(errs, fmt.Errorf("resource %q is duplicated", r.Name))
		}

		resourceNames[r.Name] = struct{}{}

		validateRequest := ValidateRequest{
			Path: fmt.Sprintf("resource %q", r.Name),
		}

		err := r.Validate(ctx, validateRequest)

		if err != nil {
			nestedErrs = append(nestedErrs, err)
		}
	}

	e := append(errs, nestedErrs...)

	return errors.Join(e...)
}
