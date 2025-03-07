// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/dialect/entsql"
	"github.com/briancabbott/entgo/schema/edge"
	"github.com/briancabbott/entgo/schema/field"

	"github.com/google/uuid"
)

// Blob holds the schema definition for the Blob entity.
type Blob struct {
	ent.Schema
}

// Fields of the Blob.
func (Blob) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(entsql.Annotation{
				Default: "uuid_generate_v4()",
			}).
			Unique(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int("count").
			Default(0),
	}
}

// Edges of the Blob.
func (Blob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Blob.Type).
			Unique(),
		edge.To("links", Blob.Type),
	}
}
