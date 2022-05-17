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
	"github.com/JonasMuehlmann/optional.go"
    "context"
	"fmt"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/null/v8"
	"container/list"
)

type Sqlite3BookmarkRepository struct {
    db *sql.DB
}
type BookmarkField string

var BookmarkFields = struct {
    ID  BookmarkField
    IsRead  BookmarkField
    Title  BookmarkField
    URL  BookmarkField
    BookmarkTypeID  BookmarkField
    IsCollection  BookmarkField
    CreatedAt  BookmarkField
    UpdatedAt  BookmarkField
    DeletedAt  BookmarkField
    
}{
    ID: "id",
    IsRead: "is_read",
    Title: "title",
    URL: "url",
    BookmarkTypeID: "bookmark_type_id",
    IsCollection: "is_collection",
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
    DeletedAt: "deleted_at",
    
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
    ID optional.Optional[model.FilterOperation[int64]]
    IsRead optional.Optional[model.FilterOperation[int64]]
    Title optional.Optional[model.FilterOperation[null.String]]
    URL optional.Optional[model.FilterOperation[string]]
    BookmarkTypeID optional.Optional[model.FilterOperation[null.Int64]]
    IsCollection optional.Optional[model.FilterOperation[int64]]
    CreatedAt optional.Optional[model.FilterOperation[string]]
    UpdatedAt optional.Optional[model.FilterOperation[string]]
    DeletedAt optional.Optional[model.FilterOperation[null.String]]
    
    BookmarkType optional.Optional[model.UpdateOperation[*BookmarkType]]
    Tags optional.Optional[model.UpdateOperation[TagSlice]]
    
}

type BookmarkFilterMapping[T any] struct {
    Field BookmarkField
    FilterOperation model.FilterOperation[T]
}

func (filter *BookmarkFilter) GetSetFilters() *list.List {
    setFilters := list.New()

    if filter.ID.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[int64]{BookmarkFields.ID, filter.ID.Wrappee})
    }
    if filter.IsRead.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[int64]{BookmarkFields.IsRead, filter.IsRead.Wrappee})
    }
    if filter.Title.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[null.String]{BookmarkFields.Title, filter.Title.Wrappee})
    }
    if filter.URL.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[string]{BookmarkFields.URL, filter.URL.Wrappee})
    }
    if filter.BookmarkTypeID.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[null.Int64]{BookmarkFields.BookmarkTypeID, filter.BookmarkTypeID.Wrappee})
    }
    if filter.IsCollection.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[int64]{BookmarkFields.IsCollection, filter.IsCollection.Wrappee})
    }
    if filter.CreatedAt.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[string]{BookmarkFields.CreatedAt, filter.CreatedAt.Wrappee})
    }
    if filter.UpdatedAt.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[string]{BookmarkFields.UpdatedAt, filter.UpdatedAt.Wrappee})
    }
    if filter.DeletedAt.HasValue {
        setFilters.PushBack(BookmarkFilterMapping[null.String]{BookmarkFields.DeletedAt, filter.DeletedAt.Wrappee})
    }
    

    return setFilters
}

type BookmarkUpdater struct {
    ID optional.Optional[model.UpdateOperation[int64]]
    IsRead optional.Optional[model.UpdateOperation[int64]]
    Title optional.Optional[model.UpdateOperation[null.String]]
    URL optional.Optional[model.UpdateOperation[string]]
    BookmarkTypeID optional.Optional[model.UpdateOperation[null.Int64]]
    IsCollection optional.Optional[model.UpdateOperation[int64]]
    CreatedAt optional.Optional[model.UpdateOperation[string]]
    UpdatedAt optional.Optional[model.UpdateOperation[string]]
    DeletedAt optional.Optional[model.UpdateOperation[null.String]]
    
    BookmarkType optional.Optional[model.UpdateOperation[*BookmarkType]]
    Tags optional.Optional[model.UpdateOperation[TagSlice]]
    
}

type BookmarkUpdaterMapping[T any] struct {
    Field BookmarkField
    Updater model.UpdateOperation[T]
}

func (updater *BookmarkUpdater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    if updater.ID.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{BookmarkFields.ID, updater.ID.Wrappee})
    }
    if updater.IsRead.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{BookmarkFields.IsRead, updater.IsRead.Wrappee})
    }
    if updater.Title.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[null.String]{BookmarkFields.Title, updater.Title.Wrappee})
    }
    if updater.URL.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[string]{BookmarkFields.URL, updater.URL.Wrappee})
    }
    if updater.BookmarkTypeID.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[null.Int64]{BookmarkFields.BookmarkTypeID, updater.BookmarkTypeID.Wrappee})
    }
    if updater.IsCollection.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[int64]{BookmarkFields.IsCollection, updater.IsCollection.Wrappee})
    }
    if updater.CreatedAt.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[string]{BookmarkFields.CreatedAt, updater.CreatedAt.Wrappee})
    }
    if updater.UpdatedAt.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[string]{BookmarkFields.UpdatedAt, updater.UpdatedAt.Wrappee})
    }
    if updater.DeletedAt.HasValue {
        setUpdaters.PushBack(BookmarkUpdaterMapping[null.String]{BookmarkFields.DeletedAt, updater.DeletedAt.Wrappee})
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

type Sqlite3BookmarkRepositoryHook func(context.Context, Sqlite3BookmarkRepository) error

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
			panic(fmt.Sprintf("Expected type %t but got %t", BookmarkFilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilterBookmark(filterMapping.Field, filterMapping.FilterOperation)

        for _, queryMod := range newQueryMod {
            queryModList = append(queryModList, queryMod)
        }
	}

	return queryModList
}

func (repo * Sqlite3BookmarkRepository) New(args ...any) (Sqlite3BookmarkRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3BookmarkRepository) Add(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *Sqlite3BookmarkRepository) Replace(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *Sqlite3BookmarkRepository) UpdateWhere(ctx context.Context, columnFilter BookmarkFilter, columnUpdater BookmarkUpdater) (numAffectedRecords int64, err error) {
    // NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
    // Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

    // TODO: Implement translator from domainColumnFilter to repositoryColumnFilter and updater
	var modelsToUpdate BookmarkSlice

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	modelsToUpdate, err = Bookmarks(queryFilters...).All(ctx, repo.db)

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

func (repo *Sqlite3BookmarkRepository) Delete(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *Sqlite3BookmarkRepository) DeleteWhere(ctx context.Context, columnFilter BookmarkFilter) (numAffectedRecords int64, err error) {
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

func (repo *Sqlite3BookmarkRepository) CountWhere(ctx context.Context, columnFilter BookmarkFilter) (int64, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Count(ctx, repo.db)
}

func (repo *Sqlite3BookmarkRepository) CountAll(ctx context.Context) (int64, error) {
	return Bookmarks().Count(ctx, repo.db)
}

func (repo *Sqlite3BookmarkRepository) DoesExist(ctx context.Context, repositoryModel Bookmark) (bool, error) {
	return BookmarkExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *Sqlite3BookmarkRepository) DoesExistWhere(ctx context.Context, columnFilter BookmarkFilter) (bool, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).Exists(ctx, repo.db)
}

func (repo *Sqlite3BookmarkRepository) GetWhere(ctx context.Context, columnFilter BookmarkFilter) ([]*Bookmark, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).All(ctx, repo.db)
}

func (repo *Sqlite3BookmarkRepository) GetFirstWhere(ctx context.Context, columnFilter BookmarkFilter) (*Bookmark, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterBookmark(setFilters)

	return Bookmarks(queryFilters...).One(ctx, repo.db)
}

func (repo *Sqlite3BookmarkRepository) GetAll(ctx context.Context) ([]*Bookmark, error) {
	return Bookmarks().All(ctx, repo.db)
}
