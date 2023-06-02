// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type SetValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}
