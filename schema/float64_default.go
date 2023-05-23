package schema

type Float64Default struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *float64       `json:"static,omitempty"`
}
