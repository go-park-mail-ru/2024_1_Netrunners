// Code generated by MockGen. DO NOT EDIT.
// Source: internal/films/api/api.go
//
// Generated by this command:
//
//	mockgen -source=internal/films/api/api.go -destination=internal/films/mocks/mock_films_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockFilmsService is a mock of FilmsService interface.
type MockFilmsService struct {
	ctrl     *gomock.Controller
	recorder *MockFilmsServiceMockRecorder
}

// MockFilmsServiceMockRecorder is the mock recorder for MockFilmsService.
type MockFilmsServiceMockRecorder struct {
	mock *MockFilmsService
}

// NewMockFilmsService creates a new mock instance.
func NewMockFilmsService(ctrl *gomock.Controller) *MockFilmsService {
	mock := &MockFilmsService{ctrl: ctrl}
	mock.recorder = &MockFilmsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmsService) EXPECT() *MockFilmsServiceMockRecorder {
	return m.recorder
}

// AddFilm mocks base method.
func (m *MockFilmsService) AddFilm(ctx context.Context, film domain.FilmDataToAdd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", ctx, film)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockFilmsServiceMockRecorder) AddFilm(ctx, film any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockFilmsService)(nil).AddFilm), ctx, film)
}

// GetActorByUuid mocks base method.
func (m *MockFilmsService) GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByUuid", ctx, actorUuid)
	ret0, _ := ret[0].(domain.ActorData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByUuid indicates an expected call of GetActorByUuid.
func (mr *MockFilmsServiceMockRecorder) GetActorByUuid(ctx, actorUuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByUuid", reflect.TypeOf((*MockFilmsService)(nil).GetActorByUuid), ctx, actorUuid)
}

// GetActorsByFilm mocks base method.
func (m *MockFilmsService) GetActorsByFilm(ctx context.Context, uuid string) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByFilm", ctx, uuid)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsByFilm indicates an expected call of GetActorsByFilm.
func (mr *MockFilmsServiceMockRecorder) GetActorsByFilm(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByFilm", reflect.TypeOf((*MockFilmsService)(nil).GetActorsByFilm), ctx, uuid)
}

// GetAllFavoriteFilms mocks base method.
func (m *MockFilmsService) GetAllFavoriteFilms(ctx context.Context, userUuid string) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFavoriteFilms", ctx, userUuid)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFavoriteFilms indicates an expected call of GetAllFavoriteFilms.
func (mr *MockFilmsServiceMockRecorder) GetAllFavoriteFilms(ctx, userUuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFavoriteFilms", reflect.TypeOf((*MockFilmsService)(nil).GetAllFavoriteFilms), ctx, userUuid)
}

// GetAllFilmComments mocks base method.
func (m *MockFilmsService) GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmComments", ctx, uuid)
	ret0, _ := ret[0].([]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmComments indicates an expected call of GetAllFilmComments.
func (mr *MockFilmsServiceMockRecorder) GetAllFilmComments(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmComments", reflect.TypeOf((*MockFilmsService)(nil).GetAllFilmComments), ctx, uuid)
}

// GetAllFilmsPreviews mocks base method.
func (m *MockFilmsService) GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmsPreviews", ctx)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmsPreviews indicates an expected call of GetAllFilmsPreviews.
func (mr *MockFilmsServiceMockRecorder) GetAllFilmsPreviews(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmsPreviews", reflect.TypeOf((*MockFilmsService)(nil).GetAllFilmsPreviews), ctx)
}

// GetFilmDataByUuid mocks base method.
func (m *MockFilmsService) GetFilmDataByUuid(ctx context.Context, uuid string) (domain.FilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmDataByUuid", ctx, uuid)
	ret0, _ := ret[0].(domain.FilmData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmDataByUuid indicates an expected call of GetFilmDataByUuid.
func (mr *MockFilmsServiceMockRecorder) GetFilmDataByUuid(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmDataByUuid", reflect.TypeOf((*MockFilmsService)(nil).GetFilmDataByUuid), ctx, uuid)
}

// GetFilmPreview mocks base method.
func (m *MockFilmsService) GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmPreview", ctx, uuid)
	ret0, _ := ret[0].(domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmPreview indicates an expected call of GetFilmPreview.
func (mr *MockFilmsServiceMockRecorder) GetFilmPreview(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmPreview", reflect.TypeOf((*MockFilmsService)(nil).GetFilmPreview), ctx, uuid)
}

// PutFavoriteFilm mocks base method.
func (m *MockFilmsService) PutFavoriteFilm(ctx context.Context, filmUuid, userUuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutFavoriteFilm", ctx, filmUuid, userUuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutFavoriteFilm indicates an expected call of PutFavoriteFilm.
func (mr *MockFilmsServiceMockRecorder) PutFavoriteFilm(ctx, filmUuid, userUuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutFavoriteFilm", reflect.TypeOf((*MockFilmsService)(nil).PutFavoriteFilm), ctx, filmUuid, userUuid)
}

// RemoveFavoriteFilm mocks base method.
func (m *MockFilmsService) RemoveFavoriteFilm(ctx context.Context, filmUuid, userUuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFavoriteFilm", ctx, filmUuid, userUuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFavoriteFilm indicates an expected call of RemoveFavoriteFilm.
func (mr *MockFilmsServiceMockRecorder) RemoveFavoriteFilm(ctx, filmUuid, userUuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFavoriteFilm", reflect.TypeOf((*MockFilmsService)(nil).RemoveFavoriteFilm), ctx, filmUuid, userUuid)
}

// RemoveFilm mocks base method.
func (m *MockFilmsService) RemoveFilm(ctx context.Context, uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFilm", ctx, uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFilm indicates an expected call of RemoveFilm.
func (mr *MockFilmsServiceMockRecorder) RemoveFilm(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFilm", reflect.TypeOf((*MockFilmsService)(nil).RemoveFilm), ctx, uuid)
}
