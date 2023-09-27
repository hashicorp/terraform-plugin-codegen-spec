// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

// StringPlanModifiers type defines StringPlanModifier types
type StringPlanModifiers []StringPlanModifier

// Equal returns true if the given StringPlanModifiers is the same
// length, and each of the StringPlanModifier entries is equal.
func (v StringPlanModifiers) Equal(other StringPlanModifiers) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var planModifiers StringPlanModifiers

	var otherPlanModifiers StringPlanModifiers

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

// StringPlanModifier type defines type and function that provides plan modification
// functionality.
type StringPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given StringPlanModifier are equal.
func (v StringPlanModifier) Equal(other StringPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
