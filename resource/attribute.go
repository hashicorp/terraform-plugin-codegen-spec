package resource

import "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type Attribute struct {
	Name string `json:"name"`

	Bool         *BoolAttribute         `json:"bool,omitempty"`
	Float64      *Float64Attribute      `json:"float64,omitempty"`
	Int64        *Int64Attribute        `json:"int64,omitempty"`
	List         *ListAttribute         `json:"list,omitempty"`
	ListNested   *ListNestedAttribute   `json:"list_nested,omitempty"`
	Map          *MapAttribute          `json:"map,omitempty"`
	MapNested    *MapNestedAttribute    `json:"map_nested,omitempty"`
	Number       *NumberAttribute       `json:"number,omitempty"`
	Object       *ObjectAttribute       `json:"object,omitempty"`
	Set          *SetAttribute          `json:"set,omitempty"`
	SetNested    *SetNestedAttribute    `json:"set_nested,omitempty"`
	SingleNested *SingleNestedAttribute `json:"single_nested,omitempty"`
	String       *StringAttribute       `json:"string,omitempty"`
}

type NestedAttributeObject struct {
	Attributes []Attribute `json:"attributes,omitempty"`

	CustomType    *schema.CustomType          `json:"custom_type,omitempty"`
	PlanModifiers []schema.ObjectPlanModifier `json:"plan_modifiers,omitempty"`
	Validators    []schema.ObjectValidator    `json:"validators,omitempty"`
}

type BoolAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.BoolDefault            `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.BoolPlanModifier      `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.BoolValidator         `json:"validators,omitempty"`
}

type Float64Attribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.Float64Default         `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.Float64PlanModifier   `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Float64Validator      `json:"validators,omitempty"`
}

type Int64Attribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.Int64Default           `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.Int64PlanModifier     `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Int64Validator        `json:"validators,omitempty"`
}

type ListAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.ListDefault            `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.ListPlanModifier      `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ListValidator         `json:"validators,omitempty"`
}

type ListNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType        `json:"custom_type,omitempty"`
	Default            *schema.ListDefault       `json:"default,omitempty"`
	DeprecationMessage *string                   `json:"deprecation_message,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	PlanModifiers      []schema.ListPlanModifier `json:"plan_modifiers,omitempty"`
	Sensitive          *bool                     `json:"sensitive,omitempty"`
	Validators         []schema.ListValidator    `json:"validators,omitempty"`
}

type MapAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.MapDefault             `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.MapPlanModifier       `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.MapValidator          `json:"validators,omitempty"`
}

type MapNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType       `json:"custom_type,omitempty"`
	Default            *schema.MapDefault       `json:"default,omitempty"`
	DeprecationMessage *string                  `json:"deprecation_message,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	PlanModifiers      []schema.MapPlanModifier `json:"plan_modifiers,omitempty"`
	Sensitive          *bool                    `json:"sensitive,omitempty"`
	Validators         []schema.MapValidator    `json:"validators,omitempty"`
}

type NumberAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.NumberDefault          `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.NumberPlanModifier    `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.NumberValidator       `json:"validators,omitempty"`
}

type ObjectAttribute struct {
	AttributeTypes           []schema.ObjectAttributeType    `json:"attribute_types"`
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.ObjectDefault          `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.ObjectPlanModifier    `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type SetAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.SetDefault             `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.SetPlanModifier       `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.SetValidator          `json:"validators,omitempty"`
}

type SetNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType       `json:"custom_type,omitempty"`
	Default            *schema.SetDefault       `json:"default,omitempty"`
	DeprecationMessage *string                  `json:"deprecation_message,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	PlanModifiers      []schema.SetPlanModifier `json:"plan_modifiers,omitempty"`
	Sensitive          *bool                    `json:"sensitive,omitempty"`
	Validators         []schema.SetValidator    `json:"validators,omitempty"`
}

type SingleNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	Attributes               []Attribute                     `json:"attributes,omitempty"`
	AssociatedExternalType   *schema.AssociatedExternalType  `json:"associated_external_type,omitempty"`
	CustomType               *schema.CustomType              `json:"custom_type,omitempty"`
	Default                  *schema.ObjectDefault           `json:"default,omitempty"`
	DeprecationMessage       *string                         `json:"deprecation_message,omitempty"`
	Description              *string                         `json:"description,omitempty"`
	PlanModifiers            []schema.ObjectPlanModifier     `json:"plan_modifiers,omitempty"`
	Sensitive                *bool                           `json:"sensitive,omitempty"`
	Validators               []schema.ObjectValidator        `json:"validators,omitempty"`
}

type StringAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	Default                *schema.StringDefault          `json:"default,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	PlanModifiers          []schema.StringPlanModifier    `json:"plan_modifiers,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.StringValidator       `json:"validators,omitempty"`
}
