// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package predicate

import (
	"github.com/briancabbott/entgo/dialect/sql"
)

// Card is the predicate function for card builders.
type Card func(*sql.Selector)

// Comment is the predicate function for comment builders.
type Comment func(*sql.Selector)

// FieldType is the predicate function for fieldtype builders.
type FieldType func(*sql.Selector)

// File is the predicate function for file builders.
type File func(*sql.Selector)

// FileType is the predicate function for filetype builders.
type FileType func(*sql.Selector)

// Goods is the predicate function for goods builders.
type Goods func(*sql.Selector)

// Group is the predicate function for group builders.
type Group func(*sql.Selector)

// GroupInfo is the predicate function for groupinfo builders.
type GroupInfo func(*sql.Selector)

// Item is the predicate function for item builders.
type Item func(*sql.Selector)

// Node is the predicate function for node builders.
type Node func(*sql.Selector)

// Pet is the predicate function for pet builders.
type Pet func(*sql.Selector)

// Spec is the predicate function for spec builders.
type Spec func(*sql.Selector)

// Task is the predicate function for enttask builders.
type Task func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
