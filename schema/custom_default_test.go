// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

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
		"import-empty-string": {
			customType: schema.CustomDefault{
				Import: pointer(""),
			},
			expected: false,
		},
		"import-string": {
			customType: schema.CustomDefault{
				Import: pointer("github.com/owner/repo/pkg"),
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
