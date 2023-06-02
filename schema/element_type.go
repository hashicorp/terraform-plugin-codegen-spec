// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ElementType struct {
	Bool    *BoolType             `json:"bool,omitempty"`
	Float64 *Float64Type          `json:"float64,omitempty"`
	Int64   *Int64Type            `json:"int64,omitempty"`
	List    *ListType             `json:"list,omitempty"`
	Map     *MapType              `json:"map,omitempty"`
	Number  *NumberType           `json:"number,omitempty"`
	Object  []ObjectAttributeType `json:"object,omitempty"`
	Set     *SetType              `json:"set,omitempty"`
	String  *StringType           `json:"string,omitempty"`
}
