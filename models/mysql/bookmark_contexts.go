// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// BookmarkContext is an object representing the database table.
type BookmarkContext struct {
	BookmarkID int `boil:"bookmark_id" json:"bookmark_id" toml:"bookmark_id" yaml:"bookmark_id"`
	TagID      int `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`

	R *bookmarkContextR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookmarkContextL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookmarkContextColumns = struct {
	BookmarkID string
	TagID      string
}{
	BookmarkID: "bookmark_id",
	TagID:      "tag_id",
}

var BookmarkContextTableColumns = struct {
	BookmarkID string
	TagID      string
}{
	BookmarkID: "bookmark_contexts.bookmark_id",
	TagID:      "bookmark_contexts.tag_id",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BookmarkContextWhere = struct {
	BookmarkID whereHelperint
	TagID      whereHelperint
}{
	BookmarkID: whereHelperint{field: "`bookmark_contexts`.`bookmark_id`"},
	TagID:      whereHelperint{field: "`bookmark_contexts`.`tag_id`"},
}

// BookmarkContextRels is where relationship names are stored.
var BookmarkContextRels = struct {
}{}

// bookmarkContextR is where relationships are stored.
type bookmarkContextR struct {
}

// NewStruct creates a new relationship struct
func (*bookmarkContextR) NewStruct() *bookmarkContextR {
	return &bookmarkContextR{}
}

// bookmarkContextL is where Load methods for each relationship are stored.
type bookmarkContextL struct{}

var (
	bookmarkContextAllColumns            = []string{"bookmark_id", "tag_id"}
	bookmarkContextColumnsWithoutDefault = []string{"bookmark_id", "tag_id"}
	bookmarkContextColumnsWithDefault    = []string{}
	bookmarkContextPrimaryKeyColumns     = []string{"tag_id", "bookmark_id"}
	bookmarkContextGeneratedColumns      = []string{}
)

type (
	// BookmarkContextSlice is an alias for a slice of pointers to BookmarkContext.
	// This should almost always be used instead of []BookmarkContext.
	BookmarkContextSlice []*BookmarkContext
	// BookmarkContextHook is the signature for custom BookmarkContext hook methods
	BookmarkContextHook func(context.Context, boil.ContextExecutor, *BookmarkContext) error

	bookmarkContextQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookmarkContextType                 = reflect.TypeOf(&BookmarkContext{})
	bookmarkContextMapping              = queries.MakeStructMapping(bookmarkContextType)
	bookmarkContextPrimaryKeyMapping, _ = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, bookmarkContextPrimaryKeyColumns)
	bookmarkContextInsertCacheMut       sync.RWMutex
	bookmarkContextInsertCache          = make(map[string]insertCache)
	bookmarkContextUpdateCacheMut       sync.RWMutex
	bookmarkContextUpdateCache          = make(map[string]updateCache)
	bookmarkContextUpsertCacheMut       sync.RWMutex
	bookmarkContextUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookmarkContextAfterSelectHooks []BookmarkContextHook

var bookmarkContextBeforeInsertHooks []BookmarkContextHook
var bookmarkContextAfterInsertHooks []BookmarkContextHook

var bookmarkContextBeforeUpdateHooks []BookmarkContextHook
var bookmarkContextAfterUpdateHooks []BookmarkContextHook

var bookmarkContextBeforeDeleteHooks []BookmarkContextHook
var bookmarkContextAfterDeleteHooks []BookmarkContextHook

var bookmarkContextBeforeUpsertHooks []BookmarkContextHook
var bookmarkContextAfterUpsertHooks []BookmarkContextHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookmarkContext) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookmarkContext) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookmarkContext) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookmarkContext) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookmarkContext) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookmarkContext) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookmarkContext) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookmarkContext) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookmarkContext) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkContextAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookmarkContextHook registers your hook function for all future operations.
func AddBookmarkContextHook(hookPoint boil.HookPoint, bookmarkContextHook BookmarkContextHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookmarkContextAfterSelectHooks = append(bookmarkContextAfterSelectHooks, bookmarkContextHook)
	case boil.BeforeInsertHook:
		bookmarkContextBeforeInsertHooks = append(bookmarkContextBeforeInsertHooks, bookmarkContextHook)
	case boil.AfterInsertHook:
		bookmarkContextAfterInsertHooks = append(bookmarkContextAfterInsertHooks, bookmarkContextHook)
	case boil.BeforeUpdateHook:
		bookmarkContextBeforeUpdateHooks = append(bookmarkContextBeforeUpdateHooks, bookmarkContextHook)
	case boil.AfterUpdateHook:
		bookmarkContextAfterUpdateHooks = append(bookmarkContextAfterUpdateHooks, bookmarkContextHook)
	case boil.BeforeDeleteHook:
		bookmarkContextBeforeDeleteHooks = append(bookmarkContextBeforeDeleteHooks, bookmarkContextHook)
	case boil.AfterDeleteHook:
		bookmarkContextAfterDeleteHooks = append(bookmarkContextAfterDeleteHooks, bookmarkContextHook)
	case boil.BeforeUpsertHook:
		bookmarkContextBeforeUpsertHooks = append(bookmarkContextBeforeUpsertHooks, bookmarkContextHook)
	case boil.AfterUpsertHook:
		bookmarkContextAfterUpsertHooks = append(bookmarkContextAfterUpsertHooks, bookmarkContextHook)
	}
}

// One returns a single bookmarkContext record from the query.
func (q bookmarkContextQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookmarkContext, error) {
	o := &BookmarkContext{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bookmark_contexts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookmarkContext records from the query.
func (q bookmarkContextQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookmarkContextSlice, error) {
	var o []*BookmarkContext

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookmarkContext slice")
	}

	if len(bookmarkContextAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookmarkContext records in the query.
func (q bookmarkContextQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bookmark_contexts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookmarkContextQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bookmark_contexts exists")
	}

	return count > 0, nil
}

// BookmarkContexts retrieves all the records using an executor.
func BookmarkContexts(mods ...qm.QueryMod) bookmarkContextQuery {
	mods = append(mods, qm.From("`bookmark_contexts`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`bookmark_contexts`.*"})
	}

	return bookmarkContextQuery{q}
}

// FindBookmarkContext retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookmarkContext(ctx context.Context, exec boil.ContextExecutor, tagID int, bookmarkID int, selectCols ...string) (*BookmarkContext, error) {
	bookmarkContextObj := &BookmarkContext{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bookmark_contexts` where `tag_id`=? AND `bookmark_id`=?", sel,
	)

	q := queries.Raw(query, tagID, bookmarkID)

	err := q.Bind(ctx, exec, bookmarkContextObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bookmark_contexts")
	}

	if err = bookmarkContextObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookmarkContextObj, err
	}

	return bookmarkContextObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookmarkContext) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_contexts provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkContextColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookmarkContextInsertCacheMut.RLock()
	cache, cached := bookmarkContextInsertCache[key]
	bookmarkContextInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookmarkContextAllColumns,
			bookmarkContextColumnsWithDefault,
			bookmarkContextColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bookmark_contexts` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bookmark_contexts` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bookmark_contexts` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bookmarkContextPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into bookmark_contexts")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TagID,
		o.BookmarkID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bookmark_contexts")
	}

CacheNoHooks:
	if !cached {
		bookmarkContextInsertCacheMut.Lock()
		bookmarkContextInsertCache[key] = cache
		bookmarkContextInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookmarkContext.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookmarkContext) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookmarkContextUpdateCacheMut.RLock()
	cache, cached := bookmarkContextUpdateCache[key]
	bookmarkContextUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookmarkContextAllColumns,
			bookmarkContextPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bookmark_contexts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bookmark_contexts` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bookmarkContextPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, append(wl, bookmarkContextPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bookmark_contexts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bookmark_contexts")
	}

	if !cached {
		bookmarkContextUpdateCacheMut.Lock()
		bookmarkContextUpdateCache[key] = cache
		bookmarkContextUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookmarkContextQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bookmark_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bookmark_contexts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookmarkContextSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bookmark_contexts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkContextPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookmarkContext slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookmarkContext")
	}
	return rowsAff, nil
}

var mySQLBookmarkContextUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookmarkContext) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bookmark_contexts provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkContextColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBookmarkContextUniqueColumns, o)

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

	bookmarkContextUpsertCacheMut.RLock()
	cache, cached := bookmarkContextUpsertCache[key]
	bookmarkContextUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookmarkContextAllColumns,
			bookmarkContextColumnsWithDefault,
			bookmarkContextColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookmarkContextAllColumns,
			bookmarkContextPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert bookmark_contexts, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`bookmark_contexts`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bookmark_contexts` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for bookmark_contexts")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bookmarkContextType, bookmarkContextMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for bookmark_contexts")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bookmark_contexts")
	}

CacheNoHooks:
	if !cached {
		bookmarkContextUpsertCacheMut.Lock()
		bookmarkContextUpsertCache[key] = cache
		bookmarkContextUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookmarkContext record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookmarkContext) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookmarkContext provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookmarkContextPrimaryKeyMapping)
	sql := "DELETE FROM `bookmark_contexts` WHERE `tag_id`=? AND `bookmark_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bookmark_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bookmark_contexts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookmarkContextQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookmarkContextQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmark_contexts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_contexts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookmarkContextSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookmarkContextBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bookmark_contexts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkContextPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookmarkContext slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bookmark_contexts")
	}

	if len(bookmarkContextAfterDeleteHooks) != 0 {
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
func (o *BookmarkContext) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookmarkContext(ctx, exec, o.TagID, o.BookmarkID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookmarkContextSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookmarkContextSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkContextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bookmark_contexts`.* FROM `bookmark_contexts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkContextPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookmarkContextSlice")
	}

	*o = slice

	return nil
}

// BookmarkContextExists checks if the BookmarkContext row exists.
func BookmarkContextExists(ctx context.Context, exec boil.ContextExecutor, tagID int, bookmarkID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bookmark_contexts` where `tag_id`=? AND `bookmark_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tagID, bookmarkID)
	}
	row := exec.QueryRowContext(ctx, sql, tagID, bookmarkID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bookmark_contexts exists")
	}

	return exists, nil
}
