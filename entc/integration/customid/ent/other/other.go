// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package other

import (
	"github.com/briancabbott/entgo/entc/integration/customid/sid"
)

const (
	// Label holds the string label denoting the other type in the database.
	Label = "other"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the other in the database.
	Table = "others"
)

// Columns holds all SQL columns for other fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() sid.ID
)
