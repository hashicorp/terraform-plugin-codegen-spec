// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// SetValidators type defines SetValidator types
type SetValidators []SetValidator

// CustomValidators returns CustomValidator for each SetValidator.
func (v SetValidators) CustomValidators() CustomValidators {
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

// SetValidator type defines type and function that provides validation
// functionality.
type SetValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given SetValidator equal.
func (v SetValidator) Equal(other SetValidator) bool {
	return v.Custom.Equal(other.Custom)
}
