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
	"github.com/briancabbott/entgo/entc/integration/customid/ent/account"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/predicate"
	"github.com/briancabbott/entgo/entc/integration/customid/ent/token"
	"github.com/briancabbott/entgo/entc/integration/customid/sid"
	"github.com/briancabbott/entgo/schema/field"
)

// AccountQuery is the builder for querying Account entities.
type AccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Account
	// eager-loading edges.
	withToken *TokenQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountQuery builder.
func (aq *AccountQuery) Where(ps ...predicate.Account) *AccountQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AccountQuery) Limit(limit int) *AccountQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AccountQuery) Offset(offset int) *AccountQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AccountQuery) Unique(unique bool) *AccountQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AccountQuery) Order(o ...OrderFunc) *AccountQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryToken chains the current query on the "token" edge.
func (aq *AccountQuery) QueryToken() *TokenQuery {
	query := &TokenQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.TokenTable, account.TokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Account entity from the query.
// Returns a *NotFoundError when no Account was found.
func (aq *AccountQuery) First(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{account.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AccountQuery) FirstX(ctx context.Context) *Account {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Account ID from the query.
// Returns a *NotFoundError when no Account ID was found.
func (aq *AccountQuery) FirstID(ctx context.Context) (id sid.ID, err error) {
	var ids []sid.ID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{account.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AccountQuery) FirstIDX(ctx context.Context) sid.ID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Account entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Account entity is found.
// Returns a *NotFoundError when no Account entities are found.
func (aq *AccountQuery) Only(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{account.Label}
	default:
		return nil, &NotSingularError{account.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AccountQuery) OnlyX(ctx context.Context) *Account {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Account ID in the query.
// Returns a *NotSingularError when more than one Account ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AccountQuery) OnlyID(ctx context.Context) (id sid.ID, err error) {
	var ids []sid.ID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{account.Label}
	default:
		err = &NotSingularError{account.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AccountQuery) OnlyIDX(ctx context.Context) sid.ID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Accounts.
func (aq *AccountQuery) All(ctx context.Context) ([]*Account, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AccountQuery) AllX(ctx context.Context) []*Account {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Account IDs.
func (aq *AccountQuery) IDs(ctx context.Context) ([]sid.ID, error) {
	var ids []sid.ID
	if err := aq.Select(account.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AccountQuery) IDsX(ctx context.Context) []sid.ID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AccountQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AccountQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AccountQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AccountQuery) Clone() *AccountQuery {
	if aq == nil {
		return nil
	}
	return &AccountQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Account{}, aq.predicates...),
		withToken:  aq.withToken.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithToken tells the query-builder to eager-load the nodes that are connected to
// the "token" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithToken(opts ...func(*TokenQuery)) *AccountQuery {
	query := &TokenQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withToken = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Account.Query().
//		GroupBy(account.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AccountQuery) GroupBy(field string, fields ...string) *AccountGroupBy {
	grbuild := &AccountGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = account.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//	}
//
//	client.Account.Query().
//		Select(account.FieldEmail).
//		Scan(ctx, &v)
//
func (aq *AccountQuery) Select(fields ...string) *AccountSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AccountSelect{AccountQuery: aq}
	selbuild.label = account.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !account.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Account, error) {
	var (
		nodes       = []*Account{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withToken != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Account).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Account{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withToken; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[sid.ID]*Account)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Token = []*Token{}
		}
		query.withFKs = true
		query.Where(predicate.Token(func(s *sql.Selector) {
			s.Where(sql.InValues(account.TokenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.account_token
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "account_token" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "account_token" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Token = append(node.Edges.Token, n)
		}
	}

	return nodes, nil
}

func (aq *AccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: account.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for i := range fields {
			if fields[i] != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(account.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = account.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountGroupBy is the group-by builder for Account entities.
type AccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AccountGroupBy) Aggregate(fns ...AggregateFunc) *AccountGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !account.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AccountGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AccountSelect is the builder for selecting fields of Account entities.
type AccountSelect struct {
	*AccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AccountQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
