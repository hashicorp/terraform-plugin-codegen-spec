// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

type Int64Validators []Int64Validator

// Equal returns true if the given Int64Validators is the same
// length, and after sorting and removal of any nil entries,
// is the same length, and each of the Int64Validator entries is
// equal.
func (v Int64Validators) Equal(other Int64Validators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var validators Int64Validators

	var otherValidators Int64Validators

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

type Int64Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the given Int64Validator.Custom field
// is equal.
func (v Int64Validator) Equal(other Int64Validator) bool {
	return v.Custom.Equal(other.Custom)
}
