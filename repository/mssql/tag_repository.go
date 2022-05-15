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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/Tag_repository.go.tpl

package repository

type MssqlTagRepository struct {
    db sql.DB
}
type TagField string

var TagFields = struct {
    ID  TagField
    Tag  TagField
    
}{
    ID: "id",
    Tag: "tag",
    
}

var TagFieldsList = []TagField{
    TagField("ID"),
    TagField("Tag"),
    
}

var TagRelationsList = []string{
    "Bookmarks",
    "ParentTagTags",
    "ChildTagTags",
    "Documents",
    "ParentTagTagParentPaths",
    "TagParentPaths",
    
}

type TagFilter struct {
    ID optional.Optional[FilterOperation[int64]]
    Tag optional.Optional[FilterOperation[string]]
    
    Bookmarks optional.Optional[UpdateOperation[repository.BookmarkSlice]]
    ParentTagTags optional.Optional[UpdateOperation[repository.TagSlice]]
    ChildTagTags optional.Optional[UpdateOperation[repository.TagSlice]]
    Documents optional.Optional[UpdateOperation[repository.DocumentSlice]]
    ParentTagTagParentPaths optional.Optional[UpdateOperation[repository.TagParentPathSlice]]
    TagParentPaths optional.Optional[UpdateOperation[repository.TagParentPathSlice]]
    
}

type TagUpdater struct {
    ID optional.Optional[UpdateOperation[int64]]
    Tag optional.Optional[UpdateOperation[string]]
    
    Bookmarks optional.Optional[UpdateOperation[repository.BookmarkSlice]]
    ParentTagTags optional.Optional[UpdateOperation[repository.TagSlice]]
    ChildTagTags optional.Optional[UpdateOperation[repository.TagSlice]]
    Documents optional.Optional[UpdateOperation[repository.DocumentSlice]]
    ParentTagTagParentPaths optional.Optional[UpdateOperation[repository.TagParentPathSlice]]
    TagParentPaths optional.Optional[UpdateOperation[repository.TagParentPathSlice]]
    
}

type MssqlTagRepositoryHook func(context.Context, MssqlTagRepository) error

func (repo * MssqlTagRepository) New(args ...any) (MssqlTagRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) Add(ctx context.Context, domainModels []domain.Tag) (numAffectedRecords int, newID int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) Replace(ctx context.Context, domainModels []domain.Tag) error {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) UpdateWhere(ctx context.Context, columnFilter domain.TagFilter, columnUpdaters map[domain.TagField]domain.TagUpdater) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) Delete(ctx context.Context, domainModels []domain.Tag) error {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) DeleteWhere(ctx context.Context, columnFilter domain.TagFilter) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) CountWhere(ctx context.Context, columnFilter domain.TagFilter) int {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) CountAll(ctx context.Context) int {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) DoesExist(ctx context.Context, domainModel domain.Tag) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) DoesExistWhere(ctx context.Context, columnFilter domain.TagFilter) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) GetWhere(ctx context.Context, columnFilter domain.TagFilter) []domain.Tag {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) GetFirstWhere(ctx context.Context, columnFilter domain.TagFilter) domain.Tag {
        panic("not implemented") // TODO: Implement
}

func (repo *MssqlTagRepository) GetAll(ctx context.Context) []domain.Tag {
        panic("not implemented") // TODO: Implement
}
