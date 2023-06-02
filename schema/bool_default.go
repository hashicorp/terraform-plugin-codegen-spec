// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type BoolDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *bool          `json:"static,omitempty"`
}
