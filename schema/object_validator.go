// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ObjectValidators type defines ObjectValidator types
type ObjectValidators []ObjectValidator

// CustomValidators returns CustomValidator for each ObjectValidator.
func (v ObjectValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given ObjectValidators is the same
// length, and each of the ObjectValidator entries is equal.
func (v ObjectValidators) Equal(other ObjectValidators) bool {
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

// ObjectValidator type defines type and function that provides validation
// functionality.
type ObjectValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given ObjectValidator equal.
func (v ObjectValidator) Equal(other ObjectValidator) bool {
	return v.Custom.Equal(other.Custom)
}
