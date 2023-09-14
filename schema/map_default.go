// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type MapDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
}

func (d *MapDefault) Equal(other *MapDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	return d.Custom.Equal(other.Custom)
}
