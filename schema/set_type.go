package schema

type SetType struct {
	ElementType

	// CustomType is a customization of the SetType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
