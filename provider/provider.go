package provider

type Provider struct {
	Name string `json:"name"`

	Schema *Schema `json:"schema,omitempty"`
}
