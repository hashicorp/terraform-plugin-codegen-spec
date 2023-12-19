// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int64PlanModifiers type defines Int64PlanModifier types
type Int64PlanModifiers []Int64PlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each Int64PlanModifier.
func (v Int64PlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given Int64PlanModifiers is the same
// length, and each of the Int64PlanModifier entries is equal.
func (v Int64PlanModifiers) Equal(other Int64PlanModifiers) bool {
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

// Int64PlanModifier type defines type and function that provides plan modification
// functionality.
type Int64PlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Int64PlanModifier are equal.
func (v Int64PlanModifier) Equal(other Int64PlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
