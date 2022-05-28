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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/bookmark.go.tpl

package domain

// TODO: Whe could remove the imports and only have one template if goimport could parse this...
import (
	"time"

	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
)

type Bookmark struct {
	CreatedAt    time.Time                    `json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt    optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	URL          string                       `json:"url" toml:"url" yaml:"url"`
	Title        optional.Optional[string]    `json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Tags         []*Tag                       `json:"tags" toml:"tags" yaml:"tags"`
	BookmarkType optional.Optional[string]    `json:"bookmark_type,omitempty" toml:"bookmark_type" yaml:"bookmark_type,omitempty"`
	ID           int64                        `json:"id" toml:"id" yaml:"id"`
	IsCollection bool                         `json:"is_collection,omitempty" toml:"is_collection" yaml:"is_collection,omitempty"`
	IsRead       bool                         `json:"is_read,omitempty" toml:"is_read" yaml:"is_read,omitempty"`
}

type BookmarkField string

var BookmarkFieldsList = []BookmarkField{
	BookmarkField("CreatedAt"),
	BookmarkField("UpdatedAt"),
	BookmarkField("DeletedAt"),
	BookmarkField("URL"),
	BookmarkField("Title"),
	BookmarkField("Tags"),
	BookmarkField("ID"),
	BookmarkField("IsCollection"),
	BookmarkField("IsRead"),
	BookmarkField("BookmarkType"),
}

var BookmarkFields = struct {
	CreatedAt    BookmarkField
	UpdatedAt    BookmarkField
	DeletedAt    BookmarkField
	URL          BookmarkField
	Title        BookmarkField
	Tags         BookmarkField
	ID           BookmarkField
	IsCollection BookmarkField
	IsRead       BookmarkField
	BookmarkType BookmarkField
}{
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	URL:          "url",
	Title:        "title",
	Tags:         "tags",
	ID:           "id",
	IsCollection: "is_collection",
	IsRead:       "is_read",
	BookmarkType: "bookmark_type",
}

type BookmarkFilter struct {
	CreatedAt    optional.Optional[model.FilterOperation[time.Time]]
	UpdatedAt    optional.Optional[model.FilterOperation[time.Time]]
	DeletedAt    optional.Optional[model.FilterOperation[optional.Optional[time.Time]]]
	URL          optional.Optional[model.FilterOperation[string]]
	Title        optional.Optional[model.FilterOperation[optional.Optional[string]]]
	Tags         optional.Optional[model.FilterOperation[*Tag]]
	ID           optional.Optional[model.FilterOperation[int64]]
	IsCollection optional.Optional[model.FilterOperation[bool]]
	IsRead       optional.Optional[model.FilterOperation[bool]]
	BookmarkType optional.Optional[model.FilterOperation[optional.Optional[string]]]
}

type BookmarkUpdater struct {
	CreatedAt    optional.Optional[model.UpdateOperation[time.Time]]
	UpdatedAt    optional.Optional[model.UpdateOperation[time.Time]]
	DeletedAt    optional.Optional[model.UpdateOperation[optional.Optional[time.Time]]]
	URL          optional.Optional[model.UpdateOperation[string]]
	Title        optional.Optional[model.UpdateOperation[optional.Optional[string]]]
	Tags         optional.Optional[model.UpdateOperation[[]*Tag]]
	BookmarkType optional.Optional[model.UpdateOperation[optional.Optional[string]]]
	ID           optional.Optional[model.UpdateOperation[int64]]
	IsCollection optional.Optional[model.UpdateOperation[bool]]
	IsRead       optional.Optional[model.UpdateOperation[bool]]
}
