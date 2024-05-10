// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// ElementType defines the type within a list, map, or set.
//
// DynamicType is intentionally not supported as it can cause confusing behavior
// or unavoidable errors for practitioners if elements cannot be conformed into
// a single element type.
type ElementType struct {
	Bool    *BoolType    `json:"bool,omitempty"`
	Float64 *Float64Type `json:"float64,omitempty"`
	Int64   *Int64Type   `json:"int64,omitempty"`
	List    *ListType    `json:"list,omitempty"`
	Map     *MapType     `json:"map,omitempty"`
	Number  *NumberType  `json:"number,omitempty"`
	Object  *ObjectType  `json:"object,omitempty"`
	Set     *SetType     `json:"set,omitempty"`
	String  *StringType  `json:"string,omitempty"`
}

// Equal returns true if all fields of the given ElementType are equal.
func (e ElementType) Equal(other ElementType) bool {
	if e.Bool == nil && other.Bool != nil {
		return false
	}

	if e.Bool != nil && other.Bool == nil {
		return false
	}

	if e.Bool != nil && other.Bool != nil {
		return e.Bool.CustomType.Equal(other.Bool.CustomType)
	}

	if e.Float64 == nil && other.Float64 != nil {
		return false
	}

	if e.Float64 != nil && other.Float64 == nil {
		return false
	}

	if e.Float64 != nil && other.Float64 != nil {
		return e.Float64.CustomType.Equal(other.Float64.CustomType)
	}

	if e.Int64 == nil && other.Int64 != nil {
		return false
	}

	if e.Int64 != nil && other.Int64 == nil {
		return false
	}

	if e.Int64 != nil && other.Int64 != nil {
		return e.Int64.CustomType.Equal(other.Int64.CustomType)
	}

	if e.List == nil && other.List != nil {
		return false
	}

	if e.List != nil && other.List == nil {
		return false
	}

	if e.List != nil && other.List != nil {
		if !e.List.CustomType.Equal(other.List.CustomType) {
			return false
		}

		return e.List.ElementType.Equal(other.List.ElementType)
	}

	if e.Map == nil && other.Map != nil {
		return false
	}

	if e.Map != nil && other.Map == nil {
		return false
	}

	if e.Map != nil && other.Map != nil {
		if !e.Map.CustomType.Equal(other.Map.CustomType) {
			return false
		}

		return e.Map.ElementType.Equal(other.Map.ElementType)
	}

	if e.Number == nil && other.Number != nil {
		return false
	}

	if e.Number != nil && other.Number == nil {
		return false
	}

	if e.Number != nil && other.Number != nil {
		return e.Number.CustomType.Equal(other.Number.CustomType)
	}

	if e.Object == nil && other.Object != nil {
		return false
	}

	if e.Object != nil && other.Object == nil {
		return false
	}

	if e.Object != nil && other.Object != nil {
		return e.Object.Equal(other.Object)
	}

	if e.Set == nil && other.Set != nil {
		return false
	}

	if e.Set != nil && other.Set == nil {
		return false
	}

	if e.Set != nil && other.Set != nil {
		if !e.Set.CustomType.Equal(other.Set.CustomType) {
			return false
		}

		return e.Set.ElementType.Equal(other.Set.ElementType)
	}

	if e.String == nil && other.String != nil {
		return false
	}

	if e.String != nil && other.String == nil {
		return false
	}

	if e.String != nil && other.String != nil {
		return e.String.CustomType.Equal(other.String.CustomType)
	}

	return true
}
