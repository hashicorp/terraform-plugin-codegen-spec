package resource

import "github.com/hashicorp/terraform-plugin-codegen-spec/schema"

type Attribute struct {
	Name string `json:"name"`

	Bool *BoolAttribute `json:"bool,omitempty"`
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
