// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type CustomPlanModifier struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}

func (c CustomPlanModifier) HasImport() bool {
	return c.Import != nil && *c.Import != ""
}
