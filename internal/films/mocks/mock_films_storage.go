// Code generated by MockGen. DO NOT EDIT.
// Source: internal/films/service/filmsservice.go

// Package mocks is a generated GoMock package.
package mocks

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

// FindActorsLong mocks base method.
func (m *MockFilmsStorage) FindActorsLong(name string, page int) ([]domain.ActorData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActorsLong", name, page)
	ret0, _ := ret[0].([]domain.ActorData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActorsLong indicates an expected call of FindActorsLong.
func (mr *MockFilmsStorageMockRecorder) FindActorsLong(name, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActorsLong", reflect.TypeOf((*MockFilmsStorage)(nil).FindActorsLong), name, page)
}

// FindActorsShort mocks base method.
func (m *MockFilmsStorage) FindActorsShort(name string, page int) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActorsShort", name, page)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActorsShort indicates an expected call of FindActorsShort.
func (mr *MockFilmsStorageMockRecorder) FindActorsShort(name, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActorsShort", reflect.TypeOf((*MockFilmsStorage)(nil).FindActorsShort), name, page)
}

// FindFilmsLong mocks base method.
func (m *MockFilmsStorage) FindFilmsLong(title string, page int) ([]domain.FilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsLong", title, page)
	ret0, _ := ret[0].([]domain.FilmData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilmsLong indicates an expected call of FindFilmsLong.
func (mr *MockFilmsStorageMockRecorder) FindFilmsLong(title, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsLong", reflect.TypeOf((*MockFilmsStorage)(nil).FindFilmsLong), title, page)
}

// FindFilmsShort mocks base method.
func (m *MockFilmsStorage) FindFilmsShort(title string, page int) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsShort", title, page)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilmsShort indicates an expected call of FindFilmsShort.
func (mr *MockFilmsStorageMockRecorder) FindFilmsShort(title, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsShort", reflect.TypeOf((*MockFilmsStorage)(nil).FindFilmsShort), title, page)
}

// FindSerialsLong mocks base method.
func (m *MockFilmsStorage) FindSerialsLong(title string, page int) ([]domain.FilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSerialsLong", title, page)
	ret0, _ := ret[0].([]domain.FilmData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSerialsLong indicates an expected call of FindSerialsLong.
func (mr *MockFilmsStorageMockRecorder) FindSerialsLong(title, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSerialsLong", reflect.TypeOf((*MockFilmsStorage)(nil).FindSerialsLong), title, page)
}

// FindSerialsShort mocks base method.
func (m *MockFilmsStorage) FindSerialsShort(title string, page int) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSerialsShort", title, page)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSerialsShort indicates an expected call of FindSerialsShort.
func (mr *MockFilmsStorageMockRecorder) FindSerialsShort(title, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSerialsShort", reflect.TypeOf((*MockFilmsStorage)(nil).FindSerialsShort), title, page)
}

// GetActorByUuid mocks base method.
func (m *MockFilmsStorage) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByUuid", actorUuid)
	ret0, _ := ret[0].(domain.ActorData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByUuid indicates an expected call of GetActorByUuid.
func (mr *MockFilmsStorageMockRecorder) GetActorByUuid(actorUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByUuid", reflect.TypeOf((*MockFilmsStorage)(nil).GetActorByUuid), actorUuid)
}

// GetActorsByFilm mocks base method.
func (m *MockFilmsStorage) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByFilm", filmUuid)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsByFilm indicates an expected call of GetActorsByFilm.
func (mr *MockFilmsStorageMockRecorder) GetActorsByFilm(filmUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByFilm", reflect.TypeOf((*MockFilmsStorage)(nil).GetActorsByFilm), filmUuid)
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

// GetAllFilmsByGenre mocks base method.
func (m *MockFilmsStorage) GetAllFilmsByGenre(genreUuid string) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmsByGenre", genreUuid)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmsByGenre indicates an expected call of GetAllFilmsByGenre.
func (mr *MockFilmsStorageMockRecorder) GetAllFilmsByGenre(genreUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmsByGenre", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllFilmsByGenre), genreUuid)
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

// GetAllGenres mocks base method.
func (m *MockFilmsStorage) GetAllGenres() ([]domain.GenreFilms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGenres")
	ret0, _ := ret[0].([]domain.GenreFilms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGenres indicates an expected call of GetAllGenres.
func (mr *MockFilmsStorageMockRecorder) GetAllGenres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGenres", reflect.TypeOf((*MockFilmsStorage)(nil).GetAllGenres))
}

// GetFilmDataByUuid mocks base method.
func (m *MockFilmsStorage) GetFilmDataByUuid(uuid string) (domain.CommonFilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmDataByUuid", uuid)
	ret0, _ := ret[0].(domain.CommonFilmData)
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
