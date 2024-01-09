// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// SetPlanModifiers type defines SetPlanModifier types
type SetPlanModifiers []SetPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each SetPlanModifier.
func (v SetPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given SetPlanModifiers is the same
// length, and each of the SetPlanModifier entries is equal.
func (v SetPlanModifiers) Equal(other SetPlanModifiers) bool {
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

// SetPlanModifier type defines type and function that provides plan modification
// functionality.
type SetPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given SetPlanModifier are equal.
func (v SetPlanModifier) Equal(other SetPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
