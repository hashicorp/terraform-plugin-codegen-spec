// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestDynamicDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		DynamicDefault *schema.DynamicDefault
		other          *schema.DynamicDefault
		expected       bool
	}{
		"both_nil": {
			expected: true,
		},
		"dynamic_default_nil_other_not_nil": {
			other:    &schema.DynamicDefault{},
			expected: false,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.DynamicDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
