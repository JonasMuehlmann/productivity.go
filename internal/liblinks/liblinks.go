// Package libtlinks implements functionality to work with links and backlinks in a database context.
package liblinks

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// AddLink adds a link from source to destination to the DB.
// Passing a transaction is optional.
func AddLink(dbConn *sqlx.DB, transaction *sql.Tx, source string, destination string) error {
	stmt := `
        INSERT INTO
            Link(
                Source,
                Destination
            )
        VALUES(
            '?',
            '?'
        );
    `

	var statement *sql.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Prepare(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	defer statement.Close()

	_, err = statement.Exec(source, destination)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// RemoveLink removes the link from source to destination from the db.
// Passing a transaction is optional.
func RemoveLink(dbConn *sqlx.DB, transaction *sql.Tx, source string, destination string) error {
	stmt := `
        DELETE FROM
            Link
        WHERE
            Source = '?'
            AND
            Destination = '?';
    `

	var statement *sql.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Prepare(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	defer statement.Close()

	_, err = statement.Exec(source, destination)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}
	return nil
}

// List all link destionations from the DB.
func ListLinks(dbConn *sqlx.DB, source string) ([]string, error) {
	stmtLinks := `
        SELECT
            Destination
        FROM
            Link
        WHERE
            Source = '?';
    `

	statementLinks, err := dbConn.Prepare(stmtLinks)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	linkRows, err := statementLinks.Query(source)
	if err != nil {
		return nil, err
	}

	stmtCountLinks := "SELECT COUNT(*) FROM Link WHERE Source = '?';"

	statementLinksCount, err := dbConn.Prepare(stmtCountLinks)
	if err != nil {
		return nil, err
	}

	defer statementLinksCount.Close()

	linksCountRow := statementLinksCount.QueryRow(source)

	var rowCountLinks int

	err = linksCountRow.Scan(&rowCountLinks)
	if err != nil {
		return nil, err
	}

	linksBuffer := make([]string, rowCountLinks)

	i := 0
	for linkRows.Next() {
		err := linkRows.Scan(&linksBuffer[i])
		if err != nil {
			return nil, err
		}
		i++
	}

	return linksBuffer, nil
}

// List all link sources from the DB.
func ListBacklinks(dbConn *sqlx.DB, destination string) ([]string, error) {
	stmtBacklinks := `
        SELECT
            Source
        FROM
            Link
        WHERE
            Destination = '?';
    `

	statementBacklinks, err := dbConn.Prepare(stmtBacklinks)
	if err != nil {
		return nil, err
	}

	defer statementBacklinks.Close()

	backlinkRows, err := statementBacklinks.Query(destination)
	if err != nil {
		return nil, err
	}

	stmtCountBacklinks := "SELECT COUNT(*) FROM Link  WHERE Destination = '?';"

	statementBacklinksCount, err := dbConn.Prepare(stmtCountBacklinks)
	if err != nil {
		return nil, err
	}

	defer statementBacklinksCount.Close()

	backLinksCountRow := statementBacklinksCount.QueryRow(destination)

	var rowCountLinks int

	err = backLinksCountRow.Scan(&rowCountLinks)
	if err != nil {
		return nil, err
	}

	backlinksBuffer := make([]string, rowCountLinks)

	i := 0
	for backlinkRows.Next() {
		err := backlinkRows.Scan(&backlinksBuffer[i])
		if err != nil {
			return nil, err
		}
		i++
	}

	return backlinksBuffer, nil
}
