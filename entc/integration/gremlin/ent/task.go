// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/briancabbott/entgo/dialect/gremlin"
	"github.com/briancabbott/entgo/entc/integration/ent/schema/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority task.Priority `json:"priority,omitempty"`
}

// FromResponse scans the gremlin response data into Task.
func (t *Task) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scant struct {
		ID       string        `json:"id,omitempty"`
		Priority task.Priority `json:"priority,omitempty"`
	}
	if err := vmap.Decode(&scant); err != nil {
		return err
	}
	t.ID = scant.ID
	t.Priority = scant.Priority
	return nil
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

// FromResponse scans the gremlin response data into Tasks.
func (t *Tasks) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scant []struct {
		ID       string        `json:"id,omitempty"`
		Priority task.Priority `json:"priority,omitempty"`
	}
	if err := vmap.Decode(&scant); err != nil {
		return err
	}
	for _, v := range scant {
		*t = append(*t, &Task{
			ID:       v.ID,
			Priority: v.Priority,
		})
	}
	return nil
}

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
