// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

type Float64Validators []Float64Validator

// Equal returns true if the given Float64Validators is the same
// length, and after sorting and removal of any nil entries,
// is the same length, and each of the Float64Validator entries is
// equal.
func (v Float64Validators) Equal(other Float64Validators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var validators Float64Validators

	var otherValidators Float64Validators

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

type Float64Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the given Float64Validator.Custom field
// is equal.
func (v Float64Validator) Equal(other Float64Validator) bool {
	return v.Custom.Equal(other.Custom)
}
