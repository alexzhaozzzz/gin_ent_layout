// Code generated by ent, DO NOT EDIT.

package ent

import (
	"back-end/merchant/internal/data/ent/predicate"
	"back-end/merchant/internal/data/ent/syspermissions"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysPermissionsQuery is the builder for querying SysPermissions entities.
type SysPermissionsQuery struct {
	config
	ctx        *QueryContext
	order      []syspermissions.OrderOption
	inters     []Interceptor
	predicates []predicate.SysPermissions
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysPermissionsQuery builder.
func (spq *SysPermissionsQuery) Where(ps ...predicate.SysPermissions) *SysPermissionsQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *SysPermissionsQuery) Limit(limit int) *SysPermissionsQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *SysPermissionsQuery) Offset(offset int) *SysPermissionsQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *SysPermissionsQuery) Unique(unique bool) *SysPermissionsQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *SysPermissionsQuery) Order(o ...syspermissions.OrderOption) *SysPermissionsQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// First returns the first SysPermissions entity from the query.
// Returns a *NotFoundError when no SysPermissions was found.
func (spq *SysPermissionsQuery) First(ctx context.Context) (*SysPermissions, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{syspermissions.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *SysPermissionsQuery) FirstX(ctx context.Context) *SysPermissions {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysPermissions ID from the query.
// Returns a *NotFoundError when no SysPermissions ID was found.
func (spq *SysPermissionsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{syspermissions.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *SysPermissionsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysPermissions entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysPermissions entity is found.
// Returns a *NotFoundError when no SysPermissions entities are found.
func (spq *SysPermissionsQuery) Only(ctx context.Context) (*SysPermissions, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{syspermissions.Label}
	default:
		return nil, &NotSingularError{syspermissions.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *SysPermissionsQuery) OnlyX(ctx context.Context) *SysPermissions {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysPermissions ID in the query.
// Returns a *NotSingularError when more than one SysPermissions ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *SysPermissionsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{syspermissions.Label}
	default:
		err = &NotSingularError{syspermissions.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *SysPermissionsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysPermissionsSlice.
func (spq *SysPermissionsQuery) All(ctx context.Context) ([]*SysPermissions, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysPermissions, *SysPermissionsQuery]()
	return withInterceptors[[]*SysPermissions](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *SysPermissionsQuery) AllX(ctx context.Context) []*SysPermissions {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysPermissions IDs.
func (spq *SysPermissionsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(syspermissions.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *SysPermissionsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *SysPermissionsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*SysPermissionsQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *SysPermissionsQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *SysPermissionsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *SysPermissionsQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysPermissionsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *SysPermissionsQuery) Clone() *SysPermissionsQuery {
	if spq == nil {
		return nil
	}
	return &SysPermissionsQuery{
		config:     spq.config,
		ctx:        spq.ctx.Clone(),
		order:      append([]syspermissions.OrderOption{}, spq.order...),
		inters:     append([]Interceptor{}, spq.inters...),
		predicates: append([]predicate.SysPermissions{}, spq.predicates...),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MenuID int64 `json:"menu_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysPermissions.Query().
//		GroupBy(syspermissions.FieldMenuID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *SysPermissionsQuery) GroupBy(field string, fields ...string) *SysPermissionsGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysPermissionsGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = syspermissions.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MenuID int64 `json:"menu_id,omitempty"`
//	}
//
//	client.SysPermissions.Query().
//		Select(syspermissions.FieldMenuID).
//		Scan(ctx, &v)
func (spq *SysPermissionsQuery) Select(fields ...string) *SysPermissionsSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &SysPermissionsSelect{SysPermissionsQuery: spq}
	sbuild.label = syspermissions.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysPermissionsSelect configured with the given aggregations.
func (spq *SysPermissionsQuery) Aggregate(fns ...AggregateFunc) *SysPermissionsSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *SysPermissionsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !syspermissions.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *SysPermissionsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysPermissions, error) {
	var (
		nodes = []*SysPermissions{}
		_spec = spq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysPermissions).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysPermissions{config: spq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (spq *SysPermissionsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *SysPermissionsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(syspermissions.Table, syspermissions.Columns, sqlgraph.NewFieldSpec(syspermissions.FieldID, field.TypeInt64))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syspermissions.FieldID)
		for i := range fields {
			if fields[i] != syspermissions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *SysPermissionsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(syspermissions.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = syspermissions.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysPermissionsGroupBy is the group-by builder for SysPermissions entities.
type SysPermissionsGroupBy struct {
	selector
	build *SysPermissionsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *SysPermissionsGroupBy) Aggregate(fns ...AggregateFunc) *SysPermissionsGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *SysPermissionsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysPermissionsQuery, *SysPermissionsGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *SysPermissionsGroupBy) sqlScan(ctx context.Context, root *SysPermissionsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysPermissionsSelect is the builder for selecting fields of SysPermissions entities.
type SysPermissionsSelect struct {
	*SysPermissionsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *SysPermissionsSelect) Aggregate(fns ...AggregateFunc) *SysPermissionsSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *SysPermissionsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysPermissionsQuery, *SysPermissionsSelect](ctx, sps.SysPermissionsQuery, sps, sps.inters, v)
}

func (sps *SysPermissionsSelect) sqlScan(ctx context.Context, root *SysPermissionsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
