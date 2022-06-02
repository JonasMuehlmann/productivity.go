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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/bookmark_repository.go.tpl

package repository

import (
    "database/sql"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
    "context"
	"fmt"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/null/v8"
	"container/list"
)

type PsqlBookmarkRepository struct {
    db *sql.DB
}
type BookmarkField string

var BookmarkFields = struct {
    CreatedAt  BookmarkField
    UpdatedAt  BookmarkField
    URL  BookmarkField
    Title  BookmarkField
    DeletedAt  BookmarkField
    BookmarkTypeID  BookmarkField
    IsCollection  BookmarkField
    ID  BookmarkField
    IsRead  BookmarkField
    
}{
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
    URL: "url",
    Title: "title",
    DeletedAt: "deleted_at",
    BookmarkTypeID: "bookmark_type_id",
    IsCollection: "is_collection",
    ID: "id",
    IsRead: "is_read",
    
}

var BookmarkFieldsList = []BookmarkField{
    BookmarkField("CreatedAt"),
    BookmarkField("UpdatedAt"),
    BookmarkField("URL"),
    BookmarkField("Title"),
    BookmarkField("DeletedAt"),
    BookmarkField("BookmarkTypeID"),
    BookmarkField("IsCollection"),
    BookmarkField("ID"),
    BookmarkField("IsRead"),
    
}

var BookmarkRelationsList = []string{
    "BookmarkType",
    "Tags",
    
}

type BookmarkFilter struct {
    CreatedAt optional.Optional[model.FilterOperation[string]]
    UpdatedAt optional.Optional[model.FilterOperation[string]]
    URL optional.Optional[model.FilterOperation[string]]
    Title optional.Optional[model.FilterOperation[null.String]]
    DeletedAt optional.Optional[model.FilterOperation[null.String]]
    BookmarkTypeID optional.Optional[model.FilterOperation[null.Int64]]
    IsCollection optional.Optional[model.FilterOperation[int64]]
    ID optional.Optional[model.FilterOperation[int64]]
    IsRead optional.Optional[model.FilterOperation[int64]]
    
    BookmarkType optional.Optional[model.FilterOperation[*BookmarkType]]
    Tags optional.Optional[model.FilterOperation[*Tag]]
    
}

type BookmarkFilterMapping[T any] struct {
    Field BookmarkField
    FilterOperation model.FilterOperation[T]
}

func (filter *BookmarkFilter) GetSetFilters() *list.List {
    setFilters := list.New()

    if filter.CreatedAt.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[string]{Field: BookmarkFields.CreatedAt, FilterOperation: filter.CreatedAt.Wrappee})
    }
    if filter.UpdatedAt.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[string]{Field: BookmarkFields.UpdatedAt, FilterOperation: filter.UpdatedAt.Wrappee})
    }
    if filter.URL.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[string]{Field: BookmarkFields.URL, FilterOperation: filter.URL.Wrappee})
    }
    if filter.Title.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[null.String]{Field: BookmarkFields.Title, FilterOperation: filter.Title.Wrappee})
    }
    if filter.DeletedAt.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[null.String]{Field: BookmarkFields.DeletedAt, FilterOperation: filter.DeletedAt.Wrappee})
    }
    if filter.BookmarkTypeID.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[null.Int64]{Field: BookmarkFields.BookmarkTypeID, FilterOperation: filter.BookmarkTypeID.Wrappee})
    }
    if filter.IsCollection.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.IsCollection, FilterOperation: filter.IsCollection.Wrappee})
    }
    if filter.ID.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.ID, FilterOperation: filter.ID.Wrappee})
    }
    if filter.IsRead.HasValue {
    setFilters.PushBack(BookmarkFilterMapping[int64]{Field: BookmarkFields.IsRead, FilterOperation: filter.IsRead.Wrappee})
    }
    

    return setFilters
}

type BookmarkUpdater struct {
    CreatedAt optional.Optional[model.UpdateOperation[string]]
    UpdatedAt optional.Optional[model.UpdateOperation[string]]
    URL optional.Optional[model.UpdateOperation[string]]
    Title optional.Optional[model.UpdateOperation[null.String]]
    DeletedAt optional.Optional[model.UpdateOperation[null.String]]
    BookmarkTypeID optional.Optional[model.UpdateOperation[null.Int64]]
    IsCollection optional.Optional[model.UpdateOperation[int64]]
    ID optional.Optional[model.UpdateOperation[int64]]
    IsRead optional.Optional[model.UpdateOperation[int64]]
    
    BookmarkType optional.Optional[model.UpdateOperation[*BookmarkType]]
    Tags optional.Optional[model.UpdateOperation[TagSlice]]
    
}

type BookmarkUpdaterMapping[T any] struct {
    Field BookmarkField
    Updater model.UpdateOperation[T]
}

func (updater *BookmarkUpdater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    if updater.CreatedAt.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[string]{Field: BookmarkFields.CreatedAt, Updater: updater.CreatedAt.Wrappee})
    }
    if updater.UpdatedAt.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[string]{Field: BookmarkFields.UpdatedAt, Updater: updater.UpdatedAt.Wrappee})
    }
    if updater.URL.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[string]{Field: BookmarkFields.URL, Updater: updater.URL.Wrappee})
    }
    if updater.Title.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[null.String]{Field: BookmarkFields.Title, Updater: updater.Title.Wrappee})
    }
    if updater.DeletedAt.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[null.String]{Field: BookmarkFields.DeletedAt, Updater: updater.DeletedAt.Wrappee})
    }
    if updater.BookmarkTypeID.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[null.Int64]{Field: BookmarkFields.BookmarkTypeID, Updater: updater.BookmarkTypeID.Wrappee})
    }
    if updater.IsCollection.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.IsCollection, Updater: updater.IsCollection.Wrappee})
    }
    if updater.ID.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.ID, Updater: updater.ID.Wrappee})
    }
    if updater.IsRead.HasValue {
    setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{Field: BookmarkFields.IsRead, Updater: updater.IsRead.Wrappee})
    }
    

    return setUpdaters
}

func (updater *BookmarkUpdater) ApplyToModel(bookmarkModel *Bookmark) {
    if updater.CreatedAt.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).CreatedAt, updater.CreatedAt.Wrappee)
    }
    if updater.UpdatedAt.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).UpdatedAt, updater.UpdatedAt.Wrappee)
    }
    if updater.URL.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).URL, updater.URL.Wrappee)
    }
    if updater.Title.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).Title, updater.Title.Wrappee)
    }
    if updater.DeletedAt.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).DeletedAt, updater.DeletedAt.Wrappee)
    }
    if updater.BookmarkTypeID.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).BookmarkTypeID, updater.BookmarkTypeID.Wrappee)
    }
    if updater.IsCollection.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).IsCollection, updater.IsCollection.Wrappee)
    }
    if updater.ID.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).ID, updater.ID.Wrappee)
    }
    if updater.IsRead.HasValue {
        model.ApplyUpdater(&(*bookmarkModel).IsRead, updater.IsRead.Wrappee)
    }
    
}

type PsqlBookmarkRepositoryHook func(context.Context, PsqlBookmarkRepository) error

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

func GetBookmarkDomainToSqlRepositoryModel(db *sql.DB) func(domainModel *domain.Bookmark) (sqlRepositoryModel *Bookmark, err error) {
    return func(domainModel *domain.Bookmark) (sqlRepositoryModel *Bookmark, err error) {
        return BookmarkDomainToSqlRepositoryModel(db, domainModel)
    }
}

func GetBookmarkSqlRepositoryToDomainModel(db *sql.DB) func(repositoryModel *Bookmark) (domainModel *domain.Bookmark, err error) {
    return func(sqlRepositoryModel *Bookmark) (domainModel *domain.Bookmark, err error) {
        return BookmarkSqlRepositoryToDomainModel(db,sqlRepositoryModel)
    }
}

type PsqlBookmarkRepositoryConstructorArgs struct {
    DB *sql.DB
}

func (repo *PsqlBookmarkRepository) New(args any) (repo PsqlBookmarkRepository, err error) {
    constructorArgs, ok := args.(PsqlBookmarkRepositoryConstructorArgs)
    if !ok {
        err = fmt.Errorf("expected type %T but got %T", PsqlBookmarkRepositoryConstructorArgs{}, args)

        return
    }


    repo.db = constructorArgs.DB

    return
}

func (repo *PsqlBookmarkRepository) Add(ctx context.Context, domainModels []*domain.Bookmark) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetBookmarkDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

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

func (repo *PsqlBookmarkRepository) Replace(ctx context.Context, domainModels []*domain.Bookmark) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetBookmarkDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

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

func (repo *PsqlBookmarkRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter, domainColumnUpdater *domain.BookmarkUpdater) (numAffectedRecords int64, err error) {
    // NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
    // Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

	var modelsToUpdate BookmarkSlice

    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    columnUpdater, err := BookmarkDomainToSqlRepositoryUpdater(repo.db, domainColumnUpdater)
    if err != nil {
        return
    }


    setFilters := *columnFilter.GetSetFilters()

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

    for _, model := range modelsToUpdate {
        columnUpdater.ApplyToModel(model)
        model.Update(ctx, tx, boil.Infer())
    }

    tx.Commit()

    return
}

func (repo *PsqlBookmarkRepository) Delete(ctx context.Context, domainModels []*domain.Bookmark) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetBookmarkDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

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

func (repo *PsqlBookmarkRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (numAffectedRecords int64, err error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Bookmarks(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *PsqlBookmarkRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (int64, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Count(ctx, repo.db)
}

func (repo *PsqlBookmarkRepository) CountAll(ctx context.Context) (int64, error) {
	return Bookmarks().Count(ctx, repo.db)
}

func (repo *PsqlBookmarkRepository) DoesExist(ctx context.Context, domainModel *domain.Bookmark) (bool, error) {
    repositoryModel, err := BookmarkDomainToSqlRepositoryModel(repo.db, domainModel)
    if err != nil {
        return false, err
    }


	return BookmarkExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *PsqlBookmarkRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (bool, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return false, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Exists(ctx, repo.db)
}

func (repo *PsqlBookmarkRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) ([]*domain.Bookmark, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return []*domain.Bookmark{}, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

    repositoryModels, err := Bookmarks(queryFilters...).All(ctx, repo.db)
    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetBookmarkSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *PsqlBookmarkRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (*domain.Bookmark, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return nil, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

    repositoryModel, err := Bookmarks(queryFilters...).One(ctx, repo.db)

    var domainModel *domain.Bookmark
    if err != nil {
        return domainModel, err
    }

    domainModel, err =BookmarkSqlRepositoryToDomainModel(repo.db, repositoryModel)

    return domainModel, err
}

func (repo *PsqlBookmarkRepository) GetAll(ctx context.Context) ([]*domain.Bookmark, error) {
    repositoryModels, err := Bookmarks().All(ctx, repo.db)
    if err != nil {
        return []*domain.Bookmark{}, err
    }

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetBookmarkSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *PsqlBookmarkRepository) AddType(ctx context.Context, type_ string) error {
    repositoryModel := BookmarkType{Type: type_}

    return repositoryModel.Insert(ctx, repo.db, boil.Infer())
}

func (repo *PsqlBookmarkRepository) DeleteType(ctx context.Context, type_ string) error {
    repositoryModel := BookmarkType{Type: type_}

    _, err := repositoryModel.Delete(ctx, repo.db)

    return err
}

func (repo *PsqlBookmarkRepository) UpdateType(ctx context.Context, oldType string, newType string) error {
    repositoryModel, err := BookmarkTypes(BookmarkTypeWhere.Type.EQ(oldType)).One(ctx, repo.db)
    if err != nil {
        return err
    }

    repositoryModel.Type = newType

    _, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

    return err
}
