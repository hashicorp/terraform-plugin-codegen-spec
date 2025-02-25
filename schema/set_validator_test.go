// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestSetValidators_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		validators schema.SetValidators
		other      schema.SetValidators
		expected   bool
	}{
		"validators_both_nil": {
			expected: true,
		},
		"validators_nil_other_not_nil": {
			other: schema.SetValidators{
				schema.SetValidator{},
			},
			expected: false,
		},
		"validators_not_nil_other_nil": {
			validators: schema.SetValidators{
				schema.SetValidator{},
			},
			expected: false,
		},
		"validators_len_diff": {
			validators: schema.SetValidators{
				schema.SetValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			other:    schema.SetValidators{},
			expected: false,
		},
		"validators_len_same": {
			validators: schema.SetValidators{
				schema.SetValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			other: schema.SetValidators{
				schema.SetValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			expected: true,
		},
		"validators_len_same_with_custom_nils": {
			validators: schema.SetValidators{
				schema.SetValidator{},
			},
			other: schema.SetValidators{
				schema.SetValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			expected: false,
		},
		"validators_schema_definition_same_order": {
			validators: schema.SetValidators{
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "two",
					},
				},
			},
			other: schema.SetValidators{
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "two",
					},
				},
			},
			expected: true,
		},
		"validators_schema_definition_different_order": {
			validators: schema.SetValidators{
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "two",
					},
				},
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "one",
					},
				},
			},
			other: schema.SetValidators{
				{
					Custom: &schema.CustomValidator{
						SchemaDefinition: "one",
					},
				},
				{
					Custom: &schema.CustomValidator{
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

			got := testCase.validators.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
