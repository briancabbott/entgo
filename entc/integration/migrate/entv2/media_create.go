// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/migrate/entv2/media"
	"github.com/briancabbott/entgo/schema/field"
)

// MediaCreate is the builder for creating a Media entity.
type MediaCreate struct {
	config
	mutation *MediaMutation
	hooks    []Hook
}

// SetSource sets the "source" field.
func (mc *MediaCreate) SetSource(s string) *MediaCreate {
	mc.mutation.SetSource(s)
	return mc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (mc *MediaCreate) SetNillableSource(s *string) *MediaCreate {
	if s != nil {
		mc.SetSource(*s)
	}
	return mc
}

// SetSourceURI sets the "source_uri" field.
func (mc *MediaCreate) SetSourceURI(s string) *MediaCreate {
	mc.mutation.SetSourceURI(s)
	return mc
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (mc *MediaCreate) SetNillableSourceURI(s *string) *MediaCreate {
	if s != nil {
		mc.SetSourceURI(*s)
	}
	return mc
}

// SetText sets the "text" field.
func (mc *MediaCreate) SetText(s string) *MediaCreate {
	mc.mutation.SetText(s)
	return mc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (mc *MediaCreate) SetNillableText(s *string) *MediaCreate {
	if s != nil {
		mc.SetText(*s)
	}
	return mc
}

// Mutation returns the MediaMutation object of the builder.
func (mc *MediaCreate) Mutation() *MediaMutation {
	return mc.mutation
}

// Save creates the Media in the database.
func (mc *MediaCreate) Save(ctx context.Context) (*Media, error) {
	var (
		err  error
		node *Media
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("entv2: uninitialized hook (forgotten import entv2/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MediaCreate) SaveX(ctx context.Context) *Media {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MediaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MediaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MediaCreate) check() error {
	return nil
}

func (mc *MediaCreate) sqlSave(ctx context.Context) (*Media, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MediaCreate) createSpec() (*Media, *sqlgraph.CreateSpec) {
	var (
		_node = &Media{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: media.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: media.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := mc.mutation.SourceURI(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSourceURI,
		})
		_node.SourceURI = value
	}
	if value, ok := mc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldText,
		})
		_node.Text = value
	}
	return _node, _spec
}

// MediaCreateBulk is the builder for creating many Media entities in bulk.
type MediaCreateBulk struct {
	config
	builders []*MediaCreate
}

// Save creates the Media entities in the database.
func (mcb *MediaCreateBulk) Save(ctx context.Context) ([]*Media, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Media, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MediaCreateBulk) SaveX(ctx context.Context) []*Media {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MediaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MediaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
