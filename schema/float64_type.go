// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package schema

// Float64Type is a representation of a 64-bit floating point number.
type Float64Type struct {
	// CustomType is a customization of the Float64Type.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
