// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/ent/filetype"
	"github.com/briancabbott/entgo/entc/integration/ent/predicate"
	"github.com/briancabbott/entgo/schema/field"
)

// FileTypeDelete is the builder for deleting a FileType entity.
type FileTypeDelete struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// Where appends a list predicates to the FileTypeDelete builder.
func (ftd *FileTypeDelete) Where(ps ...predicate.FileType) *FileTypeDelete {
	ftd.mutation.Where(ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FileTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftd.hooks) == 0 {
		affected, err = ftd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftd.mutation = mutation
			affected, err = ftd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftd.hooks) - 1; i >= 0; i-- {
			if ftd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: filetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filetype.FieldID,
			},
		},
	}
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
}

// FileTypeDeleteOne is the builder for deleting a single FileType entity.
type FileTypeDeleteOne struct {
	ftd *FileTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FileTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{filetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
