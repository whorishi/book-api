package handler

import (
	"strconv"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"book-rest-api/datastore"
	"book-rest-api/model"
)

type handler struct {
	store datastore.Book
}

func New(b datastore.Book) handler {
	return handler{store: b}
}

type response struct {
	Books []model.Book
}

func (h handler) GetByID(ctx *gofr.Context) (interface{},error) {
	resp, err := h.store.GetByID(ctx)

	if err!=nil {
		return nil,err
	}

	return response{Books: resp}, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var book model.Book

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&book); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var book model.Book
	if err = ctx.Bind(&book); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	book.ID = id

	resp, err := h.store.Update(ctx, book)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}

func validateID(id string) (int , error){
	res, err := strconv.Atoi(id);
	if err!=nil {
		return 0,err
	}
	return res,err
}

