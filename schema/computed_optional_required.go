// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

const (
	Computed         ComputedOptionalRequired = "computed"
	ComputedOptional ComputedOptionalRequired = "computed_optional"
	Optional         ComputedOptionalRequired = "optional"
	Required         ComputedOptionalRequired = "required"
)

type ComputedOptionalRequired string

func (c ComputedOptionalRequired) Equal(other ComputedOptionalRequired) bool {
	return c == other
}

type OptionalRequired = ComputedOptionalRequired
