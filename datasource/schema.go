package datasource

type Schema struct {
	Attributes []Attribute `json:"attributes,omitempty"`
	Blocks     []Block     `json:"blocks,omitempty"`
}
