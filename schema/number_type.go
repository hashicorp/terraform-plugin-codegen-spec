package schema

type NumberType struct {
	// CustomType is a customization of the NumberType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
