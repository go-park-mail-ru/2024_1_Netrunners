// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handlers/actors.go

// Package mock_service is a generated GoMock package.
package mockService

import (
	context "context"
	reflect "reflect"

	domain "github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockActorsService is a mock of ActorsService interface.
type MockActorsService struct {
	ctrl     *gomock.Controller
	recorder *MockActorsServiceMockRecorder
}

// MockActorsServiceMockRecorder is the mock recorder for MockActorsService.
type MockActorsServiceMockRecorder struct {
	mock *MockActorsService
}

// NewMockActorsService creates a new mock instance.
func NewMockActorsService(ctrl *gomock.Controller) *MockActorsService {
	mock := &MockActorsService{ctrl: ctrl}
	mock.recorder = &MockActorsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorsService) EXPECT() *MockActorsServiceMockRecorder {
	return m.recorder
}

// GetActorByUuid mocks base method.
func (m *MockActorsService) GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByUuid", ctx, actorUuid)
	ret0, _ := ret[0].(domain.ActorData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByUuid indicates an expected call of GetActorByUuid.
func (mr *MockActorsServiceMockRecorder) GetActorByUuid(ctx, actorUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
<<<<<<< HEAD
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByUuid",
		reflect.TypeOf((*MockActorsService)(nil).GetActorByUuid), ctx, actorUuid)
=======
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByUuid", reflect.TypeOf((*MockActorsService)(nil).GetActorByUuid), ctx, actorUuid)
>>>>>>> cc029ef (handlers-tests)
}

// GetActorsByFilm mocks base method.
func (m *MockActorsService) GetActorsByFilm(ctx context.Context, filmUuid string) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByFilm", ctx, filmUuid)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsByFilm indicates an expected call of GetActorsByFilm.
func (mr *MockActorsServiceMockRecorder) GetActorsByFilm(ctx, filmUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
<<<<<<< HEAD
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByFilm",
		reflect.TypeOf((*MockActorsService)(nil).GetActorsByFilm), ctx, filmUuid)
=======
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByFilm", reflect.TypeOf((*MockActorsService)(nil).GetActorsByFilm), ctx, filmUuid)
>>>>>>> cc029ef (handlers-tests)
}
