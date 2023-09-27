// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	_ "embed"
)

const (
	Version0_1 = "0.1"
)

var (
	//go:embed v0.1/schema.json
	JSONSchemaVersion0_1 []byte
)
