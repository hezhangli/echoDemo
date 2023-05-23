// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echoDemo/ent/predicate"
	"echoDemo/ent/userdemo"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDemoQuery is the builder for querying UserDemo entities.
type UserDemoQuery struct {
	config
	ctx        *QueryContext
	order      []userdemo.OrderOption
	inters     []Interceptor
	predicates []predicate.UserDemo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserDemoQuery builder.
func (udq *UserDemoQuery) Where(ps ...predicate.UserDemo) *UserDemoQuery {
	udq.predicates = append(udq.predicates, ps...)
	return udq
}

// Limit the number of records to be returned by this query.
func (udq *UserDemoQuery) Limit(limit int) *UserDemoQuery {
	udq.ctx.Limit = &limit
	return udq
}

// Offset to start from.
func (udq *UserDemoQuery) Offset(offset int) *UserDemoQuery {
	udq.ctx.Offset = &offset
	return udq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (udq *UserDemoQuery) Unique(unique bool) *UserDemoQuery {
	udq.ctx.Unique = &unique
	return udq
}

// Order specifies how the records should be ordered.
func (udq *UserDemoQuery) Order(o ...userdemo.OrderOption) *UserDemoQuery {
	udq.order = append(udq.order, o...)
	return udq
}

// First returns the first UserDemo entity from the query.
// Returns a *NotFoundError when no UserDemo was found.
func (udq *UserDemoQuery) First(ctx context.Context) (*UserDemo, error) {
	nodes, err := udq.Limit(1).All(setContextOp(ctx, udq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userdemo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (udq *UserDemoQuery) FirstX(ctx context.Context) *UserDemo {
	node, err := udq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserDemo ID from the query.
// Returns a *NotFoundError when no UserDemo ID was found.
func (udq *UserDemoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = udq.Limit(1).IDs(setContextOp(ctx, udq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userdemo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (udq *UserDemoQuery) FirstIDX(ctx context.Context) string {
	id, err := udq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserDemo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserDemo entity is found.
// Returns a *NotFoundError when no UserDemo entities are found.
func (udq *UserDemoQuery) Only(ctx context.Context) (*UserDemo, error) {
	nodes, err := udq.Limit(2).All(setContextOp(ctx, udq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userdemo.Label}
	default:
		return nil, &NotSingularError{userdemo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (udq *UserDemoQuery) OnlyX(ctx context.Context) *UserDemo {
	node, err := udq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserDemo ID in the query.
// Returns a *NotSingularError when more than one UserDemo ID is found.
// Returns a *NotFoundError when no entities are found.
func (udq *UserDemoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = udq.Limit(2).IDs(setContextOp(ctx, udq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userdemo.Label}
	default:
		err = &NotSingularError{userdemo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (udq *UserDemoQuery) OnlyIDX(ctx context.Context) string {
	id, err := udq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserDemos.
func (udq *UserDemoQuery) All(ctx context.Context) ([]*UserDemo, error) {
	ctx = setContextOp(ctx, udq.ctx, "All")
	if err := udq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserDemo, *UserDemoQuery]()
	return withInterceptors[[]*UserDemo](ctx, udq, qr, udq.inters)
}

// AllX is like All, but panics if an error occurs.
func (udq *UserDemoQuery) AllX(ctx context.Context) []*UserDemo {
	nodes, err := udq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserDemo IDs.
func (udq *UserDemoQuery) IDs(ctx context.Context) (ids []string, err error) {
	if udq.ctx.Unique == nil && udq.path != nil {
		udq.Unique(true)
	}
	ctx = setContextOp(ctx, udq.ctx, "IDs")
	if err = udq.Select(userdemo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (udq *UserDemoQuery) IDsX(ctx context.Context) []string {
	ids, err := udq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (udq *UserDemoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, udq.ctx, "Count")
	if err := udq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, udq, querierCount[*UserDemoQuery](), udq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (udq *UserDemoQuery) CountX(ctx context.Context) int {
	count, err := udq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (udq *UserDemoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, udq.ctx, "Exist")
	switch _, err := udq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (udq *UserDemoQuery) ExistX(ctx context.Context) bool {
	exist, err := udq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserDemoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (udq *UserDemoQuery) Clone() *UserDemoQuery {
	if udq == nil {
		return nil
	}
	return &UserDemoQuery{
		config:     udq.config,
		ctx:        udq.ctx.Clone(),
		order:      append([]userdemo.OrderOption{}, udq.order...),
		inters:     append([]Interceptor{}, udq.inters...),
		predicates: append([]predicate.UserDemo{}, udq.predicates...),
		// clone intermediate query.
		sql:  udq.sql.Clone(),
		path: udq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserDemo.Query().
//		GroupBy(userdemo.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (udq *UserDemoQuery) GroupBy(field string, fields ...string) *UserDemoGroupBy {
	udq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserDemoGroupBy{build: udq}
	grbuild.flds = &udq.ctx.Fields
	grbuild.label = userdemo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.UserDemo.Query().
//		Select(userdemo.FieldName).
//		Scan(ctx, &v)
func (udq *UserDemoQuery) Select(fields ...string) *UserDemoSelect {
	udq.ctx.Fields = append(udq.ctx.Fields, fields...)
	sbuild := &UserDemoSelect{UserDemoQuery: udq}
	sbuild.label = userdemo.Label
	sbuild.flds, sbuild.scan = &udq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserDemoSelect configured with the given aggregations.
func (udq *UserDemoQuery) Aggregate(fns ...AggregateFunc) *UserDemoSelect {
	return udq.Select().Aggregate(fns...)
}

func (udq *UserDemoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range udq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, udq); err != nil {
				return err
			}
		}
	}
	for _, f := range udq.ctx.Fields {
		if !userdemo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if udq.path != nil {
		prev, err := udq.path(ctx)
		if err != nil {
			return err
		}
		udq.sql = prev
	}
	return nil
}

func (udq *UserDemoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserDemo, error) {
	var (
		nodes = []*UserDemo{}
		_spec = udq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserDemo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserDemo{config: udq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, udq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (udq *UserDemoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := udq.querySpec()
	_spec.Node.Columns = udq.ctx.Fields
	if len(udq.ctx.Fields) > 0 {
		_spec.Unique = udq.ctx.Unique != nil && *udq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, udq.driver, _spec)
}

func (udq *UserDemoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userdemo.Table, userdemo.Columns, sqlgraph.NewFieldSpec(userdemo.FieldID, field.TypeString))
	_spec.From = udq.sql
	if unique := udq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if udq.path != nil {
		_spec.Unique = true
	}
	if fields := udq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdemo.FieldID)
		for i := range fields {
			if fields[i] != userdemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := udq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := udq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := udq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := udq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (udq *UserDemoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(udq.driver.Dialect())
	t1 := builder.Table(userdemo.Table)
	columns := udq.ctx.Fields
	if len(columns) == 0 {
		columns = userdemo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if udq.sql != nil {
		selector = udq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if udq.ctx.Unique != nil && *udq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range udq.predicates {
		p(selector)
	}
	for _, p := range udq.order {
		p(selector)
	}
	if offset := udq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := udq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserDemoGroupBy is the group-by builder for UserDemo entities.
type UserDemoGroupBy struct {
	selector
	build *UserDemoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (udgb *UserDemoGroupBy) Aggregate(fns ...AggregateFunc) *UserDemoGroupBy {
	udgb.fns = append(udgb.fns, fns...)
	return udgb
}

// Scan applies the selector query and scans the result into the given value.
func (udgb *UserDemoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, udgb.build.ctx, "GroupBy")
	if err := udgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserDemoQuery, *UserDemoGroupBy](ctx, udgb.build, udgb, udgb.build.inters, v)
}

func (udgb *UserDemoGroupBy) sqlScan(ctx context.Context, root *UserDemoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(udgb.fns))
	for _, fn := range udgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*udgb.flds)+len(udgb.fns))
		for _, f := range *udgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*udgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := udgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserDemoSelect is the builder for selecting fields of UserDemo entities.
type UserDemoSelect struct {
	*UserDemoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uds *UserDemoSelect) Aggregate(fns ...AggregateFunc) *UserDemoSelect {
	uds.fns = append(uds.fns, fns...)
	return uds
}

// Scan applies the selector query and scans the result into the given value.
func (uds *UserDemoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uds.ctx, "Select")
	if err := uds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserDemoQuery, *UserDemoSelect](ctx, uds.UserDemoQuery, uds, uds.inters, v)
}

func (uds *UserDemoSelect) sqlScan(ctx context.Context, root *UserDemoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uds.fns))
	for _, fn := range uds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
