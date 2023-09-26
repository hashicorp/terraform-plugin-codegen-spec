// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"context"
	"errors"
	"fmt"
	"sort"
)

// ObjectAttributeTypes type defines ObjectAttributeType types
type ObjectAttributeTypes []ObjectAttributeType

func (o ObjectAttributeTypes) Equal(other ObjectAttributeTypes) bool {
	if o == nil && other == nil {
		return true
	}

	if o == nil || other == nil {
		return false
	}

	if len(o) != len(other) {
		return false
	}

	// Name is required by the spec JSON schema.
	sort.Slice(o, func(i, j int) bool {
		return o[i].Name < o[j].Name
	})

	// Name is required by the spec JSON schema.
	sort.Slice(other, func(i, j int) bool {
		return other[i].Name < other[j].Name
	})

	for k, objectAttributeType := range o {
		if !objectAttributeType.Equal(other[k]) {
			return false
		}
	}

	return true
}

// Validate returns true if each of the object attribute names is unique within the object.
func (o ObjectAttributeTypes) Validate(ctx context.Context, req ObjectValidateRequest) error {
	attrTypeNames := make(map[string]struct{}, len(o))

	var errs, nestedErrs error

	for _, attributeType := range o {
		if _, ok := attrTypeNames[attributeType.Name]; ok {
			errs = errors.Join(errs, fmt.Errorf("%s object attribute type %q is duplicated", req.Path, attributeType.Name))
		}

		attrTypeNames[attributeType.Name] = struct{}{}

		if attributeType.Object != nil {
			objectValidateRequest := ObjectValidateRequest{
				Path: fmt.Sprintf("%s object attribute type %q", req.Path, attributeType.Name),
			}

			err := attributeType.Object.AttributeTypes.Validate(ctx, objectValidateRequest)

			nestedErrs = errors.Join(nestedErrs, err)
		}
	}

	return errors.Join(errs, nestedErrs)
}

// ObjectAttributeType defines the types within an object.
type ObjectAttributeType struct {
	Name string `json:"name"`

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

// Equal returns true if all fields of the given ObjectAttributeType are equal.
func (o ObjectAttributeType) Equal(other ObjectAttributeType) bool {
	if o.Name != other.Name {
		return false
	}

	if o.Bool == nil && other.Bool != nil {
		return false
	}

	if o.Bool != nil && other.Bool == nil {
		return false
	}

	if o.Bool != nil && other.Bool != nil {
		return o.Bool.CustomType.Equal(other.Bool.CustomType)
	}

	if o.Float64 == nil && other.Float64 != nil {
		return false
	}

	if o.Float64 != nil && other.Float64 == nil {
		return false
	}

	if o.Float64 != nil && other.Float64 != nil {
		return o.Float64.CustomType.Equal(other.Float64.CustomType)
	}

	if o.Int64 == nil && other.Int64 != nil {
		return false
	}

	if o.Int64 != nil && other.Int64 == nil {
		return false
	}

	if o.Int64 != nil && other.Int64 != nil {
		return o.Int64.CustomType.Equal(other.Int64.CustomType)
	}

	if o.List == nil && other.List != nil {
		return false
	}

	if o.List != nil && other.List == nil {
		return false
	}

	if o.List != nil && other.List != nil {
		if !o.List.CustomType.Equal(other.List.CustomType) {
			return false
		}

		return o.List.ElementType.Equal(other.List.ElementType)
	}

	if o.Map == nil && other.Map != nil {
		return false
	}

	if o.Map != nil && other.Map == nil {
		return false
	}

	if o.Map != nil && other.Map != nil {
		if !o.Map.CustomType.Equal(other.Map.CustomType) {
			return false
		}

		return o.Map.ElementType.Equal(other.Map.ElementType)
	}

	if o.Number == nil && other.Number != nil {
		return false
	}

	if o.Number != nil && other.Number == nil {
		return false
	}

	if o.Number != nil && other.Number != nil {
		return o.Number.CustomType.Equal(other.Number.CustomType)
	}

	if o.Object == nil && other.Object != nil {
		return false
	}

	if o.Object != nil && other.Object == nil {
		return false
	}

	if o.Object != nil && other.Object != nil {
		return o.Object.Equal(other.Object)
	}

	if o.Set == nil && other.Set != nil {
		return false
	}

	if o.Set != nil && other.Set == nil {
		return false
	}

	if o.Set != nil && other.Set != nil {
		if !o.Set.CustomType.Equal(other.Set.CustomType) {
			return false
		}

		return o.Set.ElementType.Equal(other.Set.ElementType)
	}

	if o.String == nil && other.String != nil {
		return false
	}

	if o.String != nil && other.String == nil {
		return false
	}

	if o.String != nil && other.String != nil {
		return o.String.CustomType.Equal(other.String.CustomType)
	}

	return true
}

type ObjectValidateRequest struct {
	Path string
}
