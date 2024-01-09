// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ListPlanModifiers type defines ListPlanModifier types
type ListPlanModifiers []ListPlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each ListPlanModifier.
func (v ListPlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given ListPlanModifiers is the same
// length, and each of the ListPlanModifier entries is equal.
func (v ListPlanModifiers) Equal(other ListPlanModifiers) bool {
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

// ListPlanModifier type defines type and function that provides plan modification
// functionality.
type ListPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given ListPlanModifier are equal.
func (v ListPlanModifier) Equal(other ListPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
