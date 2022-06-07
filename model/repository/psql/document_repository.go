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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/Document_repository.go.tpl

package repository

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"

	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"time"
)

type PsqlDocumentRepository struct {
	db *sql.DB
}
type DocumentField string

var DocumentFields = struct {
	ID             DocumentField
	Path           DocumentField
	DocumentTypeID DocumentField
	CreatedAt      DocumentField
	UpdatedAt      DocumentField
	DeletedAt      DocumentField
}{
	ID:             "id",
	Path:           "path",
	DocumentTypeID: "document_type_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

var DocumentFieldsList = []DocumentField{
	DocumentField("ID"),
	DocumentField("Path"),
	DocumentField("DocumentTypeID"),
	DocumentField("CreatedAt"),
	DocumentField("UpdatedAt"),
	DocumentField("DeletedAt"),
}

var DocumentRelationsList = []string{
	"DocumentType",
	"Tags",
	"SourceDocuments",
	"DestinationDocuments",
}

type DocumentFilter struct {
	ID             optional.Optional[model.FilterOperation[int64]]
	Path           optional.Optional[model.FilterOperation[string]]
	DocumentTypeID optional.Optional[model.FilterOperation[null.Int]]
	CreatedAt      optional.Optional[model.FilterOperation[time.Time]]
	UpdatedAt      optional.Optional[model.FilterOperation[time.Time]]
	DeletedAt      optional.Optional[model.FilterOperation[null.Time]]

	DocumentType         optional.Optional[model.FilterOperation[*DocumentType]]
	Tags                 optional.Optional[model.FilterOperation[*Tag]]
	SourceDocuments      optional.Optional[model.FilterOperation[*Document]]
	DestinationDocuments optional.Optional[model.FilterOperation[*Document]]
}

type DocumentFilterMapping[T any] struct {
	Field           DocumentField
	FilterOperation model.FilterOperation[T]
}

func (filter *DocumentFilter) GetSetFilters() *list.List {
	setFilters := list.New()

	if filter.ID.HasValue {
		setFilters.PushBack(DocumentFilterMapping[int64]{Field: DocumentFields.ID, FilterOperation: filter.ID.Wrappee})
	}
	if filter.Path.HasValue {
		setFilters.PushBack(DocumentFilterMapping[string]{Field: DocumentFields.Path, FilterOperation: filter.Path.Wrappee})
	}
	if filter.DocumentTypeID.HasValue {
		setFilters.PushBack(DocumentFilterMapping[null.Int]{Field: DocumentFields.DocumentTypeID, FilterOperation: filter.DocumentTypeID.Wrappee})
	}
	if filter.CreatedAt.HasValue {
		setFilters.PushBack(DocumentFilterMapping[time.Time]{Field: DocumentFields.CreatedAt, FilterOperation: filter.CreatedAt.Wrappee})
	}
	if filter.UpdatedAt.HasValue {
		setFilters.PushBack(DocumentFilterMapping[time.Time]{Field: DocumentFields.UpdatedAt, FilterOperation: filter.UpdatedAt.Wrappee})
	}
	if filter.DeletedAt.HasValue {
		setFilters.PushBack(DocumentFilterMapping[null.Time]{Field: DocumentFields.DeletedAt, FilterOperation: filter.DeletedAt.Wrappee})
	}

	return setFilters
}

type DocumentUpdater struct {
	ID             optional.Optional[model.UpdateOperation[int64]]
	Path           optional.Optional[model.UpdateOperation[string]]
	DocumentTypeID optional.Optional[model.UpdateOperation[null.Int]]
	CreatedAt      optional.Optional[model.UpdateOperation[time.Time]]
	UpdatedAt      optional.Optional[model.UpdateOperation[time.Time]]
	DeletedAt      optional.Optional[model.UpdateOperation[null.Time]]

	DocumentType         optional.Optional[model.UpdateOperation[*DocumentType]]
	Tags                 optional.Optional[model.UpdateOperation[TagSlice]]
	SourceDocuments      optional.Optional[model.UpdateOperation[DocumentSlice]]
	DestinationDocuments optional.Optional[model.UpdateOperation[DocumentSlice]]
}

type DocumentUpdaterMapping[T any] struct {
	Field   DocumentField
	Updater model.UpdateOperation[T]
}

func (updater *DocumentUpdater) GetSetUpdaters() *list.List {
	setUpdaters := list.New()

	if updater.ID.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[int64]{Field: DocumentFields.ID, Updater: updater.ID.Wrappee})
	}
	if updater.Path.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[string]{Field: DocumentFields.Path, Updater: updater.Path.Wrappee})
	}
	if updater.DocumentTypeID.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[null.Int]{Field: DocumentFields.DocumentTypeID, Updater: updater.DocumentTypeID.Wrappee})
	}
	if updater.CreatedAt.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[time.Time]{Field: DocumentFields.CreatedAt, Updater: updater.CreatedAt.Wrappee})
	}
	if updater.UpdatedAt.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[time.Time]{Field: DocumentFields.UpdatedAt, Updater: updater.UpdatedAt.Wrappee})
	}
	if updater.DeletedAt.HasValue {
		setUpdaters.PushBack(DocumentUpdaterMapping[null.Time]{Field: DocumentFields.DeletedAt, Updater: updater.DeletedAt.Wrappee})
	}

	return setUpdaters
}

func (updater *DocumentUpdater) ApplyToModel(documentModel *Document) {
	if updater.ID.HasValue {
		model.ApplyUpdater(&(*documentModel).ID, updater.ID.Wrappee)
	}
	if updater.Path.HasValue {
		model.ApplyUpdater(&(*documentModel).Path, updater.Path.Wrappee)
	}
	if updater.DocumentTypeID.HasValue {
		model.ApplyUpdater(&(*documentModel).DocumentTypeID, updater.DocumentTypeID.Wrappee)
	}
	if updater.CreatedAt.HasValue {
		model.ApplyUpdater(&(*documentModel).CreatedAt, updater.CreatedAt.Wrappee)
	}
	if updater.UpdatedAt.HasValue {
		model.ApplyUpdater(&(*documentModel).UpdatedAt, updater.UpdatedAt.Wrappee)
	}
	if updater.DeletedAt.HasValue {
		model.ApplyUpdater(&(*documentModel).DeletedAt, updater.DeletedAt.Wrappee)
	}

}

type PsqlDocumentRepositoryHook func(context.Context, PsqlDocumentRepository) error

type queryModSliceDocument []qm.QueryMod

func (s queryModSliceDocument) Apply(q *queries.Query) {
	qm.Apply(q, s...)
}

func buildQueryModFilterDocument[T any](filterField DocumentField, filterOperation model.FilterOperation[T]) queryModSliceDocument {
	var newQueryMod queryModSliceDocument

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
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.RHS))))
	case model.FilterAnd:
		filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
		if !ok {
			panic("Expected a scalar operand for FilterAnd operator")
		}

		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.LHS)))
		newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.RHS)))
	default:
		panic("Unhandled FilterOperator")
	}

	return newQueryMod
}

func buildQueryModListFromFilterDocument(setFilters list.List) queryModSliceDocument {
	queryModList := make(queryModSliceDocument, 0, 6)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(DocumentFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %T but got %T", DocumentFilterMapping[any]{}, filter))
		}

		newQueryMod := buildQueryModFilterDocument(filterMapping.Field, filterMapping.FilterOperation)

		queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}

func GetDocumentDomainToSqlRepositoryModel(db *sql.DB) func(domainModel *domain.Document) (sqlRepositoryModel *Document, err error) {
	return func(domainModel *domain.Document) (sqlRepositoryModel *Document, err error) {
		return DocumentDomainToSqlRepositoryModel(db, domainModel)
	}
}

func GetDocumentSqlRepositoryToDomainModel(db *sql.DB) func(repositoryModel *Document) (domainModel *domain.Document, err error) {
	return func(sqlRepositoryModel *Document) (domainModel *domain.Document, err error) {
		return DocumentSqlRepositoryToDomainModel(db, sqlRepositoryModel)
	}
}

type PsqlDocumentRepositoryConstructorArgs struct {
	DB *sql.DB
}

func (repo *PsqlDocumentRepository) New(args any) (*PsqlDocumentRepository, error) {
	constructorArgs, ok := args.(PsqlDocumentRepositoryConstructorArgs)
	if !ok {
		return repo, fmt.Errorf("expected type %T but got %T", PsqlDocumentRepositoryConstructorArgs{}, args)
	}

	repo.db = constructorArgs.DB

	return repo, nil
}

func (repo *PsqlDocumentRepository) Add(ctx context.Context, domainModels []*domain.Document) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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

func (repo *PsqlDocumentRepository) Replace(ctx context.Context, domainModels []*domain.Document) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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

func (repo *PsqlDocumentRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter, domainColumnUpdater *domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate DocumentSlice

	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return
	}

	columnUpdater, err := DocumentDomainToSqlRepositoryUpdater(ctx, repo.db, domainColumnUpdater)
	if err != nil {
		return
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	modelsToUpdate, err = Documents(queryFilters...).All(ctx, repo.db)
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

func (repo *PsqlDocumentRepository) Delete(ctx context.Context, domainModels []*domain.Document) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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

func (repo *PsqlDocumentRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (numAffectedRecords int64, err error) {
	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Documents(queryFilters...).DeleteAll(ctx, tx)
	if err != nil {
		return
	}

	tx.Commit()

	return
}

func (repo *PsqlDocumentRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (int64, error) {
	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return 0, err
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Count(ctx, repo.db)
}

func (repo *PsqlDocumentRepository) CountAll(ctx context.Context) (int64, error) {
	return Documents().Count(ctx, repo.db)
}

func (repo *PsqlDocumentRepository) DoesExist(ctx context.Context, domainModel *domain.Document) (bool, error) {
	repositoryModel, err := DocumentDomainToSqlRepositoryModel(repo.db, domainModel)
	if err != nil {
		return false, err
	}

	return DocumentExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *PsqlDocumentRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (bool, error) {
	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return false, err
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Exists(ctx, repo.db)
}

func (repo *PsqlDocumentRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) ([]*domain.Document, error) {
	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return []*domain.Document{}, err
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	repositoryModels, err := Documents(queryFilters...).All(ctx, repo.db)

	domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetDocumentSqlRepositoryToDomainModel(ctx, repo.db))

	return domainModels, err
}

func (repo *PsqlDocumentRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (*domain.Document, error) {
	columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
	if err != nil {
		return nil, err
	}

	setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	repositoryModel, err := Documents(queryFilters...).One(ctx, repo.db)

	var domainModel *domain.Document
	if err != nil {
		return domainModel, err
	}

	domainModel, err = DocumentSqlRepositoryToDomainModel(repo.db, repositoryModel)

	return domainModel, err
}

func (repo *PsqlDocumentRepository) GetAll(ctx context.Context) ([]*domain.Document, error) {
	repositoryModels, err := Documents().All(ctx, repo.db)

	if err != nil {
		return []*domain.Document{}, err
	}

	domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetDocumentSqlRepositoryToDomainModel(ctx, repo.db))

	return domainModels, err
}

func (repo *PsqlDocumentRepository) AddType(ctx context.Context, type_ string) error {
	repositoryModel := DocumentType{DocumentType: type_}

	return repositoryModel.Insert(ctx, repo.db, boil.Infer())
}

func (repo *PsqlDocumentRepository) DeleteType(ctx context.Context, type_ string) error {
	repositoryModel := DocumentType{DocumentType: type_}

	_, err := repositoryModel.Delete(ctx, repo.db)

	return err
}

func (repo *PsqlDocumentRepository) UpdateType(ctx context.Context, oldType string, newType string) error {
	repositoryModel, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(oldType)).One(ctx, repo.db)
	if err != nil {
		return err
	}

	repositoryModel.DocumentType = newType

	_, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

	return err
}
