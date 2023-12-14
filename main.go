package main

import (
	"book-rest-api/datastore"
	"book-rest-api/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()

	h := handler.New(s)

	app.GET("/books", h.GetByID)
	app.POST("/books", h.Create)
	app.PUT("/books/{id}", h.Update)
	app.DELETE("/books/{id}", h.Delete)

	app.Start()
}
