// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
)

type CustomDefault struct {
	Imports          []code.Import `json:"imports,omitempty"`
	SchemaDefinition string        `json:"schema_definition"`
}

func (c *CustomDefault) HasImport() bool {
	return len(c.Imports) > 0
}

func (c *CustomDefault) Equal(other *CustomDefault) bool {
	if c == nil && other == nil {
		return true
	}

	if c == nil || other == nil {
		return false
	}

	if c.Imports == nil && other.Imports != nil {
		return false
	}

	if c.Imports != nil && other.Imports == nil {
		return false
	}

	if len(c.Imports) != len(other.Imports) {
		return false
	}

	// Path is required by the spec JSON schema.
	sort.Slice(c.Imports, func(i, j int) bool {
		return c.Imports[i].Path < c.Imports[j].Path
	})

	// Path is required by the spec JSON schema.
	sort.Slice(other.Imports, func(i, j int) bool {
		return other.Imports[i].Path < other.Imports[j].Path
	})

	for k, v := range c.Imports {
		if v.Path != other.Imports[k].Path {
			return false
		}

		if v.Alias == nil && other.Imports[k].Alias == nil {
			continue
		}

		if v.Alias == nil || other.Imports[k].Alias == nil {
			return false
		}

		if *v.Alias != *other.Imports[k].Alias {
			return false
		}
	}

	return c.SchemaDefinition == other.SchemaDefinition
}
