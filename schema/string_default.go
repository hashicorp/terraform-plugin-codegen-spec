// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type StringDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *string        `json:"static,omitempty"`
}
