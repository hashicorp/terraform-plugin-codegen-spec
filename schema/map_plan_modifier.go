// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// MapPlanModifiers type defines MapPlanModifier types
type MapPlanModifiers []MapPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each MapPlanModifier.
func (v MapPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given MapPlanModifiers is the same
// length, and each of the MapPlanModifier entries is equal.
func (v MapPlanModifiers) Equal(other MapPlanModifiers) bool {
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

// MapPlanModifier type defines type and function that provides plan modification
// functionality.
type MapPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given MapPlanModifier are equal.
func (v MapPlanModifier) Equal(other MapPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
