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

// Package libbookmarks implements functionality to work with bookmarks in a database context.
package libbookmarks

import (
	"errors"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gocarina/gocsv"
)

func DeserializeBookmarks(serializedBookmarkList string) (bookmarkList []Bookmark, err error) {
	bookmarkList = make([]Bookmark, 0, 100)

	err = gocsv.UnmarshalString(serializedBookmarkList, &bookmarkList)
	if err != nil {
		err = helpers.DeserializationError{Inner: err}

		return
	}

	if len(bookmarkList) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Empty bookmark list")}

		return
	}

	return
}

func ImportBookmarks(dbConn *sqlx.DB, bookmarkList []Bookmark) error {
	transaction, err := dbConn.Beginx()
	if err != nil {
		return err
	}

	for _, bookmark := range bookmarkList {
		err := AddBookmark(nil, transaction, bookmark)

		if err != nil {
			return err
		}
	}

	err = transaction.Commit()

	if err != nil {
		return err
	}

	return nil
}

func SerializeBookmarks(bookmarkList []Bookmark) (serializedBookmarks string, err error) {
	if len(bookmarkList) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Empty bookmark list")}

		return
	}

	serializedBookmarks, err = gocsv.MarshalString(bookmarkList)

	return
}

func ExportBookmarks(dbConn *sqlx.DB) (bookmarkList []Bookmark, err error) {
	return GetBookmarks(dbConn, BookmarkFilter{})
}

// GetBookmarks returns all bookmarks stored in the DB which satisfy the given filter.
func GetBookmarks(dbConn *sqlx.DB, filter BookmarkFilter) ([]Bookmark, error) {
	stmtBookmarks := `
        SELECT
            Bookmark.Id AS Id,
            Bookmark.IsRead AS IsRead,
            Bookmark.Title AS Title,
            Bookmark.Url AS Url,
            Bookmark.TimeAdded AS TimeAdded,
            BookmarkType.Type AS Type,
            Bookmark.IsCollection AS IsCollection
        FROM
            Bookmark
        LEFT JOIN BookmarkType ON
            BookmarkType.Id = Bookmark.BookmarkTypeId
    `

	stmtBookmarks = ApplyBookmarkFilters(stmtBookmarks, filter)

	stmtTags := `
        SELECT
            Tag.Tag
        FROM
            Tag
        INNER JOIN BookmarkContext ON
            BookmarkContext.TagId = Tag.Id
        WHERE
            BookmarkContext.BookmarkId = ?;
    `

	bookmarksBuffer := []Bookmark{}

	err := dbConn.Select(&bookmarksBuffer, stmtBookmarks)

	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarksBuffer {
		bookmark.Tags = []string{}

		err = dbConn.Select(&bookmark.Tags, stmtTags, bookmark.Id)

		if err != nil {
			return nil, err
		}
	}

	return bookmarksBuffer, nil
}

// TODO: Allow passing string and id for type_

// AddType makes a new BookmarkType available for use in the DB.
// Passing a transaction is optional.
func AddType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	if type_ == "" {
		return errors.New("Can't add empty tag")
	}

	stmt := `
        INSERT INTO
            BookmarkType(
                Type
            )
        VALUES(
            ?
        );
    `
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(type_)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for type_

// RemoveType removes an available BookmarkType from the DB.
// Passing a transaction is optional.
func RemoveType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	stmt := `
        DELETE FROM
            BookmarkType
        WHERE
            Type = ?;
    `
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	result, err := statement.Exec(type_)
	if err != nil {
		return err
	}

	numAffectedRows, err := result.RowsAffected()
	if numAffectedRows == 0 || err != nil {
		return errors.New("Type to be deleted does not exist")
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// ListTypes lists all BookmarkTypes available in the DB.
func ListTypes(dbConn *sqlx.DB) ([]string, error) {
	stmtTypes := `
        SELECT
            Type
        FROM
            BookmarkType;
    `

	stmtNumTypes := `
        SELECT
            Count(*)
        FROM
            BookmarkType;
    `

	var numTypes int

	err := dbConn.Get(&numTypes, stmtNumTypes)
	if err != nil {
		return nil, err
	}

	types := make([]string, 0, numTypes)

	err = dbConn.Select(&types, stmtTypes)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// TODO: Allow passing string and id for type_

// AddBookmark adds a new bookmark to the DB.
// Passing a transaction is optional.
func AddBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, title string, url string, type_ optional.Optional[int]) error {
	stmt := `
        INSERT INTO
            Bookmark(
                Title,
                Url,
                TimeAdded,
                BookmarkTypeId
            )
        VALUES(
            ?,
            ?,
            ?,
            ?
        );
    `
	if title == "" || url == "" {
		return errors.New("Entry is missing a column")
	}

	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(title, url, time.Now().Format("2006-01-02"), type_)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

func RemoveBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, ID int) error {
	stmt := `
        DELETE FROM
            Bookmark
        WHERE
            ID = ?;
    `
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	result, err := statement.Exec(ID)
	if err != nil {
		return err
	}

	numAffectedRows, err := result.RowsAffected()
	if numAffectedRows == 0 || err != nil {
		return errors.New("Bookmark to be deleted does not exist")
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil

}

func EditBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, newData Bookmark) error {
	// TODO: Handle change of tags
	stmt := `
        UPDATE
            Bookmark
        SET
            Title = ?,
            Url = ?,
            TimeAdded = ?,
            BookmarkTypeId = ?,
            IsRead = ?,
            IsCollection = ?
        WHERE
            Id = ?;
    `

	var typeID optional.Optional[int]
	if newData.Type.HasValue {
		typeIDUnwrapped, err := helpers.GetIdFromBookmarkType(dbConn, transaction, newData.Type.Wrappee)
		if err != nil {
			return err
		}

		typeID.Wrappee = typeIDUnwrapped
	}

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, newData.Title, newData.Url, newData.TimeAdded, typeID, newData.IsRead, newData.IsCollection, newData.Id)

	return err
}

// editBookmarkField sets column to newVal for the bookmark with the specified id.
// Passing a transaction is optional.
func editBookmarkField(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, column string, newVal interface{}) error {
	stmt := `
        UPDATE
            Bookmark
        SET
            ` + column + ` = ?
        WHERE
            Id = ?;
    `

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, newVal, id)

	return err
}

// EditIsRead sets IsRead to true for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsRead(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, isRead bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsRead", isRead)
}

// EditTitle sets Title to newTile for the bookmark with the specified id.
// Passing a transaction is optional.
func EditTitle(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newTitle string) error {
	return editBookmarkField(dbConn, transaction, id, "Title", newTitle)
}

// EditUrl sets Url to newUrl for the bookmark with the specified id.
// Passing a transaction is optional.
func EditUrl(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newUrl string) error {
	return editBookmarkField(dbConn, transaction, id, "Url", newUrl)
}

// TODO: Allow passing string and id for type_

// EditType sets Type to newType for the bookmark with the specified id.
// Passing a transaction is optional.
func EditType(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newType string) error {
	typeId, err := helpers.GetIdFromType(dbConn, transaction, newType)
	if err != nil {
		return err
	}
	return editBookmarkField(dbConn, transaction, id, "BookmarkTypeId", typeId)
}

// EditIsCollection sets isCollection to isCollection for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsCollection(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, isCollection bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsCollection", isCollection)
}

// AddTag adds a tag newTag to the bookmark with bookmarkId.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmarkId int, newTag string) error {
	stmt := `
        INSERT INTO
            BookmarkContext(BookmarkId, TagId)
        VALUES(
            ?,
            ?
    );
    `

	var statementContext *sqlx.Stmt
	var err error

	if transaction != nil {
		statementContext, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statementContext, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	tagId, err := helpers.GetIdFromTag(dbConn, transaction, newTag)
	if err != nil {
		return err
	}

	result, err := statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	numAffectedRows, err := result.RowsAffected()
	if numAffectedRows == 0 || err != nil {
		return errors.New("Type to be deleted does not exist")
	}

	err = statementContext.Close()

	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for tag_

// RemoveTag removes a tag tag_ from the bookmark with bookmarkId.
// Passing a transaction is optional.
func RemoveTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmarkId int, tag_ string) error {
	stmt := `
        DELETE FROM
            BookmarkContext
        WHERE
            BookmarkId = ?
            AND
            TagId = ?;
    );
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	tagId, err := helpers.GetIdFromTag(dbConn, transaction, tag_)
	if err != nil {
		return err
	}

	_, err = statement.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}
