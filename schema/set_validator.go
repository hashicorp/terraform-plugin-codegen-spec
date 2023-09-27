// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

// SetValidators type defines SetValidator types
type SetValidators []SetValidator

// Equal returns true if the given SetValidators is the same
// length, and each of the SetValidator entries is equal.
func (v SetValidators) Equal(other SetValidators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var validators SetValidators

	var otherValidators SetValidators

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

// SetValidator type defines type and function that provides validation
// functionality.
type SetValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given SetValidator equal.
func (v SetValidator) Equal(other SetValidator) bool {
	return v.Custom.Equal(other.Custom)
}
