// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestBoolPlanModifiers_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		planModifiers schema.BoolPlanModifiers
		other         schema.BoolPlanModifiers
		expected      bool
	}{
		"plan_modifiers_both_nil": {
			expected: true,
		},
		"plan_modifiers_nil_other_not_nil": {
			other: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_not_nil_other_nil": {
			planModifiers: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_len_diff": {
			planModifiers: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other:    schema.BoolPlanModifiers{},
			expected: false,
		},
		"plan_modifiers_len_same": {
			planModifiers: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: true,
		},
		"plan_modifiers_len_same_with_custom_nils": {
			planModifiers: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{},
			},
			other: schema.BoolPlanModifiers{
				schema.BoolPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: false,
		},
		"plan_modifiers_schema_definition_same_order": {
			planModifiers: schema.BoolPlanModifiers{
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
			other: schema.BoolPlanModifiers{
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
			planModifiers: schema.BoolPlanModifiers{
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
			other: schema.BoolPlanModifiers{
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

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.planModifiers.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
