package datastore

import (
	"book-rest-api/model"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	//"gofr.dev/examples/using-mysql/models"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func TestCoreLayer(*testing.T) {
	app := gofr.New()
	seeder := datastore.NewSeeder(&app.DataStore,"../db")
	seeder.ResetCounter = true
	createTable(app)
}

func createTable(app *gofr.Gofr) {
	_, err := app.DB().Exec("DROP TABLE IF EXISTS books;")

	if err != nil {
		return
	}
	_, err = app.DB().Exec("CREATE TABLE IF NOT EXISTS books " +
		"(id int primary key AUTO_INCREMENT, title varchar(225) not null, author varchar(255), publisher varchar(225),price int, category varchar(225));")

	if err != nil {
		return
	}
}

func TestAddBooks(t *testing.T) {
	ctx := gofr.NewContext(nil,nil,gofr.New())
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()
	tests := []struct {
		desc string
		book model.Book
		mockErr error
		err error
	}{
		{"Valid Case",model.Book{Title: "I Thought Twice", Author: "Alen Peterson", Publisher: "Chicago Literacy", Price: 999, Category: "Fiction"}, nil, nil},
		{"DB error",model.Book{Title: "I Thought Twice", Author: "Alen Peterson", Publisher: "Chicago Literacy", Price: 999},errors.DB{},errors.DB{Err: errors.DB{}}},
	}

	for i, tc := range tests {
		mock.ExpectExec("INSERT INTO books (title,author,publisher,price,category) VALUES (?,?,?,?,?)").
		WithArgs(tc.book.Title,tc.book.Author,tc.book.Publisher,tc.book.Price,tc.book.Category).
		WillReturnResult(sqlmock.NewResult(2,1)).
		WillReturnError(tc.mockErr)

		rows := sqlmock.NewRows([]string{"id","title", "author", "publisher", "price", "category"}).
			AddRow(tc.book.ID, tc.book.Title, tc.book.Author, tc.book.Publisher, tc.book.Price, tc.book.Category)
		mock.ExpectQuery("SELECT id, title,author,publisher,price,category FROM books WHERE id = ?").
			WithArgs(tc.book.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		store := New()
		resp,err := store.Create(ctx, tc.book)

		ctx.Logger.Log(resp)
		assert.IsType(t,tc.err,err,"TEST[%d], failed.\n%s", i, tc.desc)
	} 
}

func TestGetBooks(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}
	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests:= []struct {
		desc string
		books []model.Book
		mockErr error
		err error
	}{
			{"Valid case with books", []model.Book{
			{ID: 1, Title: "John Doe", Author: "john", Publisher: "publisher_name", Price: 999, Category: "fun"},
			{ID: 2, Title: "John James", Author: "james", Publisher: "publisher_name", Price: 478, Category: "action"},
		}, nil, nil},
		{"Valid case with no books", []model.Book{}, nil, nil},
		{"Error case", nil, errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {
		rows := sqlmock.NewRows([]string{"id", "title", "author", "publisher", "price", "category"})
		for _, bk := range tc.books {
			rows.AddRow(bk.ID, bk.Title, bk.Author, bk.Publisher, bk.Price, bk.Category)
		}

		mock.ExpectQuery("SELECT id,title,author,publisher,price,category FROM books").WillReturnRows(rows).WillReturnError(tc.mockErr)

		store := New()
		resp, err := store.GetByID(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.books, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

