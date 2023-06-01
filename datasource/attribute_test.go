package datasource_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func TestAttributes_Validate(t *testing.T) {
	testCases := map[string]struct {
		attributes    datasource.Attributes
		request       datasource.AttributeValidateRequest
		expectedError error
	}{
		"attribute-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
			},
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-triplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
			},
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" is duplicated`),
		},
		"attribute-names-unique": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Bool: &datasource.BoolAttribute{},
				},
				{
					Name: "attr_two",
					Bool: &datasource.BoolAttribute{},
				},
			},
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
		},
		"list-attribute-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					ListNested: &datasource.ListNestedAttribute{
						NestedObject: datasource.NestedAttributeObject{
							Attributes: datasource.Attributes{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-triplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					ListNested: &datasource.ListNestedAttribute{
						NestedObject: datasource.NestedAttributeObject{
							Attributes: datasource.Attributes{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"list-nested-attribute-names-unique": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					ListNested: &datasource.ListNestedAttribute{
						NestedObject: datasource.NestedAttributeObject{
							Attributes: datasource.Attributes{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
		},
		"attribute-and-list-attribute-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					ListNested: &datasource.ListNestedAttribute{
						NestedObject: datasource.NestedAttributeObject{
							Attributes: datasource.Attributes{
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
					ListNested: &datasource.ListNestedAttribute{
						NestedObject: datasource.NestedAttributeObject{
							Attributes: datasource.Attributes{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" attribute "nested_attr_one" is duplicated`),
		},
		"object-attribute-type-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Object: &datasource.ObjectAttribute{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-attribute-names-and-type-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Object: &datasource.ObjectAttribute{
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
					Object: &datasource.ObjectAttribute{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated` + "\n" +
				`datasource "example" attribute "attr_one" object attribute type "obj_attr_one" is duplicated`),
		},
		"object-object-attribute-type-names-duplicated": {
			attributes: datasource.Attributes{
				{
					Name: "attr_one",
					Object: &datasource.ObjectAttribute{
						AttributeTypes: schema.ObjectAttributeTypes{
							{
								Name: "obj_attr_one",
								Object: schema.ObjectAttributeTypes{
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
			request: datasource.AttributeValidateRequest{
				Path: `datasource "example"`,
			},
			expectedError: fmt.Errorf(`datasource "example" attribute "attr_one" object attribute type "obj_attr_one" object attribute type "nested_obj_attr_one" is duplicated`),
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
