// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/dialect"
	"github.com/briancabbott/entgo/entc/integration/customid/sid"
	"github.com/briancabbott/entgo/schema/edge"
	"github.com/briancabbott/entgo/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", sid.ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(sid.New).
			Immutable(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("token", Token.Type),
	}
}
