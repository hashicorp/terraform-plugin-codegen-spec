// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestBoolValidators_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		validators schema.BoolValidators
		other      schema.BoolValidators
		expected   bool
	}{
		"validators_both_nil": {
			expected: true,
		},
		"validators_nil_other_not_nil": {
			other: schema.BoolValidators{
				schema.BoolValidator{},
			},
			expected: false,
		},
		"validators_not_nil_other_nil": {
			validators: schema.BoolValidators{
				schema.BoolValidator{},
			},
			expected: false,
		},
		"validators_len_diff": {
			validators: schema.BoolValidators{
				schema.BoolValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			other:    schema.BoolValidators{},
			expected: false,
		},
		"validators_len_same": {
			validators: schema.BoolValidators{
				schema.BoolValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			other: schema.BoolValidators{
				schema.BoolValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			expected: true,
		},
		"validators_len_same_with_custom_nils": {
			validators: schema.BoolValidators{
				schema.BoolValidator{},
			},
			other: schema.BoolValidators{
				schema.BoolValidator{
					Custom: &schema.CustomValidator{},
				},
			},
			expected: false,
		},
		"validators_schema_definition_same_order": {
			validators: schema.BoolValidators{
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
			other: schema.BoolValidators{
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
			validators: schema.BoolValidators{
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
			other: schema.BoolValidators{
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
