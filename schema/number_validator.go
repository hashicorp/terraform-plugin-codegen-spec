// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// NumberValidators type defines NumberValidator types
type NumberValidators []NumberValidator

// CustomValidators returns CustomValidator for each NumberValidator.
func (v NumberValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given NumberValidators is the same
// length, and each of the NumberValidator entries is equal.
func (v NumberValidators) Equal(other NumberValidators) bool {
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

// NumberValidator type defines type and function that provides validation
// functionality.
type NumberValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given NumberValidator equal.
func (v NumberValidator) Equal(other NumberValidator) bool {
	return v.Custom.Equal(other.Custom)
}
