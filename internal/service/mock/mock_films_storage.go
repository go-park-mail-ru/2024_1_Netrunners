// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/filmsservice.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockFilmsStorage is a mock of FilmsStorage interface.
type MockFilmsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFilmsStorageMockRecorder
}

// MockFilmsStorageMockRecorder is the mock recorder for MockFilmsStorage.
type MockFilmsStorageMockRecorder struct {
	mock *MockFilmsStorage
}

// NewMockFilmsStorage creates a new mock instance.
func NewMockFilmsStorage(ctrl *gomock.Controller) *MockFilmsStorage {
	mock := &MockFilmsStorage{ctrl: ctrl}
	mock.recorder = &MockFilmsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmsStorage) EXPECT() *MockFilmsStorageMockRecorder {
	return m.recorder
}

// AddFilm mocks base method.
func (m *MockFilmsStorage) AddFilm(film domain.FilmDataToAdd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", film)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockFilmsStorageMockRecorder) AddFilm(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockFilmsStorage)(nil).AddFilm), film)
}

// GetAllFavoriteFilms mocks base method.
func (m *MockFilmsStorage) GetAllFavoriteFilms(userUuid string) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFavoriteFilms", userUuid)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFavoriteFilms indicates an expected call of GetAllFavoriteFilms.
func (mr *MockFilmsStorageMockRecorder) GetAllFavoriteFilms(userUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFavoriteFilms", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllFavoriteFilms), userUuid)
}

// GetAllFilmActors mocks base method.
func (m *MockFilmsStorage) GetAllFilmActors(uuid string) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmActors", uuid)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmActors indicates an expected call of GetAllFilmActors.
func (mr *MockFilmsStorageMockRecorder) GetAllFilmActors(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmActors", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllFilmActors), uuid)
}

// GetAllFilmComments mocks base method.
func (m *MockFilmsStorage) GetAllFilmComments(uuid string) ([]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmComments", uuid)
	ret0, _ := ret[0].([]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmComments indicates an expected call of GetAllFilmComments.
func (mr *MockFilmsStorageMockRecorder) GetAllFilmComments(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmComments", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllFilmComments), uuid)
}

// GetAllFilmsPreviews mocks base method.
func (m *MockFilmsStorage) GetAllFilmsPreviews() ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmsPreviews")
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmsPreviews indicates an expected call of GetAllFilmsPreviews.
func (mr *MockFilmsStorageMockRecorder) GetAllFilmsPreviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmsPreviews", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllFilmsPreviews))
}

// GetFilmDataByUuid mocks base method.
func (m *MockFilmsStorage) GetFilmDataByUuid(uuid string) (domain.FilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmDataByUuid", uuid)
	ret0, _ := ret[0].(domain.FilmData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmDataByUuid indicates an expected call of GetFilmDataByUuid.
func (mr *MockFilmsStorageMockRecorder) GetFilmDataByUuid(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmDataByUuid", reflect.TypeOf((*MockFilmsStorage)(nil).GetFilmDataByUuid), uuid)
}

// GetFilmPreview mocks base method.
func (m *MockFilmsStorage) GetFilmPreview(uuid string) (domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmPreview", uuid)
	ret0, _ := ret[0].(domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmPreview indicates an expected call of GetFilmPreview.
func (mr *MockFilmsStorageMockRecorder) GetFilmPreview(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmPreview", reflect.TypeOf((*MockFilmsStorage)(nil).GetFilmPreview), uuid)
}

// PutFavoriteFilm mocks base method.
func (m *MockFilmsStorage) PutFavoriteFilm(filmUuid, userUuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutFavoriteFilm", filmUuid, userUuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutFavoriteFilm indicates an expected call of PutFavoriteFilm.
func (mr *MockFilmsStorageMockRecorder) PutFavoriteFilm(filmUuid, userUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutFavoriteFilm", reflect.TypeOf((*MockFilmsStorage)(nil).PutFavoriteFilm), filmUuid, userUuid)
}

// RemoveFavoriteFilm mocks base method.
func (m *MockFilmsStorage) RemoveFavoriteFilm(filmUuid, userUuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFavoriteFilm", filmUuid, userUuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFavoriteFilm indicates an expected call of RemoveFavoriteFilm.
func (mr *MockFilmsStorageMockRecorder) RemoveFavoriteFilm(filmUuid, userUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFavoriteFilm", reflect.TypeOf((*MockFilmsStorage)(nil).RemoveFavoriteFilm), filmUuid, userUuid)
}

// RemoveFilm mocks base method.
func (m *MockFilmsStorage) RemoveFilm(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFilm", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFilm indicates an expected call of RemoveFilm.
func (mr *MockFilmsStorageMockRecorder) RemoveFilm(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFilm", reflect.TypeOf((*MockFilmsStorage)(nil).RemoveFilm), uuid)
}
