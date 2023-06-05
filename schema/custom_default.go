// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type CustomDefault struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}

func (c CustomDefault) HasImport() bool {
	return c.Import != nil && *c.Import != ""
}
