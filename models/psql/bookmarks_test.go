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

func testBookmarks(t *testing.T) {
	

	query := Bookmarks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookmarksDelete(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
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

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksQueryDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Bookmarks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksSliceDeleteAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookmarksExists(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookmarkExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Bookmark exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookmarkExists to return true, but got false.")
	}
}

func testBookmarksFind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookmarkFound, err := FindBookmark(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookmarkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookmarksBind(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Bookmarks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookmarksOne(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Bookmarks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookmarksAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	bookmarkOne := &Bookmark{}
	bookmarkTwo := &Bookmark{}
	if err = randomize.Struct(seed, bookmarkOne, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTwo, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookmarks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookmarksCount(t *testing.T) {
	

	var err error
	seed := randomize.NewSeed()
	bookmarkOne := &Bookmark{}
	bookmarkTwo := &Bookmark{}
	if err = randomize.Struct(seed, bookmarkOne, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err = randomize.Struct(seed, bookmarkTwo, bookmarkDBTypes, false, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookmarkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookmarkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookmarkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func bookmarkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookmark) error {
	*o = Bookmark{}
	return nil
}

func testBookmarksHooks(t *testing.T) {
	

	var err error

	ctx := context.Background()
	empty := &Bookmark{}
	o := &Bookmark{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookmarkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Bookmark object: %s", err)
	}

	AddBookmarkHook(boil.BeforeInsertHook, bookmarkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeInsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterInsertHook, bookmarkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterInsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterSelectHook, bookmarkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterSelectHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeUpdateHook, bookmarkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeUpdateHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterUpdateHook, bookmarkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterUpdateHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeDeleteHook, bookmarkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeDeleteHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterDeleteHook, bookmarkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterDeleteHooks = []BookmarkHook{}

	AddBookmarkHook(boil.BeforeUpsertHook, bookmarkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkBeforeUpsertHooks = []BookmarkHook{}

	AddBookmarkHook(boil.AfterUpsertHook, bookmarkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookmarkAfterUpsertHooks = []BookmarkHook{}
}

func testBookmarksInsert(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarksInsertWhitelist(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookmarkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookmarkToManyTags(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c Tag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"bookmark_contexts\" (\"bookmark_id\", \"tag_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"bookmark_contexts\" (\"bookmark_id\", \"tag_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Tags().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookmarkSlice{&a}
	if err = a.L.LoadTags(ctx, tx, false, (*[]*Bookmark)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Tags = nil
	if err = a.L.LoadTags(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookmarkToManyAddOpTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c, d, e Tag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Tag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tagDBTypes, false, strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Tag{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTags(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Bookmarks[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Bookmarks[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Tags[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Tags[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Tags().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBookmarkToManySetOpTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c, d, e Tag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Tag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tagDBTypes, false, strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetTags(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Tags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTags(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Tags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Bookmarks) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Bookmarks) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Bookmarks[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Bookmarks[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Tags[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Tags[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBookmarkToManyRemoveOpTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c, d, e Tag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Tag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tagDBTypes, false, strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddTags(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Tags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTags(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Tags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Bookmarks) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Bookmarks) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Bookmarks[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Bookmarks[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Tags) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Tags[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Tags[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBookmarkToOneBookmarkTypeUsingBookmarkType(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Bookmark
	var foreign BookmarkType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookmarkTypeDBTypes, false, bookmarkTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookmarkType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BookmarkTypeID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BookmarkType().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BookmarkSlice{&local}
	if err = local.L.LoadBookmarkType(ctx, tx, false, (*[]*Bookmark)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookmarkType == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BookmarkType = nil
	if err = local.L.LoadBookmarkType(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookmarkType == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBookmarkToOneSetOpBookmarkTypeUsingBookmarkType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b, c BookmarkType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BookmarkType{&b, &c} {
		err = a.SetBookmarkType(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BookmarkType != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Bookmarks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BookmarkTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.BookmarkTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookmarkTypeID))
		reflect.Indirect(reflect.ValueOf(&a.BookmarkTypeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BookmarkTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.BookmarkTypeID, x.ID)
		}
	}
}

func testBookmarkToOneRemoveOpBookmarkTypeUsingBookmarkType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookmark
	var b BookmarkType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookmarkDBTypes, false, strmangle.SetComplement(bookmarkPrimaryKeyColumns, bookmarkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookmarkTypeDBTypes, false, strmangle.SetComplement(bookmarkTypePrimaryKeyColumns, bookmarkTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBookmarkType(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBookmarkType(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.BookmarkType().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.BookmarkType != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BookmarkTypeID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Bookmarks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBookmarksReload(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
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

func testBookmarksReloadAll(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookmarkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookmarksSelect(t *testing.T) {
	

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookmarks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookmarkDBTypes = map[string]string{`ID`: `integer`, `IsRead`: `integer`, `Title`: `text`, `URL`: `text`, `BookmarkTypeID`: `integer`, `IsCollection`: `integer`}
	_               = bytes.MinRead
)

func testBookmarksUpdate(t *testing.T) {
	

	if 0 == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookmarksSliceUpdateAll(t *testing.T) {
	

	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookmark{}
	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookmarkDBTypes, true, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookmarkAllColumns, bookmarkPrimaryKeyColumns) {
		fields = bookmarkAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookmarkAllColumns,
			bookmarkPrimaryKeyColumns,
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

	slice := BookmarkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookmarksUpsert(t *testing.T) {
	

	if len(bookmarkAllColumns) == len(bookmarkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Bookmark{}
	if err = randomize.Struct(seed, &o, bookmarkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookmark: %s", err)
	}

	count, err := Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookmarkDBTypes, false, bookmarkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookmark struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookmark: %s", err)
	}

	count, err = Bookmarks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
