// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ObjectValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}
