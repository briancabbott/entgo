// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package task

import (
	"fmt"

	"github.com/briancabbott/entgo"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// TeamsTable is the table that holds the teams relation/edge. The primary key declared below.
	TeamsTable = "task_teams"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "tasks"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldStatus,
	FieldUUID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tasks",
}

var (
	// TeamsPrimaryKey and TeamsColumn2 are the table columns denoting the
	// primary key for the teams relation (M2M).
	TeamsPrimaryKey = []string{"task_id", "team_id"}
)

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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/briancabbott/entgo/entc/integration/privacy/ent/runtime"
//
var (
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPlanned is the default value of the Status enum.
const DefaultStatus = StatusPlanned

// Status values.
const (
	StatusPlanned    Status = "planned"
	StatusInProgress Status = "in_progress"
	StatusClosed     Status = "closed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPlanned, StatusInProgress, StatusClosed:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for status field: %q", s)
	}
}
