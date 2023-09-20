// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSchemas_Get(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		version       string
		expected      []byte
		expectedError error
	}{
		"valid_version": {
			version:       "v1.0",
			expected:      v1_0,
			expectedError: nil,
		},
		"invalid_version": {
			version:       "v1.x",
			expected:      nil,
			expectedError: fmt.Errorf(`schema version: "v1.x" is invalid`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Schema(testCase.version)

			if err != nil {
				if testCase.expectedError == nil {
					t.Fatalf("expected no error, got: %s", err)
				}

				if !strings.Contains(err.Error(), testCase.expectedError.Error()) {
					t.Fatalf("expected error %q, got: %s", testCase.expectedError, err)
				}
			}

			if err == nil && testCase.expectedError != nil {
				t.Fatalf("got no error, expected: %s", testCase.expectedError)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
