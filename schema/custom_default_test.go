// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestCustomDefault_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType schema.CustomDefault
		expected   bool
	}{
		"import-nil": {
			customType: schema.CustomDefault{},
			expected:   false,
		},
		"import-empty": {
			customType: schema.CustomDefault{
				Imports: []code.Import{}, // disallowed by spec, but still worth checking
			},
			expected: false,
		},
		"import-empty-string": {
			customType: schema.CustomDefault{
				Imports: []code.Import{
					{
						Path: "", // disallowed by spec, but still worth checking
					},
				},
			},
			expected: true,
		},
		"import-string": {
			customType: schema.CustomDefault{
				Imports: []code.Import{
					{
						Path: "github.com/owner/repo/pkg",
					},
				},
			},
			expected: true,
		},
		"import-strings": {
			customType: schema.CustomDefault{
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

			got := testCase.customType.HasImport()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customDefault *schema.CustomDefault
		other         *schema.CustomDefault
		expected      bool
	}{
		"custom_default_both_nil": {
			expected: true,
		},
		"custom_default_nil_other_not_nil": {
			other:    &schema.CustomDefault{},
			expected: false,
		},
		"custom_default_imports_nil_other_not_nil": {
			customDefault: &schema.CustomDefault{},
			other: &schema.CustomDefault{
				Imports: []code.Import{},
			},
			expected: false,
		},
		"custom_default_imports_not_nil_other_nil": {
			customDefault: &schema.CustomDefault{
				Imports: []code.Import{},
			},
			other:    &schema.CustomDefault{},
			expected: false,
		},
		"custom_default_imports_alias_nil_other_not_nil": {
			customDefault: &schema.CustomDefault{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomDefault{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			expected: false,
		},
		"custom_default_imports_alias_not_nil_other_nil": {
			customDefault: &schema.CustomDefault{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			other: &schema.CustomDefault{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"custom_default_imports_path_empty_other_not_empty": {
			customDefault: &schema.CustomDefault{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomDefault{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			expected: false,
		},
		"custom_default_imports_path_not_empty_other_empty": {
			customDefault: &schema.CustomDefault{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			other: &schema.CustomDefault{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"custom_default_imports_same_order": {
			customDefault: &schema.CustomDefault{
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
			other: &schema.CustomDefault{
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
		"custom_default_imports_different_order": {
			customDefault: &schema.CustomDefault{
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
			other: &schema.CustomDefault{
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

			got := testCase.customDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
