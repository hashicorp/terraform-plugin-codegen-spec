// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// BoolValidators type is a slice of BoolValidator.
type BoolValidators []BoolValidator

// CustomValidators returns CustomValidator for each BoolValidator.
func (v BoolValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given BoolValidators is the same
// length, and each of the BoolValidator entries is equal.
func (v BoolValidators) Equal(other BoolValidators) bool {
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

// BoolValidator type defines type and function that provides validation
// functionality.
type BoolValidator struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given BoolValidator equal.
func (v BoolValidator) Equal(other BoolValidator) bool {
	return v.Custom.Equal(other.Custom)
}
