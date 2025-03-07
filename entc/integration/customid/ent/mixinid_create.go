// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/briancabbott/entgo/dialect"
	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/mixinid"
	"github.com/briancabbott/entgo/schema/field"
	"github.com/google/uuid"
)

// MixinIDCreate is the builder for creating a MixinID entity.
type MixinIDCreate struct {
	config
	mutation *MixinIDMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSomeField sets the "some_field" field.
func (mic *MixinIDCreate) SetSomeField(s string) *MixinIDCreate {
	mic.mutation.SetSomeField(s)
	return mic
}

// SetMixinField sets the "mixin_field" field.
func (mic *MixinIDCreate) SetMixinField(s string) *MixinIDCreate {
	mic.mutation.SetMixinField(s)
	return mic
}

// SetID sets the "id" field.
func (mic *MixinIDCreate) SetID(u uuid.UUID) *MixinIDCreate {
	mic.mutation.SetID(u)
	return mic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mic *MixinIDCreate) SetNillableID(u *uuid.UUID) *MixinIDCreate {
	if u != nil {
		mic.SetID(*u)
	}
	return mic
}

// Mutation returns the MixinIDMutation object of the builder.
func (mic *MixinIDCreate) Mutation() *MixinIDMutation {
	return mic.mutation
}

// Save creates the MixinID in the database.
func (mic *MixinIDCreate) Save(ctx context.Context) (*MixinID, error) {
	var (
		err  error
		node *MixinID
	)
	mic.defaults()
	if len(mic.hooks) == 0 {
		if err = mic.check(); err != nil {
			return nil, err
		}
		node, err = mic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixinIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mic.check(); err != nil {
				return nil, err
			}
			mic.mutation = mutation
			if node, err = mic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mic.hooks) - 1; i >= 0; i-- {
			if mic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MixinIDCreate) SaveX(ctx context.Context) *MixinID {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MixinIDCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MixinIDCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mic *MixinIDCreate) defaults() {
	if _, ok := mic.mutation.ID(); !ok {
		v := mixinid.DefaultID()
		mic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MixinIDCreate) check() error {
	if _, ok := mic.mutation.SomeField(); !ok {
		return &ValidationError{Name: "some_field", err: errors.New(`ent: missing required field "MixinID.some_field"`)}
	}
	if _, ok := mic.mutation.MixinField(); !ok {
		return &ValidationError{Name: "mixin_field", err: errors.New(`ent: missing required field "MixinID.mixin_field"`)}
	}
	return nil
}

func (mic *MixinIDCreate) sqlSave(ctx context.Context) (*MixinID, error) {
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (mic *MixinIDCreate) createSpec() (*MixinID, *sqlgraph.CreateSpec) {
	var (
		_node = &MixinID{config: mic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mixinid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		}
	)
	_spec.OnConflict = mic.conflict
	if id, ok := mic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mic.mutation.SomeField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldSomeField,
		})
		_node.SomeField = value
	}
	if value, ok := mic.mutation.MixinField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldMixinField,
		})
		_node.MixinField = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MixinID.Create().
//		SetSomeField(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MixinIDUpsert) {
//			SetSomeField(v+v).
//		}).
//		Exec(ctx)
//
func (mic *MixinIDCreate) OnConflict(opts ...sql.ConflictOption) *MixinIDUpsertOne {
	mic.conflict = opts
	return &MixinIDUpsertOne{
		create: mic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mic *MixinIDCreate) OnConflictColumns(columns ...string) *MixinIDUpsertOne {
	mic.conflict = append(mic.conflict, sql.ConflictColumns(columns...))
	return &MixinIDUpsertOne{
		create: mic,
	}
}

type (
	// MixinIDUpsertOne is the builder for "upsert"-ing
	//  one MixinID node.
	MixinIDUpsertOne struct {
		create *MixinIDCreate
	}

	// MixinIDUpsert is the "OnConflict" setter.
	MixinIDUpsert struct {
		*sql.UpdateSet
	}
)

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsert) SetSomeField(v string) *MixinIDUpsert {
	u.Set(mixinid.FieldSomeField, v)
	return u
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsert) UpdateSomeField() *MixinIDUpsert {
	u.SetExcluded(mixinid.FieldSomeField)
	return u
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsert) SetMixinField(v string) *MixinIDUpsert {
	u.Set(mixinid.FieldMixinField, v)
	return u
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsert) UpdateMixinField() *MixinIDUpsert {
	u.SetExcluded(mixinid.FieldMixinField)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mixinid.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MixinIDUpsertOne) UpdateNewValues() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mixinid.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.MixinID.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MixinIDUpsertOne) Ignore() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MixinIDUpsertOne) DoNothing() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MixinIDCreate.OnConflict
// documentation for more info.
func (u *MixinIDUpsertOne) Update(set func(*MixinIDUpsert)) *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MixinIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsertOne) SetSomeField(v string) *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetSomeField(v)
	})
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsertOne) UpdateSomeField() *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateSomeField()
	})
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsertOne) SetMixinField(v string) *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetMixinField(v)
	})
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsertOne) UpdateMixinField() *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateMixinField()
	})
}

// Exec executes the query.
func (u *MixinIDUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MixinIDCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MixinIDUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MixinIDUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MixinIDUpsertOne.ID is not supported by MySQL driver. Use MixinIDUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MixinIDUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MixinIDCreateBulk is the builder for creating many MixinID entities in bulk.
type MixinIDCreateBulk struct {
	config
	builders []*MixinIDCreate
	conflict []sql.ConflictOption
}

// Save creates the MixinID entities in the database.
func (micb *MixinIDCreateBulk) Save(ctx context.Context) ([]*MixinID, error) {
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MixinID, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MixinIDMutation)
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
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = micb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MixinIDCreateBulk) SaveX(ctx context.Context) []*MixinID {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MixinIDCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MixinIDCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MixinID.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MixinIDUpsert) {
//			SetSomeField(v+v).
//		}).
//		Exec(ctx)
//
func (micb *MixinIDCreateBulk) OnConflict(opts ...sql.ConflictOption) *MixinIDUpsertBulk {
	micb.conflict = opts
	return &MixinIDUpsertBulk{
		create: micb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (micb *MixinIDCreateBulk) OnConflictColumns(columns ...string) *MixinIDUpsertBulk {
	micb.conflict = append(micb.conflict, sql.ConflictColumns(columns...))
	return &MixinIDUpsertBulk{
		create: micb,
	}
}

// MixinIDUpsertBulk is the builder for "upsert"-ing
// a bulk of MixinID nodes.
type MixinIDUpsertBulk struct {
	create *MixinIDCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mixinid.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MixinIDUpsertBulk) UpdateNewValues() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mixinid.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MixinIDUpsertBulk) Ignore() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MixinIDUpsertBulk) DoNothing() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MixinIDCreateBulk.OnConflict
// documentation for more info.
func (u *MixinIDUpsertBulk) Update(set func(*MixinIDUpsert)) *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MixinIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsertBulk) SetSomeField(v string) *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetSomeField(v)
	})
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsertBulk) UpdateSomeField() *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateSomeField()
	})
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsertBulk) SetMixinField(v string) *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetMixinField(v)
	})
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsertBulk) UpdateMixinField() *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateMixinField()
	})
}

// Exec executes the query.
func (u *MixinIDUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MixinIDCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MixinIDCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MixinIDUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
