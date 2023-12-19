// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ListValidators type defines ListValidator types
type ListValidators []ListValidator

// CustomValidators returns CustomValidator for each ListValidator.
func (v ListValidators) CustomValidators() CustomValidators {
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

// ListValidator type defines type and function that provides validation
// functionality.
type ListValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given ListValidator equal.
func (v ListValidator) Equal(other ListValidator) bool {
	return v.Custom.Equal(other.Custom)
}
