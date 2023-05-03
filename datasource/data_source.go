package datasource

type DataSource struct {
	Name string `json:"name"`

	Schema *Schema `json:"schema,omitempty"`
}
