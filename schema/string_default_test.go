// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestStringDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		stringDefault *schema.StringDefault
		other         *schema.StringDefault
		expected      bool
	}{
		"both_nil": {
			expected: true,
		},
		"string_default_nil_other_not_nil": {
			other:    &schema.StringDefault{},
			expected: false,
		},
		"string_default_static_nil_other_not_nil": {
			stringDefault: &schema.StringDefault{},
			other: &schema.StringDefault{
				Static: pointer("str"),
			},
			expected: false,
		},
		"string_default_static_not_nil_other_nil": {
			stringDefault: &schema.StringDefault{
				Static: pointer("str"),
			},
			other:    &schema.StringDefault{},
			expected: false,
		},
		"match": {
			stringDefault: &schema.StringDefault{
				Static: pointer("str"),
			},
			other: &schema.StringDefault{
				Static: pointer("str"),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.stringDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
