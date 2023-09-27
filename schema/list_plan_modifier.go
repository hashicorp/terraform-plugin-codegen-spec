// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

// ListPlanModifiers type defines ListPlanModifier types
type ListPlanModifiers []ListPlanModifier

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

	var planModifiers ListPlanModifiers

	var otherPlanModifiers ListPlanModifiers

	// Remove nils otherwise sort will panic.
	for _, planModifier := range v {
		if planModifier.Custom != nil {
			planModifiers = append(planModifiers, planModifier)
		}
	}

	// Remove nils otherwise sort will panic.
	for _, planModifier := range other {
		if planModifier.Custom != nil {
			otherPlanModifiers = append(otherPlanModifiers, planModifier)
		}
	}

	// Compare length after removing nils.
	if len(planModifiers) != len(otherPlanModifiers) {
		return false
	}

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(planModifiers, func(i, j int) bool {
		return planModifiers[i].Custom.SchemaDefinition < planModifiers[j].Custom.SchemaDefinition
	})

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(otherPlanModifiers, func(i, j int) bool {
		return otherPlanModifiers[i].Custom.SchemaDefinition < otherPlanModifiers[j].Custom.SchemaDefinition
	})

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
