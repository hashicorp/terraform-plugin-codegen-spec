// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

// Parse returns a Specification from the JSON document contents or any validation errors.
func Parse(ctx context.Context, document []byte, schemaVersion string) (Specification, error) {
	if err := Validate(ctx, document, schemaVersion); err != nil {
		return Specification{}, err
	}

	var spec Specification

	if err := json.Unmarshal(document, &spec); err != nil {
		return spec, err
	}

	if err := spec.Validate(ctx); err != nil {
		return spec, err
	}

	return spec, nil
}

// Validate loads the schema version specified and validates the document.
func Validate(ctx context.Context, document []byte, version string) error {
	if len(document) == 0 {
		return errors.New("empty document")
	}

	documentLoader := gojsonschema.NewBytesLoader(document)

	schemaVersion, err := Schema(version)

	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewBytesLoader(schemaVersion)

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
