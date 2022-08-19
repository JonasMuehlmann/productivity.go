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

// Link is an object representing the database table.
type Link struct {
	L             linkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
	R             *linkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	SourceID      int64  `boil:"source_id" json:"source_id" toml:"source_id" yaml:"source_id"`
	DestinationID int64  `boil:"destination_id" json:"destination_id" toml:"destination_id" yaml:"destination_id"`
}

var LinkColumns = struct {
	SourceID      string
	DestinationID string
}{
	SourceID:      "source_id",
	DestinationID: "destination_id",
}

var LinkTableColumns = struct {
	SourceID      string
	DestinationID string
}{
	SourceID:      "links.source_id",
	DestinationID: "links.destination_id",
}

// Generated where

var LinkWhere = struct {
	SourceID      whereHelperint64
	DestinationID whereHelperint64
}{
	SourceID:      whereHelperint64{field: "`links`.`source_id`"},
	DestinationID: whereHelperint64{field: "`links`.`destination_id`"},
}

// LinkRels is where relationship names are stored.
var LinkRels = struct {
}{}

// linkR is where relationships are stored.
type linkR struct {
}

// NewStruct creates a new relationship struct
func (*linkR) NewStruct() *linkR {
	return &linkR{}
}

// linkL is where Load methods for each relationship are stored.
type linkL struct{}

var (
	linkAllColumns            = []string{"source_id", "destination_id"}
	linkColumnsWithoutDefault = []string{"source_id", "destination_id"}
	linkColumnsWithDefault    = []string{}
	linkPrimaryKeyColumns     = []string{"source_id", "destination_id"}
	linkGeneratedColumns      =[]string{}
)

type (
	// LinkSlice is an alias for a slice of pointers to Link.
	// This should almost always be used instead of []Link.
	LinkSlice []*Link
	// LinkHook is the signature for custom Link hook methods
	LinkHook func(context.Context, boil.ContextExecutor, *Link) error

	linkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	linkType                 = reflect.TypeOf(&Link{})
	linkMapping              = queries.MakeStructMapping(linkType)
	linkPrimaryKeyMapping, _ = queries.BindMapping(linkType, linkMapping, linkPrimaryKeyColumns)
	linkInsertCacheMut       sync.RWMutex
	linkInsertCache          = make(map[string]insertCache)
	linkUpdateCacheMut       sync.RWMutex
	linkUpdateCache          = make(map[string]updateCache)
	linkUpsertCacheMut       sync.RWMutex
	linkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var linkAfterSelectHooks []LinkHook

var linkBeforeInsertHooks []LinkHook
var linkAfterInsertHooks []LinkHook

var linkBeforeUpdateHooks []LinkHook
var linkAfterUpdateHooks []LinkHook

var linkBeforeDeleteHooks []LinkHook
var linkAfterDeleteHooks []LinkHook

var linkBeforeUpsertHooks []LinkHook
var linkAfterUpsertHooks []LinkHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Link) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Link) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Link) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Link) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Link) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Link) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Link) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Link) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Link) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range linkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLinkHook registers your hook function for all future operations.
func AddLinkHook(hookPoint boil.HookPoint, linkHook LinkHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		linkAfterSelectHooks = append(linkAfterSelectHooks, linkHook)
	case boil.BeforeInsertHook:
		linkBeforeInsertHooks = append(linkBeforeInsertHooks, linkHook)
	case boil.AfterInsertHook:
		linkAfterInsertHooks = append(linkAfterInsertHooks, linkHook)
	case boil.BeforeUpdateHook:
		linkBeforeUpdateHooks = append(linkBeforeUpdateHooks, linkHook)
	case boil.AfterUpdateHook:
		linkAfterUpdateHooks = append(linkAfterUpdateHooks, linkHook)
	case boil.BeforeDeleteHook:
		linkBeforeDeleteHooks = append(linkBeforeDeleteHooks, linkHook)
	case boil.AfterDeleteHook:
		linkAfterDeleteHooks = append(linkAfterDeleteHooks, linkHook)
	case boil.BeforeUpsertHook:
		linkBeforeUpsertHooks = append(linkBeforeUpsertHooks, linkHook)
	case boil.AfterUpsertHook:
		linkAfterUpsertHooks = append(linkAfterUpsertHooks, linkHook)
	}
}

// One returns a single link record from the query.
func (q linkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Link, error) {
	o := &Link{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for links")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Link records from the query.
func (q linkQuery) All(ctx context.Context, exec boil.ContextExecutor) (LinkSlice, error) {
	var o []*Link

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to Link slice")
	}

	if len(linkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Link records in the query.
func (q linkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count links rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q linkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if links exists")
	}

	return count > 0, nil
}

// Links retrieves all the records using an executor.
func Links(mods ...qm.QueryMod) linkQuery {
	mods = append(mods, qm.From("`links`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`links`.*"})
	}

	return linkQuery{q}
}

// FindLink retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLink(ctx context.Context, exec boil.ContextExecutor, sourceID int64, destinationID int64, selectCols ...string) (*Link, error) {
	linkObj := &Link{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `links` where `source_id`=? AND `destination_id`=?", sel,
	)

	q := queries.Raw(query, sourceID, destinationID)

	err := q.Bind(ctx, exec, linkObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: unable to select from links")
	}

	if err = linkObj.doAfterSelectHooks(ctx, exec); err != nil {
		return linkObj, err
	}

	return linkObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Link) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no links provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	linkInsertCacheMut.RLock()
	cache, cached := linkInsertCache[key]
	linkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			linkAllColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(linkType, linkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `links` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `links` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `links` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, linkPrimaryKeyColumns))
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
		return errors.Wrap(err, "repository: unable to insert into links")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.SourceID,
		o.DestinationID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "repository: unable to populate default values for links")
	}

CacheNoHooks:
	if !cached {
		linkInsertCacheMut.Lock()
		linkInsertCache[key] = cache
		linkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Link.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Link) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	linkUpdateCacheMut.RLock()
	cache, cached := linkUpdateCache[key]
	linkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			linkAllColumns,
			linkPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repository: unable to update links, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `links` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, linkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, append(wl, linkPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "repository: unable to update links row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by update for links")
	}

	if !cached {
		linkUpdateCacheMut.Lock()
		linkUpdateCache[key] = cache
		linkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q linkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all for links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected for links")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LinkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `links` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all in link slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected all in update all link")
	}
	return rowsAff, nil
}

var mySQLLinkUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Link) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no links provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLinkUniqueColumns, o)

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

	linkUpsertCacheMut.RLock()
	cache, cached := linkUpsertCache[key]
	linkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			linkAllColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			linkAllColumns,
			linkPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("repository: unable to upsert links, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`links`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `links` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(linkType, linkMapping, ret)
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
		return errors.Wrap(err, "repository: unable to upsert for links")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(linkType, linkMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "repository: unable to retrieve unique values for links")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "repository: unable to populate default values for links")
	}

CacheNoHooks:
	if !cached {
		linkUpsertCacheMut.Lock()
		linkUpsertCache[key] = cache
		linkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Link record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Link) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repository: no Link provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), linkPrimaryKeyMapping)
	sql := "DELETE FROM `links` WHERE `source_id`=? AND `destination_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete from links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by delete for links")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q linkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repository: no linkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for links")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LinkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(linkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from link slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for links")
	}

	if len(linkAfterDeleteHooks) != 0 {
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
func (o *Link) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLink(ctx, exec, o.SourceID, o.DestinationID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LinkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LinkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `links`.* FROM `links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repository: unable to reload all in LinkSlice")
	}

	*o = slice

	return nil
}

// LinkExists checks if the Link row exists.
func LinkExists(ctx context.Context, exec boil.ContextExecutor, sourceID int64, destinationID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `links` where `source_id`=? AND `destination_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, sourceID, destinationID)
	}
	row := exec.QueryRowContext(ctx, sql, sourceID, destinationID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repository: unable to check if links exists")
	}

	return exists, nil
}
