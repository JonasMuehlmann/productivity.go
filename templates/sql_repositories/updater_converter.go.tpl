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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/updater_converter.go.tpl

package repository

import (
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	 repoCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
    "github.com/volatiletech/null/v8"
    "database/sql"
    "context"
)

func BookmarkDomainToSqlRepositoryUpdater(ctx context.Context, db *sql.DB, domainUpdater *domain.{{.Entities.Bookmark}}Updater) (sqlRepositoryUpdater *{{.Entities.Bookmark}}Updater, err error)  {
    sqlRepositoryUpdater = new({{.Entities.Bookmark}}Updater)

	if domainUpdater.CreatedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater string
        convertedUpdater, err = repoCommon.TimeToStr(domainUpdater.CreatedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.CreatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.CreatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.CreatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.CreatedAt.Wrappee.Operator, Operand: domainUpdater.CreatedAt.Wrappee.Operand})
        {{ end}}
    }

	if domainUpdater.UpdatedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater string
        convertedUpdater, err = repoCommon.TimeToStr(domainUpdater.UpdatedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.UpdatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.UpdatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: domainUpdater.UpdatedAt.Wrappee.Operand})
        {{ end }}
    }

	if domainUpdater.DeletedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater null.String
        convertedUpdater, err = repoCommon.OptionalTimeToNullStr(domainUpdater.DeletedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.DeletedAt.Push(model.UpdateOperation[null.String]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.DeletedAt.Push(model.UpdateOperation[null.Time]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: domainUpdater.DeletedAt.Wrappee.Operand})
        {{ end }}
    }

	if domainUpdater.URL.HasValue {
        sqlRepositoryUpdater.URL.Push(model.UpdateOperation[string]{Operator: domainUpdater.URL.Wrappee.Operator, Operand: sqlRepositoryUpdater.URL.Wrappee.Operand})
    }

	if domainUpdater.Title.HasValue {
        var convertedUpdater null.String
        convertedUpdater, err = repoCommon.OptionalStringToNullString(domainUpdater.Title.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.DeletedAt.Push(model.UpdateOperation[null.String]{Operator: domainUpdater.Title.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.Tags.HasValue {
        var convertedTag *Tag
        convertedUpdater := make(TagSlice, 0, len(domainUpdater.Tags.Wrappee.Operand))

        for _, tag := range domainUpdater.Tags.Wrappee.Operand {
            convertedTag, err = TagDomainToSqlRepositoryModel(db, tag)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedTag)
        }

        sqlRepositoryUpdater.Tags.Push(model.UpdateOperation[TagSlice]{Operator: domainUpdater.Tags.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.ID.HasValue {
        sqlRepositoryUpdater.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: sqlRepositoryUpdater.ID.Wrappee.Operand})
    }

	if domainUpdater.IsCollection.HasValue {
        var convertedUpdater int64
        convertedUpdater, err = repoCommon.BoolToInt(domainUpdater.IsCollection.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.IsCollection.Push(model.UpdateOperation[int64]{Operator: domainUpdater.IsCollection.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.IsRead.HasValue {
        var convertedUpdater int64
        convertedUpdater, err = repoCommon.BoolToInt(domainUpdater.IsRead.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.IsRead.Push(model.UpdateOperation[int64]{Operator: domainUpdater.IsRead.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.BookmarkType.HasValue {
        var convertedUpdater *BookmarkType
        if domainUpdater.BookmarkType.Wrappee.Operand.HasValue {
            convertedUpdater, err = BookmarkTypes(BookmarkTypeWhere.Type.EQ(domainUpdater.BookmarkType.Wrappee.Operand.Wrappee)).One(context.Background(), db)
            if err != nil {
                return
            }
        } else {
            convertedUpdater = nil
        }

        sqlRepositoryUpdater.BookmarkType.Push(model.UpdateOperation[*BookmarkType]{Operator: domainUpdater.BookmarkType.Wrappee.Operator, Operand: convertedUpdater})
    }


    return

}

func DocumentDomainToSqlRepositoryUpdater(ctx context.Context, db *sql.DB, domainUpdater *domain.{{.Entities.Document}}Updater) (sqlRepositoryUpdater *{{.Entities.Document}}Updater, err error)  {
    sqlRepositoryUpdater = new({{.Entities.Document}}Updater)

	if domainUpdater.DocumentType.HasValue {
        var convertedUpdater *DocumentType
        if domainUpdater.DocumentType.Wrappee.Operand.HasValue {
            convertedUpdater, err = DocumentTypes(DocumentTypeWhere.DocumentType.EQ(domainUpdater.DocumentType.Wrappee.Operand.Wrappee)).One(context.Background(), db)
            if err != nil {
                return
		}
        } else {
            convertedUpdater = nil
        }

        sqlRepositoryUpdater.DocumentType.Push(model.UpdateOperation[*DocumentType]{Operator: domainUpdater.DocumentType.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.Path.HasValue {
        sqlRepositoryUpdater.Path.Push(model.UpdateOperation[string]{Operator: domainUpdater.Path.Wrappee.Operator, Operand: sqlRepositoryUpdater.Path.Wrappee.Operand})
    }

	if domainUpdater.CreatedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater string
        convertedUpdater, err = repoCommon.TimeToStr(domainUpdater.CreatedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.CreatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.CreatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.CreatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.CreatedAt.Wrappee.Operator, Operand: domainUpdater.CreatedAt.Wrappee.Operand})
        {{ end}}
    }

	if domainUpdater.UpdatedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater string
        convertedUpdater, err = repoCommon.TimeToStr(domainUpdater.UpdatedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.UpdatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.UpdatedAt.Push(model.UpdateOperation[string]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: domainUpdater.UpdatedAt.Wrappee.Operand})
        {{ end }}
    }

	if domainUpdater.DeletedAt.HasValue {
        {{ if eq .DatabaseName "sqlite3" }}
        var convertedUpdater null.String
        convertedUpdater, err = repoCommon.OptionalTimeToNullStr(domainUpdater.DeletedAt.Wrappee.Operand)
        if err != nil {
            return
        }

        sqlRepositoryUpdater.DeletedAt.Push(model.UpdateOperation[null.String]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: convertedUpdater})
        {{ else }}
        sqlRepositoryUpdater.DeletedAt.Push(model.UpdateOperation[null.Time]{Operator: domainUpdater.UpdatedAt.Wrappee.Operator, Operand: domainUpdater.DeletedAt.Wrappee.Operand})
        {{ end }}
    }

	if domainUpdater.Tags.HasValue {
        var convertedTag *Tag
        convertedUpdater := make(TagSlice, 0, len(domainUpdater.Tags.Wrappee.Operand))

        for _, tag := range domainUpdater.Tags.Wrappee.Operand {
            convertedTag, err = TagDomainToSqlRepositoryModel(db, tag)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedTag)
        }

        sqlRepositoryUpdater.Tags.Push(model.UpdateOperation[TagSlice]{Operator: domainUpdater.Tags.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.LinkedDocuments.HasValue {
        var convertedDocument *Document
        convertedUpdater := make(DocumentSlice, 0, len(domainUpdater.LinkedDocuments.Wrappee.Operand))

        for _, document := range domainUpdater.LinkedDocuments.Wrappee.Operand {
            convertedDocument, err =DocumentDomainToSqlRepositoryModel(db, document)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedDocument)
        }

        sqlRepositoryUpdater.DestinationDocuments.Push(model.UpdateOperation[DocumentSlice]{Operator: domainUpdater.LinkedDocuments.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.BacklinkedDocuments.HasValue {
        var convertedDocument *Document
        convertedUpdater := make(DocumentSlice, 0, len(domainUpdater.BacklinkedDocuments.Wrappee.Operand))

        for _, document := range domainUpdater.BacklinkedDocuments.Wrappee.Operand {
            convertedDocument, err =DocumentDomainToSqlRepositoryModel(db, document)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedDocument)
        }

        sqlRepositoryUpdater.SourceDocuments.Push(model.UpdateOperation[DocumentSlice]{Operator: domainUpdater.BacklinkedDocuments.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.ID.HasValue {
        sqlRepositoryUpdater.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: sqlRepositoryUpdater.ID.Wrappee.Operand})
    }

    return
}

func TagDomainToSqlRepositoryUpdater(ctx context.Context, db *sql.DB, domainUpdater *domain.{{.Entities.Tag}}Updater) (sqlRepositoryUpdater *{{.Entities.Tag}}Updater, err error)  {
    sqlRepositoryUpdater = new({{.Entities.Tag}}Updater)

	if domainUpdater.Tag.HasValue {
        sqlRepositoryUpdater.Tag.Push(model.UpdateOperation[string]{Operator: domainUpdater.Tag.Wrappee.Operator, Operand: sqlRepositoryUpdater.Tag.Wrappee.Operand})
    }

	if domainUpdater.ParentPath.HasValue {
        var convertedTag *Tag
        convertedUpdater := make(TagSlice, 0, len(domainUpdater.ParentPath.Wrappee.Operand))

        for _, tag := range domainUpdater.ParentPath.Wrappee.Operand {
            convertedTag, err = TagDomainToSqlRepositoryModel(db, tag)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedTag)
        }

        sqlRepositoryUpdater.ParentTagTags.Push(model.UpdateOperation[TagSlice]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.Subtags.HasValue {
        var convertedTag *Tag
        convertedUpdater := make(TagSlice, 0, len(domainUpdater.Subtags.Wrappee.Operand))

        for _, tag := range domainUpdater.Subtags.Wrappee.Operand {
            convertedTag, err = TagDomainToSqlRepositoryModel(db, tag)
            if err != nil {
                return
            }

            convertedUpdater = append(convertedUpdater, convertedTag)
        }

        sqlRepositoryUpdater.ChildTagTags.Push(model.UpdateOperation[TagSlice]{Operator: domainUpdater.Subtags.Wrappee.Operator, Operand: convertedUpdater})
    }

	if domainUpdater.ID.HasValue {
        sqlRepositoryUpdater.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: sqlRepositoryUpdater.ID.Wrappee.Operand})
    }

    return
}
