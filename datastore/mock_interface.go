package datastore

import (
	"book-rest-api/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// Delete implements Book.
func (*MockStore) Delete(ctx *gofr.Context, id int) error {
	panic("unimplemented")
}

// Update implements Book.
func (*MockStore) Update(ctx *gofr.Context, model model.Book) (model.Book, error) {
	panic("unimplemented")
}

type MockStoreMockRecorder struct {
	mock *MockStore
}

func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

func (m *MockStore) Create(ctx *gofr.Context, book model.Book) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStoreMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), ctx, book)
}

func (m *MockStore) GetByID(ctx *gofr.Context) ([]model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx)
	ret0, _ := ret[0].([]model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStoreMockRecorder) GetByID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStore)(nil).GetByID), ctx)
}
