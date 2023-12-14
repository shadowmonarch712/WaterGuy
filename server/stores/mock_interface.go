package stores

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "waterguy/models"
	gofr "gofr.dev/pkg/gofr"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUser) CreateUserEntry(ctx *gofr.Context, model models.UserEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserEntry", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) CreateUserEntry(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserEntry", reflect.TypeOf((*MockUser)(nil).CreateUserEntry), ctx, model)
}

// Delete mocks base method.
func (m *MockUser) DeleteUserEntry(ctx *gofr.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserEntry", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserMockRecorder) DeleteUserEntry(ctx *gofr.Context) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserEntry", reflect.TypeOf((*MockUser)(nil).DeleteUserEntry), ctx)
}

// Get mocks base method.
func (m *MockUser) FetchUserEntry(ctx *gofr.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserEntry", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserMockRecorder) FetchUserEntry(ctx *gofr.Context) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserEntry", reflect.TypeOf((*MockUser)(nil).FetchUserEntry), ctx)
}