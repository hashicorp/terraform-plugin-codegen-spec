// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "github.com/hashicorp/terraform-plugin-codegen-spec/code"

type AssociatedExternalType struct {
	Import *code.Import `json:"import,omitempty"`
	Type   string       `json:"type"`
}

func (a AssociatedExternalType) HasImport() bool {
	return a.Import != nil && a.Import.Import != ""
}
