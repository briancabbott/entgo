// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/dialect"
	"github.com/briancabbott/entgo/entc/integration/customid/sid"
	"github.com/briancabbott/entgo/schema/field"
)

// Other holds the schema definition for the Other entity.
type Other struct {
	ent.Schema
}

// Fields of the Other.
func (Other) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", sid.ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(sid.New).
			Immutable(),
	}
}
