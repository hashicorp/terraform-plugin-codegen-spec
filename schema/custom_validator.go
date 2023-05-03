package schema

type CustomValidator struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}
