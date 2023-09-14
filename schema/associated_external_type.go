// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

type AssociatedExternalType struct {
	Import *code.Import `json:"import,omitempty"`
	Type   string       `json:"type"`
}

func (a *AssociatedExternalType) HasImport() bool {
	return a.Import != nil && a.Import.Path != ""
}

func (a *AssociatedExternalType) Equal(other *AssociatedExternalType) bool {
	if a == nil && other == nil {
		return true
	}

	if a == nil && other != nil {
		return false
	}

	if a != nil && other == nil {
		return false
	}

	if a.Import == nil && other.Import != nil {
		return false
	}

	if a.Import != nil && other.Import == nil {
		return false
	}

	if a.Import != nil && other.Import != nil {
		if a.Import.Alias == nil && other.Import.Alias != nil {
			return false
		}

		if a.Import.Alias != nil && other.Import.Alias == nil {
			return false
		}

		if a.Import.Alias != nil && other.Import.Alias != nil {
			if *a.Import.Alias != *other.Import.Alias {
				return false
			}
		}

		if a.Import.Path != other.Import.Path {
			return false
		}
	}

	return a.Type == other.Type
}
