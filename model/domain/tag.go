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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/tag.go.tpl


package domain

import (
    "github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
)

type Tag struct {
    
    ID int64 `json:"id" toml:"id" yaml:"id"`
    ParentPath []*Tag `json:"parentPath" toml:"parentPath" yaml:"parentPath"`
    Tag string `json:"tag" toml:"tag" yaml:"tag"`
    Subtags []*Tag `json:"subtags" toml:"subtags" yaml:"subtags"`
}


func (t *Tag) IsDefault() bool {
    
    var IDZero int64
    if t.ID != IDZero {
    
        return false
    }
    
    if t.ParentPath != nil {
    
        return false
    }
    
    var TagZero string
    if t.Tag != TagZero {
    
        return false
    }
    
    if t.Subtags != nil {
    
        return false
    }
    

    return true
}

func (t *Tag) AddChildren(newChildren []*Tag) {
    t.Subtags = append(t.Subtags, newChildren...)

    for _, child := range newChildren {
        if len(child.ParentPath) == 0 {
            child.ParentPath = make([]*Tag, 1)
        }

        child.ParentPath[0] = t
    }
}

func (t *Tag) AddDirectParent(newParent *Tag) {
    t.ParentPath = append(t.ParentPath, newParent)
    newParent.Subtags = append(newParent.Subtags, t)

    for _, child := range t.Subtags {
        child.ParentPath = append(child.ParentPath, newParent)
    }
}

type TagField string

var TagFields = struct {
    ID  TagField
    ParentPath  TagField
    Tag  TagField
    Subtags  TagField
    
}{
    ID: "id",
    ParentPath: "parentPath",
    Tag: "tag",
    Subtags: "subtags",
    
}

type TagFilter struct {
    ID optional.Optional[model.FilterOperation[int64]]
    ParentPath optional.Optional[model.FilterOperation[*Tag]]
    Tag optional.Optional[model.FilterOperation[string]]
    Subtags optional.Optional[model.FilterOperation[*Tag]]
    
}

func (filter *TagFilter) IsDefault() bool {
    if filter.ID.HasValue {
        return false
    }
    if filter.ParentPath.HasValue {
        return false
    }
    if filter.Tag.HasValue {
        return false
    }
    if filter.Subtags.HasValue {
        return false
    }
    

    return true
}


type TagUpdater struct {
    ID optional.Optional[model.UpdateOperation[int64]]
    ParentPath optional.Optional[model.UpdateOperation[[]*Tag]]
    Tag optional.Optional[model.UpdateOperation[string]]
    Subtags optional.Optional[model.UpdateOperation[[]*Tag]]
    
}

func (updater *TagUpdater) IsDefault() bool {
    if updater.ID.HasValue {
        return false
    }
    if updater.ParentPath.HasValue {
        return false
    }
    if updater.Tag.HasValue {
        return false
    }
    if updater.Subtags.HasValue {
        return false
    }
    

    return true
}

const (
    TagFilterLeaf = "TagFilterLeaf"
    TagFilterRoot = "TagFilterRoot"
)

var PredefinedTagFilters = map[string]TagFilter {
    TagFilterLeaf: {ParentPath: optional.Make(model.FilterOperation[*Tag]{
        Operand: model.ScalarOperand[*Tag]{
            Operand: nil,
        },
        Operator: model.FilterEqual,
    })},
    TagFilterRoot: {Subtags: optional.Make(model.FilterOperation[*Tag]{
        Operand: model.ScalarOperand[*Tag]{
            Operand: nil,
        },
        Operator: model.FilterEqual,
    })},
}
