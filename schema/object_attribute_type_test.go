// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestObjectAttributeTypes_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		objectAttributeTypes schema.ObjectAttributeTypes
		other                schema.ObjectAttributeTypes
		expected             bool
	}{
		"object_attribute_types_both_nil": {
			expected: true,
		},
		"object_attribute_types_nil_other_not_nil": {
			other:    schema.ObjectAttributeTypes{},
			expected: false,
		},
		"object_attribute_types_not_nil_other_nil": {
			objectAttributeTypes: schema.ObjectAttributeTypes{},
			expected:             false,
		},
		"object_attribute_types_len_diff": {
			objectAttributeTypes: schema.ObjectAttributeTypes{
				{},
			},
			other:    schema.ObjectAttributeTypes{},
			expected: false,
		},
		"match": {
			objectAttributeTypes: schema.ObjectAttributeTypes{},
			other:                schema.ObjectAttributeTypes{},
			expected:             true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.objectAttributeTypes.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

// TestObjectAttributeType_Equal does not test for equality of field type
// CustomType as this is tested by TestCustomType_Equal.
func TestObjectAttributeType_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		objectAttributeType schema.ObjectAttributeType
		other               schema.ObjectAttributeType
		expected            bool
	}{
		"name_mismatch": {
			objectAttributeType: schema.ObjectAttributeType{
				Name: "one",
			},
			other: schema.ObjectAttributeType{
				Name: "two",
			},
			expected: false,
		},
		"bool_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Bool: &schema.BoolType{},
			},
			expected: false,
		},
		"bool_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Bool: &schema.BoolType{},
			},
			expected: false,
		},
		"dynamic_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Dynamic: &schema.DynamicType{},
			},
			expected: false,
		},
		"dynamic_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Dynamic: &schema.DynamicType{},
			},
			expected: false,
		},
		"float64_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Float64: &schema.Float64Type{},
			},
			expected: false,
		},
		"float64_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Float64: &schema.Float64Type{},
			},
			expected: false,
		},
		"int64_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Int64: &schema.Int64Type{},
			},
			expected: false,
		},
		"int64_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Int64: &schema.Int64Type{},
			},
			expected: false,
		},
		"list_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				List: &schema.ListType{},
			},
			expected: false,
		},
		"list_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				List: &schema.ListType{},
			},
			expected: false,
		},
		"map_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Map: &schema.MapType{},
			},
			expected: false,
		},
		"map_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Map: &schema.MapType{},
			},
			expected: false,
		},
		"number_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Number: &schema.NumberType{},
			},
			expected: false,
		},
		"number_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Number: &schema.NumberType{},
			},
			expected: false,
		},
		"object_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Object: &schema.ObjectType{},
			},
			expected: false,
		},
		"object_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Object: &schema.ObjectType{},
			},
			expected: false,
		},
		"set_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				Set: &schema.SetType{},
			},
			expected: false,
		},
		"set_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				Set: &schema.SetType{},
			},
			expected: false,
		},
		"string_nil_other_not_nil": {
			other: schema.ObjectAttributeType{
				String: &schema.StringType{},
			},
			expected: false,
		},
		"string_not_nil_other_nil": {
			objectAttributeType: schema.ObjectAttributeType{
				String: &schema.StringType{},
			},
			expected: false,
		},
		"match": {
			objectAttributeType: schema.ObjectAttributeType{
				Name: "one",
				Bool: &schema.BoolType{},
			},
			other: schema.ObjectAttributeType{
				Name: "one",
				Bool: &schema.BoolType{},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.objectAttributeType.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
