// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type StringValidator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}
