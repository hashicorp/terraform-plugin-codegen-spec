// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

type CustomPlanModifier struct {
	Imports          []code.Import `json:"imports,omitempty"`
	SchemaDefinition string        `json:"schema_definition"`
}

func (c CustomPlanModifier) HasImport() bool {
	return len(c.Imports) > 0
}
