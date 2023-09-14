// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestSetDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setDefault *schema.SetDefault
		other      *schema.SetDefault
		expected   bool
	}{
		"both_nil": {
			expected: true,
		},
		"set_default_nil_other_not_nil": {
			other:    &schema.SetDefault{},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.setDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
