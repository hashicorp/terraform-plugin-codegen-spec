// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// DynamicValidators type is a slice of DynamicValidator.
type DynamicValidators []DynamicValidator

// CustomValidators returns CustomValidator for each DynamicValidator.
func (v DynamicValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given DynamicValidators is the same
// length, and each of the DynamicValidator entries is equal.
func (v DynamicValidators) Equal(other DynamicValidators) bool {
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

// DynamicValidator type defines type and function that provides validation
// functionality.
type DynamicValidator struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given DynamicValidator equal.
func (v DynamicValidator) Equal(other DynamicValidator) bool {
	return v.Custom.Equal(other.Custom)
}
