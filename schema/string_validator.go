// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// StringValidators type defines StringValidator types
type StringValidators []StringValidator

// CustomValidators returns CustomValidator for each StringValidator.
func (v StringValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given StringValidators is the same
// length, and each of the StringValidator entries is equal.
func (v StringValidators) Equal(other StringValidators) bool {
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

// StringValidator type defines type and function that provides validation
// functionality.
type StringValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given StringValidator equal.
func (v StringValidator) Equal(other StringValidator) bool {
	return v.Custom.Equal(other.Custom)
}
