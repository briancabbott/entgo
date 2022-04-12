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
	"github.com/briancabbott/entgo/entc/integration/customid/ent/account"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/token"
	"github.com/briancabbott/entgo/entc/integration/customid/sid"
	"github.com/briancabbott/entgo/schema/field"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBody sets the "body" field.
func (tc *TokenCreate) SetBody(s string) *TokenCreate {
	tc.mutation.SetBody(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TokenCreate) SetID(s sid.ID) *TokenCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TokenCreate) SetNillableID(s *sid.ID) *TokenCreate {
	if s != nil {
		tc.SetID(*s)
	}
	return tc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (tc *TokenCreate) SetAccountID(id sid.ID) *TokenCreate {
	tc.mutation.SetAccountID(id)
	return tc
}

// SetAccount sets the "account" edge to the Account entity.
func (tc *TokenCreate) SetAccount(a *Account) *TokenCreate {
	return tc.SetAccountID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	var (
		err  error
		node *Token
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := token.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Token.body"`)}
	}
	if v, ok := tc.mutation.Body(); ok {
		if err := token.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Token.body": %w`, err)}
		}
	}
	if _, ok := tc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Token.account"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*sid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: token.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: token.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldBody,
		})
		_node.Body = value
	}
	if nodes := tc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeOther,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_token = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.Create().
//		SetBody(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetBody(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TokenCreate) OnConflict(opts ...sql.ConflictOption) *TokenUpsertOne {
	tc.conflict = opts
	return &TokenUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TokenCreate) OnConflictColumns(columns ...string) *TokenUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertOne{
		create: tc,
	}
}

type (
	// TokenUpsertOne is the builder for "upsert"-ing
	//  one Token node.
	TokenUpsertOne struct {
		create *TokenCreate
	}

	// TokenUpsert is the "OnConflict" setter.
	TokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetBody sets the "body" field.
func (u *TokenUpsert) SetBody(v string) *TokenUpsert {
	u.Set(token.FieldBody, v)
	return u
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *TokenUpsert) UpdateBody() *TokenUpsert {
	u.SetExcluded(token.FieldBody)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TokenUpsertOne) UpdateNewValues() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(token.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Token.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TokenUpsertOne) Ignore() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertOne) DoNothing() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreate.OnConflict
// documentation for more info.
func (u *TokenUpsertOne) Update(set func(*TokenUpsert)) *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetBody sets the "body" field.
func (u *TokenUpsertOne) SetBody(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateBody() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateBody()
	})
}

// Exec executes the query.
func (u *TokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenUpsertOne) ID(ctx context.Context) (id sid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TokenUpsertOne.ID is not supported by MySQL driver. Use TokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenUpsertOne) IDX(ctx context.Context) sid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	builders []*TokenCreate
	conflict []sql.ConflictOption
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetBody(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenUpsertBulk {
	tcb.conflict = opts
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TokenCreateBulk) OnConflictColumns(columns ...string) *TokenUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// TokenUpsertBulk is the builder for "upsert"-ing
// a bulk of Token nodes.
type TokenUpsertBulk struct {
	create *TokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TokenUpsertBulk) UpdateNewValues() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(token.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TokenUpsertBulk) Ignore() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertBulk) DoNothing() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreateBulk.OnConflict
// documentation for more info.
func (u *TokenUpsertBulk) Update(set func(*TokenUpsert)) *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetBody sets the "body" field.
func (u *TokenUpsertBulk) SetBody(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateBody() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateBody()
	})
}

// Exec executes the query.
func (u *TokenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
