// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// Parse returns a Specification from the JSON document contents, or any validation errors.
func Parse(ctx context.Context, document []byte) (Specification, error) {
	if err := Validate(ctx, document); err != nil {
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

// Validate loads the schema version specified in the document, and validates the document.
func Validate(ctx context.Context, document []byte) error {
	if len(document) == 0 {
		return errors.New("empty document")
	}

	documentLoader := gojsonschema.NewBytesLoader(document)

	// We only need to grab the version here, the JSON schema can do the remaining of the validation
	var versionedDocument struct {
		Version *string `json:"version"`
	}

	if err := json.Unmarshal(document, &versionedDocument); err != nil {
		return err
	}

	if versionedDocument.Version == nil || *versionedDocument.Version == "" {
		return errors.New("version is required")
	}

	var schemaVersion []byte

	switch *versionedDocument.Version {
	case Version0_1:
		schemaVersion = JSONSchemaVersion0_1
	default:
		return fmt.Errorf("version: %q is unsupported", *versionedDocument.Version)
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
