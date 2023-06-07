package spec

import (
	"context"
	"encoding/json"
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

// Parse returns a Specification from the JSON document contents or any validation errors.
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
