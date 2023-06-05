// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec_test

func pointer[T any](in T) *T {
	return &in
}
