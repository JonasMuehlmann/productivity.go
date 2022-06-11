// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DocumentContext is an object representing the database table.
type DocumentContext struct {
	L          documentContextL  `boil:"-" json:"-" toml:"-" yaml:"-"`
	R          *documentContextR `boil:"-" json:"-" toml:"-" yaml:"-"`
	DocumentID int64             `boil:"document_id" json:"document_id" toml:"document_id" yaml:"document_id"`
	TagID      int64             `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
}

var DocumentContextColumns = struct {
	DocumentID string
	TagID      string
}{
	DocumentID: "document_id",
	TagID:      "tag_id",
}

var DocumentContextTableColumns = struct {
	DocumentID string
	TagID      string
}{
	DocumentID: "document_contexts.document_id",
	TagID:      "document_contexts.tag_id",
}

// Generated where

var DocumentContextWhere = struct {
	DocumentID whereHelperint64
	TagID      whereHelperint64
}{
	DocumentID: whereHelperint64{field: "`document_contexts`.`document_id`"},
	TagID:      whereHelperint64{field: "`document_contexts`.`tag_id`"},
}

// DocumentContextRels is where relationship names are stored.
var DocumentContextRels = struct {
}{}

// documentContextR is where relationships are stored.
type documentContextR struct {
}

// NewStruct creates a new relationship struct
func (*documentContextR) NewStruct() *documentContextR {
	return &documentContextR{}
}

// documentContextL is where Load methods for each relationship are stored.
type documentContextL struct{}

var (
	documentContextAllColumns            = []string{"document_id", "tag_id"}
	documentContextColumnsWithoutDefault = []string{"document_id", "tag_id"}
	documentContextColumnsWithDefault    = []string{}
	documentContextPrimaryKeyColumns     = []string{"tag_id", "document_id"}
	documentContextGeneratedColumns      = []string{}
)

type (
	// DocumentContextSlice is an alias for a slice of pointers to DocumentContext.
	// This should almost always be used instead of []DocumentContext.
	DocumentContextSlice []*DocumentContext
	// DocumentContextHook is the signature for custom DocumentContext hook methods
	DocumentContextHook func(context.Context, boil.ContextExecutor, *DocumentContext) error

	documentContextQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	documentContextType                 = reflect.TypeOf(&DocumentContext{})
	documentContextMapping              = queries.MakeStructMapping(documentContextType)
	documentContextPrimaryKeyMapping, _ = queries.BindMapping(documentContextType, documentContextMapping, documentContextPrimaryKeyColumns)
	documentContextInsertCacheMut       sync.RWMutex
	documentContextInsertCache          = make(map[string]insertCache)
	documentContextUpdateCacheMut       sync.RWMutex
	documentContextUpdateCache          = make(map[string]updateCache)
	documentContextUpsertCacheMut       sync.RWMutex
	documentContextUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var documentContextAfterSelectHooks []DocumentContextHook

var documentContextBeforeInsertHooks []DocumentContextHook
var documentContextAfterInsertHooks []DocumentContextHook

var documentContextBeforeUpdateHooks []DocumentContextHook
var documentContextAfterUpdateHooks []DocumentContextHook

var documentContextBeforeDeleteHooks []DocumentContextHook
var documentContextAfterDeleteHooks []DocumentContextHook

var documentContextBeforeUpsertHooks []DocumentContextHook
var documentContextAfterUpsertHooks []DocumentContextHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DocumentContext) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DocumentContext) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DocumentContext) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DocumentContext) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DocumentContext) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DocumentContext) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DocumentContext) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DocumentContext) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DocumentContext) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentContextAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDocumentContextHook registers your hook function for all future operations.
func AddDocumentContextHook(hookPoint boil.HookPoint, documentContextHook DocumentContextHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		documentContextAfterSelectHooks = append(documentContextAfterSelectHooks, documentContextHook)
	case boil.BeforeInsertHook:
		documentContextBeforeInsertHooks = append(documentContextBeforeInsertHooks, documentContextHook)
	case boil.AfterInsertHook:
		documentContextAfterInsertHooks = append(documentContextAfterInsertHooks, documentContextHook)
	case boil.BeforeUpdateHook:
		documentContextBeforeUpdateHooks = append(documentContextBeforeUpdateHooks, documentContextHook)
	case boil.AfterUpdateHook:
		documentContextAfterUpdateHooks = append(documentContextAfterUpdateHooks, documentContextHook)
	case boil.BeforeDeleteHook:
		documentContextBeforeDeleteHooks = append(documentContextBeforeDeleteHooks, documentContextHook)
	case boil.AfterDeleteHook:
		documentContextAfterDeleteHooks = append(documentContextAfterDeleteHooks, documentContextHook)
	case boil.BeforeUpsertHook:
		documentContextBeforeUpsertHooks = append(documentContextBeforeUpsertHooks, documentContextHook)
	case boil.AfterUpsertHook:
		documentContextAfterUpsertHooks = append(documentContextAfterUpsertHooks, documentContextHook)
	}
}

// One returns a single documentContext record from the query.
func (q documentContextQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DocumentContext, error) {
	o := &DocumentContext{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for document_contexts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DocumentContext records from the query.
func (q documentContextQuery) All(ctx context.Context, exec boil.ContextExecutor) (DocumentContextSlice, error) {
	var o []*DocumentContext

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to DocumentContext slice")
	}

	if len(documentContextAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DocumentContext records in the query.
func (q documentContextQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count document_contexts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q documentContextQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if document_contexts exists")
	}

	return count > 0, nil
}

// DocumentContexts retrieves all the records using an executor.
func DocumentContexts(mods ...qm.QueryMod) documentContextQuery {
	mods = append(mods, qm.From("`document_contexts`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`document_contexts`.*"})
	}

	return documentContextQuery{q}
}

// FindDocumentContext retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDocumentContext(ctx context.Context, exec boil.ContextExecutor, tagID int64, documentID int64, selectCols ...string) (*DocumentContext, error) {
	documentContextObj := &DocumentContext{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `document_contexts` where `tag_id`=? AND `document_id`=?", sel,
	)

	q := queries.Raw(query, tagID, documentID)

	err := q.Bind(ctx, exec, documentContextObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: unable to select from document_contexts")
	}

	if err = documentContextObj.doAfterSelectHooks(ctx, exec); err != nil {
		return documentContextObj, err
	}

	return documentContextObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DocumentContext) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no document_contexts provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentContextColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	documentContextInsertCacheMut.RLock()
	cache, cached := documentContextInsertCache[key]
	documentContextInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			documentContextAllColumns,
			documentContextColumnsWithDefault,
			documentContextColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(documentContextType, documentContextMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(documentContextType, documentContextMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `document_contexts` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `document_contexts` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `document_contexts` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, documentContextPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "repository: unable to insert into document_contexts")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TagID,
		o.DocumentID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "repository: unable to populate default values for document_contexts")
	}

CacheNoHooks:
	if !cached {
		documentContextInsertCacheMut.Lock()
		documentContextInsertCache[key] = cache
		documentContextInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DocumentContext.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DocumentContext) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	documentContextUpdateCacheMut.RLock()
	cache, cached := documentContextUpdateCache[key]
	documentContextUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			documentContextAllColumns,
			documentContextPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repository: unable to update document_contexts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `document_contexts` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, documentContextPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(documentContextType, documentContextMapping, append(wl, documentContextPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update document_contexts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by update for document_contexts")
	}

	if !cached {
		documentContextUpdateCacheMut.Lock()
		documentContextUpdateCache[key] = cache
		documentContextUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q documentContextQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all for document_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected for document_contexts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DocumentContextSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("repository: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `document_contexts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentContextPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all in documentContext slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected all in update all documentContext")
	}
	return rowsAff, nil
}

var mySQLDocumentContextUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DocumentContext) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no document_contexts provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentContextColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDocumentContextUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	documentContextUpsertCacheMut.RLock()
	cache, cached := documentContextUpsertCache[key]
	documentContextUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			documentContextAllColumns,
			documentContextColumnsWithDefault,
			documentContextColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			documentContextAllColumns,
			documentContextPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("repository: unable to upsert document_contexts, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`document_contexts`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `document_contexts` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(documentContextType, documentContextMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(documentContextType, documentContextMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "repository: unable to upsert for document_contexts")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(documentContextType, documentContextMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "repository: unable to retrieve unique values for document_contexts")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "repository: unable to populate default values for document_contexts")
	}

CacheNoHooks:
	if !cached {
		documentContextUpsertCacheMut.Lock()
		documentContextUpsertCache[key] = cache
		documentContextUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DocumentContext record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DocumentContext) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repository: no DocumentContext provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), documentContextPrimaryKeyMapping)
	sql := "DELETE FROM `document_contexts` WHERE `tag_id`=? AND `document_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete from document_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by delete for document_contexts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q documentContextQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repository: no documentContextQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from document_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for document_contexts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DocumentContextSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(documentContextBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `document_contexts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentContextPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from documentContext slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for document_contexts")
	}

	if len(documentContextAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DocumentContext) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDocumentContext(ctx, exec, o.TagID, o.DocumentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DocumentContextSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DocumentContextSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `document_contexts`.* FROM `document_contexts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentContextPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repository: unable to reload all in DocumentContextSlice")
	}

	*o = slice

	return nil
}

// DocumentContextExists checks if the DocumentContext row exists.
func DocumentContextExists(ctx context.Context, exec boil.ContextExecutor, tagID int64, documentID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `document_contexts` where `tag_id`=? AND `document_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tagID, documentID)
	}
	row := exec.QueryRowContext(ctx, sql, tagID, documentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repository: unable to check if document_contexts exists")
	}

	return exists, nil
}
