package spec

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
)

type Specification struct {
	DataSources []datasource.DataSource `json:"datasources,omitempty"`
	Provider    *provider.Provider      `json:"provider,omitempty"`
	Resources   []resource.Resource     `json:"resources,omitempty"`
}
