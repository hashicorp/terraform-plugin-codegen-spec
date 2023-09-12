// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
		"datasource-attributes-only": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"attributes": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"datasource-blocks-only": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"datasource-attributes-and-blocks": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"attributes": [],
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"datasource-no-attributes-or-blocks": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
			expected: fmt.Errorf("datasources.0.schema: Must have at least 1 properties"),
		},
		"resource-attributes-only": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"attributes": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"resource-blocks-only": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"resource-attributes-and-blocks": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"attributes": [],
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
		},
		"resource-no-attributes-or-blocks": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
      }
    }
  ],
  "provider": {
    "name": "provider"
  }
}`),
			expected: fmt.Errorf("resources.0.schema: Must have at least 1 properties"),
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
