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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/{{.EntityName}}_repository.go.tpl

package repository

import (
    "database/sql"
    "github.com/JonasMuehlmann/bntp.go/repository/{{.DatabaseName}}/{{LowercaseBeginning (Pluralize .EntityName)}}.go"
)

{{template "structDefinition" .}}
{{template "repositoryHelperTypes" .}}

func (repo * {{$StructName}}) New(_ ...any) (repository.BookmarkRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) Add(_ context.Context, _ []domain.Bookmark) (numAffectedRecords int, newID int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) Replace(_ context.Context, _ []domain.Bookmark) error {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) UpdateWhere(_ context.Context, _ domain.BookmarkFilter, _ map[domain.BookmarkField]domain.BookmarkUpdateOperation) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) Delete(_ context.Context, _ []domain.Bookmark) error {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) DeleteWhere(_ context.Context, _ domain.BookmarkFilter) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) CountWhere(_ context.Context, _ domain.BookmarkFilter) int {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) CountAll(_ context.Context) int {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) DoesExist(_ context.Context, _ domain.Bookmark) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) DoesExistWhere(_ context.Context, _ domain.BookmarkFilter) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) GetWhere(_ context.Context, _ domain.BookmarkFilter) []domain.Bookmark {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) GetFirstWhere(_ context.Context, _ domain.BookmarkFilter) domain.Bookmark {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) GetAll(_ context.Context) []domain.Bookmark {
        panic("not implemented") // TODO: Implement
}
