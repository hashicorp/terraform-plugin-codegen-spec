// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float64Validators type defines Float64Validator types
type Float64Validators []Float64Validator

// CustomValidators returns CustomValidator for each Float64Validator.
func (v Float64Validators) CustomValidators() CustomValidators {
	var customValidators CustomValidators

	for _, validator := range v {
		customValidator := validator.Custom

		if customValidator == nil {
			continue
		}

		customValidators = append(customValidators, customValidator)
	}

	return customValidators
}

// Equal returns true if the given Float64Validators is the same
// length, and each of the Float64Validator entries is equal.
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

	validators := v.CustomValidators()

	otherValidators := other.CustomValidators()

	if len(validators) != len(otherValidators) {
		return false
	}

	validators.Sort()

	otherValidators.Sort()

	for k, validator := range validators {
		if !validator.Equal(otherValidators[k]) {
			return false
		}
	}

	return true
}

// Float64Validator type defines type and function that provides validation
// functionality.
type Float64Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Float64Validator equal.
func (v Float64Validator) Equal(other Float64Validator) bool {
	return v.Custom.Equal(other.Custom)
}
