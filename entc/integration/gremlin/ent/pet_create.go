// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/briancabbott/entgo/dialect/gremlin"
	"github.com/briancabbott/entgo/dialect/gremlin/graph/dsl"
	"github.com/briancabbott/entgo/dialect/gremlin/graph/dsl/__"
	"github.com/briancabbott/entgo/dialect/gremlin/graph/dsl/g"
	"github.com/briancabbott/entgo/dialect/gremlin/graph/dsl/p"
	"github.com/briancabbott/entgo/entc/integration/gremlin/ent/pet"
	"github.com/briancabbott/entgo/entc/integration/gremlin/ent/user"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (pc *PetCreate) SetAge(f float64) *PetCreate {
	pc.mutation.SetAge(f)
	return pc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (pc *PetCreate) SetNillableAge(f *float64) *PetCreate {
	if f != nil {
		pc.SetAge(*f)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetUUID sets the "uuid" field.
func (pc *PetCreate) SetUUID(u uuid.UUID) *PetCreate {
	pc.mutation.SetUUID(u)
	return pc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pc *PetCreate) SetNillableUUID(u *uuid.UUID) *PetCreate {
	if u != nil {
		pc.SetUUID(*u)
	}
	return pc
}

// SetNickname sets the "nickname" field.
func (pc *PetCreate) SetNickname(s string) *PetCreate {
	pc.mutation.SetNickname(s)
	return pc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (pc *PetCreate) SetNillableNickname(s *string) *PetCreate {
	if s != nil {
		pc.SetNickname(*s)
	}
	return pc
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (pc *PetCreate) SetTeamID(id string) *PetCreate {
	pc.mutation.SetTeamID(id)
	return pc
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (pc *PetCreate) SetNillableTeamID(id *string) *PetCreate {
	if id != nil {
		pc = pc.SetTeamID(*id)
	}
	return pc
}

// SetTeam sets the "team" edge to the User entity.
func (pc *PetCreate) SetTeam(u *User) *PetCreate {
	return pc.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PetCreate) SetOwnerID(id string) *PetCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *string) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.gremlinSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PetCreate) defaults() {
	if _, ok := pc.mutation.Age(); !ok {
		v := pet.DefaultAge
		pc.mutation.SetAge(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	return nil
}

func (pc *PetCreate) gremlinSave(ctx context.Context) (*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := pc.gremlin().Query()
	if err := pc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	pe := &Pet{config: pc.config}
	if err := pe.FromResponse(res); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pc *PetCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(pet.Label)
	if value, ok := pc.mutation.Age(); ok {
		v.Property(dsl.Single, pet.FieldAge, value)
	}
	if value, ok := pc.mutation.Name(); ok {
		v.Property(dsl.Single, pet.FieldName, value)
	}
	if value, ok := pc.mutation.UUID(); ok {
		v.Property(dsl.Single, pet.FieldUUID, value)
	}
	if value, ok := pc.mutation.Nickname(); ok {
		v.Property(dsl.Single, pet.FieldNickname, value)
	}
	for _, id := range pc.mutation.TeamIDs() {
		v.AddE(user.TeamLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(pet.Label, user.TeamLabel, id)),
		})
	}
	for _, id := range pc.mutation.OwnerIDs() {
		v.AddE(user.PetsLabel).From(g.V(id)).InV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	builders []*PetCreate
}
