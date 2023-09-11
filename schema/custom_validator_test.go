// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestCustomValidator_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customValidator *schema.CustomValidator
		expected        bool
	}{
		"import-nil": {
			customValidator: &schema.CustomValidator{},
			expected:        false,
		},
		"import-empty": {
			customValidator: &schema.CustomValidator{
				Imports: []code.Import{}, // disallowed by spec, but still worth checking
			},
			expected: false,
		},
		"import-empty-string": {
			customValidator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "", // disallowed by spec, but still worth checking
					},
				},
			},
			expected: true,
		},
		"import-string": {
			customValidator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/owner/repo/pkg",
					},
				},
			},
			expected: true,
		},
		"import-strings": {
			customValidator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/owner/repo/pkg1",
					},
					{
						Path: "github.com/owner/repo/pkg2",
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

			got := testCase.customValidator.HasImport()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomValidator_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		validator *schema.CustomValidator
		other     *schema.CustomValidator
		expected  bool
	}{
		"validator_both_nil": {
			expected: true,
		},
		"validator_nil_other_not_nil": {
			other:    &schema.CustomValidator{},
			expected: false,
		},
		"validator_imports_nil_other_not_nil": {
			validator: &schema.CustomValidator{},
			other: &schema.CustomValidator{
				Imports: []code.Import{},
			},
			expected: false,
		},
		"validator_imports_not_nil_other_nil": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{},
			},
			other:    &schema.CustomValidator{},
			expected: false,
		},
		"validator_imports_alias_nil_other_not_nil": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			expected: false,
		},
		"validator_imports_alias_not_nil_other_nil": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"validator_imports_path_empty_other_not_empty": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			expected: false,
		},
		"validator_imports_path_not_empty_other_empty": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"validator_imports_same_order": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("one"),
						Path:  "one",
					},
					{
						Alias: pointer("two"),
						Path:  "two",
					},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("one"),
						Path:  "one",
					},
					{
						Alias: pointer("two"),
						Path:  "two",
					},
				},
			},
			expected: true,
		},
		"validator_imports_different_order": {
			validator: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("one"),
						Path:  "one",
					},
					{
						Alias: pointer("two"),
						Path:  "two",
					},
				},
			},
			other: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Alias: pointer("two"),
						Path:  "two",
					},
					{
						Alias: pointer("one"),
						Path:  "one",
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

			got := testCase.validator.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
