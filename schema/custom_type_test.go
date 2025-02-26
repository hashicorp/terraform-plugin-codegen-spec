// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestCustomType_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType *schema.CustomType
		expected   bool
	}{
		"import-nil": {
			customType: &schema.CustomType{},
			expected:   false,
		},
		"import-empty-string": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Path: "",
				},
			},
			expected: false,
		},
		"import-string": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Path: "github.com/owner/repo/pkg",
				},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.customType.HasImport()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCustomType_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType *schema.CustomType
		other      *schema.CustomType
		expected   bool
	}{
		"nil_other_not_nil": {
			other:    &schema.CustomType{},
			expected: false,
		},
		"not_nil_other_nil": {
			customType: &schema.CustomType{},
			expected:   false,
		},
		"import_nil_other_import_not_nil": {
			customType: &schema.CustomType{},
			other: &schema.CustomType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_not_nil_other_import_nil": {
			customType: &schema.CustomType{
				Import: &code.Import{},
			},
			other:    &schema.CustomType{},
			expected: false,
		},
		"import_alias_nil_other_import_alias_not_nil": {
			customType: &schema.CustomType{
				Import: &code.Import{},
			},
			other: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
				},
			},
			expected: false,
		},
		"import_alias_not_nil_other_import_alias_nil": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
				},
			},
			other: &schema.CustomType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_path_empty_other_import_path_not_empty": {
			customType: &schema.CustomType{
				Import: &code.Import{},
			},
			other: &schema.CustomType{
				Import: &code.Import{
					Path: "path",
				},
			},
			expected: false,
		},
		"import_path_not_empty_other_import_path_empty": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Path: "path",
				},
			},
			other: &schema.CustomType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_match": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
			},
			other: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
			},
			expected: true,
		},
		"type_empty_other_type_not_empty": {
			customType: &schema.CustomType{},
			other: &schema.CustomType{
				Type: "type",
			},
			expected: false,
		},
		"type_not_empty_other_type_empty": {
			customType: &schema.CustomType{
				Type: "type",
			},
			other:    &schema.CustomType{},
			expected: false,
		},
		"type_match": {
			customType: &schema.CustomType{
				Type: "type",
			},
			other: &schema.CustomType{
				Type: "type",
			},
			expected: true,
		},
		"value_type_empty_other_value_type_not_empty": {
			customType: &schema.CustomType{},
			other: &schema.CustomType{
				ValueType: "valueType",
			},
			expected: false,
		},
		"value_type_not_empty_other_value_type_empty": {
			customType: &schema.CustomType{
				ValueType: "valueType",
			},
			other:    &schema.CustomType{},
			expected: false,
		},
		"value_type_match": {
			customType: &schema.CustomType{
				ValueType: "valueType",
			},
			other: &schema.CustomType{
				ValueType: "valueType",
			},
			expected: true,
		},
		"match": {
			customType: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
				Type:      "type",
				ValueType: "valueType",
			},
			other: &schema.CustomType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
				Type:      "type",
				ValueType: "valueType",
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.customType.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
