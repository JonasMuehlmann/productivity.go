// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/Tag_repository.go.tpl

package repository

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"

	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MssqlTagRepository struct {
	db *sql.DB
}
type TagField string

var TagFields = struct {
	ID  TagField
	Tag TagField
}{
	ID:  "id",
	Tag: "tag",
}

var TagFieldsList = []TagField{
	TagField("ID"),
	TagField("Tag"),
}

var TagRelationsList = []string{
	"Bookmarks",
	"ParentTagTags",
	"ChildTagTags",
	"Documents",
	"ParentTagTagParentPaths",
	"TagParentPaths",
}

type TagFilter struct {
	ID  optional.Optional[model.FilterOperation[int64]]
	Tag optional.Optional[model.FilterOperation[string]]

	Bookmarks               optional.Optional[model.UpdateOperation[BookmarkSlice]]
	ParentTagTags           optional.Optional[model.UpdateOperation[TagSlice]]
	ChildTagTags            optional.Optional[model.UpdateOperation[TagSlice]]
	Documents               optional.Optional[model.UpdateOperation[DocumentSlice]]
	ParentTagTagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
	TagParentPaths          optional.Optional[model.UpdateOperation[TagParentPathSlice]]
}

type TagFilterMapping[T any] struct {
	Field           TagField
	FilterOperation model.FilterOperation[T]
}

func (filter *TagFilter) GetSetFilters() *list.List {
	setFilters := list.New()

	if filter.ID.HasValue {
		setFilters.PushBack(TagFilterMapping[int64]{TagFields.ID, filter.ID.Wrappee})
	}
	if filter.Tag.HasValue {
		setFilters.PushBack(TagFilterMapping[string]{TagFields.Tag, filter.Tag.Wrappee})
	}

	return setFilters
}

type TagUpdater struct {
	ID  optional.Optional[model.UpdateOperation[int64]]
	Tag optional.Optional[model.UpdateOperation[string]]

	Bookmarks               optional.Optional[model.UpdateOperation[BookmarkSlice]]
	ParentTagTags           optional.Optional[model.UpdateOperation[TagSlice]]
	ChildTagTags            optional.Optional[model.UpdateOperation[TagSlice]]
	Documents               optional.Optional[model.UpdateOperation[DocumentSlice]]
	ParentTagTagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
	TagParentPaths          optional.Optional[model.UpdateOperation[TagParentPathSlice]]
}

type TagUpdaterMapping[T any] struct {
	Field   TagField
	Updater model.UpdateOperation[T]
}

func (updater *TagUpdater) GetSetUpdaters() *list.List {
	setUpdaters := list.New()

	if updater.ID.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[int64]{TagFields.ID, updater.ID.Wrappee})
	}
	if updater.Tag.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[string]{TagFields.Tag, updater.Tag.Wrappee})
	}

	return setUpdaters
}

func (updater *TagUpdater) ApplyToModel(tagModel *Tag) {
	if updater.ID.HasValue {
		model.ApplyUpdater(&(*tagModel).ID, updater.ID.Wrappee)
	}
	if updater.Tag.HasValue {
		model.ApplyUpdater(&(*tagModel).Tag, updater.Tag.Wrappee)
	}

}

type MssqlTagRepositoryHook func(context.Context, MssqlTagRepository) error

type queryModSliceTag []qm.QueryMod

func (s queryModSliceTag) Apply(q *queries.Query) {
	qm.Apply(q, s...)
}

func buildQueryModFilterTag[T any](filterField TagField, filterOperation model.FilterOperation[T]) queryModSliceTag {
	var newQueryMod queryModSliceTag

	filterOperator := filterOperation.Operator

	switch filterOperator {
	case model.FilterEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" = ?", filterOperand.Operand))
	case model.FilterNEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterNEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" != ?", filterOperand.Operand))
	case model.FilterGreaterThan:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterGreaterThan operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" > ?", filterOperand.Operand))
	case model.FilterGreaterThanEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterGreaterThanEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" >= ?", filterOperand.Operand))
	case model.FilterLessThan:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterLessThan operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" < ?", filterOperand.Operand))
	case model.FilterLessThanEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterLessThanEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" <= ?", filterOperand.Operand))
	case model.FilterIn:
		filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
		if !ok {
			panic("Expected a list operand for FilterIn operator")
		}

		newQueryMod = append(newQueryMod, qm.WhereIn(string(filterField)+" IN (?)", filterOperand.Operands))
	case model.FilterNotIn:
		filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
		if !ok {
			panic("Expected a list operand for FilterNotIn operator")
		}

		newQueryMod = append(newQueryMod, qm.WhereNotIn(string(filterField)+" IN (?)", filterOperand.Operands))
	case model.FilterBetween:
		filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterBetween operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
	case model.FilterNotBetween:
		filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterNotBetween operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
	case model.FilterLike:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterLike operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" LIKE ?", filterOperand.Operand))
	case model.FilterNotLike:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterLike operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT LIKE ?", filterOperand.Operand))
	case model.FilterOr:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterOr operator")
		}
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS))))
	case model.FilterAnd:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterAnd operator")
		}

		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS)))
	default:
		panic("Unhandled FilterOperator")
	}

	return newQueryMod
}

func buildQueryModListFromFilterTag(setFilters list.List) queryModSliceTag {
	queryModList := make(queryModSliceTag, 0, 2)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(TagFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %t but got %t", TagFilterMapping[any]{}, filter))
		}

		newQueryMod := buildQueryModFilterTag(filterMapping.Field, filterMapping.FilterOperation)

		for _, queryMod := range newQueryMod {
			queryModList = append(queryModList, queryMod)
		}
	}

	return queryModList
}

func (repo *MssqlTagRepository) New(args ...any) (MssqlTagRepository, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) Add(ctx context.Context, repositoryModels []Tag) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		err = repositoryModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlTagRepository) Replace(ctx context.Context, repositoryModels []Tag) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		_, err = repositoryModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlTagRepository) UpdateWhere(ctx context.Context, columnFilter TagFilter, columnUpdater TagUpdater) (numAffectedRecords int64, err error) {
	// NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
	// Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

	// TODO: Implement translator from domainColumnFilter to repositoryColumnFilter and updater
	var modelsToUpdate TagSlice

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	modelsToUpdate, err = Tags(queryFilters...).All(ctx, repo.db)

	numAffectedRecords = int64(len(modelsToUpdate))

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, model := range modelsToUpdate {
		columnUpdater.ApplyToModel(model)
		model.Update(ctx, tx, boil.Infer())
	}

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) Delete(ctx context.Context, repositoryModels []Tag) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		_, err = repositoryModel.Delete(ctx, tx)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlTagRepository) DeleteWhere(ctx context.Context, columnFilter TagFilter) (numAffectedRecords int64, err error) {
	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Tags(queryFilters...).DeleteAll(ctx, tx)

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) CountWhere(ctx context.Context, columnFilter TagFilter) (int64, error) {
	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Count(ctx, repo.db)
}

func (repo *MssqlTagRepository) CountAll(ctx context.Context) (int64, error) {
	return Tags().Count(ctx, repo.db)
}

func (repo *MssqlTagRepository) DoesExist(ctx context.Context, repositoryModel Tag) (bool, error) {
	return TagExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *MssqlTagRepository) DoesExistWhere(ctx context.Context, columnFilter TagFilter) (bool, error) {
	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Exists(ctx, repo.db)
}

func (repo *MssqlTagRepository) GetWhere(ctx context.Context, columnFilter TagFilter) ([]*Tag, error) {
	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).All(ctx, repo.db)
}

func (repo *MssqlTagRepository) GetFirstWhere(ctx context.Context, columnFilter TagFilter) (*Tag, error) {
	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).One(ctx, repo.db)
}

func (repo *MssqlTagRepository) GetAll(ctx context.Context) ([]*Tag, error) {
	return Tags().All(ctx, repo.db)
}
