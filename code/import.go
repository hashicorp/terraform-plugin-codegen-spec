// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package code

// Import represents a required source code import to ensure that generated code
// compiles or runs successfully using code from sources external to the
// generated code file. The syntax and semantics of an import is programming
// language specific.
type Import struct {
	// Alias is an optional string containing an alias or specialized behavior
	// for the import, based on the programming language.
	//
	// For example in Go, this may be a different package name to prevent name
	// collisions within the same source code file or an underscore to import
	// the package solely for its initialization side-effects.
	Alias *string `json:"alias,omitempty"`

	// Path contains the source of the code for importing, typically a fully
	// qualified path or name, based on the programming language.
	//
	// For example in Go, this is a package path, such as
	// github.com/hashicorp/terraform-plugin-framework/types.
	Path string `json:"path"`
}
