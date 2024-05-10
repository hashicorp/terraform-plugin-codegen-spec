// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// DynamicPlanModifiers type defines DynamicPlanModifier types
type DynamicPlanModifiers []DynamicPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each DynamicPlanModifier.
func (v DynamicPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given DynamicPlanModifiers is the same
// length, and each of the DynamicPlanModifier entries is equal.
func (v DynamicPlanModifiers) Equal(other DynamicPlanModifiers) bool {
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

// DynamicPlanModifier type defines type and function that provides plan modification
// functionality.
type DynamicPlanModifier struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given DynamicPlanModifier are equal.
func (v DynamicPlanModifier) Equal(other DynamicPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
