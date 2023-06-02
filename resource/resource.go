// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

type Resource struct {
	Name string `json:"name"`

	Schema *Schema `json:"schema,omitempty"`
}
