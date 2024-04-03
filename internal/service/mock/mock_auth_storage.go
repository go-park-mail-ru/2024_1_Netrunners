// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/authservice.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockusersStorage is a mock of usersStorage interface.
type MockusersStorage struct {
	ctrl     *gomock.Controller
	recorder *MockusersStorageMockRecorder
}

// MockusersStorageMockRecorder is the mock recorder for MockusersStorage.
type MockusersStorageMockRecorder struct {
	mock *MockusersStorage
}

// NewMockusersStorage creates a new mock instance.
func NewMockusersStorage(ctrl *gomock.Controller) *MockusersStorage {
	mock := &MockusersStorage{ctrl: ctrl}
	mock.recorder = &MockusersStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockusersStorage) EXPECT() *MockusersStorageMockRecorder {
	return m.recorder
}

// ChangeUserName mocks base method.
func (m *MockusersStorage) ChangeUserName(email, newName string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserName", email, newName)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUserName indicates an expected call of ChangeUserName.
func (mr *MockusersStorageMockRecorder) ChangeUserName(email, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserName",
		reflect.TypeOf((*MockusersStorage)(nil).ChangeUserName), email, newName)
}

// ChangeUserPassword mocks base method.
func (m *MockusersStorage) ChangeUserPassword(email, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", email, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockusersStorageMockRecorder) ChangeUserPassword(email, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockusersStorage)(nil).ChangeUserPassword), email, newPassword)
}

// CreateUser mocks base method.
func (m *MockusersStorage) CreateUser(user domain.UserSignUp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockusersStorageMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser",
		reflect.TypeOf((*MockusersStorage)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockusersStorage) GetUser(email string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", email)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockusersStorageMockRecorder) GetUser(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser",
		reflect.TypeOf((*MockusersStorage)(nil).GetUser), email)
}

// GetUserDataByUuid mocks base method.
func (m *MockusersStorage) GetUserDataByUuid(uuid string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataByUuid", uuid)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDataByUuid indicates an expected call of GetUserDataByUuid.
func (mr *MockusersStorageMockRecorder) GetUserDataByUuid(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataByUuid",
		reflect.TypeOf((*MockusersStorage)(nil).GetUserDataByUuid), uuid)
}

// GetUserPreview mocks base method.
func (m *MockusersStorage) GetUserPreview(uuid string) (domain.UserPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPreview", uuid)
	ret0, _ := ret[0].(domain.UserPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPreview indicates an expected call of GetUserPreview.
func (mr *MockusersStorageMockRecorder) GetUserPreview(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPreview",
		reflect.TypeOf((*MockusersStorage)(nil).GetUserPreview), uuid)
}

// HasUser mocks base method.
func (m *MockusersStorage) HasUser(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUser", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// HasUser indicates an expected call of HasUser.
func (mr *MockusersStorageMockRecorder) HasUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUser",
		reflect.TypeOf((*MockusersStorage)(nil).HasUser), email, password)
}

// RemoveUser mocks base method.
func (m *MockusersStorage) RemoveUser(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockusersStorageMockRecorder) RemoveUser(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser",
		reflect.TypeOf((*MockusersStorage)(nil).RemoveUser), email)
}
