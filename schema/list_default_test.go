// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestListDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		listDefault *schema.ListDefault
		other       *schema.ListDefault
		expected    bool
	}{
		"both_nil": {
			expected: true,
		},
		"list_default_nil_other_not_nil": {
			other:    &schema.ListDefault{},
			expected: false,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.listDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
