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

type Sqlite3TagRepository struct {
    db *sql.DB
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
    ID optional.Optional[model.FilterOperation[int64]]
    Tag optional.Optional[model.FilterOperation[string]]
    
    Bookmarks optional.Optional[model.UpdateOperation[BookmarkSlice]]
    ParentTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    ChildTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    Documents optional.Optional[model.UpdateOperation[DocumentSlice]]
    ParentTagTagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    TagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    
}

type TagFilterMapping[T any] struct {
    Field TagField
    FilterOperation model.FilterOperation[T]
}

func (filter *TagFilter) GetSetFilters() *list.List {
    setFilters := list.New()

    if filter.ID.HasValue {
        setFilters.PushBack(TagFilterMapping[int64]{TagFields.ID, filter.ID.Wrappee})
    }
    if filter.Tag.HasValue {
        setFilters.PushBack(TagFilterMapping[string]{TagFields.Tag, filter.Tag.Wrappee})
    }
    

    return setFilters
}

type TagUpdater struct {
    ID optional.Optional[model.UpdateOperation[int64]]
    Tag optional.Optional[model.UpdateOperation[string]]
    
    Bookmarks optional.Optional[model.UpdateOperation[BookmarkSlice]]
    ParentTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    ChildTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    Documents optional.Optional[model.UpdateOperation[DocumentSlice]]
    ParentTagTagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    TagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    
}

type TagUpdaterMapping[T any] struct {
    Field TagField
    Updater model.UpdateOperation[T]
}

func (updater *TagUpdater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    if updater.ID.HasValue {
        setUpdaters.PushBack(TagUpdaterMapping[int64]{TagFields.ID, updater.ID.Wrappee})
    }
    if updater.Tag.HasValue {
        setUpdaters.PushBack(TagUpdaterMapping[string]{TagFields.Tag, updater.Tag.Wrappee})
    }
    

    return setUpdaters
}

func (updater *TagUpdater) ApplyToModel(tagModel *Tag) {
    if updater.ID.HasValue {
        model.ApplyUpdater(&(*tagModel).ID, updater.ID.Wrappee)
    }
    if updater.Tag.HasValue {
        model.ApplyUpdater(&(*tagModel).Tag, updater.Tag.Wrappee)
    }
    
}

type Sqlite3TagRepositoryHook func(context.Context, Sqlite3TagRepository) error

type queryModSlice []qm.QueryMod

func (s queryModSlice) Apply(q *queries.Query) {
    qm.Apply(q, s...)
}

func buildQueryModFilter[T any](filterField TagField, filterOperation model.FilterOperation[T]) queryModSlice {
    var newQueryMod queryModSlice

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
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilter(filterField, filterOperand.RHS))))
    case model.FilterAnd:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterAnd operator")
        }

        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilter(filterField, filterOperand.RHS)))
    default:
        panic("Unhandled FilterOperator")
    }

    return newQueryMod
}

func buildQueryModListFromFilter(setFilters list.List) queryModSlice {
	queryModList := make(queryModSlice, 0, 2)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(TagFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %t but got %t", TagFilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilter(filterMapping.Field, filterMapping.FilterOperation)

        for _, queryMod := range newQueryMod {
            queryModList = append(queryModList, queryMod)
        }
	}

	return queryModList
}

func (repo * Sqlite3TagRepository) New(args ...any) (Sqlite3TagRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) Add(ctx context.Context, domainModels []domain.Tag) (numAffectedRecords int, newID int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) Replace(ctx context.Context, domainModels []domain.Tag) error {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) UpdateWhere(ctx context.Context, columnFilter domain.TagFilter, columnUpdaters map[domain.TagField]domain.TagUpdater) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) Delete(ctx context.Context, domainModels []domain.Tag) error {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) DeleteWhere(ctx context.Context, columnFilter domain.TagFilter) (numAffectedRecords int, err error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) CountWhere(ctx context.Context, columnFilter domain.TagFilter) int {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) CountAll(ctx context.Context) int {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) DoesExist(ctx context.Context, domainModel domain.Tag) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) DoesExistWhere(ctx context.Context, columnFilter domain.TagFilter) bool {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) GetWhere(ctx context.Context, columnFilter domain.TagFilter) []domain.Tag {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) GetFirstWhere(ctx context.Context, columnFilter domain.TagFilter) domain.Tag {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) GetAll(ctx context.Context) []domain.Tag {
        panic("not implemented") // TODO: Implement
}
