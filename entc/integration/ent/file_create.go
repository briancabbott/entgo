// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/ent/fieldtype"
	"github.com/briancabbott/entgo/entc/integration/ent/file"
	"github.com/briancabbott/entgo/entc/integration/ent/filetype"
	"github.com/briancabbott/entgo/entc/integration/ent/user"
	"github.com/briancabbott/entgo/schema/field"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetUser sets the "user" field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.mutation.SetUser(s)
	return fc
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the "group" field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.mutation.SetGroup(s)
	return fc
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOp sets the "op" field.
func (fc *FileCreate) SetOp(b bool) *FileCreate {
	fc.mutation.SetOp(b)
	return fc
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fc *FileCreate) SetNillableOp(b *bool) *FileCreate {
	if b != nil {
		fc.SetOp(*b)
	}
	return fc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fc *FileCreate) SetOwnerID(id int) *FileCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the "owner" edge to the User entity.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fc *FileCreate) SetTypeID(id int) *FileCreate {
	fc.mutation.SetTypeID(id)
	return fc
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the "type" edge to the FileType entity.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fc *FileCreate) AddFieldIDs(ids ...int) *FileCreate {
	fc.mutation.AddFieldIDs(ids...)
	return fc
}

// AddField adds the "field" edges to the FieldType entity.
func (fc *FileCreate) AddField(f ...*FieldType) *FileCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.Size(); !ok {
		v := file.DefaultSize
		fc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "File.size"`)}
	}
	if v, ok := fc.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fc.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldUser,
		})
		_node.User = &value
	}
	if value, ok := fc.mutation.Group(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldGroup,
		})
		_node.Group = value
	}
	if value, ok := fc.mutation.GetOp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: file.FieldOp,
		})
		_node.Op = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
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
		_node.user_files = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.file_type_files = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FieldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fieldtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.Create().
//		SetSize(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetSize(v+v).
//		}).
//		Exec(ctx)
//
func (fc *FileCreate) OnConflict(opts ...sql.ConflictOption) *FileUpsertOne {
	fc.conflict = opts
	return &FileUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fc *FileCreate) OnConflictColumns(columns ...string) *FileUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertOne{
		create: fc,
	}
}

type (
	// FileUpsertOne is the builder for "upsert"-ing
	//  one File node.
	FileUpsertOne struct {
		create *FileCreate
	}

	// FileUpsert is the "OnConflict" setter.
	FileUpsert struct {
		*sql.UpdateSet
	}
)

// SetSize sets the "size" field.
func (u *FileUpsert) SetSize(v int) *FileUpsert {
	u.Set(file.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsert) UpdateSize() *FileUpsert {
	u.SetExcluded(file.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *FileUpsert) AddSize(v int) *FileUpsert {
	u.Add(file.FieldSize, v)
	return u
}

// SetName sets the "name" field.
func (u *FileUpsert) SetName(v string) *FileUpsert {
	u.Set(file.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsert) UpdateName() *FileUpsert {
	u.SetExcluded(file.FieldName)
	return u
}

// SetUser sets the "user" field.
func (u *FileUpsert) SetUser(v string) *FileUpsert {
	u.Set(file.FieldUser, v)
	return u
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *FileUpsert) UpdateUser() *FileUpsert {
	u.SetExcluded(file.FieldUser)
	return u
}

// ClearUser clears the value of the "user" field.
func (u *FileUpsert) ClearUser() *FileUpsert {
	u.SetNull(file.FieldUser)
	return u
}

// SetGroup sets the "group" field.
func (u *FileUpsert) SetGroup(v string) *FileUpsert {
	u.Set(file.FieldGroup, v)
	return u
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *FileUpsert) UpdateGroup() *FileUpsert {
	u.SetExcluded(file.FieldGroup)
	return u
}

// ClearGroup clears the value of the "group" field.
func (u *FileUpsert) ClearGroup() *FileUpsert {
	u.SetNull(file.FieldGroup)
	return u
}

// SetOp sets the "op" field.
func (u *FileUpsert) SetOp(v bool) *FileUpsert {
	u.Set(file.FieldOp, v)
	return u
}

// UpdateOp sets the "op" field to the value that was provided on create.
func (u *FileUpsert) UpdateOp() *FileUpsert {
	u.SetExcluded(file.FieldOp)
	return u
}

// ClearOp clears the value of the "op" field.
func (u *FileUpsert) ClearOp() *FileUpsert {
	u.SetNull(file.FieldOp)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *FileUpsertOne) UpdateNewValues() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.File.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FileUpsertOne) Ignore() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertOne) DoNothing() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreate.OnConflict
// documentation for more info.
func (u *FileUpsertOne) Update(set func(*FileUpsert)) *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetSize sets the "size" field.
func (u *FileUpsertOne) SetSize(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertOne) AddSize(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSize() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetName sets the "name" field.
func (u *FileUpsertOne) SetName(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateName() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetUser sets the "user" field.
func (u *FileUpsertOne) SetUser(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUser(v)
	})
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUser() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUser()
	})
}

// ClearUser clears the value of the "user" field.
func (u *FileUpsertOne) ClearUser() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearUser()
	})
}

// SetGroup sets the "group" field.
func (u *FileUpsertOne) SetGroup(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateGroup() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateGroup()
	})
}

// ClearGroup clears the value of the "group" field.
func (u *FileUpsertOne) ClearGroup() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearGroup()
	})
}

// SetOp sets the "op" field.
func (u *FileUpsertOne) SetOp(v bool) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetOp(v)
	})
}

// UpdateOp sets the "op" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateOp() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateOp()
	})
}

// ClearOp clears the value of the "op" field.
func (u *FileUpsertOne) ClearOp() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearOp()
	})
}

// Exec executes the query.
func (u *FileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
	conflict []sql.ConflictOption
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetSize(v+v).
//		}).
//		Exec(ctx)
//
func (fcb *FileCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileUpsertBulk {
	fcb.conflict = opts
	return &FileUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fcb *FileCreateBulk) OnConflictColumns(columns ...string) *FileUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertBulk{
		create: fcb,
	}
}

// FileUpsertBulk is the builder for "upsert"-ing
// a bulk of File nodes.
type FileUpsertBulk struct {
	create *FileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *FileUpsertBulk) UpdateNewValues() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FileUpsertBulk) Ignore() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertBulk) DoNothing() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreateBulk.OnConflict
// documentation for more info.
func (u *FileUpsertBulk) Update(set func(*FileUpsert)) *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetSize sets the "size" field.
func (u *FileUpsertBulk) SetSize(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertBulk) AddSize(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSize() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetName sets the "name" field.
func (u *FileUpsertBulk) SetName(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateName() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetUser sets the "user" field.
func (u *FileUpsertBulk) SetUser(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUser(v)
	})
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUser() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUser()
	})
}

// ClearUser clears the value of the "user" field.
func (u *FileUpsertBulk) ClearUser() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearUser()
	})
}

// SetGroup sets the "group" field.
func (u *FileUpsertBulk) SetGroup(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateGroup() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateGroup()
	})
}

// ClearGroup clears the value of the "group" field.
func (u *FileUpsertBulk) ClearGroup() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearGroup()
	})
}

// SetOp sets the "op" field.
func (u *FileUpsertBulk) SetOp(v bool) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetOp(v)
	})
}

// UpdateOp sets the "op" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateOp() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateOp()
	})
}

// ClearOp clears the value of the "op" field.
func (u *FileUpsertBulk) ClearOp() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearOp()
	})
}

// Exec executes the query.
func (u *FileUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
