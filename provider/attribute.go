package provider

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

	CustomType *schema.CustomType       `json:"custom_type,omitempty"`
	Validators []schema.ObjectValidator `json:"validators,omitempty"`
}

type BoolAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.BoolValidator         `json:"validators,omitempty"`
}

type Float64Attribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Float64Validator      `json:"validators,omitempty"`
}

type Int64Attribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.Int64Validator        `json:"validators,omitempty"`
}

type ListAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ListValidator         `json:"validators,omitempty"`
}

type ListNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType     `json:"custom_type,omitempty"`
	DeprecationMessage *string                `json:"deprecation_message,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Sensitive          *bool                  `json:"sensitive,omitempty"`
	Validators         []schema.ListValidator `json:"validators,omitempty"`
}

type MapAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.MapValidator          `json:"validators,omitempty"`
}

type MapNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.MapValidator `json:"validators,omitempty"`
}

type NumberAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.NumberValidator       `json:"validators,omitempty"`
}

type ObjectAttribute struct {
	AttributeTypes   []schema.ObjectAttributeType `json:"attribute_types"`
	OptionalRequired schema.OptionalRequired      `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type SetAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	ElementType      schema.ElementType      `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.SetValidator          `json:"validators,omitempty"`
}

type SetNestedAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedAttributeObject   `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.SetValidator `json:"validators,omitempty"`
}

type SingleNestedAttribute struct {
	OptionalRequired       schema.OptionalRequired        `json:"optional_required"`
	Attributes             []Attribute                    `json:"attributes,omitempty"`
	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type StringAttribute struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.StringValidator       `json:"validators,omitempty"`
}
