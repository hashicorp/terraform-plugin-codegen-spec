// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MPL-2.0

package spec_test

func pointer[T any](in T) *T {
	return &in
}
