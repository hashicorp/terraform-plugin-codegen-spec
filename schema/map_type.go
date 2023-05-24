package schema

type MapType struct {
	ElementType

	// CustomType is a customization of the MapType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
