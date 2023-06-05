// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type MapValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}
