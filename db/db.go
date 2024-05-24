package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type EmailEntry struct {
	ID          int64
	Email       string
	ConfirmedAt sql.NullTime
	OptOut      bool
}

type GetEmailBatchQueryParam struct {
	Page  int
	Count int
}

func InitializeDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS emails (
		id           INTEGER NOT NULL PRIMARY KEY,
		email        TEXT UNIQUE,
		confirmed_at DATETIME,
		opt_out      INTEGER
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func emailEntryFromRows(rows *sql.Rows) (*EmailEntry, error) {
	var entry EmailEntry
	var optOutInt int

	err := rows.Scan(&entry.ID, &entry.Email, &entry.ConfirmedAt, &optOutInt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	entry.OptOut = optOutInt == 1

	return &entry, nil
}

func AddEmailEntry(ctx context.Context, db *sql.DB, email string) (int64, error) {
	existingEntry, err := GetEmailEntry(ctx, db, email)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingEntry != nil {
		return 0, errors.New("email already exists")
	}

	query := `INSERT INTO emails (email, confirmed_at, opt_out) values (?, ?, ?)`
	result, err := db.ExecContext(ctx, query, email, nil, 0)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.LastInsertId()
}

func GetEmailEntry(ctx context.Context, db *sql.DB, email string) (*EmailEntry, error) {
	query := `SELECT * FROM emails WHERE email = ?`
	rows, err := db.QueryContext(ctx, query, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, sql.ErrNoRows
	}

	entry, err := emailEntryFromRows(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return entry, nil
}

func UpdateEmailEntry(ctx context.Context, db *sql.DB, entry EmailEntry) error {
	query := `UPDATE emails SET opt_out = ?, confirmed_at = ? WHERE id = ?`

	optOutInt := 0
	if entry.OptOut {
		optOutInt = 1
	}

	var confirmedAt interface{}
	if entry.ConfirmedAt.Valid {
		confirmedAt = entry.ConfirmedAt.Time // Use the Time field from sql.NullTime
	} else {
		confirmedAt = nil
	}

	_, err := db.ExecContext(ctx, query, optOutInt, confirmedAt, entry.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteEmailEntry(ctx context.Context, db *sql.DB, email string) error {
	query := `DELETE FROM emails WHERE email = ?`
	_, err := db.ExecContext(ctx, query, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetEmailBatch(ctx context.Context, db *sql.DB, params GetEmailBatchQueryParam) ([]EmailEntry, error) {
	query := `SELECT * FROM emails WHERE opt_out = FALSE ORDER BY id LIMIT ? OFFSET ?`

	rows, err := db.QueryContext(ctx, query, params.Count, (params.Page-1)*params.Count)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	emails := make([]EmailEntry, 0, params.Count)

	for rows.Next() {
		email, err := emailEntryFromRows(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if email != nil {
			emails = append(emails, *email)
		}
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return emails, nil
}
