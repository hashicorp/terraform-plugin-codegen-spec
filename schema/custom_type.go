// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type CustomType struct {
	Import    *string `json:"import,omitempty"`
	Type      string  `json:"type"`
	ValueType string  `json:"value_type"`
}

func (c CustomType) HasImport() bool {
	return c.Import != nil && *c.Import != ""
}
