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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/{{LowercaseBeginning .EntityName}}_repository.go.tpl

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
    {{ if ne .DatabaseName "sqlite3" }}
    "time"
    {{end}}
)

{{template "structDefinition" .}}
{{template "repositoryHelperTypes" .}}

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

type {{$StructName}}ConstructorArgs struct {
    DB *sql.DB
}

func (repo *{{$StructName}}) New(args any) (*{{$StructName}}, error) {
    constructorArgs, ok := args.({{$StructName}}ConstructorArgs)
    if !ok {
        return repo, fmt.Errorf("expected type %T but got %T", {{$StructName}}ConstructorArgs{}, args)
    }

    repo.db = constructorArgs.DB

    return repo, nil
}

func (repo *{{$StructName}}) Add(ctx context.Context, domainModels []*domain.Bookmark) error {
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

func (repo *{{$StructName}}) Replace(ctx context.Context, domainModels []*domain.Bookmark) error {
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

func (repo *{{$StructName}}) UpdateWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter, domainColumnUpdater *domain.BookmarkUpdater) (numAffectedRecords int64, err error) {
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

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

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

func (repo *{{$StructName}}) Delete(ctx context.Context, domainModels []*domain.Bookmark) error {
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

func (repo *{{$StructName}}) DeleteWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (numAffectedRecords int64, err error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Bookmarks(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *{{$StructName}}) CountWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (int64, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Bookmarks(queryFilters...).Count(ctx, repo.db)
}

func (repo *{{$StructName}}) CountAll(ctx context.Context) (int64, error) {
	return Bookmarks().Count(ctx, repo.db)
}

func (repo *{{$StructName}}) DoesExist(ctx context.Context, domainModel *domain.Bookmark) (bool, error) {
    repositoryModel, err := BookmarkDomainToSqlRepositoryModel(repo.db, domainModel)
    if err != nil {
        return false, err
    }


	return BookmarkExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *{{$StructName}}) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (bool, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return false, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Bookmarks(queryFilters...).Exists(ctx, repo.db)
}

func (repo *{{$StructName}}) GetWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) ([]*domain.Bookmark, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return []*domain.Bookmark{}, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModels, err := Bookmarks(queryFilters...).All(ctx, repo.db)
    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetBookmarkSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *{{$StructName}}) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.BookmarkFilter) (*domain.Bookmark, error) {
    columnFilter, err := BookmarkDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return nil, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModel, err := Bookmarks(queryFilters...).One(ctx, repo.db)

    var domainModel *domain.Bookmark
    if err != nil {
        return domainModel, err
    }

    domainModel, err =BookmarkSqlRepositoryToDomainModel(repo.db, repositoryModel)

    return domainModel, err
}

func (repo *{{$StructName}}) GetAll(ctx context.Context) ([]*domain.Bookmark, error) {
    repositoryModels, err := Bookmarks().All(ctx, repo.db)
    if err != nil {
        return []*domain.Bookmark{}, err
    }

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetBookmarkSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *{{$StructName}}) AddType(ctx context.Context, type_ string) error {
    repositoryModel := BookmarkType{Type: type_}

    return repositoryModel.Insert(ctx, repo.db, boil.Infer())
}

func (repo *{{$StructName}}) DeleteType(ctx context.Context, type_ string) error {
    repositoryModel := BookmarkType{Type: type_}

    _, err := repositoryModel.Delete(ctx, repo.db)

    return err
}

func (repo *{{$StructName}}) UpdateType(ctx context.Context, oldType string, newType string) error {
    repositoryModel, err := BookmarkTypes(BookmarkTypeWhere.Type.EQ(oldType)).One(ctx, repo.db)
    if err != nil {
        return err
    }

    repositoryModel.Type = newType

    _, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

    return err
}
