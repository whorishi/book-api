package datastore

import (  
	"book-rest-api/model"  
    "gofr.dev/pkg/gofr"
)

type Book interface {
	GetByID(ctx *gofr.Context) ([]model.Book, error)

	Create(ctx *gofr.Context, model model.Book) (model.Book, error)
	// Update updates an existing student record with the provided information.
	Update(ctx *gofr.Context, model model.Book) (model.Book, error)
	// Delete removes a student record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}

