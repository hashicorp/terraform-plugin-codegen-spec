// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	_ "embed"
	"fmt"
)

const (
	version_1_0 = "v1.0"
)

var (
	//go:embed v1.0/schema.json
	v1_0 []byte
)

func Schema(version string) ([]byte, error) {
	switch version {
	case version_1_0:
		return v1_0, nil
	}

	return nil, fmt.Errorf("schema version: %q is invalid", version)
}
