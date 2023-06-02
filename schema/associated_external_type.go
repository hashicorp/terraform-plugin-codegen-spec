package schema

type AssociatedExternalType struct {
	Import *string `json:"import,omitempty"`
	Type   string  `json:"type"`
}

func (a AssociatedExternalType) HasImport() bool {
	return a.Import != nil && *a.Import != ""
}
