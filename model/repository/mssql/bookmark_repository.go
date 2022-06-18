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
	"fmt"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repoCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"time"
)

//******************************************************************//
//                        Types and constants                       //
//******************************************************************//
type MssqlBookmarkRepository struct {
	db *sql.DB

	tagRepository repoCommon.TagRepository
}

type BookmarkField string

var BookmarkFields = struct {
	ID             BookmarkField
	IsRead         BookmarkField
	Title          BookmarkField
	URL            BookmarkField
	BookmarkTypeID BookmarkField
	IsCollection   BookmarkField
	CreatedAt      BookmarkField
	UpdatedAt      BookmarkField
	DeletedAt      BookmarkField
}{
	ID:             "id",
	IsRead:         "is_read",
	Title:          "title",
	URL:            "url",
	BookmarkTypeID: "bookmark_type_id",
	IsCollection:   "is_collection",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

var BookmarkFieldsList = []BookmarkField{
	BookmarkField("ID"),
	BookmarkField("IsRead"),
	BookmarkField("Title"),
	BookmarkField("URL"),
	BookmarkField("BookmarkTypeID"),
	BookmarkField("IsCollection"),
	BookmarkField("CreatedAt"),
	BookmarkField("UpdatedAt"),
	BookmarkField("DeletedAt"),
}

var BookmarkRelationsList = []string{
	"BookmarkType",
	"Tags",
}

type BookmarkFilter struct {
	ID             optional.Optional[model.FilterOperation[int64]]
	IsRead         optional.Optional[model.FilterOperation[int64]]
	Title          optional.Optional[model.FilterOperation[null.String]]
	URL            optional.Optional[model.FilterOperation[string]]
	BookmarkTypeID optional.Optional[model.FilterOperation[null.Int64]]
	IsCollection   optional.Optional[model.FilterOperation[int64]]
	CreatedAt      optional.Optional[model.FilterOperation[time.Time]]
	UpdatedAt      optional.Optional[model.FilterOperation[time.Time]]
	DeletedAt      optional.Optional[model.FilterOperation[null.Time]]

	BookmarkType optional.Optional[model.FilterOperation[*BookmarkType]]
	Tags         optional.Optional[model.FilterOperation[*Tag]]
}

type BookmarkFilterMapping[T any] struct {
	Field           BookmarkField
	FilterOperation model.FilterOperation[T]
}

func (filter *BookmarkFilter) GetSetFilters() *list.List {
	setFilters := list.New()

	if filter.ID.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.ID, FilterOperation: filter.ID.Wrappee})
	}
	if filter.IsRead.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.IsRead, FilterOperation: filter.IsRead.Wrappee})
	}
	if filter.Title.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[null.String]{Field: BookmarkFields.Title, FilterOperation: filter.Title.Wrappee})
	}
	if filter.URL.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[string]{Field: BookmarkFields.URL, FilterOperation: filter.URL.Wrappee})
	}
	if filter.BookmarkTypeID.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[null.Int64]{Field: BookmarkFields.BookmarkTypeID, FilterOperation: filter.BookmarkTypeID.Wrappee})
	}
	if filter.IsCollection.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.IsCollection, FilterOperation: filter.IsCollection.Wrappee})
	}
	if filter.CreatedAt.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[time.Time]{Field: BookmarkFields.CreatedAt, FilterOperation: filter.CreatedAt.Wrappee})
	}
	if filter.UpdatedAt.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[time.Time]{Field: BookmarkFields.UpdatedAt, FilterOperation: filter.UpdatedAt.Wrappee})
	}
	if filter.DeletedAt.HasValue {
		setFilters.PushBack(BookmarkFilterMapping[null.Time]{Field: BookmarkFields.DeletedAt, FilterOperation: filter.DeletedAt.Wrappee})
	}

	return setFilters
}

type BookmarkUpdater struct {
	CreatedAt      optional.Optional[model.UpdateOperation[time.Time]]
	UpdatedAt      optional.Optional[model.UpdateOperation[time.Time]]
	BookmarkType   optional.Optional[model.UpdateOperation[*BookmarkType]]
	DeletedAt      optional.Optional[model.UpdateOperation[null.Time]]
	URL            optional.Optional[model.UpdateOperation[string]]
	Tags           optional.Optional[model.UpdateOperation[TagSlice]]
	Title          optional.Optional[model.UpdateOperation[null.String]]
	BookmarkTypeID optional.Optional[model.UpdateOperation[null.Int64]]
	IsCollection   optional.Optional[model.UpdateOperation[int64]]
	ID             optional.Optional[model.UpdateOperation[int64]]
	IsRead         optional.Optional[model.UpdateOperation[int64]]
}

type BookmarkUpdaterMapping[T any] struct {
	Field   BookmarkField
	Updater model.UpdateOperation[T]
}

func (updater *BookmarkUpdater) GetSetUpdaters() *list.List {
	setUpdaters := list.New()

	if updater.ID.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.ID, Updater: updater.ID.Wrappee})
	}
	if updater.IsRead.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.IsRead, Updater: updater.IsRead.Wrappee})
	}
	if updater.Title.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[null.String]{Field: BookmarkFields.Title, Updater: updater.Title.Wrappee})
	}
	if updater.URL.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[string]{Field: BookmarkFields.URL, Updater: updater.URL.Wrappee})
	}
	if updater.BookmarkTypeID.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[null.Int64]{Field: BookmarkFields.BookmarkTypeID, Updater: updater.BookmarkTypeID.Wrappee})
	}
	if updater.IsCollection.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.IsCollection, Updater: updater.IsCollection.Wrappee})
	}
	if updater.CreatedAt.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[time.Time]{Field: BookmarkFields.CreatedAt, Updater: updater.CreatedAt.Wrappee})
	}
	if updater.UpdatedAt.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[time.Time]{Field: BookmarkFields.UpdatedAt, Updater: updater.UpdatedAt.Wrappee})
	}
	if updater.DeletedAt.HasValue {
		setUpdaters.PushBack(BookmarkUpdaterMapping[null.Time]{Field: BookmarkFields.DeletedAt, Updater: updater.DeletedAt.Wrappee})
	}

	return setUpdaters
}

func (updater *BookmarkUpdater) ApplyToModel(bookmarkModel *Bookmark) {
	if updater.ID.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).ID, updater.ID.Wrappee)
	}
	if updater.IsRead.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).IsRead, updater.IsRead.Wrappee)
	}
	if updater.Title.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).Title, updater.Title.Wrappee)
	}
	if updater.URL.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).URL, updater.URL.Wrappee)
	}
	if updater.BookmarkTypeID.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).BookmarkTypeID, updater.BookmarkTypeID.Wrappee)
	}
	if updater.IsCollection.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).IsCollection, updater.IsCollection.Wrappee)
	}
	if updater.CreatedAt.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).CreatedAt, updater.CreatedAt.Wrappee)
	}
	if updater.UpdatedAt.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).UpdatedAt, updater.UpdatedAt.Wrappee)
	}
	if updater.DeletedAt.HasValue {
		model.ApplyUpdater(&(*bookmarkModel).DeletedAt, updater.DeletedAt.Wrappee)
	}

}

type queryModSliceBookmark []qm.QueryMod

func (s queryModSliceBookmark) Apply(q *queries.Query) {
	qm.Apply(q, s...)
}

func buildQueryModFilterBookmark[T any](filterField BookmarkField, filterOperation model.FilterOperation[T]) queryModSliceBookmark {
	var newQueryMod queryModSliceBookmark

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
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterBookmark(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterBookmark(filterField, filterOperand.RHS))))
	case model.FilterAnd:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterAnd operator")
		}

		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterBookmark(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterBookmark(filterField, filterOperand.RHS)))
	default:
		panic("Unhandled FilterOperator")
	}

	return newQueryMod
}

func buildQueryModListFromFilterBookmark(setFilters list.List) queryModSliceBookmark {
	queryModList := make(queryModSliceBookmark, 0, 9)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(BookmarkFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %T but got %T", BookmarkFilterMapping[any]{}, filter))
		}

		newQueryMod := buildQueryModFilterBookmark(filterMapping.Field, filterMapping.FilterOperation)

		queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}

type MssqlBookmarkRepositoryConstructorArgs struct {
	DB *sql.DB

	TagRepository repoCommon.TagRepository
}

func (repo *MssqlBookmarkRepository) New(args any) (newRepo repoCommon.BookmarkRepository, err error) {
	constructorArgs, ok := args.(MssqlBookmarkRepositoryConstructorArgs)
	if !ok {
		err = fmt.Errorf("expected type %T but got %T", MssqlBookmarkRepositoryConstructorArgs{}, args)

		return
	}

	repo.db = constructorArgs.DB

	repo.tagRepository = constructorArgs.TagRepository

	newRepo = repo

	return
}

//******************************************************************//
//                              Methods                             //
//******************************************************************//
func (repo *MssqlBookmarkRepository) Add(ctx context.Context, domainModels []*domain.Bookmark) error {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)
		return nil
	}

	err := goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Bookmark](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetBookmarkDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Bookmark)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		err = repoModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlBookmarkRepository) Replace(ctx context.Context, domainModels []*domain.Bookmark) error {

	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetBookmarkDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Bookmark)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		_, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}
func (repo *MssqlBookmarkRepository) Upsert(ctx context.Context, domainModels []*domain.Bookmark) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetBookmarkDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Bookmark)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		err = repoModel.Upsert(ctx, tx, boil.Infer(), boil.Infer())

		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlBookmarkRepository) Update(ctx context.Context, domainModels []*domain.Bookmark, domainColumnUpdater *domain.BookmarkUpdater) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetBookmarkDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	repositoryUpdater, err := repo.BookmarkDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Bookmark)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		repoUpdater, ok := repositoryUpdater.(*BookmarkUpdater)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		repoUpdater.ApplyToModel(repoModel)
		repoModel.Update(ctx, tx, boil.Infer())
	}

	tx.Commit()

	return err
}

func (repo *MssqlBookmarkRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter, domainColumnUpdater *domain.BookmarkUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate BookmarkSlice

	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repositoryUpdater, err := repo.BookmarkDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return
	}

	repoUpdater, ok := repositoryUpdater.(*BookmarkUpdater)
	if !ok {
		err = fmt.Errorf("Expected type *BookmarkUpdater but got %T", repoUpdater)

		return
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		err = fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)

		return
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	modelsToUpdate, err = Bookmarks(queryFilters...).All(ctx, repo.db)
	if err != nil {
		return
	}

	numAffectedRecords = int64(len(modelsToUpdate))

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	for _, repoModel := range modelsToUpdate {
		repoUpdater.ApplyToModel(repoModel)
		repoModel.Update(ctx, tx, boil.Infer())
	}

	tx.Commit()

	return
}

func (repo *MssqlBookmarkRepository) Delete(ctx context.Context, domainModels []*domain.Bookmark) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetBookmarkDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Bookmark)
		if !ok {
			return fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
		}

		_, err = repoModel.Delete(ctx, tx)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlBookmarkRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (numAffectedRecords int64, err error) {
	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		err = fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)

		return
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Bookmarks(queryFilters...).DeleteAll(ctx, tx)

	tx.Commit()

	return
}

func (repo *MssqlBookmarkRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (int64, error) {
	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return 0, err
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		return 0, fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)

	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Count(ctx, repo.db)
}

func (repo *MssqlBookmarkRepository) CountAll(ctx context.Context) (int64, error) {
	return Bookmarks().Count(ctx, repo.db)
}

func (repo *MssqlBookmarkRepository) DoesExist(ctx context.Context, domainModel *domain.Bookmark) (bool, error) {
	repositoryModel, err := repo.BookmarkDomainToRepositoryModel(ctx, domainModel)
	if err != nil {
		return false, err
	}

	repoModel, ok := repositoryModel.(*Bookmark)
	if !ok {
		return false, fmt.Errorf("Expected type *Bookmark but got %T", repoModel)
	}

	return BookmarkExists(ctx, repo.db, repoModel.ID)
}

func (repo *MssqlBookmarkRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (bool, error) {
	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return false, err
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		return false, fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Exists(ctx, repo.db)
}

func (repo *MssqlBookmarkRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) ([]*domain.Bookmark, error) {
	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return []*domain.Bookmark{}, err
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		return []*domain.Bookmark{}, fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	repositoryModels, err := Bookmarks(queryFilters...).All(ctx, repo.db)

	domainModels := make([]*domain.Bookmark, 0, len(repositoryModels))

	for _, repoModel := range repositoryModels {
		domainModel, err := repo.BookmarkRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			return domainModels, err
		}

		domainModels = append(domainModels, domainModel)
	}

	return domainModels, err
}

func (repo *MssqlBookmarkRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (*domain.Bookmark, error) {
	repositoryFilter, err := repo.BookmarkDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return nil, err
	}

	repoFilter, ok := repositoryFilter.(*BookmarkFilter)
	if !ok {
		return nil, fmt.Errorf("Expected type *BookmarkFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	repositoryModel, err := Bookmarks(queryFilters...).One(ctx, repo.db)

	var domainModel *domain.Bookmark
	if err != nil {
		return domainModel, err
	}

	domainModel, err = repo.BookmarkRepositoryToDomainModel(ctx, repositoryModel)

	return domainModel, err
}

func (repo *MssqlBookmarkRepository) GetAll(ctx context.Context) ([]*domain.Bookmark, error) {
	repositoryModels, err := Bookmarks().All(ctx, repo.db)
	if err != nil {
		return []*domain.Bookmark{}, err
	}

	domainModels := make([]*domain.Bookmark, 0, len(repositoryModels))

	for _, repoModel := range repositoryModels {
		domainModel, err := repo.BookmarkRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			return domainModels, err
		}

		domainModels = append(domainModels, domainModel)
	}

	return domainModels, err
}

func (repo *MssqlBookmarkRepository) AddType(ctx context.Context, types []string) error {
	for _, type_ := range types {
		repositoryModel := BookmarkType{BookmarkType: type_}

		err := repositoryModel.Insert(ctx, repo.db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *MssqlBookmarkRepository) DeleteType(ctx context.Context, types []string) error {
	_, err := BookmarkTypes(BookmarkTypeWhere.BookmarkType.IN(types)).DeleteAll(ctx, repo.db)

	return err
}

func (repo *MssqlBookmarkRepository) UpdateType(ctx context.Context, oldType string, newType string) error {
	repositoryModel, err := BookmarkTypes(BookmarkTypeWhere.BookmarkType.EQ(oldType)).One(ctx, repo.db)
	if err != nil {
		return err
	}

	repositoryModel.BookmarkType = newType

	_, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

	return err
}

func (repo *MssqlBookmarkRepository) GetTagRepository() repoCommon.TagRepository {
	return repo.tagRepository
}

//******************************************************************//
//                            Converters                            //
//******************************************************************//
func (repo *MssqlBookmarkRepository) GetBookmarkDomainToRepositoryModel(ctx context.Context) func(domainModel *domain.Bookmark) (repositoryModel any, err error) {
	return func(domainModel *domain.Bookmark) (repositoryModel any, err error) {
		return repo.BookmarkDomainToRepositoryModel(ctx, domainModel)
	}
}

func (repo *MssqlBookmarkRepository) GetBookmarkRepositoryToDomainModel(ctx context.Context) func(repositoryModel any) (domainModel *domain.Bookmark, err error) {
	return func(repositoryModel any) (domainModel *domain.Bookmark, err error) {

		return repo.BookmarkRepositoryToDomainModel(ctx, repositoryModel)
	}
}

//******************************************************************//
//                          Model Converter                         //
//******************************************************************//

func (repo *MssqlBookmarkRepository) BookmarkDomainToRepositoryModel(ctx context.Context, domainModel *domain.Bookmark) (repositoryModel any, err error) {
	repositoryModelConcrete := new(Bookmark)
	repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()

	repositoryModelConcrete.URL = domainModel.URL
	repositoryModelConcrete.ID = domainModel.ID

	//**********************    Set Timestamps    **********************//

	repositoryModelConcrete.CreatedAt = domainModel.CreatedAt
	repositoryModelConcrete.UpdatedAt = domainModel.UpdatedAt

	if domainModel.DeletedAt.HasValue {
		var convertedTime null.Time
		convertedTime, err = repoCommon.OptionalTimeToNullTime(domainModel.DeletedAt)
		if err != nil {
			return
		}

		repositoryModelConcrete.DeletedAt = convertedTime
	}

	//*************************    Set Title    ************************//
	if domainModel.Title.HasValue {
		repositoryModelConcrete.Title.Valid = true
		repositoryModelConcrete.Title.String = domainModel.Title.Wrappee
	}

	//******************    Set IsRead/IsCollection    *****************//
	if domainModel.IsRead {
		repositoryModelConcrete.IsRead = 1
	}

	if domainModel.IsCollection {
		repositoryModelConcrete.IsCollection = 1
	}

	//*************************    Set Tags    *************************//
	var repositoryTag *Tag

	if domainModel.Tags != nil {
		repositoryModelConcrete.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
		for _, domainTag := range domainModel.Tags {

			repositoryTag, err = Tags(TagWhere.Tag.EQ(domainTag.Tag)).One(ctx, repo.db)
			if err != nil {
				return
			}

			repositoryModelConcrete.R.Tags = append(repositoryModelConcrete.R.Tags, &Tag{Tag: repositoryTag.Tag, ID: repositoryTag.ID})
		}
	}

	//*************************    Set Type    *************************//
	if domainModel.BookmarkType.HasValue {
		var repositoryBookmarkType *BookmarkType

		repositoryModelConcrete.R.BookmarkType = &BookmarkType{BookmarkType: domainModel.BookmarkType.Wrappee}
		repositoryBookmarkType, err = BookmarkTypes(BookmarkTypeWhere.BookmarkType.EQ(domainModel.BookmarkType.Wrappee)).One(ctx, repo.db)
		if err != nil {
			return
		}

		if repositoryBookmarkType != nil {
			repositoryModelConcrete.BookmarkTypeID = null.NewInt64(repositoryBookmarkType.ID, true)
			repositoryModelConcrete.R.BookmarkType.ID = repositoryBookmarkType.ID
		} else {
			repositoryModelConcrete.R.BookmarkType = nil
		}
	}

	repositoryModel = repositoryModelConcrete

	return
}

func (repo *MssqlBookmarkRepository) BookmarkRepositoryToDomainModel(ctx context.Context, repositoryModel any) (domainModel *domain.Bookmark, err error) {
	domainModel = new(domain.Bookmark)

	repositoryModelConcrete := repositoryModel.(Bookmark)

	domainModel.URL = repositoryModelConcrete.URL
	domainModel.ID = repositoryModelConcrete.ID

	if repositoryModelConcrete.R == nil {
		repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()
	}

	if repositoryModelConcrete.R.BookmarkType != nil {
		domainModel.BookmarkType = optional.Make(repositoryModelConcrete.R.BookmarkType.BookmarkType)
	}

	//**********************    Set Timestamps    **********************//

	domainModel.CreatedAt = repositoryModelConcrete.CreatedAt
	domainModel.UpdatedAt = repositoryModelConcrete.UpdatedAt

	if repositoryModelConcrete.DeletedAt.Valid {
		domainModel.DeletedAt.Push(repositoryModelConcrete.DeletedAt.Time)
	}

	//*************************    Set Title    ************************//
	if repositoryModelConcrete.Title.Valid {
		domainModel.Title.Push(repositoryModelConcrete.Title.String)
	}

	//******************    Set IsRead/IsCollection    *****************//
	domainModel.IsRead = repositoryModelConcrete.IsRead > 0
	domainModel.IsCollection = repositoryModelConcrete.IsCollection > 0

	//*************************    Set Tags    *************************//
	var domainTag *domain.Tag

	domainModel.Tags = make([]*domain.Tag, 0, len(repositoryModelConcrete.R.Tags))
	for _, repositoryTag := range repositoryModelConcrete.R.Tags {
		domainTag, err = repo.GetTagRepository().TagRepositoryToDomainModel(ctx, repositoryTag)
		if err != nil {
			return
		}

		domainModel.Tags = append(domainModel.Tags, domainTag)
	}

	return
}

//******************************************************************//
//                         Filter Converter                         //
//******************************************************************//

func (repo *MssqlBookmarkRepository) BookmarkDomainToRepositoryFilter(ctx context.Context, domainFilter *domain.BookmarkFilter) (repositoryFilter any, err error) {
	repositoryFilterConcrete := new(BookmarkFilter)

	repositoryFilterConcrete.URL = domainFilter.URL
	repositoryFilterConcrete.ID = domainFilter.ID

	//**********************    Set Timestamps    **********************//

	repositoryFilterConcrete.CreatedAt = domainFilter.CreatedAt
	repositoryFilterConcrete.UpdatedAt = domainFilter.UpdatedAt

	if domainFilter.DeletedAt.HasValue {
		var convertedFilter model.FilterOperation[null.Time]

		convertedFilter, err = model.ConvertFilter[null.Time, optional.Optional[time.Time]](domainFilter.DeletedAt.Wrappee, repoCommon.OptionalTimeToNullTime)
		if err != nil {
			return
		}

		repositoryFilterConcrete.DeletedAt.Push(convertedFilter)
	}

	//*************************    Set Title    ************************//
	if domainFilter.Title.HasValue {
		var convertedFilter model.FilterOperation[null.String]

		convertedFilter, err = model.ConvertFilter[null.String, optional.Optional[string]](domainFilter.Title.Wrappee, repoCommon.OptionalStringToNullString)
		if err != nil {
			return
		}

		repositoryFilterConcrete.Title.Push(convertedFilter)
	}

	//******************    Set IsRead/IsCollection    *****************//
	if domainFilter.IsRead.HasValue {
		var convertedFilter model.FilterOperation[int64]

		convertedFilter, err = model.ConvertFilter[int64, bool](domainFilter.IsRead.Wrappee, repoCommon.BoolToInt)
		if err != nil {
			return
		}

		repositoryFilterConcrete.IsRead.Push(convertedFilter)
	}

	if domainFilter.IsCollection.HasValue {
		var convertedFilter model.FilterOperation[int64]

		convertedFilter, err = model.ConvertFilter[int64, bool](domainFilter.IsCollection.Wrappee, repoCommon.BoolToInt)
		if err != nil {
			return
		}

		repositoryFilterConcrete.IsCollection.Push(convertedFilter)
	}

	//*************************    Set Tags    *************************//

	if domainFilter.Tags.HasValue {
		var convertedFilter model.FilterOperation[*Tag]

		convertedFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.Tags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverterGeneric[domain.Tag, Tag](ctx, repo.GetTagRepository().TagDomainToRepositoryModel))
		if err != nil {
			return
		}

		repositoryFilterConcrete.Tags.Push(convertedFilter)
	}

	//*************************    Set Type    *************************//

	if domainFilter.BookmarkType.HasValue {
		var convertedTypeIDFilter model.FilterOperation[null.Int64]
		var convertedTypeFilter model.FilterOperation[*BookmarkType]

		convertedTypeFilter, err = model.ConvertFilter[*BookmarkType, optional.Optional[string]](domainFilter.BookmarkType.Wrappee, func(type_ optional.Optional[string]) (*BookmarkType, error) {
			if !type_.HasValue {
				return nil, nil
			}

			bookmarkType, err := BookmarkTypes(BookmarkTypeWhere.BookmarkType.EQ(type_.Wrappee)).One(ctx, repo.db)

			return bookmarkType, err
		})
		if err != nil {
			return
		}

		convertedTypeIDFilter, err = model.ConvertFilter[null.Int64, optional.Optional[string]](domainFilter.BookmarkType.Wrappee, func(type_ optional.Optional[string]) (null.Int64, error) {
			if !type_.HasValue {
				return null.NewInt64(-1, false), nil
			}

			bookmarkType, err := BookmarkTypes(BookmarkTypeWhere.BookmarkType.EQ(type_.Wrappee)).One(ctx, repo.db)

			return null.NewInt64(bookmarkType.ID, true), err
		})
		if err != nil {
			return
		}

		repositoryFilterConcrete.BookmarkType.Push(convertedTypeFilter)
		repositoryFilterConcrete.BookmarkTypeID.Push(convertedTypeIDFilter)
	}

	repositoryFilter = repositoryFilterConcrete

	return
}

//******************************************************************//
//                         Updater Converter                        //
//******************************************************************//

func (repo *MssqlBookmarkRepository) BookmarkDomainToRepositoryUpdater(ctx context.Context, domainUpdater *domain.BookmarkUpdater) (repositoryUpdater any, err error) {
	repositoryUpdaterConcrete := new(BookmarkUpdater)

	if domainUpdater.CreatedAt.HasValue {

		repositoryUpdaterConcrete.CreatedAt.Push(model.UpdateOperation[time.Time]{Operator: domainUpdater.CreatedAt.Wrappee.Operator, Operand: domainUpdater.CreatedAt.Wrappee.Operand})

	}

	if domainUpdater.UpdatedAt.HasValue {

		repositoryUpdaterConcrete.UpdatedAt.Push(model.UpdateOperation[time.Time]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: domainUpdater.UpdatedAt.Wrappee.Operand})

	}

	if domainUpdater.DeletedAt.HasValue {

		var convertedTime null.Time
		convertedTime, err = repoCommon.OptionalTimeToNullTime(domainUpdater.DeletedAt.Wrappee.Operand)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.DeletedAt.Push(model.UpdateOperation[null.Time]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: convertedTime})

	}

	if domainUpdater.URL.HasValue {
		repositoryUpdaterConcrete.URL.Push(model.UpdateOperation[string]{Operator: domainUpdater.URL.Wrappee.Operator, Operand: repositoryUpdaterConcrete.URL.Wrappee.Operand})
	}

	if domainUpdater.Title.HasValue {
		var convertedUpdater null.String
		convertedUpdater, err = repoCommon.OptionalStringToNullString(domainUpdater.Title.Wrappee.Operand)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.Title.Push(model.UpdateOperation[null.String]{Operator: domainUpdater.Title.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.Tags.HasValue {
		var rawTag any
		convertedUpdater := make(TagSlice, 0, len(domainUpdater.Tags.Wrappee.Operand))

		for _, tag := range domainUpdater.Tags.Wrappee.Operand {
			rawTag, err = repo.GetTagRepository().TagDomainToRepositoryModel(ctx, tag)
			if err != nil {
				return
			}

			convertedUpdater = append(convertedUpdater, rawTag.(*Tag))
		}

		repositoryUpdaterConcrete.Tags.Push(model.UpdateOperation[TagSlice]{Operator: domainUpdater.Tags.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.ID.HasValue {
		repositoryUpdaterConcrete.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: repositoryUpdaterConcrete.ID.Wrappee.Operand})
	}

	if domainUpdater.IsCollection.HasValue {
		var convertedUpdater int64
		convertedUpdater, err = repoCommon.BoolToInt(domainUpdater.IsCollection.Wrappee.Operand)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.IsCollection.Push(model.UpdateOperation[int64]{Operator: domainUpdater.IsCollection.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.IsRead.HasValue {
		var convertedUpdater int64
		convertedUpdater, err = repoCommon.BoolToInt(domainUpdater.IsRead.Wrappee.Operand)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.IsRead.Push(model.UpdateOperation[int64]{Operator: domainUpdater.IsRead.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.BookmarkType.HasValue {
		var convertedBookmarkType *BookmarkType
		if domainUpdater.BookmarkType.Wrappee.Operand.HasValue {
			convertedBookmarkType, err = BookmarkTypes(BookmarkTypeWhere.BookmarkType.EQ(domainUpdater.BookmarkType.Wrappee.Operand.Wrappee)).One(context.Background(), repo.db)
			if err != nil {
				return
			}
		} else {
			convertedBookmarkType = nil
		}

		repositoryUpdaterConcrete.BookmarkType.Push(model.UpdateOperation[*BookmarkType]{Operator: domainUpdater.BookmarkType.Wrappee.Operator, Operand: convertedBookmarkType})
	}

	repositoryUpdater = repositoryUpdaterConcrete

	return

}
