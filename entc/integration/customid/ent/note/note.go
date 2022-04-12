// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package note

import (
	"github.com/briancabbott/entgoriancabbott/entgo/entc/integration/customid/ent/schema"
)

const (
	// Label holds the string label denoting the note type in the database.
	Label = "note"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the note in the database.
	Table = "notes"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "notes"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "note_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "notes"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "note_children"
)

// Columns holds all SQL columns for note fields.
var Columns = []string{
	FieldID,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"note_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() schema.NoteID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
