package schema

type ListType struct {
	ElementType

	// CustomType is a customization of the ListType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
