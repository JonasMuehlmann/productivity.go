package test

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/stretchr/testify/assert"
)

// ######################
// # ImportMinimalCSV() #
// ######################
func TestImportMinimalCSVEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := ""
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVNoHeaderButData(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "Foo;Bar"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoTitle(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "dss;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoUrl(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "dss;Url"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOnlyHeader(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "Url;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOneEntry(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVManyEntries(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteUrl(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteTitle(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

// ###############
// # ExportCSV() #
// ###############
func TestExportCSVEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.ExportCSV(bookmarks, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestExportCSV(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", sql.NullInt16{Valid: false})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo2", "Bar2", sql.NullInt16{Valid: false})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo3", "Bar3", sql.NullInt16{Valid: false})
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.ExportCSV(bookmarks, filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)
}

// #############
// # AddType() #
// #############
func TestAddTypeEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "")
	assert.Error(t, err)
}

func TestAddType(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestAddTypeTransaction(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

// ################
// # RemoveType() #
// ################
func TestRemoveTypeEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "")
	assert.Error(t, err)
}

func TestRemoveTypeNonExistent(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "Foo")
	assert.Error(t, err)
}

func TestRemoveType(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTypeTransaction(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

// ###############
// # ListTypes() #
// ###############
func TestListTypes(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 1)
	assert.Equal(t, "Foo", types[0])
}

func TestListTypesMany(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo2")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo3")
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 3)
	assert.Equal(t, []string{"Foo", "Foo2", "Foo3"}, types)
}

func TestListTypesEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 0)
}
