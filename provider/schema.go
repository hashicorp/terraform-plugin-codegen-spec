// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

type Schema struct {
	Attributes []Attribute `json:"attributes,omitempty"`
	Blocks     []Block     `json:"blocks,omitempty"`
}
