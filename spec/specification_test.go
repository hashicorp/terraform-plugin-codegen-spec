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
													Name:   "obj_string_attr",
													String: &schema.StringType{},
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
													Object: []schema.ObjectAttributeType{
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
												Name:   "obj_string_attr",
												String: &schema.StringType{},
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
														Object: []schema.ObjectAttributeType{
															{
																Name:   "obj_list_obj_attr",
																String: &schema.StringType{},
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
										Default: &schema.BoolDefault{
											Static: pointer(true),
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
