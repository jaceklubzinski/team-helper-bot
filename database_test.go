package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestgetRow(t *testing.T) {
	dbClient, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbClient.Close()
	db := newDB(dbClient)
	rows := sqlmock.NewRows([]string{"title", "description"}).AddRow("Test", "TestDesc")
	mock.ExpectQuery("select * from helper").WillReturnRows(rows)
	db.getRow()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when select helper table", err)
	}

	assert.Equal(t, hellperMessages["Test"], "TestDesc")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when select incidents", err)
	}

}
