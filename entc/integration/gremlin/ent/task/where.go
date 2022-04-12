// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package enttask

import (
	"github.com/briancabbott/entgoriancabbott/entgo/dialect/gremlin/graph/dsl"
	"github.com/briancabbott/entgoriancabbott/entgo/dialect/gremlin/graph/dsl/__"
	"github.com/briancabbott/entgoriancabbott/entgo/dialect/gremlin/graph/dsl/p"
	"github.com/briancabbott/entgoriancabbott/entgo/entc/integration/ent/schema/task"
	"github.com/briancabbott/entgoriancabbott/entgo/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.EQ(vc))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.EQ(vc))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.NEQ(vc))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...task.Priority) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.Within(v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...task.Priority) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.Without(v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.GT(vc))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.GTE(vc))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.LT(vc))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.LTE(vc))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
