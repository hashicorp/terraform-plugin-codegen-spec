// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestDynamicPlanModifiers_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		planModifiers schema.DynamicPlanModifiers
		other         schema.DynamicPlanModifiers
		expected      bool
	}{
		"plan_modifiers_both_nil": {
			expected: true,
		},
		"plan_modifiers_nil_other_not_nil": {
			other: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_not_nil_other_nil": {
			planModifiers: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{},
			},
			expected: false,
		},
		"plan_modifiers_len_diff": {
			planModifiers: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other:    schema.DynamicPlanModifiers{},
			expected: false,
		},
		"plan_modifiers_len_same": {
			planModifiers: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			other: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: true,
		},
		"plan_modifiers_len_same_with_custom_nils": {
			planModifiers: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{},
			},
			other: schema.DynamicPlanModifiers{
				schema.DynamicPlanModifier{
					Custom: &schema.CustomPlanModifier{},
				},
			},
			expected: false,
		},
		"plan_modifiers_schema_definition_same_order": {
			planModifiers: schema.DynamicPlanModifiers{
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
			other: schema.DynamicPlanModifiers{
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
			planModifiers: schema.DynamicPlanModifiers{
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
			other: schema.DynamicPlanModifiers{
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
