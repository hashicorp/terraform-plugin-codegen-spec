package provider

import "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type Block struct {
	Name string `json:"name"`

	ListNested   *ListNestedBlock   `json:"list_nested,omitempty"`
	SetNested    *SetNestedBlock    `json:"set_nested,omitempty"`
	SingleNested *SingleNestedBlock `json:"single_nested,omitempty"`
}

type NestedBlockObject struct {
	Attributes []Attribute `json:"attributes,omitempty"`
	Blocks     []Block     `json:"blocks,omitempty"`

	CustomType *schema.CustomType       `json:"custom_type,omitempty"`
	Validators []schema.ObjectValidator `json:"validators,omitempty"`
}

type ListNestedBlock struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedBlockObject       `json:"nested_object"`

	CustomType         *schema.CustomType     `json:"custom_type,omitempty"`
	DeprecationMessage *string                `json:"deprecation_message,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Sensitive          *bool                  `json:"sensitive,omitempty"`
	Validators         []schema.ListValidator `json:"validators,omitempty"`
}

type SetNestedBlock struct {
	OptionalRequired schema.OptionalRequired `json:"optional_required"`
	NestedObject     NestedBlockObject       `json:"nested_object"`

	CustomType         *schema.CustomType    `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Sensitive          *bool                 `json:"sensitive,omitempty"`
	Validators         []schema.SetValidator `json:"validators,omitempty"`
}

type SingleNestedBlock struct {
	Attributes       []Attribute             `json:"attributes,omitempty"`
	Blocks           []Block                 `json:"blocks,omitempty"`
	OptionalRequired schema.OptionalRequired `json:"optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}
