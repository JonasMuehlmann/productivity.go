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

// TagParentPath is an object representing the database table.
type TagParentPath struct {
	L           tagParentPathL  `boil:"-" json:"-" toml:"-" yaml:"-"`
	R           *tagParentPathR `boil:"-" json:"-" toml:"-" yaml:"-"`
	TagID       int64           `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	ParentTagID int64           `boil:"parent_tag_id" json:"parent_tag_id" toml:"parent_tag_id" yaml:"parent_tag_id"`
	Distance    int64           `boil:"distance" json:"distance" toml:"distance" yaml:"distance"`
}

var TagParentPathColumns = struct {
	TagID       string
	ParentTagID string
	Distance    string
}{
	TagID:       "tag_id",
	ParentTagID: "parent_tag_id",
	Distance:    "distance",
}

var TagParentPathTableColumns = struct {
	TagID       string
	ParentTagID string
	Distance    string
}{
	TagID:       "tag_parent_paths.tag_id",
	ParentTagID: "tag_parent_paths.parent_tag_id",
	Distance:    "tag_parent_paths.distance",
}

// Generated where

var TagParentPathWhere = struct {
	TagID       whereHelperint64
	ParentTagID whereHelperint64
	Distance    whereHelperint64
}{
	TagID:       whereHelperint64{field: "\"tag_parent_paths\".\"tag_id\""},
	ParentTagID: whereHelperint64{field: "\"tag_parent_paths\".\"parent_tag_id\""},
	Distance:    whereHelperint64{field: "\"tag_parent_paths\".\"distance\""},
}

// TagParentPathRels is where relationship names are stored.
var TagParentPathRels = struct {
	ParentTag string
	Tag       string
}{
	ParentTag: "ParentTag",
	Tag:       "Tag",
}

// tagParentPathR is where relationships are stored.
type tagParentPathR struct {
	ParentTag *Tag `boil:"ParentTag" json:"ParentTag" toml:"ParentTag" yaml:"ParentTag"`
	Tag       *Tag `boil:"Tag" json:"Tag" toml:"Tag" yaml:"Tag"`
}

// NewStruct creates a new relationship struct
func (*tagParentPathR) NewStruct() *tagParentPathR {
	return &tagParentPathR{}
}

func (r *tagParentPathR) GetParentTag() *Tag {
	if r == nil {
		return nil
	}
	return r.ParentTag
}

func (r *tagParentPathR) GetTag() *Tag {
	if r == nil {
		return nil
	}
	return r.Tag
}

// tagParentPathL is where Load methods for each relationship are stored.
type tagParentPathL struct{}

var (
	tagParentPathAllColumns            = []string{"tag_id", "parent_tag_id", "distance"}
	tagParentPathColumnsWithoutDefault = []string{"tag_id", "distance"}
	tagParentPathColumnsWithDefault    = []string{"parent_tag_id"}
	tagParentPathPrimaryKeyColumns     = []string{"tag_id", "parent_tag_id"}
	tagParentPathGeneratedColumns      = []string{"parent_tag_id"}
)

type (
	// TagParentPathSlice is an alias for a slice of pointers to TagParentPath.
	// This should almost always be used instead of []TagParentPath.
	TagParentPathSlice []*TagParentPath
	// TagParentPathHook is the signature for custom TagParentPath hook methods
	TagParentPathHook func(context.Context, boil.ContextExecutor, *TagParentPath) error

	tagParentPathQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagParentPathType                 = reflect.TypeOf(&TagParentPath{})
	tagParentPathMapping              = queries.MakeStructMapping(tagParentPathType)
	tagParentPathPrimaryKeyMapping, _ = queries.BindMapping(tagParentPathType, tagParentPathMapping, tagParentPathPrimaryKeyColumns)
	tagParentPathInsertCacheMut       sync.RWMutex
	tagParentPathInsertCache          = make(map[string]insertCache)
	tagParentPathUpdateCacheMut       sync.RWMutex
	tagParentPathUpdateCache          = make(map[string]updateCache)
	tagParentPathUpsertCacheMut       sync.RWMutex
	tagParentPathUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagParentPathAfterSelectHooks []TagParentPathHook

var tagParentPathBeforeInsertHooks []TagParentPathHook
var tagParentPathAfterInsertHooks []TagParentPathHook

var tagParentPathBeforeUpdateHooks []TagParentPathHook
var tagParentPathAfterUpdateHooks []TagParentPathHook

var tagParentPathBeforeDeleteHooks []TagParentPathHook
var tagParentPathAfterDeleteHooks []TagParentPathHook

var tagParentPathBeforeUpsertHooks []TagParentPathHook
var tagParentPathAfterUpsertHooks []TagParentPathHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TagParentPath) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TagParentPath) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TagParentPath) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TagParentPath) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TagParentPath) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TagParentPath) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TagParentPath) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TagParentPath) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TagParentPath) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagParentPathAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagParentPathHook registers your hook function for all future operations.
func AddTagParentPathHook(hookPoint boil.HookPoint, tagParentPathHook TagParentPathHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tagParentPathAfterSelectHooks = append(tagParentPathAfterSelectHooks, tagParentPathHook)
	case boil.BeforeInsertHook:
		tagParentPathBeforeInsertHooks = append(tagParentPathBeforeInsertHooks, tagParentPathHook)
	case boil.AfterInsertHook:
		tagParentPathAfterInsertHooks = append(tagParentPathAfterInsertHooks, tagParentPathHook)
	case boil.BeforeUpdateHook:
		tagParentPathBeforeUpdateHooks = append(tagParentPathBeforeUpdateHooks, tagParentPathHook)
	case boil.AfterUpdateHook:
		tagParentPathAfterUpdateHooks = append(tagParentPathAfterUpdateHooks, tagParentPathHook)
	case boil.BeforeDeleteHook:
		tagParentPathBeforeDeleteHooks = append(tagParentPathBeforeDeleteHooks, tagParentPathHook)
	case boil.AfterDeleteHook:
		tagParentPathAfterDeleteHooks = append(tagParentPathAfterDeleteHooks, tagParentPathHook)
	case boil.BeforeUpsertHook:
		tagParentPathBeforeUpsertHooks = append(tagParentPathBeforeUpsertHooks, tagParentPathHook)
	case boil.AfterUpsertHook:
		tagParentPathAfterUpsertHooks = append(tagParentPathAfterUpsertHooks, tagParentPathHook)
	}
}

// One returns a single tagParentPath record from the query.
func (q tagParentPathQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TagParentPath, error) {
	o := &TagParentPath{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tag_parent_paths")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TagParentPath records from the query.
func (q tagParentPathQuery) All(ctx context.Context, exec boil.ContextExecutor) (TagParentPathSlice, error) {
	var o []*TagParentPath

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TagParentPath slice")
	}

	if len(tagParentPathAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TagParentPath records in the query.
func (q tagParentPathQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tag_parent_paths rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagParentPathQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tag_parent_paths exists")
	}

	return count > 0, nil
}

// ParentTag pointed to by the foreign key.
func (o *TagParentPath) ParentTag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ParentTagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// Tag pointed to by the foreign key.
func (o *TagParentPath) Tag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// LoadParentTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tagParentPathL) LoadParentTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagParentPath interface{}, mods queries.Applicator) error {
	var slice []*TagParentPath
	var object *TagParentPath

	if singular {
		object = maybeTagParentPath.(*TagParentPath)
	} else {
		slice = *maybeTagParentPath.(*[]*TagParentPath)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagParentPathR{}
		}
		args = append(args, object.ParentTagID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagParentPathR{}
			}

			for _, a := range args {
				if a == obj.ParentTagID {
					continue Outer
				}
			}

			args = append(args, obj.ParentTagID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(tagParentPathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ParentTag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.ParentTagTagParentPaths = append(foreign.R.ParentTagTagParentPaths, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ParentTagID == foreign.ID {
				local.R.ParentTag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.ParentTagTagParentPaths = append(foreign.R.ParentTagTagParentPaths, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tagParentPathL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagParentPath interface{}, mods queries.Applicator) error {
	var slice []*TagParentPath
	var object *TagParentPath

	if singular {
		object = maybeTagParentPath.(*TagParentPath)
	} else {
		slice = *maybeTagParentPath.(*[]*TagParentPath)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagParentPathR{}
		}
		args = append(args, object.TagID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagParentPathR{}
			}

			for _, a := range args {
				if a == obj.TagID {
					continue Outer
				}
			}

			args = append(args, obj.TagID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(tagParentPathAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.TagParentPaths = append(foreign.R.TagParentPaths, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TagID == foreign.ID {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.TagParentPaths = append(foreign.R.TagParentPaths, local)
				break
			}
		}
	}

	return nil
}

// SetParentTag of the tagParentPath to the related item.
// Sets o.R.ParentTag to related.
// Adds o to related.R.ParentTagTagParentPaths.
func (o *TagParentPath) SetParentTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tag_parent_paths\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"parent_tag_id"}),
		strmangle.WhereClause("\"", "\"", 0, tagParentPathPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.TagID, o.ParentTagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ParentTagID = related.ID
	if o.R == nil {
		o.R = &tagParentPathR{
			ParentTag: related,
		}
	} else {
		o.R.ParentTag = related
	}

	if related.R == nil {
		related.R = &tagR{
			ParentTagTagParentPaths: TagParentPathSlice{o},
		}
	} else {
		related.R.ParentTagTagParentPaths = append(related.R.ParentTagTagParentPaths, o)
	}

	return nil
}

// SetTag of the tagParentPath to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.TagParentPaths.
func (o *TagParentPath) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tag_parent_paths\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"tag_id"}),
		strmangle.WhereClause("\"", "\"", 0, tagParentPathPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.TagID, o.ParentTagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TagID = related.ID
	if o.R == nil {
		o.R = &tagParentPathR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &tagR{
			TagParentPaths: TagParentPathSlice{o},
		}
	} else {
		related.R.TagParentPaths = append(related.R.TagParentPaths, o)
	}

	return nil
}

// TagParentPaths retrieves all the records using an executor.
func TagParentPaths(mods ...qm.QueryMod) tagParentPathQuery {
	mods = append(mods, qm.From("\"tag_parent_paths\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tag_parent_paths\".*"})
	}

	return tagParentPathQuery{q}
}

// FindTagParentPath retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTagParentPath(ctx context.Context, exec boil.ContextExecutor, tagID int64, parentTagID int64, selectCols ...string) (*TagParentPath, error) {
	tagParentPathObj := &TagParentPath{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tag_parent_paths\" where \"tag_id\"=? AND \"parent_tag_id\"=?", sel,
	)

	q := queries.Raw(query, tagID, parentTagID)

	err := q.Bind(ctx, exec, tagParentPathObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tag_parent_paths")
	}

	if err = tagParentPathObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tagParentPathObj, err
	}

	return tagParentPathObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TagParentPath) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag_parent_paths provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagParentPathColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagParentPathInsertCacheMut.RLock()
	cache, cached := tagParentPathInsertCache[key]
	tagParentPathInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagParentPathAllColumns,
			tagParentPathColumnsWithDefault,
			tagParentPathColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, tagParentPathGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(tagParentPathType, tagParentPathMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagParentPathType, tagParentPathMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tag_parent_paths\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tag_parent_paths\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tag_parent_paths")
	}

	if !cached {
		tagParentPathInsertCacheMut.Lock()
		tagParentPathInsertCache[key] = cache
		tagParentPathInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TagParentPath.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TagParentPath) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tagParentPathUpdateCacheMut.RLock()
	cache, cached := tagParentPathUpdateCache[key]
	tagParentPathUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagParentPathAllColumns,
			tagParentPathPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tagParentPathGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tag_parent_paths, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tag_parent_paths\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, tagParentPathPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagParentPathType, tagParentPathMapping, append(wl, tagParentPathPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tag_parent_paths row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tag_parent_paths")
	}

	if !cached {
		tagParentPathUpdateCacheMut.Lock()
		tagParentPathUpdateCache[key] = cache
		tagParentPathUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tagParentPathQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tag_parent_paths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tag_parent_paths")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagParentPathSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagParentPathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tag_parent_paths\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagParentPathPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tagParentPath slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tagParentPath")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TagParentPath) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag_parent_paths provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagParentPathColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagParentPathUpsertCacheMut.RLock()
	cache, cached := tagParentPathUpsertCache[key]
	tagParentPathUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tagParentPathAllColumns,
			tagParentPathColumnsWithDefault,
			tagParentPathColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tagParentPathAllColumns,
			tagParentPathPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tag_parent_paths, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tagParentPathPrimaryKeyColumns))
			copy(conflict, tagParentPathPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"tag_parent_paths\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tagParentPathType, tagParentPathMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagParentPathType, tagParentPathMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert tag_parent_paths")
	}

	if !cached {
		tagParentPathUpsertCacheMut.Lock()
		tagParentPathUpsertCache[key] = cache
		tagParentPathUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TagParentPath record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TagParentPath) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TagParentPath provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagParentPathPrimaryKeyMapping)
	sql := "DELETE FROM \"tag_parent_paths\" WHERE \"tag_id\"=? AND \"parent_tag_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tag_parent_paths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tag_parent_paths")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagParentPathQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tagParentPathQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tag_parent_paths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tag_parent_paths")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagParentPathSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tagParentPathBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagParentPathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tag_parent_paths\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagParentPathPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tagParentPath slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tag_parent_paths")
	}

	if len(tagParentPathAfterDeleteHooks) != 0 {
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
func (o *TagParentPath) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTagParentPath(ctx, exec, o.TagID, o.ParentTagID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagParentPathSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagParentPathSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagParentPathPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tag_parent_paths\".* FROM \"tag_parent_paths\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagParentPathPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TagParentPathSlice")
	}

	*o = slice

	return nil
}

// TagParentPathExists checks if the TagParentPath row exists.
func TagParentPathExists(ctx context.Context, exec boil.ContextExecutor, tagID int64, parentTagID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tag_parent_paths\" where \"tag_id\"=? AND \"parent_tag_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tagID, parentTagID)
	}
	row := exec.QueryRowContext(ctx, sql, tagID, parentTagID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tag_parent_paths exists")
	}

	return exists, nil
}
