package schema

type CustomValidator struct {
	Import           *string `json:"import,omitempty"`
	SchemaDefinition string  `json:"schema_definition"`
}

func (c CustomValidator) HasImport() bool {
	return c.Import != nil && *c.Import != ""
}
