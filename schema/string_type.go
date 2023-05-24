package schema

type StringType struct {
	// CustomType is a customization of the StringType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
