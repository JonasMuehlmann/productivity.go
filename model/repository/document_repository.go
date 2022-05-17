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
	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
)

type DocumentRepository interface {
	New(...any) (DocumentRepository, error)

	Add(ctx context.Context, domainModels []domain.Document) (numAffectedRecords int, newID int, err error)
	Replace(ctx context.Context, domainModels []domain.Document) error
	UpdateWhere(ctx context.Context, columnFilter domain.DocumentFilter, columnUpdaters map[domain.DocumentField]domain.DocumentUpdater) (numAffectedRecords int, err error)
	Delete(ctx context.Context, domainModels []domain.Document) error
	DeleteWhere(ctx context.Context, columnFilter domain.DocumentFilter) (numAffectedRecords int, err error)
	CountWhere(ctx context.Context, columnFilter domain.DocumentFilter) int
	CountAll(ctx context.Context) int
	DoesExist(ctx context.Context, domainModel domain.Document) bool
	DoesExistWhere(ctx context.Context, columnFilter domain.DocumentFilter) bool
	GetWhere(ctx context.Context, columnFilter domain.DocumentFilter) []domain.Document
	GetFirstWhere(ctx context.Context, columnFilter domain.DocumentFilter) domain.Document
	GetAll(ctx context.Context) []domain.Document
}
