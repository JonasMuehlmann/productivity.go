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

type MssqlDocumentRepository struct {
    db sql.DB
}
type DocumentField string

var DocumentFields = struct {
    ID  DocumentField
    Path  DocumentField
    DocumentTypeID  DocumentField
    CreatedAt  DocumentField
    UpdatedAt  DocumentField
    DeletedAt  DocumentField
    
}{
    ID: "id",
    Path: "path",
    DocumentTypeID: "document_type_id",
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
    DeletedAt: "deleted_at",
    
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
    ID optional.Optional[FilterOperation[int64]]
    Path optional.Optional[FilterOperation[string]]
    DocumentTypeID optional.Optional[FilterOperation[int64]]
    CreatedAt optional.Optional[FilterOperation[string]]
    UpdatedAt optional.Optional[FilterOperation[string]]
    DeletedAt optional.Optional[FilterOperation[null.String]]
    
    DocumentType optional.Optional[UpdateOperation[*repository.DocumentType]]
    Tags optional.Optional[UpdateOperation[repository.TagSlice]]
    SourceDocuments optional.Optional[UpdateOperation[repository.DocumentSlice]]
    DestinationDocuments optional.Optional[UpdateOperation[repository.DocumentSlice]]
    
}

type DocumentUpdater struct {
    ID optional.Optional[UpdateOperation[int64]]
    Path optional.Optional[UpdateOperation[string]]
    DocumentTypeID optional.Optional[UpdateOperation[int64]]
    CreatedAt optional.Optional[UpdateOperation[string]]
    UpdatedAt optional.Optional[UpdateOperation[string]]
    DeletedAt optional.Optional[UpdateOperation[null.String]]
    
    DocumentType optional.Optional[UpdateOperation[*repository.DocumentType]]
    Tags optional.Optional[UpdateOperation[repository.TagSlice]]
    SourceDocuments optional.Optional[UpdateOperation[repository.DocumentSlice]]
    DestinationDocuments optional.Optional[UpdateOperation[repository.DocumentSlice]]
    
}

type MssqlDocumentRepositoryHook func(context.Context, MssqlDocumentRepository) error

func (repo *MssqlDocumentRepository) New(args ...any) (MssqlDocumentRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) Add(ctx context.Context, domainModels []domain.Document) (numAffectedRecords int, newID int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) Replace(ctx context.Context, domainModels []domain.Document) error {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) UpdateWhere(ctx context.Context, columnFilter domain.DocumentFilter, columnUpdaters map[domain.DocumentField]domain.DocumentUpdater) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) Delete(ctx context.Context, domainModels []domain.Document) error {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) DeleteWhere(ctx context.Context, columnFilter domain.DocumentFilter) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) CountWhere(ctx context.Context, columnFilter domain.DocumentFilter) int {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) CountAll(ctx context.Context) int {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) DoesExist(ctx context.Context, domainModel domain.Document) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) DoesExistWhere(ctx context.Context, columnFilter domain.DocumentFilter) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) GetWhere(ctx context.Context, columnFilter domain.DocumentFilter) []domain.Document {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) GetFirstWhere(ctx context.Context, columnFilter domain.DocumentFilter) domain.Document {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlDocumentRepository) GetAll(ctx context.Context) []domain.Document {
        panic("not implemented") // TODO: Implement
}
