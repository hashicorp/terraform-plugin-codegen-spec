// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// NumberPlanModifiers type defines NumberPlanModifier types
type NumberPlanModifiers []NumberPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each NumberPlanModifier.
func (v NumberPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given NumberPlanModifiers is the same
// length, and each of the NumberPlanModifier entries is equal.
func (v NumberPlanModifiers) Equal(other NumberPlanModifiers) bool {
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

// NumberPlanModifier type defines type and function that provides plan modification
// functionality.
type NumberPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given NumberPlanModifier are equal.
func (v NumberPlanModifier) Equal(other NumberPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
