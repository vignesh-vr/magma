// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/comment"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// CommentQuery is the builder for querying Comment entities.
type CommentQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Comment
	withFKs    bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (cq *CommentQuery) Where(ps ...predicate.Comment) *CommentQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CommentQuery) Limit(limit int) *CommentQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CommentQuery) Offset(offset int) *CommentQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *CommentQuery) Order(o ...Order) *CommentQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Comment entity in the query. Returns *NotFoundError when no comment was found.
func (cq *CommentQuery) First(ctx context.Context) (*Comment, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &NotFoundError{comment.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommentQuery) FirstX(ctx context.Context) *Comment {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first Comment id in the query. Returns *NotFoundError when no id was found.
func (cq *CommentQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comment.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *CommentQuery) FirstXID(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Comment entity in the query, returns an error if not exactly one entity was returned.
func (cq *CommentQuery) Only(ctx context.Context) (*Comment, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &NotFoundError{comment.Label}
	default:
		return nil, &NotSingularError{comment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommentQuery) OnlyX(ctx context.Context) *Comment {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only Comment id in the query, returns an error if not exactly one id was returned.
func (cq *CommentQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comment.Label}
	default:
		err = &NotSingularError{comment.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (cq *CommentQuery) OnlyXID(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Comments.
func (cq *CommentQuery) All(ctx context.Context) ([]*Comment, error) {
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommentQuery) AllX(ctx context.Context) []*Comment {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of Comment ids.
func (cq *CommentQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := cq.Select(comment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommentQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommentQuery) Count(ctx context.Context) (int, error) {
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommentQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommentQuery) Exist(ctx context.Context) (bool, error) {
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommentQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommentQuery) Clone() *CommentQuery {
	return &CommentQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]Order{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.Comment{}, cq.predicates...),
		// clone intermediate query.
		sql: cq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comment.Query().
//		GroupBy(comment.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CommentQuery) GroupBy(field string, fields ...string) *CommentGroupBy {
	group := &CommentGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = cq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Comment.Query().
//		Select(comment.FieldCreateTime).
//		Scan(ctx, &v)
//
func (cq *CommentQuery) Select(field string, fields ...string) *CommentSelect {
	selector := &CommentSelect{config: cq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = cq.sqlQuery()
	return selector
}

func (cq *CommentQuery) sqlAll(ctx context.Context) ([]*Comment, error) {
	var (
		nodes   = []*Comment{}
		withFKs = cq.withFKs
		_spec   = cq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, comment.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Comment{config: cq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *CommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: comment.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CommentQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(comment.Table)
	selector := builder.Select(t1.Columns(comment.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(comment.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommentGroupBy is the builder for group-by Comment entities.
type CommentGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentGroupBy) Aggregate(fns ...Aggregate) *CommentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *CommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CommentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *CommentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CommentGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *CommentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CommentGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *CommentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CommentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *CommentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CommentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CommentGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// CommentSelect is the builder for select fields of Comment entities.
type CommentSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (cs *CommentSelect) Scan(ctx context.Context, v interface{}) error {
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *CommentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cs *CommentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *CommentSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cs *CommentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *CommentSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cs *CommentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *CommentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cs *CommentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *CommentSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *CommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *CommentSelect) sqlQuery() sql.Querier {
	selector := cs.sql
	selector.Select(selector.Columns(cs.fields...)...)
	return selector
}