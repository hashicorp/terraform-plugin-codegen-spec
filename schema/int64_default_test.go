// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestInt64Default_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		int64Default *schema.Int64Default
		other        *schema.Int64Default
		expected     bool
	}{
		"both_nil": {
			expected: true,
		},
		"int64_default_nil_other_not_nil": {
			other:    &schema.Int64Default{},
			expected: false,
		},
		"int64_default_static_nil_other_not_nil": {
			int64Default: &schema.Int64Default{},
			other: &schema.Int64Default{
				Static: pointer(int64(1234)),
			},
			expected: false,
		},
		"int64_default_static_not_nil_other_nil": {
			int64Default: &schema.Int64Default{
				Static: pointer(int64(1234)),
			},
			other:    &schema.Int64Default{},
			expected: false,
		},
		"match": {
			int64Default: &schema.Int64Default{
				Static: pointer(int64(1234)),
			},
			other: &schema.Int64Default{
				Static: pointer(int64(1234)),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.int64Default.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
