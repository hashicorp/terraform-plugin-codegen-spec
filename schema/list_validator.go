// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

// ListValidators type defines ListValidator types
type ListValidators []ListValidator

// Equal returns true if the given ListValidators is the same
// length, and each of the ListValidator entries is equal.
func (v ListValidators) Equal(other ListValidators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var validators ListValidators

	var otherValidators ListValidators

	// Remove nils otherwise sort will panic.
	for _, validator := range v {
		if validator.Custom != nil {
			validators = append(validators, validator)
		}
	}

	// Remove nils otherwise sort will panic.
	for _, validator := range other {
		if validator.Custom != nil {
			otherValidators = append(otherValidators, validator)
		}
	}

	// Compare length after removing nils.
	if len(validators) != len(otherValidators) {
		return false
	}

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].Custom.SchemaDefinition < validators[j].Custom.SchemaDefinition
	})

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(otherValidators, func(i, j int) bool {
		return otherValidators[i].Custom.SchemaDefinition < otherValidators[j].Custom.SchemaDefinition
	})

	for k, validator := range validators {
		if !validator.Equal(otherValidators[k]) {
			return false
		}
	}

	return true
}

// ListValidator type defines type and function that provides validation
// functionality.
type ListValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given ListValidator equal.
func (v ListValidator) Equal(other ListValidator) bool {
	return v.Custom.Equal(other.Custom)
}
