// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type Float64Default struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *float64       `json:"static,omitempty"`
}
