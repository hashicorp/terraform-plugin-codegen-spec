// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestAssociatedExternalType_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType schema.AssociatedExternalType
		expected   bool
	}{
		"import-nil": {
			customType: schema.AssociatedExternalType{},
			expected:   false,
		},
		"import-empty-string": {
			customType: schema.AssociatedExternalType{
				Import: &code.Import{
					Path: "",
				},
			},
			expected: false,
		},
		"import-string": {
			customType: schema.AssociatedExternalType{
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

func TestAssociatedExternalType_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		assocExtType *schema.AssociatedExternalType
		other        *schema.AssociatedExternalType
		expected     bool
	}{
		"nil_other_not_nil": {
			other:    &schema.AssociatedExternalType{},
			expected: false,
		},
		"not_nil_other_nil": {
			assocExtType: &schema.AssociatedExternalType{},
			expected:     false,
		},
		"import_nil_other_import_not_nil": {
			assocExtType: &schema.AssociatedExternalType{},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_not_nil_other_import_nil": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			other:    &schema.AssociatedExternalType{},
			expected: false,
		},
		"import_alias_nil_other_import_alias_not_nil": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
				},
			},
			expected: false,
		},
		"import_alias_not_nil_other_import_alias_nil": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
				},
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_path_empty_other_import_path_not_empty": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{
					Path: "path",
				},
			},
			expected: false,
		},
		"import_path_not_empty_other_import_path_empty": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{
					Path: "path",
				},
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{},
			},
			expected: false,
		},
		"import_match": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
			},
			expected: true,
		},
		"type_empty_other_type_not_empty": {
			assocExtType: &schema.AssociatedExternalType{},
			other: &schema.AssociatedExternalType{
				Type: "type",
			},
			expected: false,
		},
		"type_not_empty_other_type_empty": {
			assocExtType: &schema.AssociatedExternalType{
				Type: "type",
			},
			other:    &schema.AssociatedExternalType{},
			expected: false,
		},
		"type_match": {
			assocExtType: &schema.AssociatedExternalType{
				Type: "type",
			},
			other: &schema.AssociatedExternalType{
				Type: "type",
			},
			expected: true,
		},
		"match": {
			assocExtType: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
				Type: "type",
			},
			other: &schema.AssociatedExternalType{
				Import: &code.Import{
					Alias: pointer("alias"),
					Path:  "path",
				},
				Type: "type",
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.assocExtType.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
