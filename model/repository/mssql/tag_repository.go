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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/sql_repository.go.tpl

package repository

import (
	"container/list"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repoCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
	log "github.com/sirupsen/logrus"
	"github.com/stoewer/go-strcase"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"strconv"
)

//******************************************************************//
//                        Types and constants                       //
//******************************************************************//
type MssqlTagRepository struct {
	db *sql.DB
}

type TagField string

var TagFields = struct {
	Tag       TagField
	Path      TagField
	Children  TagField
	ParentTag TagField
	ID        TagField
}{
	Tag:       "tag",
	Path:      "path",
	Children:  "children",
	ParentTag: "parent_tag",
	ID:        "id",
}

var TagFieldsList = []TagField{
	TagField("Tag"),
	TagField("Path"),
	TagField("Children"),
	TagField("ParentTag"),
	TagField("ID"),
}

var TagRelationsList = []string{
	"ParentTagTag",
	"Bookmarks",
	"Documents",
	"ParentTagTags",
}

type TagFilter struct {
	Tag       optional.Optional[model.FilterOperation[string]]
	Path      optional.Optional[model.FilterOperation[string]]
	Children  optional.Optional[model.FilterOperation[string]]
	ParentTag optional.Optional[model.FilterOperation[null.Int64]]
	ID        optional.Optional[model.FilterOperation[int64]]

	ParentTagTag  optional.Optional[model.FilterOperation[*Tag]]
	Bookmarks     optional.Optional[model.FilterOperation[*Bookmark]]
	Documents     optional.Optional[model.FilterOperation[*Document]]
	ParentTagTags optional.Optional[model.FilterOperation[*Tag]]
}

type TagUpdater struct {
	ParentTagTag  optional.Optional[model.UpdateOperation[*Tag]]
	Tag           optional.Optional[model.UpdateOperation[string]]
	Path          optional.Optional[model.UpdateOperation[string]]
	Children      optional.Optional[model.UpdateOperation[string]]
	Bookmarks     optional.Optional[model.UpdateOperation[BookmarkSlice]]
	Documents     optional.Optional[model.UpdateOperation[DocumentSlice]]
	ParentTagTags optional.Optional[model.UpdateOperation[TagSlice]]
	ParentTag     optional.Optional[model.UpdateOperation[null.Int64]]
	ID            optional.Optional[model.UpdateOperation[int64]]
}

type TagUpdaterMapping[T any] struct {
	Field   TagField
	Updater model.UpdateOperation[T]
}

func (updater *TagUpdater) GetSetUpdaters() *list.List {
	setUpdaters := list.New()

	if updater.Tag.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Tag, Updater: updater.Tag.Wrappee})
	}
	if updater.Path.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Path, Updater: updater.Path.Wrappee})
	}
	if updater.Children.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Children, Updater: updater.Children.Wrappee})
	}
	if updater.ParentTag.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[null.Int64]{Field: TagFields.ParentTag, Updater: updater.ParentTag.Wrappee})
	}
	if updater.ID.HasValue {
		setUpdaters.PushBack(TagUpdaterMapping[int64]{Field: TagFields.ID, Updater: updater.ID.Wrappee})
	}

	return setUpdaters
}

func (updater *TagUpdater) ApplyToModel(tagModel *Tag) {
	if updater.Tag.HasValue {
		model.ApplyUpdater(&(*tagModel).Tag, updater.Tag.Wrappee)
	}
	if updater.Path.HasValue {
		model.ApplyUpdater(&(*tagModel).Path, updater.Path.Wrappee)
	}
	if updater.Children.HasValue {
		model.ApplyUpdater(&(*tagModel).Children, updater.Children.Wrappee)
	}
	if updater.ParentTag.HasValue {
		model.ApplyUpdater(&(*tagModel).ParentTag, updater.ParentTag.Wrappee)
	}
	if updater.ID.HasValue {
		model.ApplyUpdater(&(*tagModel).ID, updater.ID.Wrappee)
	}

}

type queryModSliceTag []qm.QueryMod

func (s queryModSliceTag) Apply(q *queries.Query) {
	qm.Apply(q, s...)
}

func buildQueryModFilterTag[T any](filterField TagField, filterOperation model.FilterOperation[T]) queryModSliceTag {
	var newQueryMod queryModSliceTag

	filterOperator := filterOperation.Operator

	switch filterOperator {
	case model.FilterEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" = ?", filterOperand.Operand))
	case model.FilterNEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterNEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" != ?", filterOperand.Operand))
	case model.FilterGreaterThan:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterGreaterThan operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" > ?", filterOperand.Operand))
	case model.FilterGreaterThanEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterGreaterThanEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" >= ?", filterOperand.Operand))
	case model.FilterLessThan:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterLessThan operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" < ?", filterOperand.Operand))
	case model.FilterLessThanEqual:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterLessThanEqual operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" <= ?", filterOperand.Operand))
	case model.FilterIn:
		filterOperand, ok := filterOperation.Operand.(model.ListOperand[T])
		if !ok {
			panic("expected a list operand for FilterIn operator")
		}

		newQueryMod = append(newQueryMod, qm.WhereIn(strcase.SnakeCase(string(filterField))+" IN (?)", filterOperand.Operands))
	case model.FilterNotIn:
		filterOperand, ok := filterOperation.Operand.(model.ListOperand[T])
		if !ok {
			panic("expected a list operand for FilterNotIn operator")
		}

		newQueryMod = append(newQueryMod, qm.WhereNotIn(strcase.SnakeCase(string(filterField))+" IN (?)", filterOperand.Operands))
	case model.FilterBetween:
		filterOperand, ok := filterOperation.Operand.(model.RangeOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterBetween operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
	case model.FilterNotBetween:
		filterOperand, ok := filterOperation.Operand.(model.RangeOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterNotBetween operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" NOT BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
	case model.FilterLike:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterLike operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" LIKE ?", filterOperand.Operand))
	case model.FilterNotLike:
		filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterLike operator")
		}

		newQueryMod = append(newQueryMod, qm.Where(strcase.SnakeCase(string(filterField))+" NOT LIKE ?", filterOperand.Operand))
	case model.FilterOr:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterOr operator")
		}
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS))))
	case model.FilterAnd:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[T])
		if !ok {
			panic("expected a scalar operand for FilterAnd operator")
		}

		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS)))
	default:
		panic("Unhandled FilterOperator")
	}

	return newQueryMod
}

func buildQueryModListFromFilterTag(filter *TagFilter) queryModSliceTag {
	queryModList := make(queryModSliceTag, 0, 5)

	if filter.Tag.HasValue {
		newQueryMod := buildQueryModFilterTag("Tag", filter.Tag.Wrappee)
		queryModList = append(queryModList, newQueryMod...)
	}
	if filter.Path.HasValue {
		newQueryMod := buildQueryModFilterTag("Path", filter.Path.Wrappee)
		queryModList = append(queryModList, newQueryMod...)
	}
	if filter.Children.HasValue {
		newQueryMod := buildQueryModFilterTag("Children", filter.Children.Wrappee)
		queryModList = append(queryModList, newQueryMod...)
	}
	if filter.ParentTag.HasValue {
		newQueryMod := buildQueryModFilterTag("ParentTag", filter.ParentTag.Wrappee)
		queryModList = append(queryModList, newQueryMod...)
	}
	if filter.ID.HasValue {
		newQueryMod := buildQueryModFilterTag("ID", filter.ID.Wrappee)
		queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}

type MssqlTagRepositoryConstructorArgs struct {
	DB *sql.DB
}

func (repo *MssqlTagRepository) New(args any) (newRepo repoCommon.TagRepository, err error) {
	constructorArgs, ok := args.(MssqlTagRepositoryConstructorArgs)
	if !ok {
		err = fmt.Errorf("expected type %T but got %T", MssqlTagRepositoryConstructorArgs{}, args)

		return
	}

	repo.db = constructorArgs.DB

	newRepo = repo

	return
}

//******************************************************************//
//                              Methods                             //
//******************************************************************//
func (repo *MssqlTagRepository) Add(ctx context.Context, domainModels []*domain.Tag) (err error) {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)

		err = helper.IneffectiveOperationError{Inner: helper.EmptyInputError}

		return
	}

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryModels []any
	repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return
	}

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Tag)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		err = repoModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				err = helper.DuplicateInsertionError{Inner: err}
			}

			return
		}
	}

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) Replace(ctx context.Context, domainModels []*domain.Tag) (err error) {

	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)

		err = helper.IneffectiveOperationError{Inner: helper.EmptyInputError}

		return
	}

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryModels []any
	repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return
	}

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Tag)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		var numAffectedRecords int64
		numAffectedRecords, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				err = helper.DuplicateInsertionError{Inner: err}
			}

			return
		}

		if numAffectedRecords == 0 {
			err = helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError}

			return
		}
	}

	tx.Commit()

	return
}
func (repo *MssqlTagRepository) Upsert(ctx context.Context, domainModels []*domain.Tag) (err error) {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)

		err = helper.IneffectiveOperationError{Inner: helper.EmptyInputError}

		return
	}

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryModels []any
	repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return
	}

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Tag)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		err = repoModel.Upsert(ctx, tx, boil.Infer(), boil.Infer())

		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				err = helper.DuplicateInsertionError{Inner: err}
			}

			return
		}
	}

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) Update(ctx context.Context, domainModels []*domain.Tag, domainColumnUpdater *domain.TagUpdater) (err error) {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)

		err = helper.IneffectiveOperationError{Inner: helper.EmptyInputError}

		return
	}

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	if domainColumnUpdater == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	if domainColumnUpdater == (&domain.TagUpdater{}) {
		err = helper.IneffectiveOperationError{Inner: helper.NopUpdaterError}
		log.Error(err)

		return
	}

	var repositoryModels []any
	repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return
	}

	var repositoryUpdater any
	repositoryUpdater, err = repo.TagDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return
	}

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	var numAffectedRecords int64
	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Tag)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		repoUpdater, ok := repositoryUpdater.(*TagUpdater)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		repoUpdater.ApplyToModel(repoModel)
		numAffectedRecords, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				err = helper.DuplicateInsertionError{Inner: err}
			}

			return
		}

		if numAffectedRecords == 0 {
			err = helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError}

			return
		}
	}

	err = tx.Commit()

	return
}

func (repo *MssqlTagRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.TagFilter, domainColumnUpdater *domain.TagUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate TagSlice

	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	if domainColumnUpdater == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	if domainColumnUpdater == (&domain.TagUpdater{}) {
		err = helper.IneffectiveOperationError{Inner: helper.NopUpdaterError}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	var repositoryUpdater any
	repositoryUpdater, err = repo.TagDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return
	}

	repoUpdater, ok := repositoryUpdater.(*TagUpdater)
	if !ok {
		err = fmt.Errorf("expected type *TagUpdater but got %T", repoUpdater)

		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	modelsToUpdate, err = Tags(queryFilters...).All(ctx, repo.db)
	if err != nil {
		return
	}

	if len(modelsToUpdate) == 0 {
		err = helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError}

		return
	}

	numAffectedRecords = int64(len(modelsToUpdate))

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repoModel := range modelsToUpdate {
		repoUpdater.ApplyToModel(repoModel)
		_, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				err = helper.DuplicateInsertionError{Inner: err}
			}

			return
		}

	}

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) Delete(ctx context.Context, domainModels []*domain.Tag) (err error) {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)

		err = helper.IneffectiveOperationError{Inner: helper.EmptyInputError}

		return
	}

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryModels []any
	repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return
	}

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Tag)
		if !ok {
			err = fmt.Errorf("expected type *Tag but got %T", repoModel)

			return
		}

		_, err = repoModel.Delete(ctx, tx)
		if err != nil {
			return
		}
	}

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numAffectedRecords int64, err error) {
	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Tags(queryFilters...).DeleteAll(ctx, tx)

	tx.Commit()

	return
}

func (repo *MssqlTagRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numRecords int64, err error) {
	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	return Tags(queryFilters...).Count(ctx, repo.db)
}

func (repo *MssqlTagRepository) CountAll(ctx context.Context) (numRecords int64, err error) {
	return Tags().Count(ctx, repo.db)
}

func (repo *MssqlTagRepository) DoesExist(ctx context.Context, domainModel *domain.Tag) (doesExist bool, err error) {
	if domainModel == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryModel any
	repositoryModel, err = repo.TagDomainToRepositoryModel(ctx, domainModel)
	if err != nil {
		return
	}

	repoModel, ok := repositoryModel.(*Tag)
	if !ok {
		err = fmt.Errorf("expected type *Tag but got %T", repoModel)

		return
	}

	return TagExists(ctx, repo.db, repoModel.ID)
}

func (repo *MssqlTagRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (doesExist bool, err error) {
	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	return Tags(queryFilters...).Exists(ctx, repo.db)
}

func (repo *MssqlTagRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (records []*domain.Tag, err error) {
	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	var repositoryModels TagSlice
	repositoryModels, err = Tags(queryFilters...).All(ctx, repo.db)
	if errors.Is(err, sql.ErrNoRows) {
		err = helper.IneffectiveOperationError{Inner: err}
	}

	records = make([]*domain.Tag, 0, len(repositoryModels))

	var domainModel *domain.Tag
	for _, repoModel := range repositoryModels {
		domainModel, err = repo.TagRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			return
		}

		records = append(records, domainModel)
	}

	return
}

func (repo *MssqlTagRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (record *domain.Tag, err error) {
	if domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
	}

	var repositoryFilter any
	repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*TagFilter)
	if !ok {
		err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

		return
	}

	queryFilters := buildQueryModListFromFilterTag(repoFilter)

	var repositoryModel *Tag
	repositoryModel, err = Tags(queryFilters...).One(ctx, repo.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = helper.IneffectiveOperationError{Inner: err}
		}

		return
	}

	record, err = repo.TagRepositoryToDomainModel(ctx, repositoryModel)

	return
}

func (repo *MssqlTagRepository) GetAll(ctx context.Context) (records []*domain.Tag, err error) {
	var repositoryModels TagSlice
	repositoryModels, err = Tags().All(ctx, repo.db)
	if err != nil {
		return
	}

	records = make([]*domain.Tag, 0, len(repositoryModels))

	var domainModel *domain.Tag
	for _, repoModel := range repositoryModels {
		domainModel, err = repo.TagRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = helper.IneffectiveOperationError{Inner: err}
			}

			return
		}

		records = append(records, domainModel)
	}

	return
}

//******************************************************************//
//                            Converters                            //
//******************************************************************//
func (repo *MssqlTagRepository) GetTagDomainToRepositoryModel(ctx context.Context) func(domainModel *domain.Tag) (repositoryModel any, err error) {
	return func(domainModel *domain.Tag) (repositoryModel any, err error) {
		return repo.TagDomainToRepositoryModel(ctx, domainModel)
	}
}

func (repo *MssqlTagRepository) GetTagRepositoryToDomainModel(ctx context.Context) func(repositoryModel any) (domainModel *domain.Tag, err error) {
	return func(repositoryModel any) (domainModel *domain.Tag, err error) {

		return repo.TagRepositoryToDomainModel(ctx, repositoryModel)
	}
}

//******************************************************************//
//                          Model Converter                         //
//******************************************************************//

func (repo *MssqlTagRepository) TagDomainToRepositoryModel(ctx context.Context, domainModel *domain.Tag) (repositoryModel any, err error) {

	// TODO: make sure to insert all tags in ParentPath and Subtags into db
	repositoryModelConcrete := new(Tag)
	repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()

	repositoryModelConcrete.ID = domainModel.ID
	repositoryModelConcrete.Tag = domainModel.Tag

	//***********************    Set ParentTag    **********************//
	if len(domainModel.ParentPath) > 0 {
		repositoryModelConcrete.ParentTag = null.NewInt64(domainModel.ParentPath[len(domainModel.ParentPath)-1].ID, true)
	}

	//*************************    Set Path    *************************//
	for _, tag := range domainModel.ParentPath {
		repositoryModelConcrete.Path += strconv.FormatInt(tag.ID, 10) + ";"
	}

	repositoryModelConcrete.Path += strconv.FormatInt(domainModel.ID, 10)

	//************************    Set Children  ************************//
	for _, tag := range domainModel.Subtags {
		repositoryModelConcrete.Children += strconv.FormatInt(tag.ID, 10) + ";"
	}

	repositoryModel = repositoryModelConcrete

	return
}

// TODO: These functions should be context aware
func (repo *MssqlTagRepository) TagRepositoryToDomainModel(ctx context.Context, repositoryModel any) (domainModel *domain.Tag, err error) {
	// TODO: make sure to insert all tags in ParentPath and Subtags into db
	domainModel = new(domain.Tag)

	repositoryModelConcrete := repositoryModel.(*Tag)

	domainModel.ID = repositoryModelConcrete.ID
	domainModel.Tag = repositoryModelConcrete.Tag

	//***********************    Set ParentPath    **********************//
	var parentTagID int64
	var parentTag *Tag
	var domainParentTag *domain.Tag

	for _, parentTagIDRaw := range strings.Split(repositoryModelConcrete.Path, ";")[:len(repositoryModelConcrete.Path)-2] {
		parentTagID, err = strconv.ParseInt(parentTagIDRaw, 10, 64)
		if err != nil {
			return
		}

		parentTag, err = Tags(TagWhere.ID.EQ(parentTagID)).One(ctx, repo.db)
		if err != nil {
			return
		}

		domainParentTag, err = repo.TagRepositoryToDomainModel(ctx, parentTag)
		if err != nil {
			return
		}

		domainModel.ParentPath = append(domainModel.ParentPath, domainParentTag)
	}

	//************************    Set Subtags ************************//
	var childTagID int64
	var childTag *Tag
	var domainChildTag *domain.Tag

	for _, childTagIDRaw := range strings.Split(repositoryModelConcrete.Children, ";")[:len(repositoryModelConcrete.Children)-2] {
		childTagID, err = strconv.ParseInt(childTagIDRaw, 10, 64)
		if err != nil {
			return
		}

		childTag, err = Tags(TagWhere.ID.EQ(childTagID)).One(ctx, repo.db)
		if err != nil {
			return
		}

		domainChildTag, err = repo.TagRepositoryToDomainModel(ctx, childTag)
		if err != nil {
			return
		}

		domainModel.Subtags = append(domainModel.Subtags, domainChildTag)
	}

	repositoryModel = repositoryModelConcrete

	return
}

//******************************************************************//
//                         Filter Converter                         //
//******************************************************************//

func (repo *MssqlTagRepository) TagDomainToRepositoryFilter(ctx context.Context, domainFilter *domain.TagFilter) (repositoryFilter any, err error) {
	repositoryFilterConcrete := new(TagFilter)

	repositoryFilterConcrete.ID = domainFilter.ID
	repositoryFilterConcrete.Tag = domainFilter.Tag

	if domainFilter.ParentPath.HasValue {

		//*********************    Set ParentPath    *********************//
		var convertedParentTagTagFilter model.FilterOperation[*Tag]

		convertedParentTagTagFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.ParentPath.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverterGeneric[domain.Tag, Tag](ctx, repo.TagDomainToRepositoryModel))
		if err != nil {
			return
		}

		repositoryFilterConcrete.ParentTagTag.Push(convertedParentTagTagFilter)
		//*************************    Set Path    *************************//
		var convertedPathFilter model.FilterOperation[string]

		convertedPathFilter, err = model.ConvertFilter[string, *domain.Tag](domainFilter.ParentPath.Wrappee, func(tag *domain.Tag) (string, error) { return strconv.FormatInt(tag.ID, 10), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.Path.Push(convertedPathFilter)
		//**********************    Set ParentTag    ***********************//
		var convertedParentTagFilter model.FilterOperation[null.Int64]

		convertedParentTagFilter, err = model.ConvertFilter[null.Int64, *domain.Tag](domainFilter.ParentPath.Wrappee, func(tag *domain.Tag) (null.Int64, error) { return null.NewInt64(tag.ID, true), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.ParentTag.Push(convertedParentTagFilter)
	}

	//**********************    Set child tags *********************//
	if domainFilter.Subtags.HasValue {
		var convertedFilter model.FilterOperation[string]

		convertedFilter, err = model.ConvertFilter[string, *domain.Tag](domainFilter.Subtags.Wrappee, func(tag *domain.Tag) (string, error) { return strconv.FormatInt(tag.ID, 10), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.Children.Push(convertedFilter)
	}

	repositoryFilter = repositoryFilterConcrete

	return
}

//******************************************************************//
//                         Updater Converter                        //
//******************************************************************//

func (repo *MssqlTagRepository) TagDomainToRepositoryUpdater(ctx context.Context, domainUpdater *domain.TagUpdater) (repositoryUpdater any, err error) {
	repositoryUpdaterConcrete := new(TagUpdater)

	//**************************    Set tag    *************************//
	if domainUpdater.Tag.HasValue {
		repositoryUpdaterConcrete.Tag.Push(model.UpdateOperation[string]{Operator: domainUpdater.Tag.Wrappee.Operator, Operand: repositoryUpdaterConcrete.Tag.Wrappee.Operand})
	}

	//***********    Set ParentPath    ***********//
	if domainUpdater.ParentPath.HasValue {
		var convertedTagRaw any
		tag := domainUpdater.ParentPath.Wrappee.Operand[len(domainUpdater.ParentPath.Wrappee.Operand)-1]
		convertedTagRaw, err = repo.TagDomainToRepositoryModel(ctx, tag)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.ParentTagTag.Push(model.UpdateOperation[*Tag]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: convertedTagRaw.(*Tag)})
		repositoryUpdaterConcrete.ParentTag.Push(model.UpdateOperation[null.Int64]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: null.NewInt64(convertedTagRaw.(*Tag).ID, true)})

		pathIDs := make([]string, 0, len(domainUpdater.ParentPath.Wrappee.Operand)+1)
		for _, tag := range domainUpdater.ParentPath.Wrappee.Operand {
			pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))
		}

		pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))

		repositoryUpdaterConcrete.Path.Push(model.UpdateOperation[string]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: strings.Join(pathIDs, ";")})
	}

	//***********************    Set Children    ***********************//
	if domainUpdater.Subtags.HasValue {
		pathIDs := make([]string, 0, len(domainUpdater.Subtags.Wrappee.Operand)+1)
		for _, tag := range domainUpdater.Subtags.Wrappee.Operand {
			pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))
		}

		repositoryUpdaterConcrete.Children.Push(model.UpdateOperation[string]{Operator: domainUpdater.Subtags.Wrappee.Operator, Operand: strings.Join(pathIDs, ";")})
	}

	//**************************    Set ID    **************************//
	if domainUpdater.ID.HasValue {
		repositoryUpdaterConcrete.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: repositoryUpdaterConcrete.ID.Wrappee.Operand})
	}

	repositoryUpdater = repositoryUpdaterConcrete

	return
}
