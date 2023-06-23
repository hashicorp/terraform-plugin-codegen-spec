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
