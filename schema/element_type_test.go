// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

// TestElementType_Equal does not test for equality of element CustomType
// as this is tested by TestCustomType_Equal.
func TestElementType_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		elementType schema.ElementType
		other       schema.ElementType
		expected    bool
	}{
		"bool_nil_other_not_nil": {
			other: schema.ElementType{
				Bool: &schema.BoolType{},
			},
			expected: false,
		},
		"bool_not_nil_other_nil": {
			elementType: schema.ElementType{
				Bool: &schema.BoolType{},
			},
			expected: false,
		},
		"float64_nil_other_not_nil": {
			other: schema.ElementType{
				Float64: &schema.Float64Type{},
			},
			expected: false,
		},
		"float64_not_nil_other_nil": {
			elementType: schema.ElementType{
				Float64: &schema.Float64Type{},
			},
			expected: false,
		},
		"int64_nil_other_not_nil": {
			other: schema.ElementType{
				Int64: &schema.Int64Type{},
			},
			expected: false,
		},
		"int64_not_nil_other_nil": {
			elementType: schema.ElementType{
				Int64: &schema.Int64Type{},
			},
			expected: false,
		},
		"list_nil_other_not_nil": {
			other: schema.ElementType{
				List: &schema.ListType{},
			},
			expected: false,
		},
		"list_not_nil_other_nil": {
			elementType: schema.ElementType{
				List: &schema.ListType{},
			},
			expected: false,
		},
		"map_nil_other_not_nil": {
			other: schema.ElementType{
				Map: &schema.MapType{},
			},
			expected: false,
		},
		"map_not_nil_other_nil": {
			elementType: schema.ElementType{
				Map: &schema.MapType{},
			},
			expected: false,
		},
		"number_nil_other_not_nil": {
			other: schema.ElementType{
				Number: &schema.NumberType{},
			},
			expected: false,
		},
		"number_not_nil_other_nil": {
			elementType: schema.ElementType{
				Number: &schema.NumberType{},
			},
			expected: false,
		},
		"object_nil_other_not_nil": {
			other: schema.ElementType{
				Object: &schema.ObjectType{},
			},
			expected: false,
		},
		"object_not_nil_other_nil": {
			elementType: schema.ElementType{
				Object: &schema.ObjectType{},
			},
			expected: false,
		},
		"set_nil_other_not_nil": {
			other: schema.ElementType{
				Set: &schema.SetType{},
			},
			expected: false,
		},
		"set_not_nil_other_nil": {
			elementType: schema.ElementType{
				Set: &schema.SetType{},
			},
			expected: false,
		},
		"string_nil_other_not_nil": {
			other: schema.ElementType{
				String: &schema.StringType{},
			},
			expected: false,
		},
		"string_not_nil_other_nil": {
			elementType: schema.ElementType{
				String: &schema.StringType{},
			},
			expected: false,
		},
		"match": {
			elementType: schema.ElementType{
				String: &schema.StringType{},
			},
			other: schema.ElementType{
				String: &schema.StringType{},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.elementType.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
