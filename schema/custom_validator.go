// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type CustomValidator struct {
	Imports          []string `json:"imports,omitempty"`
	SchemaDefinition string   `json:"schema_definition"`
}

func (c CustomValidator) HasImport() bool {
	return len(c.Imports) > 0
}
