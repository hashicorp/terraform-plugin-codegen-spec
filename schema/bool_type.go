// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package schema

// BoolType is a representation of a boolean.
type BoolType struct {
	// CustomType is a customization of the BoolType.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
