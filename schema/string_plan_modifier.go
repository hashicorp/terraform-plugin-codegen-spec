// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// StringPlanModifiers type defines StringPlanModifier types
type StringPlanModifiers []StringPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each StringPlanModifier.
func (v StringPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
	var customPlanModifiers CustomPlanModifiers

	for _, planModifier := range v {
		customPlanModifier := planModifier.Custom

		if customPlanModifier == nil {
			continue
		}

		customPlanModifiers = append(customPlanModifiers, customPlanModifier)
	}

	return customPlanModifiers
}

// Equal returns true if the given StringPlanModifiers is the same
// length, and each of the StringPlanModifier entries is equal.
func (v StringPlanModifiers) Equal(other StringPlanModifiers) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	planModifiers := v.CustomPlanModifiers()

	otherPlanModifiers := other.CustomPlanModifiers()

	if len(planModifiers) != len(otherPlanModifiers) {
		return false
	}

	planModifiers.Sort()

	otherPlanModifiers.Sort()

	for k, planModifier := range planModifiers {
		if !planModifier.Equal(otherPlanModifiers[k]) {
			return false
		}
	}

	return true
}

// StringPlanModifier type defines type and function that provides plan modification
// functionality.
type StringPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given StringPlanModifier are equal.
func (v StringPlanModifier) Equal(other StringPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
