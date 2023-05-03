package schema

type CustomPlanModifier struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}
