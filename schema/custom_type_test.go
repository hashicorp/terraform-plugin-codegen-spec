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
		customType schema.CustomType
		expected   bool
	}{
		"import-nil": {
			customType: schema.CustomType{},
			expected:   false,
		},
		"import-empty-string": {
			customType: schema.CustomType{
				Import: &code.Import{
					Path: "",
				},
			},
			expected: false,
		},
		"import-string": {
			customType: schema.CustomType{
				Import: &code.Import{
					Path: "github.com/owner/repo/pkg",
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
