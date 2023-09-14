// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// TestObjectType_Equal does not test for equality of schema.ObjectType.CustomType
// as that is tested by TestCustomType_Equal.
// TestObjectType_Equal does not test for equality of schema.ObjectType.ObjectAttributeTypes
// other than Name as element equality is tested by TestElementType_Equal.
func TestObjectType_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		objectType *schema.ObjectType
		other      *schema.ObjectType
		expected   bool
	}{
		"object_nil_other_not_nil": {
			other:    &schema.ObjectType{},
			expected: false,
		},
		"object_not_nil_other_nil": {
			objectType: &schema.ObjectType{},
			expected:   false,
		},
		"object_attribute_types_nil_other_not_nil": {
			objectType: &schema.ObjectType{},
			other: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{},
			},
			expected: false,
		},
		"object_attribute_types_not_nil_other_nil": {
			objectType: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{},
			},
			other:    &schema.ObjectType{},
			expected: false,
		},
		"object_attribute_types_name_mismatch": {
			objectType: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{
					{
						Name: "one",
					},
				},
			},
			other: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{
					{
						Name: "two",
					},
				},
			},
			expected: false,
		},
		"object_match": {
			objectType: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{
					{
						Name: "one",
					},
				},
			},
			other: &schema.ObjectType{
				AttributeTypes: schema.ObjectAttributeTypes{
					{
						Name: "one",
					},
				},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.objectType.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
