// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

type CustomDefault struct {
	Imports          []code.Import `json:"imports,omitempty"`
	SchemaDefinition string        `json:"schema_definition"`
}

func (c CustomDefault) HasImport() bool {
	return len(c.Imports) > 0
}
