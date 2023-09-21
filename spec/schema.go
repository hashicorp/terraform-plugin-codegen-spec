// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	_ "embed"
)

const (
	Version1_0 = "1.0"
)

var (
	//go:embed v1.0/schema.json
	JSONSchemaVersion1_0 []byte
)
