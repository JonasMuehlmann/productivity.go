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

{{$StructName := .StructName}}
package domain

import (
    "github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
)

{{$StructName :=  .StructName -}}

type {{.StructName}} struct {
    {{range $field := .StructFields}}
    {{.FieldName}} {{.FieldType}} {{.FieldTags -}}
    {{end}}
}

type {{.StructName}}Field string

var {{.StructName}}Fields = struct {
    {{range $field := .StructFields -}}
    {{.FieldName}}  {{$StructName}}Field
    {{end}}
}{
    {{range $field := .StructFields -}}
    {{.FieldName}}: "{{.LogicalFieldName -}}",
    {{end}}
}

type {{.StructName}}Filter struct {
    {{range $field := .StructFields -}}
    {{.FieldName}} optional.Optional[model.FilterOperation[{{.FieldType}}]]
    {{end}}
}

type {{$.StructName}}Updater struct {
    {{range $field := .StructFields -}}
    {{.FieldName}} optional.Optional[model.UpdateOperation[{{.FieldType}}]]
    {{end}}
}
