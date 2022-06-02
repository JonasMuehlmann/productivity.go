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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/repository.go.tpl

package repository

import (
    "context"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
)

type DocumentRepository interface {
	New(args any) (repo DocumentRepository, err error)

	Add(ctx context.Context, domainModels []*domain.Document) error
	Replace(ctx context.Context, domainModels []*domain.Document) error
	UpdateWhere(ctx context.Context, columnFilter *domain.DocumentFilter, columnUpdaters *domain.DocumentUpdater) (numAffectedRecords int64, err error)
	Delete(ctx context.Context, domainModels []*domain.Document) error
	DeleteWhere(ctx context.Context, columnFilter *domain.DocumentFilter) (numAffectedRecords int64, err error)
	CountWhere(ctx context.Context, columnFilter *domain.DocumentFilter) (numRecords int64, err error)
	CountAll(ctx context.Context) (numRecords int64, err error)
	DoesExist(ctx context.Context, domainModel *domain.Document) (doesExist bool, err error)
	DoesExistWhere(ctx context.Context, columnFilter *domain.DocumentFilter) (doesExist bool, err error)
	GetWhere(ctx context.Context, columnFilter *domain.DocumentFilter) (records []*domain.Document, err error)
	GetFirstWhere(ctx context.Context, columnFilter *domain.DocumentFilter) (record *domain.Document, err error)
	GetAll(ctx context.Context) (records []*domain.Document, err error)
    
    AddType(ctx context.Context, type_ string) error
    DeleteType(ctx context.Context, type_ string) error
    UpdateType(ctx context.Context, oldType string, newType string) error
    
}
