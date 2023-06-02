// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type Int64Default struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *int64         `json:"static,omitempty"`
}
