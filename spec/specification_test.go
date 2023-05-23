package spec_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestSpecification_JSONUnmarshal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		data          []byte
		expected      spec.Specification
		expectedError error
	}{
		"example": {
			data: testReadFile("example.json"),
			expected: spec.Specification{
				DataSources: []datasource.DataSource{
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
									Name: "float64_attribute",
									Float64: &datasource.Float64Attribute{
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
									Name: "list_attribute",
									List: &datasource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
										ElementType: schema.ElementType{
											String: &schema.StringType{},
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
											Object: []schema.ObjectAttributeType{
												{
													Name: "obj_string_attr",
													ElementType: schema.ElementType{
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
											Object: []schema.ObjectAttributeType{
												{
													Name: "obj_obj_attr",
													ElementType: schema.ElementType{
														Object: []schema.ObjectAttributeType{
															{
																Name: "obj_obj_string_attr",
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
									Name: "number_attribute",
									Number: &datasource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_attribute",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												ElementType: schema.ElementType{
													List: &schema.ListType{
														ElementType: schema.ElementType{
															String: &schema.StringType{},
														},
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
												ElementType: schema.ElementType{
													List: &schema.ListType{
														ElementType: schema.ElementType{
															Object: []schema.ObjectAttributeType{
																{
																	Name: "obj_list_obj_attr",
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
									Name: "set_attribute",
									Set: &datasource.SetAttribute{
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
									Name: "single_nested_bool_attribute",
									SingleNested: &datasource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: pointer("example.com/apisdk"),
											Type:   "*apisdk.DataSourceProperty",
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
									Name: "string_attribute",
									String: &datasource.StringAttribute{
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
																Name: "obj_string_attr",
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
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
															Name: "obj_string_attr",
															ElementType: schema.ElementType{
																String: &schema.StringType{},
															},
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
						},
					},
				},
				Provider: &provider.Provider{
					Name: "provider",
				},
				Resources: []resource.Resource{
					{
						Name: "resource",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "bool_attribute",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: schema.Computed,
										CustomType: &schema.CustomType{
											Import:    pointer(""),
											Type:      "",
											ValueType: "",
										},
										PlanModifiers: []schema.BoolPlanModifier{
											{},
											{
												Custom: &schema.CustomPlanModifier{
													Import:           pointer("github.com/my_account/my_project/boolplanmodifier"),
													SchemaDefinition: "myboolplanmodifier.Modify()",
												},
											},
										},
										Validators: []schema.BoolValidator{
											{
												Custom: &schema.CustomValidator{
													Import:           pointer("github.com/my_account/my_project/myboolvalidator"),
													SchemaDefinition: "myboolvalidator.Validate()",
												},
											},
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
									Name: "float64_attribute",
									Float64: &resource.Float64Attribute{
										ComputedOptionalRequired: schema.Computed,
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
									Name: "int64_attribute",
									Int64: &resource.Int64Attribute{
										ComputedOptionalRequired: schema.Computed,
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
									Name: "list_attribute",
									List: &resource.ListAttribute{
										ComputedOptionalRequired: schema.Computed,
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
												Import:           pointer("github.com/hashicorp/terraform-plugin-framework/types"),
												SchemaDefinition: "types.ListValueMust(types.String, []attr.Value{types.StringValue(\"example\")})",
											},
										},
										ElementType: schema.ElementType{
											String: &schema.StringType{},
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
											Object: []schema.ObjectAttributeType{
												{
													Name: "obj_string_attr",
													ElementType: schema.ElementType{
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
											Object: []schema.ObjectAttributeType{
												{
													Name: "obj_obj_attr",
													ElementType: schema.ElementType{
														Object: []schema.ObjectAttributeType{
															{
																Name: "obj_obj_string_attr",
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
									Name: "number_attribute",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "number_attribute_default_custom",
									Number: &resource.NumberAttribute{
										ComputedOptionalRequired: schema.Optional,
										Default: &schema.NumberDefault{
											Custom: &schema.CustomDefault{
												Import:           pointer("math/big"),
												SchemaDefinition: "big.NewFloat(123.45)",
											},
										},
									},
								},
								{
									Name: "object_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_string_attr",
												ElementType: schema.ElementType{
													String: &schema.StringType{},
												},
											},
										},
										ComputedOptionalRequired: schema.Computed,
									},
								},
								{
									Name: "object_list_attribute",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_list_attr",
												ElementType: schema.ElementType{
													List: &schema.ListType{
														ElementType: schema.ElementType{
															String: &schema.StringType{},
														},
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
												ElementType: schema.ElementType{
													List: &schema.ListType{
														ElementType: schema.ElementType{
															Object: []schema.ObjectAttributeType{
																{
																	Name: "obj_list_obj_attr",
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
									Name: "set_attribute",
									Set: &resource.SetAttribute{
										ComputedOptionalRequired: schema.Computed,
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
												Import:           pointer("github.com/hashicorp/terraform-plugin-framework/types"),
												SchemaDefinition: "types.SetValueMust(types.String, []attr.Value{types.StringValue(\"example\")})",
											},
										},
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
									Name: "single_nested_bool_attribute",
									SingleNested: &resource.SingleNestedAttribute{
										AssociatedExternalType: &schema.AssociatedExternalType{
											Import: pointer("example.com/apisdk"),
											Type:   "*apisdk.DataSourceProperty",
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
									Name: "string_attribute",
									String: &resource.StringAttribute{
										ComputedOptionalRequired: schema.Computed,
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
																Name: "obj_string_attr",
																ElementType: schema.ElementType{
																	String: &schema.StringType{},
																},
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
															Name: "obj_string_attr",
															ElementType: schema.ElementType{
																String: &schema.StringType{},
															},
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
