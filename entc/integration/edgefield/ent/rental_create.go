// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/car"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/rental"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/user"
	"github.com/briancabbott/entgo/schema/field"
	"github.com/google/uuid"
)

// RentalCreate is the builder for creating a Rental entity.
type RentalCreate struct {
	config
	mutation *RentalMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (rc *RentalCreate) SetDate(t time.Time) *RentalCreate {
	rc.mutation.SetDate(t)
	return rc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (rc *RentalCreate) SetNillableDate(t *time.Time) *RentalCreate {
	if t != nil {
		rc.SetDate(*t)
	}
	return rc
}

// SetUserID sets the "user_id" field.
func (rc *RentalCreate) SetUserID(i int) *RentalCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetCarID sets the "car_id" field.
func (rc *RentalCreate) SetCarID(u uuid.UUID) *RentalCreate {
	rc.mutation.SetCarID(u)
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RentalCreate) SetUser(u *User) *RentalCreate {
	return rc.SetUserID(u.ID)
}

// SetCar sets the "car" edge to the Car entity.
func (rc *RentalCreate) SetCar(c *Car) *RentalCreate {
	return rc.SetCarID(c.ID)
}

// Mutation returns the RentalMutation object of the builder.
func (rc *RentalCreate) Mutation() *RentalMutation {
	return rc.mutation
}

// Save creates the Rental in the database.
func (rc *RentalCreate) Save(ctx context.Context) (*Rental, error) {
	var (
		err  error
		node *Rental
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RentalCreate) SaveX(ctx context.Context) *Rental {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RentalCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RentalCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RentalCreate) defaults() {
	if _, ok := rc.mutation.Date(); !ok {
		v := rental.DefaultDate()
		rc.mutation.SetDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RentalCreate) check() error {
	if _, ok := rc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Rental.date"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Rental.user_id"`)}
	}
	if _, ok := rc.mutation.CarID(); !ok {
		return &ValidationError{Name: "car_id", err: errors.New(`ent: missing required field "Rental.car_id"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Rental.user"`)}
	}
	if _, ok := rc.mutation.CarID(); !ok {
		return &ValidationError{Name: "car", err: errors.New(`ent: missing required edge "Rental.car"`)}
	}
	return nil
}

func (rc *RentalCreate) sqlSave(ctx context.Context) (*Rental, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RentalCreate) createSpec() (*Rental, *sqlgraph.CreateSpec) {
	var (
		_node = &Rental{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rental.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rental.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rental.FieldDate,
		})
		_node.Date = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CarID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RentalCreateBulk is the builder for creating many Rental entities in bulk.
type RentalCreateBulk struct {
	config
	builders []*RentalCreate
}

// Save creates the Rental entities in the database.
func (rcb *RentalCreateBulk) Save(ctx context.Context) ([]*Rental, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rental, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RentalMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RentalCreateBulk) SaveX(ctx context.Context) []*Rental {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RentalCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RentalCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
