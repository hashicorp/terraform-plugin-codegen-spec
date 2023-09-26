// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

// CustomType defines a custom type, and value type for a schema type,
// and associated value.
type CustomType struct {
	// Import defines a path, and optional alias for imported code.
	Import *code.Import `json:"import,omitempty"`

	// Type defines the type for use in the schema.
	Type string `json:"type"`

	// ValueType defines the type for use with the value associated
	// with the schema type.
	ValueType string `json:"value_type"`
}

// HasImport returns true if the CustomType has a non-empty import path.
func (c *CustomType) HasImport() bool {
	return c.Import != nil && c.Import.Path != ""
}

// Equal returns true if all fields of the given CustomType are equal.
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
		if c.Import.Alias == nil && other.Import.Alias != nil {
			return false
		}

		if c.Import.Alias != nil && other.Import.Alias == nil {
			return false
		}

		if c.Import.Alias != nil && other.Import.Alias != nil {
			if *c.Import.Alias != *other.Import.Alias {
				return false
			}
		}

		if c.Import.Path != other.Import.Path {
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
