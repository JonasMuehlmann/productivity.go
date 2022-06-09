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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/model_converter.go.tpl

package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
)

func BookmarkDomainToSqlRepositoryModel(ctx context.Context, db *sql.DB, domainModel *domain.Bookmark) (sqlRepositoryModel *Bookmark, err error) {
	sqlRepositoryModel = new(Bookmark)

	sqlRepositoryModel.URL = domainModel.URL
	sqlRepositoryModel.ID = domainModel.ID

	//**********************    Set Timestamps    **********************//
	// TODO: Extract constant here

	sqlRepositoryModel.CreatedAt = domainModel.CreatedAt
	sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt

	if domainModel.DeletedAt.HasValue {
		sqlRepositoryModel.DeletedAt.Valid = true
		sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt
	}

	//*************************    Set Title    ************************//
	if domainModel.Title.HasValue {
		sqlRepositoryModel.Title.Valid = true
		sqlRepositoryModel.Title.String = domainModel.Title.Wrappee
	}

	//******************    Set IsRead/IsCollection    *****************//

	sqlRepositoryModel.IsRead = domainModel.IsRead
	sqlRepositoryModel.IsCollection = domainModel.IsCollection

	//*************************    Set Tags    *************************//
	var repositoryTag *Tag

	sqlRepositoryModel.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
	for _, domainTag := range domainModel.Tags {
		repositoryTag, err = Tags(TagWhere.Tag.EQ(domainTag.Tag)).One(ctx, db)
		if err != nil {
			return
		}

		sqlRepositoryModel.R.Tags = append(sqlRepositoryModel.R.Tags, &Tag{Tag: repositoryTag.Tag, ID: repositoryTag.ID})
	}

	//*************************    Set Type    *************************//
	if domainModel.BookmarkType.HasValue {
		var repositoryBookmarkType *BookmarkType

		sqlRepositoryModel.R.BookmarkType.Type = domainModel.BookmarkType.Wrappee
		repositoryBookmarkType, err = BookmarkTypes(BookmarkTypeWhere.Type.EQ(domainModel.BookmarkType.Wrappee)).One(ctx, db)
		if err != nil {
			return
		}

		if repositoryBookmarkType != nil {
			sqlRepositoryModel.BookmarkTypeID = null.NewInt64(repositoryBookmarkType.ID, true)
			sqlRepositoryModel.R.BookmarkType.ID = repositoryBookmarkType.ID
		} else {
			sqlRepositoryModel.R.BookmarkType = nil
		}
	}

	return
}

func BookmarkSqlRepositoryToDomainModel(ctx context.Context, db *sql.DB, sqlRepositoryModel *Bookmark) (domainModel *domain.Bookmark, err error) {
	domainModel = new(domain.Bookmark)

	domainModel.URL = sqlRepositoryModel.URL
	domainModel.ID = sqlRepositoryModel.ID
	domainModel.BookmarkType = optional.Make(sqlRepositoryModel.R.BookmarkType.Type)

	//**********************    Set Timestamps    **********************//

	domainModel.CreatedAt = sqlRepositoryModel.CreatedAt
	domainModel.UpdatedAt = sqlRepositoryModel.UpdatedAt

	if sqlRepositoryModel.DeletedAt.Valid {
		domainModel.DeletedAt.Push(sqlRepositoryModel.DeletedAt.Time)
	}

	//*************************    Set Title    ************************//
	if sqlRepositoryModel.Title.Valid {
		domainModel.Title.Push(sqlRepositoryModel.Title.String)
	}

	//******************    Set IsRead/IsCollection    *****************//

	domainModel.IsRead = sqlRepositoryModel.IsRead
	domainModel.IsCollection = sqlRepositoryModel.IsCollection

	//*************************    Set Tags    *************************//
	var domainTag *domain.Tag

	domainModel.Tags = make([]*domain.Tag, 0, len(sqlRepositoryModel.R.Tags))
	for _, repositoryTag := range sqlRepositoryModel.R.Tags {
		domainTag, err = TagSqlRepositoryToDomainModel(db, repositoryTag)
		if err != nil {
			return
		}

		domainModel.Tags = append(domainModel.Tags, domainTag)
	}

	return
}

func DocumentDomainToSqlRepositoryModel(ctx context.Context, db *sql.DB, domainModel *domain.Document) (sqlRepositoryModel *Document, err error) {
	sqlRepositoryModel = new(Document)

	sqlRepositoryModel.Path = domainModel.Path
	sqlRepositoryModel.ID = domainModel.ID

	//**********************    Set Timestamps    **********************//

	sqlRepositoryModel.CreatedAt = domainModel.CreatedAt
	sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt

	if domainModel.DeletedAt.HasValue {
		sqlRepositoryModel.DeletedAt.Valid = true
		sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt
	}

	//*************************    Set Tags    *************************//
	var repositoryTag *Tag

	sqlRepositoryModel.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
	for _, modelTag := range domainModel.Tags {
		repositoryTag, err = Tags(TagWhere.Tag.EQ(modelTag.Tag)).One(ctx, db)
		if err != nil {
			return
		}

		sqlRepositoryModel.R.Tags = append(sqlRepositoryModel.R.Tags, &Tag{Tag: modelTag.Tag, ID: repositoryTag.ID})
	}

	//*************************    Set Type    *************************//
	var repositoryDocumentType *DocumentType

	if domainModel.DocumentType.HasValue {
		sqlRepositoryModel.R.DocumentType.DocumentType = domainModel.DocumentType.Wrappee
		repositoryDocumentType, err = DocumentTypes(DocumentTypeWhere.DocumentType.EQ(domainModel.DocumentType.Wrappee)).One(ctx, db)
		if err != nil {
			return
		}

		if repositoryDocumentType != nil {
			sqlRepositoryModel.DocumentTypeID = null.NewInt64(repositoryDocumentType.ID, true)
			sqlRepositoryModel.R.DocumentType.ID = repositoryDocumentType.ID
		} else {
			sqlRepositoryModel.R.DocumentType = nil
		}
	}

	//**************    Set linked/backlinked documents    *************//
	var repositoryDocument *Document

	sqlRepositoryModel.R.SourceDocuments = make(DocumentSlice, 0, len(domainModel.LinkedDocuments))
	sqlRepositoryModel.R.DestinationDocuments = make(DocumentSlice, 0, len(domainModel.BacklinkedDocuments))

	for _, link := range domainModel.LinkedDocuments {
		repositoryDocument, err = DocumentDomainToSqlRepositoryModel(ctx, db, link)
		if err != nil {
			return
		}

		sqlRepositoryModel.R.SourceDocuments = append(sqlRepositoryModel.R.SourceDocuments, repositoryDocument)
	}

	for _, backlink := range domainModel.BacklinkedDocuments {
		repositoryDocument, err = DocumentDomainToSqlRepositoryModel(ctx, db, backlink)
		if err != nil {
			return
		}

		sqlRepositoryModel.R.DestinationDocuments = append(sqlRepositoryModel.R.DestinationDocuments, repositoryDocument)
	}

	return
}

func DocumentSqlRepositoryToDomainModel(ctx context.Context, db *sql.DB, sqlRepositoryModel *Document) (domainModel *domain.Document, err error) {
	domainModel = new(domain.Document)

	domainModel.Path = sqlRepositoryModel.Path
	domainModel.ID = sqlRepositoryModel.ID
	domainModel.DocumentType = optional.Make(sqlRepositoryModel.R.DocumentType.DocumentType)

	//**********************    Set Timestamps    **********************//

	domainModel.CreatedAt = sqlRepositoryModel.CreatedAt
	domainModel.UpdatedAt = sqlRepositoryModel.UpdatedAt

	if sqlRepositoryModel.DeletedAt.Valid {
		domainModel.DeletedAt.Push(sqlRepositoryModel.DeletedAt.Time)
	}

	//*************************    Set Tags    *************************//
	var domainTag *domain.Tag

	domainModel.Tags = make([]*domain.Tag, 0, len(sqlRepositoryModel.R.Tags))
	for _, repositoryTag := range sqlRepositoryModel.R.Tags {
		domainTag, err = TagSqlRepositoryToDomainModel(ctx, db, repositoryTag)
		if err != nil {
			return
		}

		domainModel.Tags = append(domainModel.Tags, domainTag)
	}

	//**************    Set linked/backlinked documents    *************//
	var domainDocument *domain.Document

	domainModel.LinkedDocuments = make([]*domain.Document, 0, len(sqlRepositoryModel.R.SourceDocuments))
	domainModel.BacklinkedDocuments = make([]*domain.Document, 0, len(sqlRepositoryModel.R.DestinationDocuments))

	for _, link := range sqlRepositoryModel.R.SourceDocuments {
		domainDocument, err = DocumentSqlRepositoryToDomainModel(ctx, db, link)
		if err != nil {
			return
		}

		domainModel.LinkedDocuments = append(domainModel.LinkedDocuments, domainDocument)
	}

	for _, backlink := range sqlRepositoryModel.R.DestinationDocuments {
		domainDocument, err = DocumentSqlRepositoryToDomainModel(ctx, db, backlink)
		if err != nil {
			return
		}

		domainModel.BacklinkedDocuments = append(domainModel.BacklinkedDocuments, domainDocument)
	}

	return
}

func TagDomainToSqlRepositoryModel(ctx context.Context, db *sql.DB, domainModel *domain.Tag) (sqlRepositoryModel *Tag, err error) {
	// TODO: make sure to insert all tags in ParentPath and Subtags into db
	sqlRepositoryModel = new(Tag)

	sqlRepositoryModel.ID = domainModel.ID
	sqlRepositoryModel.Tag = domainModel.Tag

	//***********************    Set ParentTag    **********************//
	if len(domainModel.ParentPath) > 0 {
		sqlRepositoryModel.ParentTag = null.NewInt64(domainModel.ParentPath[len(domainModel.ParentPath)-1].ID, true)
	}

	//*************************    Set Path    *************************//
	for _, tag := range domainModel.ParentPath {
		sqlRepositoryModel.Path += strconv.FormatInt(tag.ID, 10) + ";"
	}

	sqlRepositoryModel.Path += strconv.FormatInt(domainModel.ID, 10)

	//************************    Set Children  ************************//
	for _, tag := range domainModel.Subtags {
		sqlRepositoryModel.Children += strconv.FormatInt(tag.ID, 10) + ";"
	}

	return
}

// TODO: These functions should be context aware
func TagSqlRepositoryToDomainModel(ctx context.Context, db *sql.DB, sqlRepositoryModel *Tag) (domainModel *domain.Tag, err error) {
	// TODO: make sure to insert all tags in ParentPath and Subtags into db
	domainModel = new(domain.Tag)

	domainModel.ID = sqlRepositoryModel.ID
	domainModel.Tag = sqlRepositoryModel.Tag

	//***********************    Set ParentPath    **********************//
	var parentTagID int64
	var parentTag *Tag
	var domainParentTag *domain.Tag

	for _, parentTagIDRaw := range strings.Split(sqlRepositoryModel.Path, ";")[:len(sqlRepositoryModel.Path)-2] {
		parentTagID, err = strconv.ParseInt(parentTagIDRaw, 10, 64)
		if err != nil {
			return
		}

		parentTag, err = Tags(TagWhere.ID.EQ(parentTagID)).One(ctx, db)
		if err != nil {
			return
		}

		domainParentTag, err = TagSqlRepositoryToDomainModel(ctx, db, parentTag)
		if err != nil {
			return
		}

		domainModel.ParentPath = append(domainModel.ParentPath, domainParentTag)
	}

	//************************    Set Subtags ************************//
	var childTagID int64
	var childTag *Tag
	var domainChildTag *domain.Tag

	for _, childTagIDRaw := range strings.Split(sqlRepositoryModel.Children, ";")[:len(sqlRepositoryModel.Children)-2] {
		childTagID, err = strconv.ParseInt(childTagIDRaw, 10, 64)
		if err != nil {
			return
		}

		childTag, err = Tags(TagWhere.ID.EQ(childTagID)).One(ctx, db)
		if err != nil {
			return
		}

		domainChildTag, err = TagSqlRepositoryToDomainModel(ctx, db, childTag)
		if err != nil {
			return
		}

		domainModel.Subtags = append(domainModel.Subtags, domainChildTag)
	}

	return
}
