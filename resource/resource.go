package resource

type Resource struct {
	Name string `json:"name"`

	Schema *Schema `json:"schema,omitempty"`
}
