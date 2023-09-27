// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"errors"
	"fmt"
)

// ValidateRequest defines the Path of the data source that is
// being validated.
type ValidateRequest struct {
	Path string
}

// DataSource defines an individual data source.
type DataSource struct {
	// Name is the string identifier for the data source.
	Name string `json:"name"`

	// Schema defines the Attributes and Blocks for the data source.
	Schema *Schema `json:"schema,omitempty"`
}

// Validate delegates to Schema.Validate.
func (r DataSource) Validate(ctx context.Context, req ValidateRequest) error {
	if r.Schema == nil {
		return nil
	}

	schemaValidateRequest := SchemaValidateRequest(req)

	return r.Schema.Validate(ctx, schemaValidateRequest)
}

// DataSourcesValidateRequest defines the request sent during validation of DataSources.
type DataSourcesValidateRequest struct{}

// DataSources type defines DataSource types.
type DataSources []DataSource

// Validate checks for duplicated data source names and delegates to DataSource.Validate
// for each data source.
func (rs DataSources) Validate(ctx context.Context, req DataSourcesValidateRequest) error {
	datasourceNames := make(map[string]struct{}, len(rs))

	var errs, nestedErrs []error

	for _, r := range rs {
		if _, ok := datasourceNames[r.Name]; ok {
			errs = append(errs, fmt.Errorf("data source %q is duplicated", r.Name))
		}

		datasourceNames[r.Name] = struct{}{}

		validateRequest := ValidateRequest{
			Path: fmt.Sprintf("data source %q", r.Name),
		}

		err := r.Validate(ctx, validateRequest)

		if err != nil {
			nestedErrs = append(nestedErrs, err)
		}
	}

	e := append(errs, nestedErrs...)

	return errors.Join(e...)
}
