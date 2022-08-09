// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hexiaopi/blog-service/internal/store (interfaces: Factory,ArticleStore,TagStore,AuthStore)

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/hexiaopi/blog-service/internal/entity"
	model "github.com/hexiaopi/blog-service/internal/model"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Articles mocks base method.
func (m *MockFactory) Articles() ArticleStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Articles")
	ret0, _ := ret[0].(ArticleStore)
	return ret0
}

// Articles indicates an expected call of Articles.
func (mr *MockFactoryMockRecorder) Articles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Articles", reflect.TypeOf((*MockFactory)(nil).Articles))
}

// Auths mocks base method.
func (m *MockFactory) Auths() AuthStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auths")
	ret0, _ := ret[0].(AuthStore)
	return ret0
}

// Auths indicates an expected call of Auths.
func (mr *MockFactoryMockRecorder) Auths() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auths", reflect.TypeOf((*MockFactory)(nil).Auths))
}

// Close mocks base method.
func (m *MockFactory) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFactoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFactory)(nil).Close))
}

// Tags mocks base method.
func (m *MockFactory) Tags() TagStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(TagStore)
	return ret0
}

// Tags indicates an expected call of Tags.
func (mr *MockFactoryMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockFactory)(nil).Tags))
}

// MockArticleStore is a mock of ArticleStore interface.
type MockArticleStore struct {
	ctrl     *gomock.Controller
	recorder *MockArticleStoreMockRecorder
}

// MockArticleStoreMockRecorder is the mock recorder for MockArticleStore.
type MockArticleStoreMockRecorder struct {
	mock *MockArticleStore
}

// NewMockArticleStore creates a new mock instance.
func NewMockArticleStore(ctrl *gomock.Controller) *MockArticleStore {
	mock := &MockArticleStore{ctrl: ctrl}
	mock.recorder = &MockArticleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleStore) EXPECT() *MockArticleStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleStore) Create(arg0 context.Context, arg1 *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockArticleStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockArticleStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockArticleStore) Get(arg0 context.Context, arg1 int) (*entity.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockArticleStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArticleStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockArticleStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]*entity.Article, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Article)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockArticleStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockArticleStore) Update(arg0 context.Context, arg1 *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleStore)(nil).Update), arg0, arg1)
}

// MockTagStore is a mock of TagStore interface.
type MockTagStore struct {
	ctrl     *gomock.Controller
	recorder *MockTagStoreMockRecorder
}

// MockTagStoreMockRecorder is the mock recorder for MockTagStore.
type MockTagStoreMockRecorder struct {
	mock *MockTagStore
}

// NewMockTagStore creates a new mock instance.
func NewMockTagStore(ctrl *gomock.Controller) *MockTagStore {
	mock := &MockTagStore{ctrl: ctrl}
	mock.recorder = &MockTagStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagStore) EXPECT() *MockTagStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTagStore) Create(arg0 context.Context, arg1 *entity.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTagStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTagStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTagStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTagStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTagStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockTagStore) Get(arg0 context.Context, arg1 int) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTagStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTagStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockTagStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]model.Tag, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.Tag)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockTagStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTagStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockTagStore) Update(arg0 context.Context, arg1 *entity.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTagStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTagStore)(nil).Update), arg0, arg1)
}

// MockAuthStore is a mock of AuthStore interface.
type MockAuthStore struct {
	ctrl     *gomock.Controller
	recorder *MockAuthStoreMockRecorder
}

// MockAuthStoreMockRecorder is the mock recorder for MockAuthStore.
type MockAuthStoreMockRecorder struct {
	mock *MockAuthStore
}

// NewMockAuthStore creates a new mock instance.
func NewMockAuthStore(ctrl *gomock.Controller) *MockAuthStore {
	mock := &MockAuthStore{ctrl: ctrl}
	mock.recorder = &MockAuthStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthStore) EXPECT() *MockAuthStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAuthStore) Get(arg0 context.Context, arg1, arg2 string) (*model.Auth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAuthStoreMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAuthStore)(nil).Get), arg0, arg1, arg2)
}
