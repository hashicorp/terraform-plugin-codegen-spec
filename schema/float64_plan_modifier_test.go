// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestFloat64PlanModifiers_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		planModifiers schema.Float64PlanModifiers
		other         schema.Float64PlanModifiers
		expected      bool
	}{
		"plan_modifiers_both_nil": {
			expected: true,
		},
		"plan_modifiers_nil_other_not_nil": {
			other: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_not_nil_other_nil": {
			planModifiers: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_len_diff": {
			planModifiers: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other:    schema.Float64PlanModifiers{},
			expected: false,
		},
		"plan_modifiers_len_same": {
			planModifiers: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: true,
		},
		"plan_modifiers_len_same_with_custom_nils": {
			planModifiers: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{},
			},
			other: schema.Float64PlanModifiers{
				schema.Float64PlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: false,
		},
		"plan_modifiers_schema_definition_same_order": {
			planModifiers: schema.Float64PlanModifiers{
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "two",
					},
				},
			},
			other: schema.Float64PlanModifiers{
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "two",
					},
				},
			},
			expected: true,
		},
		"plan_modifiers_schema_definition_different_order": {
			planModifiers: schema.Float64PlanModifiers{
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "two",
					},
				},
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "one",
					},
				},
			},
			other: schema.Float64PlanModifiers{
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomPlanModifier{
						SchemaDefinition: "two",
					},
				},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.planModifiers.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
