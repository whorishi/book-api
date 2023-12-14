package datastore

import (
	"book-rest-api/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type book struct{}

func New() *book {
	return &book{}
}

func (b book) GetByID(ctx *gofr.Context) ([]model.Book, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT id,title,author,publisher,price,category from books")
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer rows.Close()

	books := make([]model.Book, 0)

	for rows.Next() {
		var c model.Book
		err = rows.Scan(&c.ID, &c.Title, &c.Author, &c.Publisher, &c.Price, &c.Category)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		books = append(books, c)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.DB{Err: err}
	}
	return books, nil
}

func (b *book) Create(ctx *gofr.Context, book model.Book) (model.Book, error) {
	var resp model.Book
	queryInsert := "INSERT INTO books (title,author,publisher,price,category) VALUES (?,?,?,?,?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, book.Title, book.Author, book.Publisher, book.Price, book.Category)

	if err != nil {
		return model.Book{}, errors.DB{Err: err}
	}
	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return model.Book{}, errors.DB{Err: err}
	}

	querySelect := "SELECT id,title,author,publisher,price,category FROM books WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).Scan(&resp.ID, &resp.Title, &resp.Author, &resp.Publisher, &resp.Price, &resp.Category)

	if err != nil {
		return model.Book{}, errors.DB{Err: err}
	}
	return resp, nil
}

func (b *book) Update(ctx *gofr.Context, book model.Book) (model.Book, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE books SET title=?, author=?, publisher=?, price=?, category=? WHERE id=?", book.Title, book.Author, book.Publisher, book.Price, book.Category, book.ID)
	if err != nil {
		return model.Book{}, errors.DB{Err: err}
	}
	return book, nil
}

func (b *book) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM books WHERE id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
