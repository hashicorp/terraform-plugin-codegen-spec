package spec

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
)

type Specification struct {
	DataSources datasource.DataSources `json:"datasources,omitempty"`
	Provider    *provider.Provider     `json:"provider,omitempty"`
	Resources   resource.Resources     `json:"resources,omitempty"`
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
