// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestFloat64Default_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		float64Default *schema.Float64Default
		other          *schema.Float64Default
		expected       bool
	}{
		"both_nil": {
			expected: true,
		},
		"float64_default_nil_other_not_nil": {
			other:    &schema.Float64Default{},
			expected: false,
		},
		"float64_default_static_nil_other_not_nil": {
			float64Default: &schema.Float64Default{},
			other: &schema.Float64Default{
				Static: pointer(1.234),
			},
			expected: false,
		},
		"float64_default_static_not_nil_other_nil": {
			float64Default: &schema.Float64Default{
				Static: pointer(1.234),
			},
			other:    &schema.Float64Default{},
			expected: false,
		},
		"match": {
			float64Default: &schema.Float64Default{
				Static: pointer(1.234),
			},
			other: &schema.Float64Default{
				Static: pointer(1.234),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.float64Default.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
