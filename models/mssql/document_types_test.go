// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDocumentTypes(t *testing.T) {
	

	query := DocumentTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDocumentTypesDelete(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDocumentTypesQueryDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DocumentTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDocumentTypesSliceDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DocumentTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDocumentTypesExists(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DocumentTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DocumentType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DocumentTypeExists to return true, but got false.")
	}
}

func testDocumentTypesFind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	documentTypeFound, err := FindDocumentType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if documentTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDocumentTypesBind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DocumentTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDocumentTypesOne(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DocumentTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDocumentTypesAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	documentTypeOne := &DocumentType{}
	documentTypeTwo := &DocumentType{}
	if err = randomize.Struct(seed, documentTypeOne, documentTypeDBTypes, false, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}
	if err = randomize.Struct(seed, documentTypeTwo, documentTypeDBTypes, false, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = documentTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = documentTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DocumentTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDocumentTypesCount(t *testing.T) {
	

	var err error
	seed := randomize.NewSeed()
	documentTypeOne := &DocumentType{}
	documentTypeTwo := &DocumentType{}
	if err = randomize.Struct(seed, documentTypeOne, documentTypeDBTypes, false, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}
	if err = randomize.Struct(seed, documentTypeTwo, documentTypeDBTypes, false, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = documentTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = documentTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func documentTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func documentTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DocumentType) error {
	*o = DocumentType{}
	return nil
}

func testDocumentTypesHooks(t *testing.T) {
	

	var err error

	ctx := context.Background()
	empty := &DocumentType{}
	o := &DocumentType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, documentTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DocumentType object: %s", err)
	}

	AddDocumentTypeHook(boil.BeforeInsertHook, documentTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	documentTypeBeforeInsertHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.AfterInsertHook, documentTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	documentTypeAfterInsertHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.AfterSelectHook, documentTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	documentTypeAfterSelectHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.BeforeUpdateHook, documentTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	documentTypeBeforeUpdateHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.AfterUpdateHook, documentTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	documentTypeAfterUpdateHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.BeforeDeleteHook, documentTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	documentTypeBeforeDeleteHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.AfterDeleteHook, documentTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	documentTypeAfterDeleteHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.BeforeUpsertHook, documentTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	documentTypeBeforeUpsertHooks = []DocumentTypeHook{}

	AddDocumentTypeHook(boil.AfterUpsertHook, documentTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	documentTypeAfterUpsertHooks = []DocumentTypeHook{}
}

func testDocumentTypesInsert(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDocumentTypesInsertWhitelist(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(documentTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDocumentTypeToManyDocuments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DocumentType
	var b, c Document

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.DocumentTypeID = a.ID
	c.DocumentTypeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Documents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.DocumentTypeID == b.DocumentTypeID {
			bFound = true
		}
		if v.DocumentTypeID == c.DocumentTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DocumentTypeSlice{&a}
	if err = a.L.LoadDocuments(ctx, tx, false, (*[]*DocumentType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Documents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Documents = nil
	if err = a.L.LoadDocuments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Documents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDocumentTypeToManyAddOpDocuments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DocumentType
	var b, c, d, e Document

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentTypeDBTypes, false, strmangle.SetComplement(documentTypePrimaryKeyColumns, documentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Document{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Document{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDocuments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.DocumentTypeID {
			t.Error("foreign key was wrong value", a.ID, first.DocumentTypeID)
		}
		if a.ID != second.DocumentTypeID {
			t.Error("foreign key was wrong value", a.ID, second.DocumentTypeID)
		}

		if first.R.DocumentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DocumentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Documents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Documents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Documents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDocumentTypesReload(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDocumentTypesReloadAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DocumentTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDocumentTypesSelect(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DocumentTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	documentTypeDBTypes = map[string]string{`ID`: `int`, `DocumentType`: `varchar`}
	_                   = bytes.MinRead
)

func testDocumentTypesUpdate(t *testing.T) {
	

	if 0 == len(documentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(documentTypeAllColumns) == len(documentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDocumentTypesSliceUpdateAll(t *testing.T) {
	

	if len(documentTypeAllColumns) == len(documentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DocumentType{}
	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, documentTypeDBTypes, true, documentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(documentTypeAllColumns, documentTypePrimaryKeyColumns) {
		fields = documentTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			documentTypeAllColumns,
			documentTypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := DocumentTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDocumentTypesUpsert(t *testing.T) {
	

	if len(documentTypeAllColumns) == len(documentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DocumentType{}
	if err = randomize.Struct(seed, &o, documentTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DocumentType: %s", err)
	}

	count, err := DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, documentTypeDBTypes, false, documentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DocumentType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DocumentType: %s", err)
	}

	count, err = DocumentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
