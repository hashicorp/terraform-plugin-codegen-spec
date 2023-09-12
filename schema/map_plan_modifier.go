// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

type MapPlanModifiers []MapPlanModifier

// Equal returns true if the given MapPlanModifiers is the same
// length, and after sorting and removal of any nil entries,
// is the same length, and each of the MapPlanModifier entries is
// equal.
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

	var planModifiers MapPlanModifiers

	var otherPlanModifiers MapPlanModifiers

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

type MapPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the given MapPlanModifier.Custom field
// is equal.
func (v MapPlanModifier) Equal(other MapPlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
