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

package subcommands

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
)

var usageBookmark string = `bntp bookmark - Interact with bookmarks.

Usage:
    bntp bookmark -h | --help
    bntp bookmark -l [--filter FILTER]
    bntp bookmark --import FILE
    bntp bookmark --export FILE [--filter FILTER]
    bntp bookmark (--add-type | --remove-type) TYPE
    bntp bookmark --list-types
    bntp bookmark -a DATA
    bntp bookmark -r BOOKMARK_ID
    bntp bookmark -e NEW_DATA
    bntp bookmark --edit-is-read BOOKMARK_ID IS_READ
    bntp bookmark --edit-title BOOKMARK_ID TITLE
    bntp bookmark --edit-url BOOKMARK_ID TITLE
    bntp bookmark --edit-type BOOKMARK_ID TITLE
    bntp bookmark --edit-is-collection BOOKMARK_ID IS_COLLECTION
    bntp bookmark (--add-tag | --remove-tag) BOOKMARK_ID TAG

Options:
    -h --help               Show this screen.
    -i --import             Import bookmarks.
    -e --export             Export bookmarks.
    -l --list               List bookmarks.
    --filter                A filter to apply when searching for bookmarks.
    --add-type              Add a bookmark type.
    --remove-type           Remove a bookmark type.
    --list-types            List the bookmark types.
    -a --add                Add a bookmark.
    -r --remove             Remove a bookmark.
    -e --edit               Edit a bookmark.
    --edit-is-read          Change the is read status of a bookmark.
    --edit-title            Change the title of a bookmark.
    --edit-url              Change the url of a bookmark.
    --edit-type             Change the type of a bookmark.
    --edit-is-collection    Change the is collection status of a bookmark.
    --add-tag               Add a tag to a bookmark.
    --remove-tag            Remove a tag from a bookmark.
`

func BookmarkMain(db *sqlx.DB) {
	arguments, err := docopt.ParseDoc(usageBookmark)
	helpers.OnError(err, log.Panic)

	// ******************************************************************//
	if _, ok := arguments["--import"]; ok {
		source, err := arguments.String("FILE")
		helpers.OnError(err, log.Panic)

		err = libbookmarks.ImportMinimalCSV(db, source)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--export"]; ok {
		source, err := arguments.String("FILE")
		helpers.OnError(err, log.Panic)

		filter := libbookmarks.BookmarkFilter{}
		if _, ok := arguments["--filter"]; ok {
			filterRaw, err := arguments.String("FILTER")
			helpers.OnError(err, log.Panic)

			err = json.Unmarshal([]byte(filterRaw), &filter)
			helpers.OnError(err, log.Panic)
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.ExportCSV(bookmarks, source)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--list"]; ok {
		filter := libbookmarks.BookmarkFilter{}
		if _, ok := arguments["--filter"]; ok {
			filterRaw, err := arguments.String("FILTER")
			helpers.OnError(err, log.Panic)

			err = json.Unmarshal([]byte(filterRaw), &filter)
			helpers.OnError(err, log.Panic)
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		helpers.OnError(err, log.Panic)

		for bookmark := range bookmarks {
			println(bookmark)
		}
		// ******************************************************************//
	} else if _, ok := arguments["--add-type"]; ok {
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, log.Panic)

		err = libbookmarks.AddType(db, nil, type_)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--remove-type"]; ok {
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, log.Panic)

		err = libbookmarks.RemoveType(db, nil, type_)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--list-types"]; ok {
		types, err := libbookmarks.ListTypes(db)
		helpers.OnError(err, log.Panic)

		for type_ := range types {
			println(type_)
		}
		// ******************************************************************//
	} else if _, ok := arguments["--add"]; ok {
		var data map[string]string
		dataRaw, err := arguments.String("DATA")
		helpers.OnError(err, log.Panic)

		err = json.Unmarshal([]byte(dataRaw), &data)
		helpers.OnError(err, log.Panic)

		title, ok := data["title"]
		if !ok {
			log.Panic("Missing parameter title in DATA")
		}

		url, ok := data["url"]
		if !ok {
			log.Panic("Missing parameter url in DATA")
		}

		var type_ sql.NullInt32
		typeRaw, ok := data["type"]
		if !ok {
			type_.Valid = false
		}

		typeInt, err := strconv.ParseInt(typeRaw, 10, 32)

		type_.Int32 = int32(typeInt)

		err = libbookmarks.AddBookmark(db, nil, title, url, type_)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit"]; ok {
		var data libbookmarks.Bookmark
		dataRaw, err := arguments.String("NEW_DATA")
		helpers.OnError(err, log.Panic)

		err = json.Unmarshal([]byte(dataRaw), &data)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditBookmark(db, nil, data)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit-is-read"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		isReadRaw, err := arguments.String("IS_READ")
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		isRead, err := strconv.ParseBool(isReadRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditIsRead(db, nil, ID, isRead)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit-title"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		title, err := arguments.String("TITLE")
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditTitle(db, nil, ID, title)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit-url"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		url, err := arguments.String("URL")
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditUrl(db, nil, ID, url)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit-type"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		type_, err := arguments.String("TYPE")
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditType(db, nil, ID, type_)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--edit-is-collection"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		isCollectionRaw, err := arguments.String("IS_COLLECTION")
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		isCollection, err := strconv.ParseBool(isCollectionRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.EditIsCollection(db, nil, ID, isCollection)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--add-tag"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		err = libbookmarks.AddType(db, nil, tag)
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.AddTag(db, nil, ID, tag)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--remove-tag"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, log.Panic)

		err = libbookmarks.RemoveType(db, nil, tag)
		helpers.OnError(err, log.Panic)

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, log.Panic)

		err = libbookmarks.RemoveTag(db, nil, ID, tag)
		helpers.OnError(err, log.Panic)
	}
}
