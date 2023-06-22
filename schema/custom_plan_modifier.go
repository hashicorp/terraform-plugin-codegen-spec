// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type CustomPlanModifier struct {
	Imports          []string `json:"imports,omitempty"`
	SchemaDefinition string   `json:"schema_definition"`
}

func (c CustomPlanModifier) HasImport() bool {
	return len(c.Imports) > 0
}
