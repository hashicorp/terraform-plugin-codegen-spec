// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type BoolDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *bool          `json:"static,omitempty"`
}

func (d *BoolDefault) Equal(other *BoolDefault) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	if !d.Custom.Equal(other.Custom) {
		return false
	}

	if d.Static == nil && other.Static != nil {
		return false
	}

	if d.Static != nil && other.Static == nil {
		return false
	}

	return *d.Static == *other.Static
}
