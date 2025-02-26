// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestBoolDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		boolDefault *schema.BoolDefault
		other       *schema.BoolDefault
		expected    bool
	}{
		"both_nil": {
			expected: true,
		},
		"bool_default_nil_other_not_nil": {
			other:    &schema.BoolDefault{},
			expected: false,
		},
		"bool_default_static_nil_other_not_nil": {
			boolDefault: &schema.BoolDefault{},
			other: &schema.BoolDefault{
				Static: pointer(true),
			},
			expected: false,
		},
		"bool_default_static_not_nil_other_nil": {
			boolDefault: &schema.BoolDefault{
				Static: pointer(true),
			},
			other:    &schema.BoolDefault{},
			expected: false,
		},
		"match": {
			boolDefault: &schema.BoolDefault{
				Static: pointer(true),
			},
			other: &schema.BoolDefault{
				Static: pointer(true),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.boolDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
