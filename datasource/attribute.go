package datasource

import "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type Attribute struct {
	Name string `json:"name"`

	Bool         *BoolAttribute         `json:"bool,omitempty"`
	List         *ListAttribute         `json:"list,omitempty"`
	ListNested   *ListNestedAttribute   `json:"list_nested,omitempty"`
	Object       *ObjectAttribute       `json:"object,omitempty"`
	SingleNested *SingleNestedAttribute `json:"single_nested,omitempty"`
}

type NestedAttributeObject struct {
	Attributes []Attribute `json:"attributes,omitempty"`

	CustomType *schema.CustomType       `json:"custom_type,omitempty"`
	Validators []schema.ObjectValidator `json:"validators,omitempty"`
}

type BoolAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.BoolValidator         `json:"validators,omitempty"`
}

type ListAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	ElementType              schema.ElementType              `json:"element_type"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ListValidator         `json:"validators,omitempty"`
}

type ListNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	NestedObject             NestedAttributeObject           `json:"nested_object"`

	CustomType         *schema.CustomType `json:"custom_type,omitempty"`
	DeprecationMessage *string            `json:"deprecation_message,omitempty"`
	Description        *string            `json:"description,omitempty"`
	Sensitive          *bool              `json:"sensitive,omitempty"`
}

type ObjectAttribute struct {
	AttributeTypes           []schema.ObjectAttributeType    `json:"attribute_types"`
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`

	AssociatedExternalType *schema.AssociatedExternalType `json:"associated_external_type,omitempty"`
	CustomType             *schema.CustomType             `json:"custom_type,omitempty"`
	DeprecationMessage     *string                        `json:"deprecation_message,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	Sensitive              *bool                          `json:"sensitive,omitempty"`
	Validators             []schema.ObjectValidator       `json:"validators,omitempty"`
}

type SingleNestedAttribute struct {
	ComputedOptionalRequired schema.ComputedOptionalRequired `json:"computed_optional_required"`
	Attributes               []Attribute                     `json:"attributes,omitempty"`
	AssociatedExternalType   *schema.AssociatedExternalType  `json:"associated_external_type,omitempty"`
	CustomType               *schema.CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                         `json:"deprecation_message,omitempty"`
	Description              *string                         `json:"description,omitempty"`
	Sensitive                *bool                           `json:"sensitive,omitempty"`
	Validators               []schema.ObjectValidator        `json:"validators,omitempty"`
}
