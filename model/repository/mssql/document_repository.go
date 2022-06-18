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
type MssqlDocumentRepository struct {
	db *sql.DB

	tagRepository repoCommon.TagRepository
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
	DocumentTypeID optional.Optional[model.FilterOperation[null.Int64]]
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
		setFilters.PushBack(DocumentFilterMapping[null.Int64]{Field: DocumentFields.DocumentTypeID, FilterOperation: filter.DocumentTypeID.Wrappee})
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
	CreatedAt            optional.Optional[model.UpdateOperation[time.Time]]
	UpdatedAt            optional.Optional[model.UpdateOperation[time.Time]]
	DocumentType         optional.Optional[model.UpdateOperation[*DocumentType]]
	DeletedAt            optional.Optional[model.UpdateOperation[null.Time]]
	Path                 optional.Optional[model.UpdateOperation[string]]
	SourceDocuments      optional.Optional[model.UpdateOperation[DocumentSlice]]
	Tags                 optional.Optional[model.UpdateOperation[TagSlice]]
	DestinationDocuments optional.Optional[model.UpdateOperation[DocumentSlice]]
	DocumentTypeID       optional.Optional[model.UpdateOperation[null.Int64]]
	ID                   optional.Optional[model.UpdateOperation[int64]]
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
		setUpdaters.PushBack(DocumentUpdaterMapping[null.Int64]{Field: DocumentFields.DocumentTypeID, Updater: updater.DocumentTypeID.Wrappee})
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

type MssqlDocumentRepositoryConstructorArgs struct {
	DB *sql.DB

	TagRepository repoCommon.TagRepository
}

func (repo *MssqlDocumentRepository) New(args any) (newRepo repoCommon.DocumentRepository, err error) {
	constructorArgs, ok := args.(MssqlDocumentRepositoryConstructorArgs)
	if !ok {
		err = fmt.Errorf("expected type %T but got %T", MssqlDocumentRepositoryConstructorArgs{}, args)

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
func (repo *MssqlDocumentRepository) Add(ctx context.Context, domainModels []*domain.Document) error {
	if len(domainModels) == 0 {
		log.Debug(helper.LogMessageEmptyInput)
		return nil
	}

	err := goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Document](nil))
	if err == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetDocumentDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Document)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		err = repoModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlDocumentRepository) Replace(ctx context.Context, domainModels []*domain.Document) error {

	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetDocumentDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Document)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		_, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}
func (repo *MssqlDocumentRepository) Upsert(ctx context.Context, domainModels []*domain.Document) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetDocumentDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Document)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		err = repoModel.Upsert(ctx, tx, boil.Infer(), boil.Infer())

		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlDocumentRepository) Update(ctx context.Context, domainModels []*domain.Document, domainColumnUpdater *domain.DocumentUpdater) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetDocumentDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	repositoryUpdater, err := repo.DocumentDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Document)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		repoUpdater, ok := repositoryUpdater.(*DocumentUpdater)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		repoUpdater.ApplyToModel(repoModel)
		repoModel.Update(ctx, tx, boil.Infer())
	}

	tx.Commit()

	return err
}

func (repo *MssqlDocumentRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter, domainColumnUpdater *domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate DocumentSlice

	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repositoryUpdater, err := repo.DocumentDomainToRepositoryUpdater(ctx, domainColumnUpdater)
	if err != nil {
		return
	}

	repoUpdater, ok := repositoryUpdater.(*DocumentUpdater)
	if !ok {
		err = fmt.Errorf("Expected type *DocumentUpdater but got %T", repoUpdater)

		return
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		err = fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)

		return
	}

	setFilters := *repoFilter.GetSetFilters()

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

	for _, repoModel := range modelsToUpdate {
		repoUpdater.ApplyToModel(repoModel)
		repoModel.Update(ctx, tx, boil.Infer())
	}

	tx.Commit()

	return
}

func (repo *MssqlDocumentRepository) Delete(ctx context.Context, domainModels []*domain.Document) error {
	repositoryModels, err := goaoi.TransformCopySlice(domainModels, repo.GetDocumentDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		repoModel, ok := repositoryModel.(*Document)
		if !ok {
			return fmt.Errorf("Expected type *Document but got %T", repoModel)
		}

		_, err = repoModel.Delete(ctx, tx)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo *MssqlDocumentRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (numAffectedRecords int64, err error) {
	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		err = fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)

		return
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Documents(queryFilters...).DeleteAll(ctx, tx)

	tx.Commit()

	return
}

func (repo *MssqlDocumentRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (int64, error) {
	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return 0, err
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		return 0, fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)

	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Count(ctx, repo.db)
}

func (repo *MssqlDocumentRepository) CountAll(ctx context.Context) (int64, error) {
	return Documents().Count(ctx, repo.db)
}

func (repo *MssqlDocumentRepository) DoesExist(ctx context.Context, domainModel *domain.Document) (bool, error) {
	repositoryModel, err := repo.DocumentDomainToRepositoryModel(ctx, domainModel)
	if err != nil {
		return false, err
	}

	repoModel, ok := repositoryModel.(*Document)
	if !ok {
		return false, fmt.Errorf("Expected type *Document but got %T", repoModel)
	}

	return DocumentExists(ctx, repo.db, repoModel.ID)
}

func (repo *MssqlDocumentRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (bool, error) {
	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return false, err
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		return false, fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Exists(ctx, repo.db)
}

func (repo *MssqlDocumentRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) ([]*domain.Document, error) {
	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return []*domain.Document{}, err
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		return []*domain.Document{}, fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	repositoryModels, err := Documents(queryFilters...).All(ctx, repo.db)

	domainModels := make([]*domain.Document, 0, len(repositoryModels))

	for _, repoModel := range repositoryModels {
		domainModel, err := repo.DocumentRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			return domainModels, err
		}

		domainModels = append(domainModels, domainModel)
	}

	return domainModels, err
}

func (repo *MssqlDocumentRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (*domain.Document, error) {
	repositoryFilter, err := repo.DocumentDomainToRepositoryFilter(ctx, domainColumnFilter)
	if err != nil {
		return nil, err
	}

	repoFilter, ok := repositoryFilter.(*DocumentFilter)
	if !ok {
		return nil, fmt.Errorf("Expected type *DocumentFilter but got %T", repoFilter)
	}

	setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	repositoryModel, err := Documents(queryFilters...).One(ctx, repo.db)

	var domainModel *domain.Document
	if err != nil {
		return domainModel, err
	}

	domainModel, err = repo.DocumentRepositoryToDomainModel(ctx, repositoryModel)

	return domainModel, err
}

func (repo *MssqlDocumentRepository) GetAll(ctx context.Context) ([]*domain.Document, error) {
	repositoryModels, err := Documents().All(ctx, repo.db)
	if err != nil {
		return []*domain.Document{}, err
	}

	domainModels := make([]*domain.Document, 0, len(repositoryModels))

	for _, repoModel := range repositoryModels {
		domainModel, err := repo.DocumentRepositoryToDomainModel(ctx, repoModel)
		if err != nil {
			return domainModels, err
		}

		domainModels = append(domainModels, domainModel)
	}

	return domainModels, err
}

func (repo *MssqlDocumentRepository) AddType(ctx context.Context, types []string) error {
	for _, type_ := range types {
		repositoryModel := DocumentType{DocumentType: type_}

		err := repositoryModel.Insert(ctx, repo.db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *MssqlDocumentRepository) DeleteType(ctx context.Context, types []string) error {
	_, err := DocumentTypes(DocumentTypeWhere.DocumentType.IN(types)).DeleteAll(ctx, repo.db)

	return err
}

func (repo *MssqlDocumentRepository) UpdateType(ctx context.Context, oldType string, newType string) error {
	repositoryModel, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(oldType)).One(ctx, repo.db)
	if err != nil {
		return err
	}

	repositoryModel.DocumentType = newType

	_, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

	return err
}

func (repo *MssqlDocumentRepository) GetTagRepository() repoCommon.TagRepository {
	return repo.tagRepository
}

//******************************************************************//
//                            Converters                            //
//******************************************************************//
func (repo *MssqlDocumentRepository) GetDocumentDomainToRepositoryModel(ctx context.Context) func(domainModel *domain.Document) (repositoryModel any, err error) {
	return func(domainModel *domain.Document) (repositoryModel any, err error) {
		return repo.DocumentDomainToRepositoryModel(ctx, domainModel)
	}
}

func (repo *MssqlDocumentRepository) GetDocumentRepositoryToDomainModel(ctx context.Context) func(repositoryModel any) (domainModel *domain.Document, err error) {
	return func(repositoryModel any) (domainModel *domain.Document, err error) {

		return repo.DocumentRepositoryToDomainModel(ctx, repositoryModel)
	}
}

//******************************************************************//
//                          Model Converter                         //
//******************************************************************//

func (repo *MssqlDocumentRepository) DocumentDomainToRepositoryModel(ctx context.Context, domainModel *domain.Document) (repositoryModel any, err error) {
	repositoryModelConcrete := new(Document)
	repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()

	repositoryModelConcrete.Path = domainModel.Path
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

	//*************************    Set Tags    *************************//
	var repositoryTag *Tag

	repositoryModelConcrete.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
	for _, modelTag := range domainModel.Tags {
		repositoryTag, err = Tags(TagWhere.Tag.EQ(modelTag.Tag)).One(ctx, repo.db)
		if err != nil {
			return
		}

		repositoryModelConcrete.R.Tags = append(repositoryModelConcrete.R.Tags, &Tag{Tag: modelTag.Tag, ID: repositoryTag.ID})
	}

	//*************************    Set Type    *************************//
	var repositoryDocumentType *DocumentType

	if domainModel.DocumentType.HasValue {
		repositoryModelConcrete.R.DocumentType = &DocumentType{DocumentType: domainModel.DocumentType.Wrappee}
		repositoryDocumentType, err = DocumentTypes(DocumentTypeWhere.DocumentType.EQ(domainModel.DocumentType.Wrappee)).One(ctx, repo.db)
		if err != nil {
			return
		}

		if repositoryDocumentType != nil {
			repositoryModelConcrete.DocumentTypeID = null.NewInt64(repositoryDocumentType.ID, true)
			repositoryModelConcrete.R.DocumentType.ID = repositoryDocumentType.ID
		} else {
			repositoryModelConcrete.R.DocumentType = nil
		}
	}

	//**************    Set linked/backlinked documents    *************//
	var repositoryDocumentRaw any

	repositoryModelConcrete.R.SourceDocuments = make(DocumentSlice, 0, len(domainModel.LinkedDocuments))
	repositoryModelConcrete.R.DestinationDocuments = make(DocumentSlice, 0, len(domainModel.BacklinkedDocuments))

	for _, link := range domainModel.LinkedDocuments {
		repositoryDocumentRaw, err = repo.DocumentDomainToRepositoryModel(ctx, link)
		if err != nil {
			return
		}

		repositoryModelConcrete.R.SourceDocuments = append(repositoryModelConcrete.R.SourceDocuments, repositoryDocumentRaw.(*Document))
	}

	for _, backlink := range domainModel.BacklinkedDocuments {
		repositoryDocumentRaw, err = repo.DocumentDomainToRepositoryModel(ctx, backlink)
		if err != nil {
			return
		}

		repositoryModelConcrete.R.DestinationDocuments = append(repositoryModelConcrete.R.DestinationDocuments, repositoryDocumentRaw.(*Document))
	}

	repositoryModel = repositoryModelConcrete

	return
}

func (repo *MssqlDocumentRepository) DocumentRepositoryToDomainModel(ctx context.Context, repositoryModel any) (domainModel *domain.Document, err error) {
	domainModel = new(domain.Document)

	repositoryModelConcrete := repositoryModel.(Document)

	domainModel.Path = repositoryModelConcrete.Path
	domainModel.ID = repositoryModelConcrete.ID

	if repositoryModelConcrete.R == nil {
		repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()
	}

	if repositoryModelConcrete.R.DocumentType != nil {
		domainModel.DocumentType = optional.Make(repositoryModelConcrete.R.DocumentType.DocumentType)
	}

	//**********************    Set Timestamps    **********************//

	domainModel.CreatedAt = repositoryModelConcrete.CreatedAt
	domainModel.UpdatedAt = repositoryModelConcrete.UpdatedAt

	if repositoryModelConcrete.DeletedAt.Valid {
		domainModel.DeletedAt.Push(repositoryModelConcrete.DeletedAt.Time)
	}

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

	//**************    Set linked/backlinked documents    *************//
	var domainDocument *domain.Document

	domainModel.LinkedDocuments = make([]*domain.Document, 0, len(repositoryModelConcrete.R.SourceDocuments))
	domainModel.BacklinkedDocuments = make([]*domain.Document, 0, len(repositoryModelConcrete.R.DestinationDocuments))

	for _, link := range repositoryModelConcrete.R.SourceDocuments {
		domainDocument, err = repo.DocumentRepositoryToDomainModel(ctx, link)
		if err != nil {
			return
		}

		domainModel.LinkedDocuments = append(domainModel.LinkedDocuments, domainDocument)
	}

	for _, backlink := range repositoryModelConcrete.R.DestinationDocuments {
		domainDocument, err = repo.DocumentRepositoryToDomainModel(ctx, backlink)
		if err != nil {
			return
		}

		domainModel.BacklinkedDocuments = append(domainModel.BacklinkedDocuments, domainDocument)
	}

	return
}

//******************************************************************//
//                         Filter Converter                         //
//******************************************************************//

func (repo *MssqlDocumentRepository) DocumentDomainToRepositoryFilter(ctx context.Context, domainFilter *domain.DocumentFilter) (repositoryFilter any, err error) {

	repositoryFilterConcrete := new(DocumentFilter)

	repositoryFilterConcrete.Path = domainFilter.Path
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
	if domainFilter.DocumentType.HasValue {
		var convertedTypeIDFilter model.FilterOperation[null.Int64]
		var convertedTypeFilter model.FilterOperation[*DocumentType]

		convertedTypeFilter, err = model.ConvertFilter[*DocumentType, optional.Optional[string]](domainFilter.DocumentType.Wrappee, func(type_ optional.Optional[string]) (*DocumentType, error) {
			if !type_.HasValue {
				return nil, nil
			}

			bookmarkType, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(type_.Wrappee)).One(ctx, repo.db)

			return bookmarkType, err
		})
		if err != nil {
			return
		}

		convertedTypeIDFilter, err = model.ConvertFilter[null.Int64, optional.Optional[string]](domainFilter.DocumentType.Wrappee, func(type_ optional.Optional[string]) (null.Int64, error) {
			if !type_.HasValue {
				return null.NewInt64(-1, false), nil
			}

			bookmarkType, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(type_.Wrappee)).One(ctx, repo.db)

			return null.NewInt64(bookmarkType.ID, true), err
		})
		if err != nil {
			return
		}

		repositoryFilterConcrete.DocumentType.Push(convertedTypeFilter)
		repositoryFilterConcrete.DocumentTypeID.Push(convertedTypeIDFilter)
	}

	//**************    Set linked/backlinked documents    *************//
	if domainFilter.LinkedDocuments.HasValue {
		var convertedFilter model.FilterOperation[*Document]

		convertedFilter, err = model.ConvertFilter[*Document, *domain.Document](domainFilter.LinkedDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverterGeneric[domain.Document, Document](ctx, repo.DocumentDomainToRepositoryModel))
		if err != nil {
			return
		}

		repositoryFilterConcrete.SourceDocuments.Push(convertedFilter)
	}
	if domainFilter.BacklinkedDocuments.HasValue {
		var convertedFilter model.FilterOperation[*Document]

		convertedFilter, err = model.ConvertFilter[*Document, *domain.Document](domainFilter.BacklinkedDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverterGeneric[domain.Document, Document](ctx, repo.DocumentDomainToRepositoryModel))
		if err != nil {
			return
		}

		repositoryFilterConcrete.DestinationDocuments.Push(convertedFilter)
	}

	repositoryFilter = repositoryFilterConcrete

	return
}

//******************************************************************//
//                         Updater Converter                        //
//******************************************************************//

func (repo *MssqlDocumentRepository) DocumentDomainToRepositoryUpdater(ctx context.Context, domainUpdater *domain.DocumentUpdater) (repositoryUpdater any, err error) {
	repositoryUpdaterConcrete := new(DocumentUpdater)

	if domainUpdater.DocumentType.HasValue {
		var convertedUpdater *DocumentType
		if domainUpdater.DocumentType.Wrappee.Operand.HasValue {
			convertedUpdater, err = DocumentTypes(DocumentTypeWhere.DocumentType.EQ(domainUpdater.DocumentType.Wrappee.Operand.Wrappee)).One(context.Background(), repo.db)
			if err != nil {
				return
			}
		} else {
			convertedUpdater = nil
		}

		repositoryUpdaterConcrete.DocumentType.Push(model.UpdateOperation[*DocumentType]{Operator: domainUpdater.DocumentType.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.Path.HasValue {
		repositoryUpdaterConcrete.Path.Push(model.UpdateOperation[string]{Operator: domainUpdater.Path.Wrappee.Operator, Operand: repositoryUpdaterConcrete.Path.Wrappee.Operand})
	}

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

	if domainUpdater.LinkedDocuments.HasValue {
		var convertedDocumentRaw any
		convertedUpdater := make(DocumentSlice, 0, len(domainUpdater.LinkedDocuments.Wrappee.Operand))

		for _, document := range domainUpdater.LinkedDocuments.Wrappee.Operand {
			convertedDocumentRaw, err = repo.DocumentDomainToRepositoryModel(ctx, document)
			if err != nil {
				return
			}

			convertedUpdater = append(convertedUpdater, convertedDocumentRaw.(*Document))
		}

		repositoryUpdaterConcrete.DestinationDocuments.Push(model.UpdateOperation[DocumentSlice]{Operator: domainUpdater.LinkedDocuments.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.BacklinkedDocuments.HasValue {
		var convertedDocumentRaw any
		convertedUpdater := make(DocumentSlice, 0, len(domainUpdater.BacklinkedDocuments.Wrappee.Operand))

		for _, document := range domainUpdater.BacklinkedDocuments.Wrappee.Operand {
			convertedDocumentRaw, err = repo.DocumentDomainToRepositoryModel(ctx, document)
			if err != nil {
				return
			}

			convertedUpdater = append(convertedUpdater, convertedDocumentRaw.(*Document))
		}

		repositoryUpdaterConcrete.SourceDocuments.Push(model.UpdateOperation[DocumentSlice]{Operator: domainUpdater.BacklinkedDocuments.Wrappee.Operator, Operand: convertedUpdater})
	}

	if domainUpdater.ID.HasValue {
		repositoryUpdaterConcrete.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: repositoryUpdaterConcrete.ID.Wrappee.Operand})
	}

	repositoryUpdater = repositoryUpdaterConcrete

	return
}
