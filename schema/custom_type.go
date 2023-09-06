// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

type CustomType struct {
	Import    *code.Import `json:"import,omitempty"`
	Type      string       `json:"type"`
	ValueType string       `json:"value_type"`
}

func (c *CustomType) HasImport() bool {
	return c.Import != nil && c.Import.Path != ""
}

func (c *CustomType) Equal(other *CustomType) bool {
	if c == nil && other == nil {
		return true
	}

	if c == nil && other != nil {
		return false
	}

	if c != nil && other == nil {
		return false
	}

	if c.Import == nil && other.Import != nil {
		return false
	}

	if c.Import != nil && other.Import == nil {
		return false
	}

	if c.Import != nil && other.Import != nil {
		cImport := *c.Import
		otherImport := *other.Import

		if cImport.Alias == nil && otherImport.Alias != nil {
			return false
		}

		if cImport.Alias != nil && otherImport.Alias == nil {
			return false
		}

		if cImport.Alias != nil && otherImport.Alias != nil {
			if *cImport.Alias != *otherImport.Alias {
				return false
			}
		}

		if cImport.Path != otherImport.Path {
			return false
		}
	}

	if c.Type != other.Type {
		return false
	}

	if c.ValueType != other.ValueType {
		return false
	}

	return true
}
