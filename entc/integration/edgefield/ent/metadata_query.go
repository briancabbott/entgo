// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/metadata"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/predicate"
	"github.com/briancabbott/entgo/entc/integration/edgefield/ent/user"
	"github.com/briancabbott/entgo/schema/field"
)

// MetadataQuery is the builder for querying Metadata entities.
type MetadataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Metadata
	// eager-loading edges.
	withUser     *UserQuery
	withChildren *MetadataQuery
	withParent   *MetadataQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetadataQuery builder.
func (mq *MetadataQuery) Where(ps ...predicate.Metadata) *MetadataQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MetadataQuery) Limit(limit int) *MetadataQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MetadataQuery) Offset(offset int) *MetadataQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MetadataQuery) Unique(unique bool) *MetadataQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MetadataQuery) Order(o ...OrderFunc) *MetadataQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryUser chains the current query on the "user" edge.
func (mq *MetadataQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, metadata.UserTable, metadata.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (mq *MetadataQuery) QueryChildren() *MetadataQuery {
	query := &MetadataQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, metadata.ChildrenTable, metadata.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (mq *MetadataQuery) QueryParent() *MetadataQuery {
	query := &MetadataQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, metadata.ParentTable, metadata.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Metadata entity from the query.
// Returns a *NotFoundError when no Metadata was found.
func (mq *MetadataQuery) First(ctx context.Context) (*Metadata, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MetadataQuery) FirstX(ctx context.Context) *Metadata {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Metadata ID from the query.
// Returns a *NotFoundError when no Metadata ID was found.
func (mq *MetadataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MetadataQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Metadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Metadata entity is found.
// Returns a *NotFoundError when no Metadata entities are found.
func (mq *MetadataQuery) Only(ctx context.Context) (*Metadata, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metadata.Label}
	default:
		return nil, &NotSingularError{metadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MetadataQuery) OnlyX(ctx context.Context) *Metadata {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Metadata ID in the query.
// Returns a *NotSingularError when more than one Metadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MetadataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metadata.Label}
	default:
		err = &NotSingularError{metadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MetadataQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetadataSlice.
func (mq *MetadataQuery) All(ctx context.Context) ([]*Metadata, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MetadataQuery) AllX(ctx context.Context) []*Metadata {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Metadata IDs.
func (mq *MetadataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(metadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MetadataQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MetadataQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MetadataQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MetadataQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MetadataQuery) Clone() *MetadataQuery {
	if mq == nil {
		return nil
	}
	return &MetadataQuery{
		config:       mq.config,
		limit:        mq.limit,
		offset:       mq.offset,
		order:        append([]OrderFunc{}, mq.order...),
		predicates:   append([]predicate.Metadata{}, mq.predicates...),
		withUser:     mq.withUser.Clone(),
		withChildren: mq.withChildren.Clone(),
		withParent:   mq.withParent.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithUser(opts ...func(*UserQuery)) *MetadataQuery {
	query := &UserQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withUser = query
	return mq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithChildren(opts ...func(*MetadataQuery)) *MetadataQuery {
	query := &MetadataQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withChildren = query
	return mq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithParent(opts ...func(*MetadataQuery)) *MetadataQuery {
	query := &MetadataQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withParent = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Age int `json:"age,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Metadata.Query().
//		GroupBy(metadata.FieldAge).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MetadataQuery) GroupBy(field string, fields ...string) *MetadataGroupBy {
	grbuild := &MetadataGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = metadata.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Age int `json:"age,omitempty"`
//	}
//
//	client.Metadata.Query().
//		Select(metadata.FieldAge).
//		Scan(ctx, &v)
//
func (mq *MetadataQuery) Select(fields ...string) *MetadataSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &MetadataSelect{MetadataQuery: mq}
	selbuild.label = metadata.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

func (mq *MetadataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !metadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Metadata, error) {
	var (
		nodes       = []*Metadata{}
		_spec       = mq.querySpec()
		loadedTypes = [3]bool{
			mq.withUser != nil,
			mq.withChildren != nil,
			mq.withParent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Metadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Metadata{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Metadata)
		for i := range nodes {
			fk := nodes[i].ID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := mq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Metadata)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*Metadata{}
		}
		query.Where(predicate.Metadata(func(s *sql.Selector) {
			s.Where(sql.InValues(metadata.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ParentID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := mq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Metadata)
		for i := range nodes {
			fk := nodes[i].ParentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(metadata.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (mq *MetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MetadataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mq *MetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metadata.Table,
			Columns: metadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadata.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for i := range fields {
			if fields[i] != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(metadata.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = metadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetadataGroupBy is the group-by builder for Metadata entities.
type MetadataGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MetadataGroupBy) Aggregate(fns ...AggregateFunc) *MetadataGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MetadataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *MetadataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgb.fields {
		if !metadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MetadataGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MetadataSelect is the builder for selecting fields of Metadata entities.
type MetadataSelect struct {
	*MetadataQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MetadataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MetadataQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *MetadataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
