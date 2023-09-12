// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestMapDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		mapDefault *schema.MapDefault
		other      *schema.MapDefault
		expected   bool
	}{
		"both_nil": {
			expected: true,
		},
		"map_default_nil_other_not_nil": {
			other:    &schema.MapDefault{},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.mapDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
