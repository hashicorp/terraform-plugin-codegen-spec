// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type StringPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}
