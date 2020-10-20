package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)

	mock.ExpectExec("CREATE TABLE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = db.createTable()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when add row to helper table", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAll(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)

	mock.ExpectExec("DELETE FROM helper").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = db.deleteAll()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when add row to helper table", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteRow(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)

	title := "TestDel"

	mock.ExpectPrepare("DELETE FROM helper").ExpectExec().
		WithArgs(title).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = db.deleteRow(title)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when add row to helper table", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddRow(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)

	title := "TestAdd"
	desc := "TestDesc"

	mock.ExpectPrepare("REPLACE INTO helper").ExpectExec().
		WithArgs(title, desc).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = db.addRow(title, desc)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when add row to helper table", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRow(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)

	title := "TestGet"
	desc := "TestDesc"

	rows := sqlmock.NewRows([]string{"title", "description"}).AddRow(title, desc)
	mock.ExpectQuery("select (.+) from helper").WillReturnRows(rows)
	err = db.getRow()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when select helper table", err)
	}

	assert.Equal(t, hellperMessages[title], desc)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
