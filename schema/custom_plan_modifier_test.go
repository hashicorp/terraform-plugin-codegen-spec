package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestCustomPlanModifier_HasImport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		customType schema.CustomPlanModifier
		expected   bool
	}{
		"import-nil": {
			customType: schema.CustomPlanModifier{},
			expected:   false,
		},
		"import-empty-string": {
			customType: schema.CustomPlanModifier{
				Import: pointer(""),
			},
			expected: false,
		},
		"import-string": {
			customType: schema.CustomPlanModifier{
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
