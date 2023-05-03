package schema

type CustomType struct {
	Import    *string `json:"import,omitempty"`
	Type      string  `json:"type"`
	ValueType string  `json:"value_type"`
}
