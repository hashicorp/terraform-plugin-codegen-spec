// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float64PlanModifiers type defines Float64PlanModifier types
type Float64PlanModifiers []Float64PlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each Float64PlanModifier.
func (v Float64PlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
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

// Equal returns true if the given Float64PlanModifiers is the same
// length, and each of the Float64PlanModifier entries is equal.
func (v Float64PlanModifiers) Equal(other Float64PlanModifiers) bool {
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

// Float64PlanModifier type defines type and function that provides plan modification
// functionality.
type Float64PlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Float64PlanModifier are equal.
func (v Float64PlanModifier) Equal(other Float64PlanModifier) bool {
	return v.Custom.Equal(other.Custom)

}
