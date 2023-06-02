// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ObjectDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
}
