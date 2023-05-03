package schema

type CustomDefault struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}
