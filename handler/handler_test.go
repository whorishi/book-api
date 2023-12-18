package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"book-rest-api/datastore"
	"book-rest-api/model"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
)

func initializeHandlerTest(t *testing.T) (*datastore.MockStore, handler, *gofr.Gofr) {
	ctrl := gomock.NewController(t)

	mockStore := datastore.NewMockStore(ctrl)
	h := New(mockStore)
	app := gofr.New()

	return mockStore, h, app
}

func TestGet(t *testing.T) {
	tests := []struct {
		desc string
		resp []model.Book
		err  error
	}{
		{"success case", []model.Book{{ID: 0, Title: "sample", Author: "email@gmail.com", Publisher: "930098800",
			Price: 990 ,Category: "kolkata"}}, nil},
		{"error case", nil, errors.Error("error fetching book listing")},
	}

	mockStore, h, app := initializeHandlerTest(t)

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		mockStore.EXPECT().GetByID(ctx).Return(tc.resp, tc.err)

		result, err := h.GetByID(ctx)

		if tc.err == nil {
			// Assert successful response
			assert.Nil(t, err)
			assert.NotNil(t, result)

			res, ok := result.(response)
			assert.True(t, ok)
			assert.Equal(t, tc.resp, res.Books)
		} else {
			// Assert error response
			assert.NotNil(t, err)
			assert.Equal(t, tc.err, err)
			assert.Nil(t, result)
		}
	}
}

func TestCreate(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	input := `{"title":"mahak","author":"msjce","publisher":"928902","price":784,"category":"asdf"}`
	expResp := model.Book{Title: "mahak", Author: "msjce", Publisher: "928902", Price: 784, Category: "asdf"}

	in := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodPost, "/books", in)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	var bk model.Book

	_ = ctx.Bind(&bk)

	mockStore.EXPECT().GetByID(ctx).Return(nil, nil).MaxTimes(2)
	mockStore.EXPECT().Create(ctx, bk).Return(expResp, nil).MaxTimes(1)

	resp, err := h.Create(ctx)

	assert.Nil(t, err, "TEST,failed :success case")

	assert.Equal(t, expResp, resp, "TEST, failed:success case")
}

func TestCreate_Error(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	tests := []struct {
		desc    string
		input   string
		expResp interface{}
		err     error
	}{{"create invalid body", `{"title":"mahak","author":"msjce","publisher":"928902","price":784,"category":"asdf"}`, model.Book{},
		errors.InvalidParam{Param: []string{"body"}}},
		{"create invalid body", `}`, model.Book{}, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range tests {
		in := strings.NewReader(tc.input)
		req := httptest.NewRequest(http.MethodPost, "/books", in)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		var emp model.Book

		_ = ctx.Bind(&emp)

		mockStore.EXPECT().Create(ctx, emp).Return(tc.expResp.(model.Book), tc.err).MaxTimes(1)

		resp, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Nil(t, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}