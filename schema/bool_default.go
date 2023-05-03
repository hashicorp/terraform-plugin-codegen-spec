package schema

type BoolDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *bool          `json:"static,omitempty"`
}
