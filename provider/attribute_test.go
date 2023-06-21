// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestAttributes_Validate(t *testing.T) {
	testCases := map[string]struct {
		attributes    provider.Attributes
		request       provider.AttributeValidateRequest
		expectedError error
	}{
		"attribute-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
			},
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-triplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
			},
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-unique": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Bool: &provider.BoolAttribute{},
				},
				{
					Name: "attr_two",
					Bool: &provider.BoolAttribute{},
				},
			},
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
		},
		"list-attribute-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					ListNested: &provider.ListNestedAttribute{
						NestedObject: provider.NestedAttributeObject{
							Attributes: provider.Attributes{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-triplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					ListNested: &provider.ListNestedAttribute{
						NestedObject: provider.NestedAttributeObject{
							Attributes: provider.Attributes{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-unique": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					ListNested: &provider.ListNestedAttribute{
						NestedObject: provider.NestedAttributeObject{
							Attributes: provider.Attributes{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
		},
		"attribute-and-list-attribute-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					ListNested: &provider.ListNestedAttribute{
						NestedObject: provider.NestedAttributeObject{
							Attributes: provider.Attributes{
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
					ListNested: &provider.ListNestedAttribute{
						NestedObject: provider.NestedAttributeObject{
							Attributes: provider.Attributes{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"object-attribute-type-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Object: &provider.ObjectAttribute{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-attribute-names-and-type-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Object: &provider.ObjectAttribute{
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
					Object: &provider.ObjectAttribute{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated` + "\n" +
				`provider "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-object-attribute-type-names-duplicated": {
			attributes: provider.Attributes{
				{
					Name: "attr_one",
					Object: &provider.ObjectAttribute{
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
			request: provider.AttributeValidateRequest{
				Path: `provider "example"`,
			},
			expectedError: fmt.Errorf(`provider "example" attribute "attr_one" object attribute type "obj_attr_one" object attribute type "nested_obj_attr_one" is duplicated`),
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
