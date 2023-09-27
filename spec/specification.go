// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
)

// Specification defines the data source(s), provider, and resource(s) for
// a [Terraform provider].
//
// [Terraform provider]: https://developer.hashicorp.com/terraform/language/providers
type Specification struct {
	// DataSources defines a slice of datasource.DataSource type.
	DataSources datasource.DataSources `json:"datasources,omitempty"`

	// Provider defines an instance of the provider.Provider type.
	Provider *provider.Provider `json:"provider,omitempty"`

	// Resources defines a slice of resource.Resource type.
	Resources resource.Resources `json:"resources,omitempty"`

	// Version defines the Provider Code Specification JSON schema version
	Version *string `json:"version,omitempty"`
}

// Validate delegates validation to each of datasource.DataSources,
// *provider.Provider and resource.Resources.
func (s Specification) Validate(ctx context.Context) error {
	var errs []error

	datasourcesValidateReq := datasource.DataSourcesValidateRequest{}

	err := s.DataSources.Validate(ctx, datasourcesValidateReq)

	if err != nil {
		errs = append(errs, err)
	}

	if s.Provider != nil {
		providerValidateReq := provider.ValidateRequest{}

		err = s.Provider.Validate(ctx, providerValidateReq)

		if err != nil {
			errs = append(errs, err)
		}
	}

	resourcesValidateReq := resource.ResourcesValidateRequest{}

	err = s.Resources.Validate(ctx, resourcesValidateReq)

	if err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
