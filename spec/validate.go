// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

func Validate(ctx context.Context, document []byte) error {
	if len(document) == 0 {
		return errors.New("empty document")
	}

	documentLoader := gojsonschema.NewBytesLoader(document)
	schemaLoader := gojsonschema.NewBytesLoader(schema)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return err
	}

	var errs error

	if !result.Valid() {
		for _, resultError := range result.Errors() {
			errs = errors.Join(errors.New(resultError.String()))
		}
	}

	return errs
}
