// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

func TestSpecification_JSONMarshal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec          spec.Specification
		expected      []byte
		expectedError error
	}{
		"empty": {
			spec:     spec.Specification{},
			expected: []byte(`{}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(testCase.spec)

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

func TestSpecification_JSONUnmarshal_Version0_1(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		data          []byte
		expected      spec.Specification
		expectedError error
	}{
		"example": {
			data: testReadFile("./v0.1/example.json"),
			expected: spec.Specification{
				Version: spec.Version0_1,
				DataSources: datasource.DataSources{
					{
						Name: "datasource",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "bool_attribute_custom_type",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.BoolType",
											ValueType: "basetypes.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_custom_type_import_alias",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.BoolType",
											ValueType: "fwtype.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_associated_external_type",
									Bool: &datasource.BoolAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_associated_external_type",
									Dynamic: &datasource.DynamicAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_custom_type",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.DynamicType",
											ValueType: "basetypes.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_custom_type_import_alias",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.DynamicType",
											ValueType: "fwtype.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_validators",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Optional,
										Validators: schema.DynamicValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/mydynamicvalidator",
														},
													},
													SchemaDefinition: "mydynamicvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "float64_attribute",
									Float64: &datasource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "float64_attribute_custom_type",
									Float64: &datasource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Float64Type",
											ValueType: "basetypes.Float64Value",
										},
									},
								},
								{
									Name: "float64_attribute_associated_external_type",
									Float64: &datasource.Float64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute",
									Int64: &datasource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute_custom_type",
									Int64: &datasource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Int64Type",
											ValueType: "basetypes.Int64Value",
										},
									},
								},
								{
									Name: "int64_attribute_associated_external_type",
									Int64: &datasource.Int64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_custom_type",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ListType",
											ValueType: "basetypes.ListValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_element_type_string_custom_type",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "list_map_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Map: &schema.MapType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "list_object_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_object_object_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_obj_string_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_attribute_associated_external_type",
									List: &datasource.ListAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_custom_type",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.MapType",
											ValueType: "basetypes.MapValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_element_type_string_custom_type",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "map_attribute_associated_external_type",
									Map: &datasource.MapAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute",
									MapNested: &datasource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute_associated_external_type",
									MapNested: &datasource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "number_attribute",
									Number: &datasource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "number_attribute_custom_type",
									Number: &datasource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.NumberType",
											ValueType: "basetypes.NumberValue",
										},
									},
								},
								{
									Name: "number_attribute_associated_external_type",
									Number: &datasource.NumberAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_dynamic",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:    "object_dynamic_attribute",
												Dynamic: &schema.DynamicType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_object_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_object_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_object_string_attr",
															String: &schema.StringType{},
														},
													},
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.ObjectType",
														ValueType: "basetypes.ObjectValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_string_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												String: &schema.StringType{
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.StringType",
														ValueType: "basetypes.StringValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ObjectType",
											ValueType: "basetypes.ObjectValue",
										},
									},
								},
								{
									Name: "object_list_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														String: &schema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_object_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_list_obj_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_associated_external_type",
									Object: &datasource.ObjectAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_nested_bool_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_bool_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &datasource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: datasource.NestedAttributeObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &datasource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_list_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &datasource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: datasource.NestedAttributeObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "list_attribute",
																	List: &datasource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_bool_attribute_associated_external_type",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_attribute",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_custom_type",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.SetType",
											ValueType: "basetypes.SetValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_element_type_string_custom_type",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "set_attribute_associated_external_type",
									Set: &datasource.SetAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute",
									SetNested: &datasource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute_associated_external_type",
									SetNested: &datasource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_bool_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_bool_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &datasource.SingleNestedAttribute{
													Attributes: []datasource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &datasource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_list_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &datasource.SingleNestedAttribute{
													Attributes: []datasource.Attribute{
														{
															Name: "list_attribute",
															List: &datasource.ListAttribute{
																ComputedOptionalRequired: schema.Computed,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_bool_attribute_associated_external_type",
									SingleNested: &datasource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute",
									String: &datasource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute_custom_type",
									String: &datasource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.StringType",
											ValueType: "basetypes.StringValue",
										},
									},
								},
								{
									Name: "string_attribute_associated_external_type",
									String: &datasource.StringAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
							},
							Blocks: []datasource.Block{
								{
									Name: "list_nested_block_bool_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_block_bool_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "list_nested_block",
													ListNested: &datasource.ListNestedBlock{
														NestedObject: datasource.NestedBlockObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &datasource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "object_attribute",
													Object: &datasource.ObjectAttribute{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_string_attr",
																String: &schema.StringType{},
															},
														},
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
											Blocks: []datasource.Block{
												{
													Name: "list_nested_block",
													ListNested: &datasource.ListNestedBlock{
														NestedObject: datasource.NestedBlockObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "list_attribute",
																	List: &datasource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_bool_attribute_associated_external_type",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute_associated_external_type",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
								{
									Name: "single_nested_single_nested_block_bool_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "single_nested_block",
												SingleNested: &datasource.SingleNestedBlock{
													Attributes: []datasource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &datasource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Attributes: []datasource.Attribute{
											{
												Name: "object_attribute",
												Object: &datasource.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										Blocks: []datasource.Block{
											{
												Name: "list_nested_block",
												ListNested: &datasource.ListNestedBlock{
													NestedObject: datasource.NestedBlockObject{
														Attributes: []datasource.Attribute{
															{
																Name: "list_attribute",
																List: &datasource.ListAttribute{
																	ComputedOptionalRequired: schema.Computed,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute_associated_external_type",
									SingleNested: &datasource.SingleNestedBlock{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
							},
							MarkdownDescription: pointer("*This* is a description"),
							Description:         pointer("This is a description"),
							DeprecationMessage:  pointer("This data source is deprecated"),
						},
					},
				},
				Provider: &provider.Provider{
					Name: "provider",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "bool_attribute",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "bool_attribute_custom_type",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.BoolType",
										ValueType: "basetypes.BoolValue",
									},
								},
							},
							{
								Name: "bool_attribute_custom_type_import_alias",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Alias: pointer("fwtype"),
											Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "fwtype.BoolType",
										ValueType: "fwtype.BoolValue",
									},
								},
							},
							{
								Name: "bool_attribute_associated_external_type",
								Bool: &provider.BoolAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute_associated_external_type",
								Dynamic: &provider.DynamicAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute_custom_type",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.DynamicType",
										ValueType: "basetypes.DynamicValue",
									},
								},
							},
							{
								Name: "dynamic_attribute_custom_type_import_alias",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Alias: pointer("fwtype"),
											Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "fwtype.DynamicType",
										ValueType: "fwtype.DynamicValue",
									},
								},
							},
							{
								Name: "dynamic_attribute_validators",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									Validators: schema.DynamicValidators{
										{
											Custom: &schema.CustomValidator{
												Imports: []code.Import{
													{
														Path: "github.com/my_account/my_project/mydynamicvalidator",
													},
												},
												SchemaDefinition: "mydynamicvalidator.Validate()",
											},
										},
									},
								},
							},
							{
								Name: "float64_attribute",
								Float64: &provider.Float64Attribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "float64_attribute_custom_type",
								Float64: &provider.Float64Attribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.Float64Type",
										ValueType: "basetypes.Float64Value",
									},
								},
							},
							{
								Name: "float64_attribute_associated_external_type",
								Float64: &provider.Float64Attribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "int64_attribute",
								Int64: &provider.Int64Attribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "int64_attribute_custom_type",
								Int64: &provider.Int64Attribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.Int64Type",
										ValueType: "basetypes.Int64Value",
									},
								},
							},
							{
								Name: "int64_attribute_associated_external_type",
								Int64: &provider.Int64Attribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "list_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "list_attribute_custom_type",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.ListType",
										ValueType: "basetypes.ListValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "list_attribute_element_type_string_custom_type",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "list_map_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Map: &schema.MapType{
											ElementType: schema.ElementType{
												String: &schema.StringType{},
											},
										},
									},
								},
							},
							{
								Name: "list_object_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Object: &schema.ObjectType{
											AttributeTypes: []schema.ObjectAttributeType{
												{
													Name:   "obj_string_attr",
													String: &schema.StringType{},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_object_object_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Object: &schema.ObjectType{
											AttributeTypes: []schema.ObjectAttributeType{
												{
													Name: "obj_obj_attr",
													Object: &schema.ObjectType{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_obj_string_attr",
																String: &schema.StringType{},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_attribute_associated_external_type",
								List: &provider.ListAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute_custom_type",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.MapType",
										ValueType: "basetypes.MapValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute_element_type_string_custom_type",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "map_attribute_associated_external_type",
								Map: &provider.MapAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_nested_bool_attribute",
								MapNested: &provider.MapNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "map_nested_bool_attribute_associated_external_type",
								MapNested: &provider.MapNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "number_attribute",
								Number: &provider.NumberAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "number_attribute_custom_type",
								Number: &provider.NumberAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.NumberType",
										ValueType: "basetypes.NumberValue",
									},
								},
							},
							{
								Name: "number_attribute_associated_external_type",
								Number: &provider.NumberAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_dynamic",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:    "object_dynamic_attribute",
											Dynamic: &schema.DynamicType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_object_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_object_attr",
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_object_string_attr",
														String: &schema.StringType{},
													},
												},
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.ObjectType",
													ValueType: "basetypes.ObjectValue",
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_string_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_string_attr",
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.ObjectType",
										ValueType: "basetypes.ObjectValue",
									},
								},
							},
							{
								Name: "object_list_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_list_attr",
											List: &schema.ListType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_list_object_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_list_attr",
											List: &schema.ListType{
												ElementType: schema.ElementType{
													Object: &schema.ObjectType{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_list_obj_attr",
																String: &schema.StringType{},
															},
														},
													},
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_associated_external_type",
								Object: &provider.ObjectAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "list_nested_bool_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_bool_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "list_nested_attribute",
												ListNested: &provider.ListNestedAttribute{
													OptionalRequired: schema.Optional,
													NestedObject: provider.NestedAttributeObject{
														Attributes: []provider.Attribute{
															{
																Name: "bool_attribute",
																Bool: &provider.BoolAttribute{
																	OptionalRequired: schema.Optional,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_list_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "list_nested_attribute",
												ListNested: &provider.ListNestedAttribute{
													OptionalRequired: schema.Optional,
													NestedObject: provider.NestedAttributeObject{
														Attributes: []provider.Attribute{
															{
																Name: "list_attribute",
																List: &provider.ListAttribute{
																	OptionalRequired: schema.Optional,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_bool_attribute_associated_external_type",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_attribute",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_attribute_custom_type",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.SetType",
										ValueType: "basetypes.SetValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_attribute_element_type_string_custom_type",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "set_attribute_associated_external_type",
								Set: &provider.SetAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_nested_bool_attribute",
								SetNested: &provider.SetNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_bool_attribute_associated_external_type",
								SetNested: &provider.SetNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_bool_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_single_nested_bool_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "single_nested_attribute",
											SingleNested: &provider.SingleNestedAttribute{
												Attributes: []provider.Attribute{
													{
														Name: "bool_attribute",
														Bool: &provider.BoolAttribute{
															OptionalRequired: schema.Optional,
														},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_single_nested_list_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "single_nested_attribute",
											SingleNested: &provider.SingleNestedAttribute{
												Attributes: []provider.Attribute{
													{
														Name: "list_attribute",
														List: &provider.ListAttribute{
															OptionalRequired: schema.Optional,
															ElementType: schema.ElementType{
																String: &schema.StringType{},
															},
														},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_bool_attribute_associated_external_type",
								SingleNested: &provider.SingleNestedAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "string_attribute",
								String: &provider.StringAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "string_attribute_custom_type",
								String: &provider.StringAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.StringType",
										ValueType: "basetypes.StringValue",
									},
								},
							},
							{
								Name: "string_attribute_associated_external_type",
								String: &provider.StringAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
						},
						Blocks: []provider.Block{
							{
								Name: "list_nested_block_bool_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_block_bool_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "list_nested_block",
												ListNested: &provider.ListNestedBlock{
													NestedObject: provider.NestedBlockObject{
														Attributes: []provider.Attribute{
															{
																Name: "bool_attribute",
																Bool: &provider.BoolAttribute{
																	OptionalRequired: schema.Optional,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "object_attribute",
												Object: &provider.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													OptionalRequired: schema.Optional,
												},
											},
										},
										Blocks: []provider.Block{
											{
												Name: "list_nested_block",
												ListNested: &provider.ListNestedBlock{
													NestedObject: provider.NestedBlockObject{
														Attributes: []provider.Attribute{
															{
																Name: "list_attribute",
																List: &provider.ListAttribute{
																	OptionalRequired: schema.Optional,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_block_bool_attribute_associated_external_type",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_block_bool_attribute",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_block_bool_attribute_associated_external_type",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_bool_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
								},
							},
							{
								Name: "single_nested_single_nested_block_bool_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Blocks: []provider.Block{
										{
											Name: "single_nested_block",
											SingleNested: &provider.SingleNestedBlock{
												Attributes: []provider.Attribute{
													{
														Name: "bool_attribute",
														Bool: &provider.BoolAttribute{
															OptionalRequired: schema.Optional,
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Attributes: []provider.Attribute{
										{
											Name: "object_attribute",
											Object: &provider.ObjectAttribute{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									Blocks: []provider.Block{
										{
											Name: "list_nested_block",
											ListNested: &provider.ListNestedBlock{
												NestedObject: provider.NestedBlockObject{
													Attributes: []provider.Attribute{
														{
															Name: "list_attribute",
															List: &provider.ListAttribute{
																OptionalRequired: schema.Optional,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_bool_attribute_associated_external_type",
								SingleNested: &provider.SingleNestedBlock{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
								},
							},
						},
						MarkdownDescription: pointer("*This* is a description"),
						Description:         pointer("This is a description"),
						DeprecationMessage:  pointer("This provider is deprecated"),
					},
				},
				Resources: resource.Resources{
					{
						Name: "resource",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										PlanModifiers: schema.BoolPlanModifiers{
											{
												Custom: &schema.CustomPlanModifier{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/boolplanmodifier",
														},
													},
													SchemaDefinition: "myboolplanmodifier.Modify()",
												},
											},
										},
										Validators: schema.BoolValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/myboolvalidator",
														},
													},
													SchemaDefinition: "myboolvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "bool_attribute_custom_type",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.BoolType",
											ValueType: "basetypes.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_custom_type_import_alias",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.BoolType",
											ValueType: "fwtype.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_default_static",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.BoolDefault{
											Static: pointer(true),
										},
									},
								},
								{
									Name: "bool_attribute_associated_external_type",
									Bool: &resource.BoolAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Optional,
									},
								},
								{
									Name: "dynamic_attribute",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_associated_external_type",
									Dynamic: &resource.DynamicAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Optional,
									},
								},
								{
									Name: "dynamic_attribute_custom_type",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.DynamicType",
											ValueType: "basetypes.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_custom_type_import_alias",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.DynamicType",
											ValueType: "fwtype.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_default",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.DynamicDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "dynamicdefault.StaticValue(types.StringValue(\"example\"))",
											},
										},
									},
								},
								{
									Name: "dynamic_attribute_plan_modifiers",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										PlanModifiers: schema.DynamicPlanModifiers{
											{
												Custom: &schema.CustomPlanModifier{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/dynamicplanmodifier",
														},
													},
													SchemaDefinition: "mydynamicplanmodifier.Modify()",
												},
											},
										},
									},
								},
								{
									Name: "dynamic_attribute_validators",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										Validators: schema.DynamicValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/mydynamicvalidator",
														},
													},
													SchemaDefinition: "mydynamicvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "float64_attribute",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "float64_attribute_custom_type",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Float64Type",
											ValueType: "basetypes.Float64Value",
										},
									},
								},
								{
									Name: "float64_attribute_default_static",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.Float64Default{
											Static: pointer(123.45),
										},
									},
								},
								{
									Name: "float64_attribute_associated_external_type",
									Float64: &resource.Float64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute_custom_type",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Int64Type",
											ValueType: "basetypes.Int64Value",
										},
									},
								},
								{
									Name: "int64_attribute_default_static",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.Int64Default{
											Static: pointer(int64(123)),
										},
									},
								},
								{
									Name: "int64_attribute_associated_external_type",
									Int64: &resource.Int64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_custom_type",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ListType",
											ValueType: "basetypes.ListValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_default_custom",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.ListDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/attr",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "listdefault.StaticValue(types.ListValueMust(types.String, []attr.Value{types.StringValue(\"example\")}))",
											},
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_element_type_string_custom_type",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "list_map_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Map: &schema.MapType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "list_object_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_object_object_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_obj_string_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_attribute_associated_external_type",
									List: &resource.ListAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_custom_type",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.MapType",
											ValueType: "basetypes.MapValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_element_type_string_custom_type",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "map_attribute_associated_external_type",
									Map: &resource.MapAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute",
									MapNested: &resource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute_associated_external_type",
									MapNested: &resource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "number_attribute",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "number_attribute_custom_type",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.NumberType",
											ValueType: "basetypes.NumberValue",
										},
									},
								},
								{
									Name: "number_attribute_default_custom",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.NumberDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "math/big",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault",
													},
												},
												SchemaDefinition: "numberdefault.StaticBigFloat(big.NewFloat(123.45))",
											},
										},
									},
								},
								{
									Name: "number_attribute_associated_external_type",
									Number: &resource.NumberAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_dynamic",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:    "object_dynamic_attribute",
												Dynamic: &schema.DynamicType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_object_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_object_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_object_string_attr",
															String: &schema.StringType{},
														},
													},
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.ObjectType",
														ValueType: "basetypes.ObjectValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_string_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												String: &schema.StringType{
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.StringType",
														ValueType: "basetypes.StringValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ObjectType",
											ValueType: "basetypes.ObjectValue",
										},
									},
								},
								{
									Name: "object_list_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														String: &schema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_list_obj_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_associated_external_type",
									Object: &resource.ObjectAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_nested_bool_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_bool_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &resource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: resource.NestedAttributeObject{
															Attributes: []resource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &resource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_list_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &resource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: resource.NestedAttributeObject{
															Attributes: []resource.Attribute{
																{
																	Name: "list_attribute",
																	List: &resource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_bool_attribute_associated_external_type",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_attribute",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_custom_type",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.SetType",
											ValueType: "basetypes.SetValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_default_custom",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.SetDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/attr",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "setdefault.StaticValue(types.SetValueMust(types.String, []attr.Value{types.StringValue(\"example\")}))",
											},
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_element_type_string_custom_type",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "set_attribute_associated_external_type",
									Set: &resource.SetAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute",
									SetNested: &resource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute_associated_external_type",
									SetNested: &resource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_bool_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_bool_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &resource.SingleNestedAttribute{
													Attributes: []resource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &resource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_list_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &resource.SingleNestedAttribute{
													Attributes: []resource.Attribute{
														{
															Name: "list_attribute",
															List: &resource.ListAttribute{
																ComputedOptionalRequired: schema.Computed,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_bool_attribute_associated_external_type",
									SingleNested: &resource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute_custom_type",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.StringType",
											ValueType: "basetypes.StringValue",
										},
									},
								},
								{
									Name: "string_attribute_default_static",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.StringDefault{
											Static: pointer("example"),
										},
									},
								},
								{
									Name: "string_attribute_associated_external_type",
									String: &resource.StringAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
							},
							Blocks: []resource.Block{
								{
									Name: "list_nested_block_bool_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_block_bool_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "list_nested_block",
													ListNested: &resource.ListNestedBlock{
														NestedObject: resource.NestedBlockObject{
															Attributes: []resource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &resource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "object_attribute",
													Object: &resource.ObjectAttribute{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_string_attr",
																String: &schema.StringType{},
															},
														},
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
											Blocks: []resource.Block{
												{
													Name: "list_nested_block",
													ListNested: &resource.ListNestedBlock{
														NestedObject: resource.NestedBlockObject{
															Attributes: []resource.Attribute{
																{
																	Name: "list_attribute",
																	List: &resource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_bool_attribute_associated_external_type",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute_associated_external_type",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
								{
									Name: "single_nested_single_nested_block_bool_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "single_nested_block",
												SingleNested: &resource.SingleNestedBlock{
													Attributes: []resource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &resource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Attributes: []resource.Attribute{
											{
												Name: "object_attribute",
												Object: &resource.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										Blocks: []resource.Block{
											{
												Name: "list_nested_block",
												ListNested: &resource.ListNestedBlock{
													NestedObject: resource.NestedBlockObject{
														Attributes: []resource.Attribute{
															{
																Name: "list_attribute",
																List: &resource.ListAttribute{
																	ComputedOptionalRequired: schema.Computed,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute_associated_external_type",
									SingleNested: &resource.SingleNestedBlock{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
							},
							MarkdownDescription: pointer("*This* is a description"),
							Description:         pointer("This is a description"),
							DeprecationMessage:  pointer("This resource is deprecated"),
							Version:             pointer(int64(1)),
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got spec.Specification

			err := json.Unmarshal(testCase.data, &got)

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

func TestSpecification_Validate_DataSources(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec          spec.Specification
		expectedError error
	}{
		"data-source-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
					},
					{
						Name: "example",
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" is duplicated`),
		},
		"data-source-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
					},
					{
						Name: "different",
					},
				},
			},
		},
		"data-source-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "first_attr",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" is duplicated`),
		},
		"data-source-attribute-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "second_attr",
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "first_block",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" is duplicated`),
		},
		"data-source-attribute-and-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first",
								},
							},
							Blocks: []datasource.Block{
								{
									Name: "first",
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "second_block",
								},
							},
						},
					},
				},
			},
		},
		"data-source-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"data-source-list-nested-attribute-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"data-source-map-nested-attribute-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"data-source-set-nested-attribute-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"data-source-single-nested-attribute-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block "nested_block" is duplicated`),
		},
		"data-source-list-nested-block-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block "nested_block" is duplicated`),
		},
		"data-source-set-nested-block-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block "nested_block" is duplicated`),
		},
		"data-source-single-nested-block-names-unique": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
											},
											{
												Name: "obj_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" object attribute type "obj_attr" is duplicated`),
		},
		"data-source-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name: "obj_obj_attr",
														},
														{
															Name: "obj_obj_attr",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" object attribute type "obj_attr" object attribute type "obj_obj_attr" is duplicated`),
		},
		"data-source-object-and-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				DataSources: datasource.DataSources{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "obj_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name: "obj_attr",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := testCase.spec.Validate(context.Background())

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

func TestSpecification_Validate_Provider(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec          spec.Specification
		expectedError error
	}{
		"provider-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
							},
							{
								Name: "first_attr",
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" is duplicated`),
		},
		"provider-attribute-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
							},
							{
								Name: "second_attr",
							},
						},
					},
				},
			},
		},
		"provider-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
							},
							{
								Name: "first_block",
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" block "first_block" is duplicated`),
		},
		"provider-attribute-and-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first",
							},
						},
						Blocks: []provider.Block{
							{
								Name: "first",
							},
						},
					},
				},
			},
		},
		"provider-block-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
							},
							{
								Name: "second_block",
							},
						},
					},
				},
			},
		},
		"provider-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								ListNested: &provider.ListNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"provider-list-nested-attribute-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								ListNested: &provider.ListNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-attribute-and-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								ListNested: &provider.ListNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								MapNested: &provider.MapNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"provider-map-nested-attribute-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								MapNested: &provider.MapNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-attribute-and-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								MapNested: &provider.MapNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SetNested: &provider.SetNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"provider-set-nested-attribute-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SetNested: &provider.SetNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-attribute-and-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SetNested: &provider.SetNestedAttribute{
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "nested_attr",
										},
										{
											Name: "nested_attr",
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"provider-single-nested-attribute-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "nested_first_attr",
										},
										{
											Name: "nested_second_attr",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-attribute-and-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "first_attr",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" block "first_block" block "nested_block" is duplicated`),
		},
		"provider-list-nested-block-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-block-and-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" block "first_block" block "nested_block" is duplicated`),
		},
		"provider-set-nested-block-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-block-and-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SingleNested: &provider.SingleNestedBlock{
									Blocks: []provider.Block{
										{
											Name: "nested_block",
										},
										{
											Name: "nested_block",
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" block "first_block" block "nested_block" is duplicated`),
		},
		"provider-single-nested-block-names-unique": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SingleNested: &provider.SingleNestedBlock{
									Blocks: []provider.Block{
										{
											Name: "nested_first_block",
										},
										{
											Name: "nested_second_block",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-block-and-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Blocks: []provider.Block{
							{
								Name: "first_block",
								SingleNested: &provider.SingleNestedBlock{
									Blocks: []provider.Block{
										{
											Name: "first_block",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"provider-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_attr",
										},
										{
											Name: "obj_attr",
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" object attribute type "obj_attr" is duplicated`),
		},
		"provider-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "first_attr",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_attr",
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
													},
													{
														Name: "obj_obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`provider "example" attribute "first_attr" object attribute type "obj_attr" object attribute type "obj_obj_attr" is duplicated`),
		},
		"provider-object-and-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Provider: &provider.Provider{
					Name: "example",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "obj_attr",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_attr",
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := testCase.spec.Validate(context.Background())

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

func TestSpecification_Validate_Resources(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		spec          spec.Specification
		expectedError error
	}{
		"resource-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
					},
					{
						Name: "example",
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" is duplicated`),
		},
		"resource-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
					},
					{
						Name: "different",
					},
				},
			},
		},
		"resource-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "first_attr",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" is duplicated`),
		},
		"resource-attribute-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "second_attr",
								},
							},
						},
					},
				},
			},
		},
		"resource-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "first_block",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" is duplicated`),
		},
		"resource-attribute-and-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first",
								},
							},
							Blocks: []resource.Block{
								{
									Name: "first",
								},
							},
						},
					},
				},
			},
		},
		"resource-block-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "second_block",
								},
							},
						},
					},
				},
			},
		},
		"resource-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"resource-list-nested-attribute-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-list-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"resource-map-nested-attribute-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-map-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"resource-set-nested-attribute-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-set-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute "nested_attr" is duplicated`),
		},
		"resource-single-nested-attribute-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-single-nested-attribute-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block "nested_block" is duplicated`),
		},
		"resource-list-nested-block-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-list-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block "nested_block" is duplicated`),
		},
		"resource-set-nested-block-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-set-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block "nested_block" is duplicated`),
		},
		"resource-single-nested-block-names-unique": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-single-nested-block-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
											},
											{
												Name: "obj_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" object attribute type "obj_attr" is duplicated`),
		},
		"resource-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name: "obj_obj_attr",
														},
														{
															Name: "obj_obj_attr",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" object attribute type "obj_attr" object attribute type "obj_obj_attr" is duplicated`),
		},
		"resource-object-and-object-object-attribute-type-names-duplicated": {
			spec: spec.Specification{
				Resources: resource.Resources{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "obj_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name: "obj_attr",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := testCase.spec.Validate(context.Background())

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

func TestSpecification_Generate_Version0_1(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		data          []byte
		expected      spec.Specification
		expectedError error
	}{
		"example": {
			data: testReadFile("./v0.1/example.json"),
			expected: spec.Specification{
				Version: spec.Version0_1,
				DataSources: datasource.DataSources{
					{
						Name: "datasource",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "bool_attribute_custom_type",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.BoolType",
											ValueType: "basetypes.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_custom_type_import_alias",
									Bool: &datasource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.BoolType",
											ValueType: "fwtype.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_associated_external_type",
									Bool: &datasource.BoolAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_associated_external_type",
									Dynamic: &datasource.DynamicAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_custom_type",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.DynamicType",
											ValueType: "basetypes.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_custom_type_import_alias",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.DynamicType",
											ValueType: "fwtype.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_validators",
									Dynamic: &datasource.DynamicAttribute{
										ComputedOptionalRequired: schema.Optional,
										Validators: schema.DynamicValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/mydynamicvalidator",
														},
													},
													SchemaDefinition: "mydynamicvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "float64_attribute",
									Float64: &datasource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "float64_attribute_custom_type",
									Float64: &datasource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Float64Type",
											ValueType: "basetypes.Float64Value",
										},
									},
								},
								{
									Name: "float64_attribute_associated_external_type",
									Float64: &datasource.Float64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute",
									Int64: &datasource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute_custom_type",
									Int64: &datasource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Int64Type",
											ValueType: "basetypes.Int64Value",
										},
									},
								},
								{
									Name: "int64_attribute_associated_external_type",
									Int64: &datasource.Int64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_custom_type",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ListType",
											ValueType: "basetypes.ListValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_element_type_string_custom_type",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "list_map_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Map: &schema.MapType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "list_object_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_object_object_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_obj_string_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_attribute_associated_external_type",
									List: &datasource.ListAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_custom_type",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.MapType",
											ValueType: "basetypes.MapValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_element_type_string_custom_type",
									Map: &datasource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "map_attribute_associated_external_type",
									Map: &datasource.MapAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute",
									MapNested: &datasource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute_associated_external_type",
									MapNested: &datasource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "number_attribute",
									Number: &datasource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "number_attribute_custom_type",
									Number: &datasource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.NumberType",
											ValueType: "basetypes.NumberValue",
										},
									},
								},
								{
									Name: "number_attribute_associated_external_type",
									Number: &datasource.NumberAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_dynamic",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:    "object_dynamic_attribute",
												Dynamic: &schema.DynamicType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_object_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_object_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_object_string_attr",
															String: &schema.StringType{},
														},
													},
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.ObjectType",
														ValueType: "basetypes.ObjectValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_string_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												String: &schema.StringType{
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.StringType",
														ValueType: "basetypes.StringValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_custom_type",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ObjectType",
											ValueType: "basetypes.ObjectValue",
										},
									},
								},
								{
									Name: "object_list_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														String: &schema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_object_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_list_obj_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_associated_external_type",
									Object: &datasource.ObjectAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_nested_bool_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_bool_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &datasource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: datasource.NestedAttributeObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &datasource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_list_attribute",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &datasource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: datasource.NestedAttributeObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "list_attribute",
																	List: &datasource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_bool_attribute_associated_external_type",
									ListNested: &datasource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_attribute",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_custom_type",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.SetType",
											ValueType: "basetypes.SetValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_element_type_string_custom_type",
									Set: &datasource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "set_attribute_associated_external_type",
									Set: &datasource.SetAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute",
									SetNested: &datasource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute_associated_external_type",
									SetNested: &datasource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: datasource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_bool_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_bool_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &datasource.SingleNestedAttribute{
													Attributes: []datasource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &datasource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_list_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &datasource.SingleNestedAttribute{
													Attributes: []datasource.Attribute{
														{
															Name: "list_attribute",
															List: &datasource.ListAttribute{
																ComputedOptionalRequired: schema.Computed,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_bool_attribute_associated_external_type",
									SingleNested: &datasource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute",
									String: &datasource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute_custom_type",
									String: &datasource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.StringType",
											ValueType: "basetypes.StringValue",
										},
									},
								},
								{
									Name: "string_attribute_associated_external_type",
									String: &datasource.StringAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
							},
							Blocks: []datasource.Block{
								{
									Name: "list_nested_block_bool_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_block_bool_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "list_nested_block",
													ListNested: &datasource.ListNestedBlock{
														NestedObject: datasource.NestedBlockObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &datasource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "object_attribute",
													Object: &datasource.ObjectAttribute{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_string_attr",
																String: &schema.StringType{},
															},
														},
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
											Blocks: []datasource.Block{
												{
													Name: "list_nested_block",
													ListNested: &datasource.ListNestedBlock{
														NestedObject: datasource.NestedBlockObject{
															Attributes: []datasource.Attribute{
																{
																	Name: "list_attribute",
																	List: &datasource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_bool_attribute_associated_external_type",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute_associated_external_type",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []datasource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &datasource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
								{
									Name: "single_nested_single_nested_block_bool_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "single_nested_block",
												SingleNested: &datasource.SingleNestedBlock{
													Attributes: []datasource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &datasource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
									SingleNested: &datasource.SingleNestedBlock{
										Attributes: []datasource.Attribute{
											{
												Name: "object_attribute",
												Object: &datasource.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										Blocks: []datasource.Block{
											{
												Name: "list_nested_block",
												ListNested: &datasource.ListNestedBlock{
													NestedObject: datasource.NestedBlockObject{
														Attributes: []datasource.Attribute{
															{
																Name: "list_attribute",
																List: &datasource.ListAttribute{
																	ComputedOptionalRequired: schema.Computed,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute_associated_external_type",
									SingleNested: &datasource.SingleNestedBlock{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []datasource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &datasource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
							},
							MarkdownDescription: pointer("*This* is a description"),
							Description:         pointer("This is a description"),
							DeprecationMessage:  pointer("This data source is deprecated"),
						},
					},
				},
				Provider: &provider.Provider{
					Name: "provider",
					Schema: &provider.Schema{
						Attributes: []provider.Attribute{
							{
								Name: "bool_attribute",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "bool_attribute_custom_type",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.BoolType",
										ValueType: "basetypes.BoolValue",
									},
								},
							},
							{
								Name: "bool_attribute_custom_type_import_alias",
								Bool: &provider.BoolAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Alias: pointer("fwtype"),
											Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "fwtype.BoolType",
										ValueType: "fwtype.BoolValue",
									},
								},
							},
							{
								Name: "bool_attribute_associated_external_type",
								Bool: &provider.BoolAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute_associated_external_type",
								Dynamic: &provider.DynamicAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "dynamic_attribute_custom_type",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.DynamicType",
										ValueType: "basetypes.DynamicValue",
									},
								},
							},
							{
								Name: "dynamic_attribute_custom_type_import_alias",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Alias: pointer("fwtype"),
											Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "fwtype.DynamicType",
										ValueType: "fwtype.DynamicValue",
									},
								},
							},
							{
								Name: "dynamic_attribute_validators",
								Dynamic: &provider.DynamicAttribute{
									OptionalRequired: schema.Optional,
									Validators: schema.DynamicValidators{
										{
											Custom: &schema.CustomValidator{
												Imports: []code.Import{
													{
														Path: "github.com/my_account/my_project/mydynamicvalidator",
													},
												},
												SchemaDefinition: "mydynamicvalidator.Validate()",
											},
										},
									},
								},
							},
							{
								Name: "float64_attribute",
								Float64: &provider.Float64Attribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "float64_attribute_custom_type",
								Float64: &provider.Float64Attribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.Float64Type",
										ValueType: "basetypes.Float64Value",
									},
								},
							},
							{
								Name: "float64_attribute_associated_external_type",
								Float64: &provider.Float64Attribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "int64_attribute",
								Int64: &provider.Int64Attribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "int64_attribute_custom_type",
								Int64: &provider.Int64Attribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.Int64Type",
										ValueType: "basetypes.Int64Value",
									},
								},
							},
							{
								Name: "int64_attribute_associated_external_type",
								Int64: &provider.Int64Attribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "list_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "list_attribute_custom_type",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.ListType",
										ValueType: "basetypes.ListValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "list_attribute_element_type_string_custom_type",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "list_map_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Map: &schema.MapType{
											ElementType: schema.ElementType{
												String: &schema.StringType{},
											},
										},
									},
								},
							},
							{
								Name: "list_object_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Object: &schema.ObjectType{
											AttributeTypes: []schema.ObjectAttributeType{
												{
													Name:   "obj_string_attr",
													String: &schema.StringType{},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_object_object_attribute",
								List: &provider.ListAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										Object: &schema.ObjectType{
											AttributeTypes: []schema.ObjectAttributeType{
												{
													Name: "obj_obj_attr",
													Object: &schema.ObjectType{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_obj_string_attr",
																String: &schema.StringType{},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_attribute_associated_external_type",
								List: &provider.ListAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute_custom_type",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.MapType",
										ValueType: "basetypes.MapValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_attribute_element_type_string_custom_type",
								Map: &provider.MapAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "map_attribute_associated_external_type",
								Map: &provider.MapAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "map_nested_bool_attribute",
								MapNested: &provider.MapNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "map_nested_bool_attribute_associated_external_type",
								MapNested: &provider.MapNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "number_attribute",
								Number: &provider.NumberAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "number_attribute_custom_type",
								Number: &provider.NumberAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.NumberType",
										ValueType: "basetypes.NumberValue",
									},
								},
							},
							{
								Name: "number_attribute_associated_external_type",
								Number: &provider.NumberAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_dynamic",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:    "object_dynamic_attribute",
											Dynamic: &schema.DynamicType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_object_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_object_attr",
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_object_string_attr",
														String: &schema.StringType{},
													},
												},
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.ObjectType",
													ValueType: "basetypes.ObjectValue",
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_attribute_types_string_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_string_attr",
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_custom_type",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.ObjectType",
										ValueType: "basetypes.ObjectValue",
									},
								},
							},
							{
								Name: "object_list_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_list_attr",
											List: &schema.ListType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_list_object_attribute",
								Object: &provider.ObjectAttribute{
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name: "obj_list_attr",
											List: &schema.ListType{
												ElementType: schema.ElementType{
													Object: &schema.ObjectType{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_list_obj_attr",
																String: &schema.StringType{},
															},
														},
													},
												},
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "object_attribute_associated_external_type",
								Object: &provider.ObjectAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									AttributeTypes: []schema.ObjectAttributeType{
										{
											Name:   "obj_string_attr",
											String: &schema.StringType{},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "list_nested_bool_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_bool_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "list_nested_attribute",
												ListNested: &provider.ListNestedAttribute{
													OptionalRequired: schema.Optional,
													NestedObject: provider.NestedAttributeObject{
														Attributes: []provider.Attribute{
															{
																Name: "bool_attribute",
																Bool: &provider.BoolAttribute{
																	OptionalRequired: schema.Optional,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_list_attribute",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "list_nested_attribute",
												ListNested: &provider.ListNestedAttribute{
													OptionalRequired: schema.Optional,
													NestedObject: provider.NestedAttributeObject{
														Attributes: []provider.Attribute{
															{
																Name: "list_attribute",
																List: &provider.ListAttribute{
																	OptionalRequired: schema.Optional,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_bool_attribute_associated_external_type",
								ListNested: &provider.ListNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_attribute",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_attribute_custom_type",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.SetType",
										ValueType: "basetypes.SetValue",
									},
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_attribute_element_type_string_custom_type",
								Set: &provider.SetAttribute{
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{
											CustomType: &schema.CustomType{
												Import: &code.Import{
													Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
												},
												Type:      "basetypes.StringType",
												ValueType: "basetypes.StringValue",
											},
										},
									},
								},
							},
							{
								Name: "set_attribute_associated_external_type",
								Set: &provider.SetAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
									ElementType: schema.ElementType{
										String: &schema.StringType{},
									},
								},
							},
							{
								Name: "set_nested_bool_attribute",
								SetNested: &provider.SetNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_bool_attribute_associated_external_type",
								SetNested: &provider.SetNestedAttribute{
									OptionalRequired: schema.Optional,
									NestedObject: provider.NestedAttributeObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_bool_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_single_nested_bool_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "single_nested_attribute",
											SingleNested: &provider.SingleNestedAttribute{
												Attributes: []provider.Attribute{
													{
														Name: "bool_attribute",
														Bool: &provider.BoolAttribute{
															OptionalRequired: schema.Optional,
														},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_single_nested_list_attribute",
								SingleNested: &provider.SingleNestedAttribute{
									Attributes: []provider.Attribute{
										{
											Name: "single_nested_attribute",
											SingleNested: &provider.SingleNestedAttribute{
												Attributes: []provider.Attribute{
													{
														Name: "list_attribute",
														List: &provider.ListAttribute{
															OptionalRequired: schema.Optional,
															ElementType: schema.ElementType{
																String: &schema.StringType{},
															},
														},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "single_nested_bool_attribute_associated_external_type",
								SingleNested: &provider.SingleNestedAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "string_attribute",
								String: &provider.StringAttribute{
									OptionalRequired: schema.Optional,
								},
							},
							{
								Name: "string_attribute_custom_type",
								String: &provider.StringAttribute{
									OptionalRequired: schema.Optional,
									CustomType: &schema.CustomType{
										Import: &code.Import{
											Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
										},
										Type:      "basetypes.StringType",
										ValueType: "basetypes.StringValue",
									},
								},
							},
							{
								Name: "string_attribute_associated_external_type",
								String: &provider.StringAttribute{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									OptionalRequired: schema.Optional,
								},
							},
						},
						Blocks: []provider.Block{
							{
								Name: "list_nested_block_bool_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_list_nested_block_bool_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Blocks: []provider.Block{
											{
												Name: "list_nested_block",
												ListNested: &provider.ListNestedBlock{
													NestedObject: provider.NestedBlockObject{
														Attributes: []provider.Attribute{
															{
																Name: "bool_attribute",
																Bool: &provider.BoolAttribute{
																	OptionalRequired: schema.Optional,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "object_attribute",
												Object: &provider.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													OptionalRequired: schema.Optional,
												},
											},
										},
										Blocks: []provider.Block{
											{
												Name: "list_nested_block",
												ListNested: &provider.ListNestedBlock{
													NestedObject: provider.NestedBlockObject{
														Attributes: []provider.Attribute{
															{
																Name: "list_attribute",
																List: &provider.ListAttribute{
																	OptionalRequired: schema.Optional,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "list_nested_block_bool_attribute_associated_external_type",
								ListNested: &provider.ListNestedBlock{
									NestedObject: provider.NestedBlockObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_block_bool_attribute",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "set_nested_block_bool_attribute_associated_external_type",
								SetNested: &provider.SetNestedBlock{
									NestedObject: provider.NestedBlockObject{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []provider.Attribute{
											{
												Name: "bool_attribute",
												Bool: &provider.BoolAttribute{
													OptionalRequired: schema.Optional,
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_bool_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
								},
							},
							{
								Name: "single_nested_single_nested_block_bool_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Blocks: []provider.Block{
										{
											Name: "single_nested_block",
											SingleNested: &provider.SingleNestedBlock{
												Attributes: []provider.Attribute{
													{
														Name: "bool_attribute",
														Bool: &provider.BoolAttribute{
															OptionalRequired: schema.Optional,
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
								SingleNested: &provider.SingleNestedBlock{
									Attributes: []provider.Attribute{
										{
											Name: "object_attribute",
											Object: &provider.ObjectAttribute{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
												OptionalRequired: schema.Optional,
											},
										},
									},
									Blocks: []provider.Block{
										{
											Name: "list_nested_block",
											ListNested: &provider.ListNestedBlock{
												NestedObject: provider.NestedBlockObject{
													Attributes: []provider.Attribute{
														{
															Name: "list_attribute",
															List: &provider.ListAttribute{
																OptionalRequired: schema.Optional,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Name: "single_nested_block_bool_attribute_associated_external_type",
								SingleNested: &provider.SingleNestedBlock{
									AssociatedExternalType: &schema.AssociatedExternalType{
										Import: &code.Import{
											Path: "example.com/apisdk",
										},
										Type: "*apisdk.Type",
									},
									Attributes: []provider.Attribute{
										{
											Name: "bool_attribute",
											Bool: &provider.BoolAttribute{
												OptionalRequired: schema.Optional,
											},
										},
									},
								},
							},
						},
						MarkdownDescription: pointer("*This* is a description"),
						Description:         pointer("This is a description"),
						DeprecationMessage:  pointer("This provider is deprecated"),
					},
				},
				Resources: resource.Resources{
					{
						Name: "resource",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										PlanModifiers: schema.BoolPlanModifiers{
											{
												Custom: &schema.CustomPlanModifier{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/boolplanmodifier",
														},
													},
													SchemaDefinition: "myboolplanmodifier.Modify()",
												},
											},
										},
										Validators: schema.BoolValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/myboolvalidator",
														},
													},
													SchemaDefinition: "myboolvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "bool_attribute_custom_type",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.BoolType",
											ValueType: "basetypes.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_custom_type_import_alias",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.BoolType",
											ValueType: "fwtype.BoolValue",
										},
									},
								},
								{
									Name: "bool_attribute_default_static",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.BoolDefault{
											Static: pointer(true),
										},
									},
								},
								{
									Name: "bool_attribute_associated_external_type",
									Bool: &resource.BoolAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Optional,
									},
								},
								{
									Name: "dynamic_attribute",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "dynamic_attribute_associated_external_type",
									Dynamic: &resource.DynamicAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Optional,
									},
								},
								{
									Name: "dynamic_attribute_custom_type",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.DynamicType",
											ValueType: "basetypes.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_custom_type_import_alias",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Alias: pointer("fwtype"),
												Path:  "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "fwtype.DynamicType",
											ValueType: "fwtype.DynamicValue",
										},
									},
								},
								{
									Name: "dynamic_attribute_default",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.DynamicDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "dynamicdefault.StaticValue(types.StringValue(\"example\"))",
											},
										},
									},
								},
								{
									Name: "dynamic_attribute_plan_modifiers",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										PlanModifiers: schema.DynamicPlanModifiers{
											{
												Custom: &schema.CustomPlanModifier{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/dynamicplanmodifier",
														},
													},
													SchemaDefinition: "mydynamicplanmodifier.Modify()",
												},
											},
										},
									},
								},
								{
									Name: "dynamic_attribute_validators",
									Dynamic: &resource.DynamicAttribute{
										ComputedOptionalRequired: schema.Computed,
										Validators: schema.DynamicValidators{
											{
												Custom: &schema.CustomValidator{
													Imports: []code.Import{
														{
															Path: "github.com/my_account/my_project/mydynamicvalidator",
														},
													},
													SchemaDefinition: "mydynamicvalidator.Validate()",
												},
											},
										},
									},
								},
								{
									Name: "float64_attribute",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "float64_attribute_custom_type",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Float64Type",
											ValueType: "basetypes.Float64Value",
										},
									},
								},
								{
									Name: "float64_attribute_default_static",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.Float64Default{
											Static: pointer(123.45),
										},
									},
								},
								{
									Name: "float64_attribute_associated_external_type",
									Float64: &resource.Float64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "int64_attribute_custom_type",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.Int64Type",
											ValueType: "basetypes.Int64Value",
										},
									},
								},
								{
									Name: "int64_attribute_default_static",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.Int64Default{
											Static: pointer(int64(123)),
										},
									},
								},
								{
									Name: "int64_attribute_associated_external_type",
									Int64: &resource.Int64Attribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_custom_type",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ListType",
											ValueType: "basetypes.ListValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_default_custom",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.ListDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/attr",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "listdefault.StaticValue(types.ListValueMust(types.String, []attr.Value{types.StringValue(\"example\")}))",
											},
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "list_attribute_element_type_string_custom_type",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "list_map_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Map: &schema.MapType{
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
									},
								},
								{
									Name: "list_object_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name:   "obj_string_attr",
														String: &schema.StringType{},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_object_object_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											Object: &schema.ObjectType{
												AttributeTypes: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_obj_string_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_attribute_associated_external_type",
									List: &resource.ListAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_custom_type",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.MapType",
											ValueType: "basetypes.MapValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_attribute_element_type_string_custom_type",
									Map: &resource.MapAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "map_attribute_associated_external_type",
									Map: &resource.MapAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute",
									MapNested: &resource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "map_nested_bool_attribute_associated_external_type",
									MapNested: &resource.MapNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "number_attribute",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "number_attribute_custom_type",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.NumberType",
											ValueType: "basetypes.NumberValue",
										},
									},
								},
								{
									Name: "number_attribute_default_custom",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.NumberDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "math/big",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault",
													},
												},
												SchemaDefinition: "numberdefault.StaticBigFloat(big.NewFloat(123.45))",
											},
										},
									},
								},
								{
									Name: "number_attribute_associated_external_type",
									Number: &resource.NumberAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_dynamic",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:    "object_dynamic_attribute",
												Dynamic: &schema.DynamicType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_object_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_object_attr",
												Object: &schema.ObjectType{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_object_string_attr",
															String: &schema.StringType{},
														},
													},
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.ObjectType",
														ValueType: "basetypes.ObjectValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_attribute_types_string_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												String: &schema.StringType{
													CustomType: &schema.CustomType{
														Import: &code.Import{
															Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
														},
														Type:      "basetypes.StringType",
														ValueType: "basetypes.StringValue",
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_custom_type",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.ObjectType",
											ValueType: "basetypes.ObjectValue",
										},
									},
								},
								{
									Name: "object_list_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														String: &schema.StringType{},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												List: &schema.ListType{
													ElementType: schema.ElementType{
														Object: &schema.ObjectType{
															AttributeTypes: []schema.ObjectAttributeType{
																{
																	Name:   "obj_list_obj_attr",
																	String: &schema.StringType{},
																},
															},
														},
													},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute_associated_external_type",
									Object: &resource.ObjectAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name:   "obj_string_attr",
												String: &schema.StringType{},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "list_nested_bool_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_bool_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &resource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: resource.NestedAttributeObject{
															Attributes: []resource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &resource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_list_attribute",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "list_nested_attribute",
													ListNested: &resource.ListNestedAttribute{
														ComputedOptionalRequired: schema.Computed,
														NestedObject: resource.NestedAttributeObject{
															Attributes: []resource.Attribute{
																{
																	Name: "list_attribute",
																	List: &resource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_bool_attribute_associated_external_type",
									ListNested: &resource.ListNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_attribute",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_custom_type",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.SetType",
											ValueType: "basetypes.SetValue",
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_default_custom",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.SetDefault{
											Custom: &schema.CustomDefault{
												Imports: []code.Import{
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/attr",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault",
													},
													{
														Path: "github.com/hashicorp/terraform-plugin-framework/types",
													},
												},
												SchemaDefinition: "setdefault.StaticValue(types.SetValueMust(types.String, []attr.Value{types.StringValue(\"example\")}))",
											},
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_attribute_element_type_string_custom_type",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{
												CustomType: &schema.CustomType{
													Import: &code.Import{
														Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
													},
													Type:      "basetypes.StringType",
													ValueType: "basetypes.StringValue",
												},
											},
										},
									},
								},
								{
									Name: "set_attribute_associated_external_type",
									Set: &resource.SetAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute",
									SetNested: &resource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_bool_attribute_associated_external_type",
									SetNested: &resource.SetNestedAttribute{
										ComputedOptionalRequired: schema.Computed,
										NestedObject: resource.NestedAttributeObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_bool_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_bool_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &resource.SingleNestedAttribute{
													Attributes: []resource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &resource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_single_nested_list_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "single_nested_attribute",
												SingleNested: &resource.SingleNestedAttribute{
													Attributes: []resource.Attribute{
														{
															Name: "list_attribute",
															List: &resource.ListAttribute{
																ComputedOptionalRequired: schema.Computed,
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
															},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "single_nested_bool_attribute_associated_external_type",
									SingleNested: &resource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "string_attribute_custom_type",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import: &code.Import{
												Path: "github.com/hashicorp/terraform-plugin-framework/types/basetypes",
											},
											Type:      "basetypes.StringType",
											ValueType: "basetypes.StringValue",
										},
									},
								},
								{
									Name: "string_attribute_default_static",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.StringDefault{
											Static: pointer("example"),
										},
									},
								},
								{
									Name: "string_attribute_associated_external_type",
									String: &resource.StringAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
							},
							Blocks: []resource.Block{
								{
									Name: "list_nested_block_bool_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_list_nested_block_bool_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "list_nested_block",
													ListNested: &resource.ListNestedBlock{
														NestedObject: resource.NestedBlockObject{
															Attributes: []resource.Attribute{
																{
																	Name: "bool_attribute",
																	Bool: &resource.BoolAttribute{
																		ComputedOptionalRequired: schema.Computed,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_object_attribute_list_nested_nested_block_list_attribute",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "object_attribute",
													Object: &resource.ObjectAttribute{
														AttributeTypes: []schema.ObjectAttributeType{
															{
																Name:   "obj_string_attr",
																String: &schema.StringType{},
															},
														},
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
											Blocks: []resource.Block{
												{
													Name: "list_nested_block",
													ListNested: &resource.ListNestedBlock{
														NestedObject: resource.NestedBlockObject{
															Attributes: []resource.Attribute{
																{
																	Name: "list_attribute",
																	List: &resource.ListAttribute{
																		ComputedOptionalRequired: schema.Computed,
																		ElementType: schema.ElementType{
																			String: &schema.StringType{},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "list_nested_block_bool_attribute_associated_external_type",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "set_nested_block_bool_attribute_associated_external_type",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											AssociatedExternalType: &schema.AssociatedExternalType{
												Import: &code.Import{
													Path: "example.com/apisdk",
												},
												Type: "*apisdk.Type",
											},
											Attributes: []resource.Attribute{
												{
													Name: "bool_attribute",
													Bool: &resource.BoolAttribute{
														ComputedOptionalRequired: schema.Computed,
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
								{
									Name: "single_nested_single_nested_block_bool_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "single_nested_block",
												SingleNested: &resource.SingleNestedBlock{
													Attributes: []resource.Attribute{
														{
															Name: "bool_attribute",
															Bool: &resource.BoolAttribute{
																ComputedOptionalRequired: schema.Computed,
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_object_attribute_single_nested_list_nested_block_list_attribute",
									SingleNested: &resource.SingleNestedBlock{
										Attributes: []resource.Attribute{
											{
												Name: "object_attribute",
												Object: &resource.ObjectAttribute{
													AttributeTypes: []schema.ObjectAttributeType{
														{
															Name:   "obj_string_attr",
															String: &schema.StringType{},
														},
													},
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
										Blocks: []resource.Block{
											{
												Name: "list_nested_block",
												ListNested: &resource.ListNestedBlock{
													NestedObject: resource.NestedBlockObject{
														Attributes: []resource.Attribute{
															{
																Name: "list_attribute",
																List: &resource.ListAttribute{
																	ComputedOptionalRequired: schema.Computed,
																	ElementType: schema.ElementType{
																		String: &schema.StringType{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Name: "single_nested_block_bool_attribute_associated_external_type",
									SingleNested: &resource.SingleNestedBlock{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: &code.Import{
												Path: "example.com/apisdk",
											},
											Type: "*apisdk.Type",
										},
										Attributes: []resource.Attribute{
											{
												Name: "bool_attribute",
												Bool: &resource.BoolAttribute{
													ComputedOptionalRequired: schema.Computed,
												},
											},
										},
									},
								},
							},
							MarkdownDescription: pointer("*This* is a description"),
							Description:         pointer("This is a description"),
							DeprecationMessage:  pointer("This resource is deprecated"),
							Version:             pointer(int64(1)),
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := spec.Parse(context.Background(), testCase.data)

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
