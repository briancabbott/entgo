// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/entc/integration/ent/spec"
)

// Spec is the model entity for the Spec schema.
type Spec struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecQuery when eager-loading is set.
	Edges SpecEdges `json:"edges"`
}

// SpecEdges holds the relations/edges for other nodes in the graph.
type SpecEdges struct {
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e SpecEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Spec) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case spec.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Spec", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Spec fields.
func (s *Spec) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spec.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		}
	}
	return nil
}

// QueryCard queries the "card" edge of the Spec entity.
func (s *Spec) QueryCard() *CardQuery {
	return (&SpecClient{config: s.config}).QueryCard(s)
}

// Update returns a builder for updating this Spec.
// Note that you need to call Spec.Unwrap() before calling this method if this Spec
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Spec) Update() *SpecUpdateOne {
	return (&SpecClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Spec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Spec) Unwrap() *Spec {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Spec is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Spec) String() string {
	var builder strings.Builder
	builder.WriteString("Spec(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Specs is a parsable slice of Spec.
type Specs []*Spec

func (s Specs) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
