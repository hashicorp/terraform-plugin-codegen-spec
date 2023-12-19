// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int64Validators type defines Int64Validator types
type Int64Validators []Int64Validator

// CustomValidators returns CustomValidator for each Float64Validator.
func (v Int64Validators) CustomValidators() CustomValidators {
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

// Equal returns true if the given Int64Validators is the same
// length, and each of the Int64Validator entries is equal.
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

// Int64Validator type defines type and function that provides validation
// functionality.
type Int64Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Int64Validator equal.
func (v Int64Validator) Equal(other Int64Validator) bool {
	return v.Custom.Equal(other.Custom)
}
