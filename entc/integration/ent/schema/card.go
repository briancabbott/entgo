// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/entc/integration/ent/template"
	"github.com/briancabbott/entgo/schema"
	"github.com/briancabbott/entgo/schema/edge"
	"github.com/briancabbott/entgo/schema/field"
	"github.com/briancabbott/entgo/schema/index"
	"github.com/briancabbott/entgo/schema/mixin"
)

type CardMixin struct {
	mixin.Schema
}

func (CardMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `mashraki:"edges"`,
		},
		field.Annotation{
			StructTag: map[string]string{
				"id":     `yaml:"-"`,
				"number": `json:"-"`,
			},
		},
	}
}

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		CardMixin{},
	}
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.Annotation{
			StructTag: map[string]string{
				"id": `json:"-"`,
			},
		},
	}
}

// Fields of the Comment.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Float("balance").
			Default(0),
		field.String("number").
			Immutable().
			NotEmpty().
			Annotations(&template.Extension{
				Type: "string",
			}),
		field.String("name").
			Optional().
			Comment("Exact name written on card").
			NotEmpty().
			Annotations(&template.Extension{
				Type: "string",
			}),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Comment("O2O inverse edge").
			Ref("card").
			Unique(),
		edge.From("spec", Spec.Type).
			Ref("card").
			Annotations(&template.Extension{
				Type: "int",
			}),
	}
}

// Indexes of the Card.
func (Card) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("number").
			Unique(),
		index.Fields("id", "name", "number"),
	}
}
