// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// MapValidators type defines MapValidator types
type MapValidators []MapValidator

// CustomValidators returns CustomValidator for each MapValidator.
func (v MapValidators) CustomValidators() CustomValidators {
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

// Equal returns true if the given MapValidators is the same
// length, and each of the MapValidator entries is equal.
func (v MapValidators) Equal(other MapValidators) bool {
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

// MapValidator type defines type and function that provides validation
// functionality.
type MapValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given MapValidator equal.
func (v MapValidator) Equal(other MapValidator) bool {
	return v.Custom.Equal(other.Custom)
}
