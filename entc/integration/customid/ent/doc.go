// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/doc"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/schema"
)

// Doc is the model entity for the Doc schema.
type Doc struct {
	config `json:"-"`
	// ID of the ent.
	ID schema.DocID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DocQuery when eager-loading is set.
	Edges        DocEdges `json:"edges"`
	doc_children *schema.DocID
}

// DocEdges holds the relations/edges for other nodes in the graph.
type DocEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Doc `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Doc `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocEdges) ParentOrErr() (*Doc, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doc.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DocEdges) ChildrenOrErr() ([]*Doc, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doc) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case doc.FieldID:
			values[i] = new(schema.DocID)
		case doc.FieldText:
			values[i] = new(sql.NullString)
		case doc.ForeignKeys[0]: // doc_children
			values[i] = &sql.NullScanner{S: new(schema.DocID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Doc", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doc fields.
func (d *Doc) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case doc.FieldID:
			if value, ok := values[i].(*schema.DocID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case doc.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				d.Text = value.String
			}
		case doc.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field doc_children", values[i])
			} else if value.Valid {
				d.doc_children = new(schema.DocID)
				*d.doc_children = *value.S.(*schema.DocID)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Doc entity.
func (d *Doc) QueryParent() *DocQuery {
	return (&DocClient{config: d.config}).QueryParent(d)
}

// QueryChildren queries the "children" edge of the Doc entity.
func (d *Doc) QueryChildren() *DocQuery {
	return (&DocClient{config: d.config}).QueryChildren(d)
}

// Update returns a builder for updating this Doc.
// Note that you need to call Doc.Unwrap() before calling this method if this Doc
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Doc) Update() *DocUpdateOne {
	return (&DocClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Doc entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Doc) Unwrap() *Doc {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doc is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Doc) String() string {
	var builder strings.Builder
	builder.WriteString("Doc(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", text=")
	builder.WriteString(d.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Docs is a parsable slice of Doc.
type Docs []*Doc

func (d Docs) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
