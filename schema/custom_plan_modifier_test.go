// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestCustomPlanModifier_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType schema.CustomPlanModifier
		expected   bool
	}{
		"import-nil": {
			customType: schema.CustomPlanModifier{},
			expected:   false,
		},
		"import-empty": {
			customType: schema.CustomPlanModifier{
				Imports: []code.Import{}, // disallowed by spec, but still worth checking
			},
			expected: false,
		},
		"import-empty-string": {
			customType: schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Path: "", // disallowed by spec, but still worth checking
					},
				},
			},
			expected: true,
		},
		"import-string": {
			customType: schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Path: "github.com/owner/repo/pkg",
					},
				},
			},
			expected: true,
		},
		"import-strings": {
			customType: schema.CustomPlanModifier{
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

func TestCustomPlanModifier_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		planModifier *schema.CustomPlanModifier
		other        *schema.CustomPlanModifier
		expected     bool
	}{
		"plan_modifier_both_nil": {
			expected: true,
		},
		"plan_modifier_nil_other_not_nil": {
			other:    &schema.CustomPlanModifier{},
			expected: false,
		},
		"plan_modifier_imports_nil_other_not_nil": {
			planModifier: &schema.CustomPlanModifier{},
			other: &schema.CustomPlanModifier{
				Imports: []code.Import{},
			},
			expected: false,
		},
		"plan_modifier_imports_not_nil_other_nil": {
			planModifier: &schema.CustomPlanModifier{
				Imports: []code.Import{},
			},
			other:    &schema.CustomPlanModifier{},
			expected: false,
		},
		"plan_modifier_imports_alias_nil_other_not_nil": {
			planModifier: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			expected: false,
		},
		"plan_modifier_imports_alias_not_nil_other_nil": {
			planModifier: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Alias: pointer("alias"),
					},
				},
			},
			other: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"plan_modifier_imports_path_empty_other_not_empty": {
			planModifier: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{},
				},
			},
			other: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			expected: false,
		},
		"plan_modifier_imports_path_not_empty_other_empty": {
			planModifier: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{
						Path: "path",
					},
				},
			},
			other: &schema.CustomPlanModifier{
				Imports: []code.Import{
					{},
				},
			},
			expected: false,
		},
		"plan_modifier_imports_same_order": {
			planModifier: &schema.CustomPlanModifier{
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
			other: &schema.CustomPlanModifier{
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
		"plan_modifier_imports_different_order": {
			planModifier: &schema.CustomPlanModifier{
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
			other: &schema.CustomPlanModifier{
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

			got := testCase.planModifier.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomPlanModifiers_Sort(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customPlanModifiers schema.CustomPlanModifiers
		expected            schema.CustomPlanModifiers
	}{
		"custom-validators-nil": {
			expected: nil,
		},
		"custom-validators-nil-entry": {
			customPlanModifiers: schema.CustomPlanModifiers{nil},
			expected:            schema.CustomPlanModifiers{nil},
		},
		"custom-validators-nil-entries": {
			customPlanModifiers: schema.CustomPlanModifiers{nil, nil},
			expected:            schema.CustomPlanModifiers{nil, nil},
		},
		"custom-validators-non-nil-with-nil-entry": {
			customPlanModifiers: schema.CustomPlanModifiers{&schema.CustomPlanModifier{}, nil},
			expected:            schema.CustomPlanModifiers{&schema.CustomPlanModifier{}, nil},
		},
		"custom-validators-nil-with-non-nil-entry": {
			customPlanModifiers: schema.CustomPlanModifiers{nil, &schema.CustomPlanModifier{}},
			expected:            schema.CustomPlanModifiers{&schema.CustomPlanModifier{}, nil},
		},
		"custom-validators-non-nil-entries-sorted": {
			customPlanModifiers: schema.CustomPlanModifiers{&schema.CustomPlanModifier{SchemaDefinition: "x"}, &schema.CustomPlanModifier{SchemaDefinition: "y"}},
			expected:            schema.CustomPlanModifiers{&schema.CustomPlanModifier{SchemaDefinition: "x"}, &schema.CustomPlanModifier{SchemaDefinition: "y"}},
		},
		"custom-validators-non-nil-entries-unsorted": {
			customPlanModifiers: schema.CustomPlanModifiers{&schema.CustomPlanModifier{SchemaDefinition: "y"}, &schema.CustomPlanModifier{SchemaDefinition: "x"}},
			expected:            schema.CustomPlanModifiers{&schema.CustomPlanModifier{SchemaDefinition: "x"}, &schema.CustomPlanModifier{SchemaDefinition: "y"}},
		},
		"custom-validators-multiple-entries": {
			customPlanModifiers: schema.CustomPlanModifiers{nil, &schema.CustomPlanModifier{SchemaDefinition: "y"}, &schema.CustomPlanModifier{SchemaDefinition: "z"}, nil, &schema.CustomPlanModifier{SchemaDefinition: "x"}},
			expected:            schema.CustomPlanModifiers{&schema.CustomPlanModifier{SchemaDefinition: "x"}, &schema.CustomPlanModifier{SchemaDefinition: "y"}, &schema.CustomPlanModifier{SchemaDefinition: "z"}, nil, nil},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testCase.customPlanModifiers.Sort()

			if diff := cmp.Diff(testCase.customPlanModifiers, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
