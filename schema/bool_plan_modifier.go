// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// BoolPlanModifiers type defines BoolPlanModifier types
type BoolPlanModifiers []BoolPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each BoolPlanModifier.
func (v BoolPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given BoolPlanModifiers is the same
// length, and each of the BoolPlanModifier entries is equal.
func (v BoolPlanModifiers) Equal(other BoolPlanModifiers) bool {
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

// BoolPlanModifier type defines type and function that provides plan modification
// functionality.
type BoolPlanModifier struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given BoolPlanModifier are equal.
func (v BoolPlanModifier) Equal(other BoolPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
