// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ObjectPlanModifiers type defines ObjectPlanModifier types
type ObjectPlanModifiers []ObjectPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each ObjectPlanModifier.
func (v ObjectPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given ObjectPlanModifiers is the same
// length, and each of the ObjectPlanModifier entries is equal.
func (v ObjectPlanModifiers) Equal(other ObjectPlanModifiers) bool {
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

// ObjectPlanModifier type defines type and function that provides plan modification
// functionality.
type ObjectPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given ObjectPlanModifier are equal.
func (v ObjectPlanModifier) Equal(other ObjectPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
