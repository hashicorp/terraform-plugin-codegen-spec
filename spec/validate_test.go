package spec_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		document []byte
		expected error
	}{
		"nil": {
			document: nil,
			expected: fmt.Errorf("empty document"),
		},
		"empty": {
			document: []byte{},
			expected: fmt.Errorf("empty document"),
		},
		"example": {
			document: testReadFile("example.json"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := spec.Validate(context.Background(), testCase.document)

			if err != nil {
				if testCase.expected == nil {
					t.Fatalf("expected no error, got: %s", err)
				}

				if !strings.Contains(err.Error(), testCase.expected.Error()) {
					t.Fatalf("expected error %q, got: %s", testCase.expected, err)
				}
			}

			if err == nil && testCase.expected != nil {
				t.Fatalf("got no error, expected: %s", testCase.expected)
			}
		})
	}
}
