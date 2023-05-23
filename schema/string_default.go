package schema

type StringDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *string        `json:"static,omitempty"`
}
