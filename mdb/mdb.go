package mdb

import (
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

type EmailEntry struct {
	Id        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	OptOut    bool
}

func tryCreate(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS email_entries (id INTEGER PRIMARY KEY, email TEXT, created_at DATETIME, updated_at DATETIME, opt_out BOOLEAN)")
	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			// Code 1 = database already exists
			if sqlError.Code != 1 {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}

func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var entry EmailEntry
	err := row.Scan(&entry.Id, &entry.Email, &entry.CreatedAt, &entry.UpdatedAt, &entry.OptOut)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}
