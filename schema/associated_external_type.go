package schema

type AssociatedExternalType struct {
	Import *string `json:"import,omitempty"`
	Type   string  `json:"type"`
}
