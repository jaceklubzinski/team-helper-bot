package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type store struct {
	db *sql.DB
}

func newRepository(db *sql.DB) *store {
	return &store{db}
}
func connectDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "helperbot.db")
}

func (d *store) createTable() error {
	helperTableCreate := `
	CREATE TABLE IF NOT EXISTS helper(
			title TEXT NOT NULL,
			description TEXT NOT NULL
	);
	`
	_, err := d.db.Exec(helperTableCreate)
	if err != nil {
		return err
	}
	return nil
}

func (d *store) deleteAll() error {
	helperTableDelete := `
	DELETE FROM helper;
	`
	_, err := d.db.Exec(helperTableDelete)
	if err != nil {
		return err
	}
	return nil
}

func (d *store) deleteRow(title string) error {
	stmt, err := d.db.Prepare("DELETE FROM helper where title=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

func (d *store) addRow(title, desc string) error {
	stmt, err := d.db.Prepare("REPLACE INTO helper(title,description) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, desc)
	if err != nil {
		return err
	}
	return nil
}

func (d *store) getRow() error {
	var title, desc string
	r, err := d.db.Query("select * from helper;")
	if err != nil {
		return err
	}
	for r.Next() {
		err := r.Scan(&title, &desc)
		if err != nil {
			return err
		}
		if _, ok := hellperMessages[title]; !ok {
			hellperMessages[title] = desc
		}
	}
	return nil
}
