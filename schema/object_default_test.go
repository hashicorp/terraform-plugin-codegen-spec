// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestObjectDefault_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		objectDefault *schema.ObjectDefault
		other         *schema.ObjectDefault
		expected      bool
	}{
		"both_nil": {
			expected: true,
		},
		"object_default_nil_other_not_nil": {
			other:    &schema.ObjectDefault{},
			expected: false,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.objectDefault.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
