// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestAttributes_Validate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		attributes    resource.Attributes
		request       resource.AttributeValidateRequest
		expectedError error
	}{
		"attribute-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-triplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-unique": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Bool: &resource.BoolAttribute{},
				},
				{
					Name: "attr_two",
					Bool: &resource.BoolAttribute{},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
		},
		"list-attribute-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					ListNested: &resource.ListNestedAttribute{
						NestedObject: resource.NestedAttributeObject{
							Attributes: resource.Attributes{
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_one",
								},
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-triplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					ListNested: &resource.ListNestedAttribute{
						NestedObject: resource.NestedAttributeObject{
							Attributes: resource.Attributes{
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_one",
								},
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-unique": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					ListNested: &resource.ListNestedAttribute{
						NestedObject: resource.NestedAttributeObject{
							Attributes: resource.Attributes{
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_two",
								},
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
		},
		"attribute-and-list-attribute-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					ListNested: &resource.ListNestedAttribute{
						NestedObject: resource.NestedAttributeObject{
							Attributes: resource.Attributes{
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_one",
								},
							},
						},
					},
				},
				{
					Name: "attr_one",
					ListNested: &resource.ListNestedAttribute{
						NestedObject: resource.NestedAttributeObject{
							Attributes: resource.Attributes{
								{
									Name: "nested_attr_one",
								},
								{
									Name: "nested_attr_one",
								},
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"object-attribute-type-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Object: &resource.ObjectAttribute{
						AttributeTypes: schema.ObjectAttributeTypes{
							{
								Name: "obj_attr_one",
							},
							{
								Name: "obj_attr_one",
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-attribute-names-and-type-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Object: &resource.ObjectAttribute{
						AttributeTypes: schema.ObjectAttributeTypes{
							{
								Name: "obj_attr_one",
							},
							{
								Name: "obj_attr_one",
							},
						},
					},
				},
				{
					Name: "attr_one",
					Object: &resource.ObjectAttribute{
						AttributeTypes: schema.ObjectAttributeTypes{
							{
								Name: "obj_attr_one",
							},
							{
								Name: "obj_attr_one",
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated` + "\n" +
				`resource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-object-attribute-type-names-duplicated": {
			attributes: resource.Attributes{
				{
					Name: "attr_one",
					Object: &resource.ObjectAttribute{
						AttributeTypes: schema.ObjectAttributeTypes{
							{
								Name: "obj_attr_one",
								Object: &schema.ObjectType{
									AttributeTypes: schema.ObjectAttributeTypes{
										{
											Name: "nested_obj_attr_one",
										},
										{
											Name: "nested_obj_attr_one",
										},
									},
								},
							},
						},
					},
				},
			},
			request: resource.AttributeValidateRequest{
				Path: `resource "example"`,
			},
			expectedError: fmt.Errorf(`resource "example" attribute "attr_one" object attribute type "obj_attr_one" object attribute type "nested_obj_attr_one" is duplicated`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := testCase.attributes.Validate(context.Background(), testCase.request)

			if err != nil {
				if testCase.expectedError == nil {
					t.Fatalf("expected no error, got: %s", err)
				}

				if err.Error() != testCase.expectedError.Error() {
					t.Fatalf("expected error %q, got: %s", testCase.expectedError, err)
				}
			}

			if err == nil && testCase.expectedError != nil {
				t.Fatalf("got no error, expected: %s", testCase.expectedError)
			}
		})
	}
}
