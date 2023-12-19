// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hexiaopi/blog-service/internal/store (interfaces: Factory,ArticleStore,TagStore,UserStore,RoleStore,SystemConfigStore,ResourceStore,OperationStore,UserRoleStore,SysRestStore,SysMenuStore)

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

// Operations mocks base method.
func (m *MockFactory) Operations() OperationStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operations")
	ret0, _ := ret[0].(OperationStore)
	return ret0
}

// Operations indicates an expected call of Operations.
func (mr *MockFactoryMockRecorder) Operations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operations", reflect.TypeOf((*MockFactory)(nil).Operations))
}

// Resources mocks base method.
func (m *MockFactory) Resources() ResourceStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].(ResourceStore)
	return ret0
}

// Resources indicates an expected call of Resources.
func (mr *MockFactoryMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockFactory)(nil).Resources))
}

// Roles mocks base method.
func (m *MockFactory) Roles() RoleStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Roles")
	ret0, _ := ret[0].(RoleStore)
	return ret0
}

// Roles indicates an expected call of Roles.
func (mr *MockFactoryMockRecorder) Roles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Roles", reflect.TypeOf((*MockFactory)(nil).Roles))
}

// SysMenus mocks base method.
func (m *MockFactory) SysMenus() SysMenuStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SysMenus")
	ret0, _ := ret[0].(SysMenuStore)
	return ret0
}

// SysMenus indicates an expected call of SysMenus.
func (mr *MockFactoryMockRecorder) SysMenus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SysMenus", reflect.TypeOf((*MockFactory)(nil).SysMenus))
}

// SysRests mocks base method.
func (m *MockFactory) SysRests() SysRestStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SysRests")
	ret0, _ := ret[0].(SysRestStore)
	return ret0
}

// SysRests indicates an expected call of SysRests.
func (mr *MockFactoryMockRecorder) SysRests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SysRests", reflect.TypeOf((*MockFactory)(nil).SysRests))
}

// Systems mocks base method.
func (m *MockFactory) Systems() SystemConfigStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Systems")
	ret0, _ := ret[0].(SystemConfigStore)
	return ret0
}

// Systems indicates an expected call of Systems.
func (mr *MockFactoryMockRecorder) Systems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Systems", reflect.TypeOf((*MockFactory)(nil).Systems))
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

// Tx mocks base method.
func (m *MockFactory) Tx(arg0 context.Context, arg1 func(context.Context, Factory) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tx indicates an expected call of Tx.
func (mr *MockFactoryMockRecorder) Tx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockFactory)(nil).Tx), arg0, arg1)
}

// UserRole mocks base method.
func (m *MockFactory) UserRole() UserRoleStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRole")
	ret0, _ := ret[0].(UserRoleStore)
	return ret0
}

// UserRole indicates an expected call of UserRole.
func (mr *MockFactoryMockRecorder) UserRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRole", reflect.TypeOf((*MockFactory)(nil).UserRole))
}

// Users mocks base method.
func (m *MockFactory) Users() UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(UserStore)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockFactoryMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockFactory)(nil).Users))
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

// Count mocks base method.
func (m *MockArticleStore) Count(arg0 context.Context, arg1 *model.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockArticleStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockArticleStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockArticleStore) Create(arg0 context.Context, arg1 *model.Article) error {
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
func (m *MockArticleStore) Get(arg0 context.Context, arg1 int) (*model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockArticleStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArticleStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockArticleStore) List(arg0 context.Context, arg1 *model.ListOption) ([]model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArticleStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockArticleStore) Update(arg0 context.Context, arg1 *model.Article) error {
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

// Count mocks base method.
func (m *MockTagStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockTagStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTagStore)(nil).Count), arg0, arg1)
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
func (m *MockTagStore) Get(arg0 context.Context, arg1 int) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTagStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTagStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockTagStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
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

// MockUserStore is a mock of UserStore interface.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore.
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance.
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockUserStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUserStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockUserStore) Create(arg0 context.Context, arg1 *entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUserStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), arg0, arg1)
}

// GetById mocks base method.
func (m *MockUserStore) GetById(arg0 context.Context, arg1 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserStoreMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserStore)(nil).GetById), arg0, arg1)
}

// GetByName mocks base method.
func (m *MockUserStore) GetByName(arg0 context.Context, arg1 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockUserStoreMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockUserStore)(nil).GetByName), arg0, arg1)
}

// List mocks base method.
func (m *MockUserStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserStore) Update(arg0 context.Context, arg1 *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserStore)(nil).Update), arg0, arg1)
}

// MockRoleStore is a mock of RoleStore interface.
type MockRoleStore struct {
	ctrl     *gomock.Controller
	recorder *MockRoleStoreMockRecorder
}

// MockRoleStoreMockRecorder is the mock recorder for MockRoleStore.
type MockRoleStoreMockRecorder struct {
	mock *MockRoleStore
}

// NewMockRoleStore creates a new mock instance.
func NewMockRoleStore(ctrl *gomock.Controller) *MockRoleStore {
	mock := &MockRoleStore{ctrl: ctrl}
	mock.recorder = &MockRoleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleStore) EXPECT() *MockRoleStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockRoleStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRoleStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRoleStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockRoleStore) Create(arg0 context.Context, arg1 *entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRoleStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRoleStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockRoleStore) Get(arg0 context.Context, arg1 int) (*entity.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockRoleStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]entity.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockRoleStore) Update(arg0 context.Context, arg1 *entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRoleStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleStore)(nil).Update), arg0, arg1)
}

// MockSystemConfigStore is a mock of SystemConfigStore interface.
type MockSystemConfigStore struct {
	ctrl     *gomock.Controller
	recorder *MockSystemConfigStoreMockRecorder
}

// MockSystemConfigStoreMockRecorder is the mock recorder for MockSystemConfigStore.
type MockSystemConfigStoreMockRecorder struct {
	mock *MockSystemConfigStore
}

// NewMockSystemConfigStore creates a new mock instance.
func NewMockSystemConfigStore(ctrl *gomock.Controller) *MockSystemConfigStore {
	mock := &MockSystemConfigStore{ctrl: ctrl}
	mock.recorder = &MockSystemConfigStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemConfigStore) EXPECT() *MockSystemConfigStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSystemConfigStore) Get(arg0 context.Context, arg1 string) (*entity.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSystemConfigStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSystemConfigStore)(nil).Get), arg0, arg1)
}

// MockResourceStore is a mock of ResourceStore interface.
type MockResourceStore struct {
	ctrl     *gomock.Controller
	recorder *MockResourceStoreMockRecorder
}

// MockResourceStoreMockRecorder is the mock recorder for MockResourceStore.
type MockResourceStoreMockRecorder struct {
	mock *MockResourceStore
}

// NewMockResourceStore creates a new mock instance.
func NewMockResourceStore(ctrl *gomock.Controller) *MockResourceStore {
	mock := &MockResourceStore{ctrl: ctrl}
	mock.recorder = &MockResourceStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceStore) EXPECT() *MockResourceStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockResourceStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockResourceStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockResourceStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockResourceStore) Create(arg0 context.Context, arg1 *entity.Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResourceStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockResourceStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResourceStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockResourceStore) Get(arg0 context.Context, arg1 int) (*entity.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResourceStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockResourceStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]entity.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]entity.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourceStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockResourceStore) Update(arg0 context.Context, arg1 *entity.Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResourceStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceStore)(nil).Update), arg0, arg1)
}

// MockOperationStore is a mock of OperationStore interface.
type MockOperationStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperationStoreMockRecorder
}

// MockOperationStoreMockRecorder is the mock recorder for MockOperationStore.
type MockOperationStoreMockRecorder struct {
	mock *MockOperationStore
}

// NewMockOperationStore creates a new mock instance.
func NewMockOperationStore(ctrl *gomock.Controller) *MockOperationStore {
	mock := &MockOperationStore{ctrl: ctrl}
	mock.recorder = &MockOperationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationStore) EXPECT() *MockOperationStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockOperationStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOperationStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOperationStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockOperationStore) Create(arg0 context.Context, arg1 *entity.OperationLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOperationStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOperationStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockOperationStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOperationStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOperationStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockOperationStore) Get(arg0 context.Context, arg1 int) (*entity.OperationLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.OperationLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOperationStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOperationStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockOperationStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]model.OperationLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.OperationLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOperationStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOperationStore)(nil).List), arg0, arg1)
}

// MockUserRoleStore is a mock of UserRoleStore interface.
type MockUserRoleStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserRoleStoreMockRecorder
}

// MockUserRoleStoreMockRecorder is the mock recorder for MockUserRoleStore.
type MockUserRoleStoreMockRecorder struct {
	mock *MockUserRoleStore
}

// NewMockUserRoleStore creates a new mock instance.
func NewMockUserRoleStore(ctrl *gomock.Controller) *MockUserRoleStore {
	mock := &MockUserRoleStore{ctrl: ctrl}
	mock.recorder = &MockUserRoleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRoleStore) EXPECT() *MockUserRoleStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRoleStore) Create(arg0 context.Context, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRoleStoreMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRoleStore)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockUserRoleStore) Delete(arg0 context.Context, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRoleStoreMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRoleStore)(nil).Delete), arg0, arg1, arg2)
}

// DeleteByUser mocks base method.
func (m *MockUserRoleStore) DeleteByUser(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUser indicates an expected call of DeleteByUser.
func (mr *MockUserRoleStoreMockRecorder) DeleteByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUser", reflect.TypeOf((*MockUserRoleStore)(nil).DeleteByUser), arg0, arg1)
}

// ListUserRole mocks base method.
func (m *MockUserRoleStore) ListUserRole(arg0 context.Context, arg1 int) ([]entity.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserRole", arg0, arg1)
	ret0, _ := ret[0].([]entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserRole indicates an expected call of ListUserRole.
func (mr *MockUserRoleStoreMockRecorder) ListUserRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserRole", reflect.TypeOf((*MockUserRoleStore)(nil).ListUserRole), arg0, arg1)
}

// MockSysRestStore is a mock of SysRestStore interface.
type MockSysRestStore struct {
	ctrl     *gomock.Controller
	recorder *MockSysRestStoreMockRecorder
}

// MockSysRestStoreMockRecorder is the mock recorder for MockSysRestStore.
type MockSysRestStoreMockRecorder struct {
	mock *MockSysRestStore
}

// NewMockSysRestStore creates a new mock instance.
func NewMockSysRestStore(ctrl *gomock.Controller) *MockSysRestStore {
	mock := &MockSysRestStore{ctrl: ctrl}
	mock.recorder = &MockSysRestStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSysRestStore) EXPECT() *MockSysRestStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSysRestStore) Count(arg0 context.Context, arg1 *entity.ListOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSysRestStoreMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSysRestStore)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockSysRestStore) Create(arg0 context.Context, arg1 *model.SysRest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSysRestStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSysRestStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSysRestStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSysRestStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSysRestStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockSysRestStore) Get(arg0 context.Context, arg1 int) (*model.SysRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.SysRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSysRestStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSysRestStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockSysRestStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]model.SysRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.SysRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSysRestStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSysRestStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockSysRestStore) Update(arg0 context.Context, arg1 *model.SysRest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSysRestStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSysRestStore)(nil).Update), arg0, arg1)
}

// MockSysMenuStore is a mock of SysMenuStore interface.
type MockSysMenuStore struct {
	ctrl     *gomock.Controller
	recorder *MockSysMenuStoreMockRecorder
}

// MockSysMenuStoreMockRecorder is the mock recorder for MockSysMenuStore.
type MockSysMenuStoreMockRecorder struct {
	mock *MockSysMenuStore
}

// NewMockSysMenuStore creates a new mock instance.
func NewMockSysMenuStore(ctrl *gomock.Controller) *MockSysMenuStore {
	mock := &MockSysMenuStore{ctrl: ctrl}
	mock.recorder = &MockSysMenuStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSysMenuStore) EXPECT() *MockSysMenuStoreMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSysMenuStore) List(arg0 context.Context, arg1 *entity.ListOption) ([]model.SysMenu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.SysMenu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSysMenuStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSysMenuStore)(nil).List), arg0, arg1)
}
