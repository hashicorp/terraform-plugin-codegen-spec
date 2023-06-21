// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"context"
	"errors"
	"fmt"
)

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

type ObjectValidateRequest struct {
	Path string
}

type ObjectAttributeTypes []ObjectAttributeType

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
