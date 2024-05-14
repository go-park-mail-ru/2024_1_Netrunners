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

// AddComment mocks base method.
func (m *MockFilmsService) AddComment(ctx context.Context, comment domain.CommentToAdd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", ctx, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddComment indicates an expected call of AddComment.
func (mr *MockFilmsServiceMockRecorder) AddComment(ctx, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockFilmsService)(nil).AddComment), ctx, comment)
}

// AddFilm mocks base method.
func (m *MockFilmsService) AddFilm(ctx context.Context, film domain.FilmToAdd) error {
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

// FindActorsLong mocks base method.
func (m *MockFilmsService) FindActorsLong(ctx context.Context, name string, page int) (domain.SearchActors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActorsLong", ctx, name, page)
	ret0, _ := ret[0].(domain.SearchActors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActorsLong indicates an expected call of FindActorsLong.
func (mr *MockFilmsServiceMockRecorder) FindActorsLong(ctx, name, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActorsLong", reflect.TypeOf((*MockFilmsService)(nil).FindActorsLong), ctx, name, page)
}

// FindActorsShort mocks base method.
func (m *MockFilmsService) FindActorsShort(ctx context.Context, name string, page int) ([]domain.ActorPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActorsShort", ctx, name, page)
	ret0, _ := ret[0].([]domain.ActorPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActorsShort indicates an expected call of FindActorsShort.
func (mr *MockFilmsServiceMockRecorder) FindActorsShort(ctx, name, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActorsShort", reflect.TypeOf((*MockFilmsService)(nil).FindActorsShort), ctx, name, page)
}

// FindFilmsLong mocks base method.
func (m *MockFilmsService) FindFilmsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsLong", ctx, title, page)
	ret0, _ := ret[0].(domain.SearchFilms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilmsLong indicates an expected call of FindFilmsLong.
func (mr *MockFilmsServiceMockRecorder) FindFilmsLong(ctx, title, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsLong", reflect.TypeOf((*MockFilmsService)(nil).FindFilmsLong), ctx, title, page)
}

// FindFilmsShort mocks base method.
func (m *MockFilmsService) FindFilmsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsShort", ctx, title, page)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilmsShort indicates an expected call of FindFilmsShort.
func (mr *MockFilmsServiceMockRecorder) FindFilmsShort(ctx, title, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsShort", reflect.TypeOf((*MockFilmsService)(nil).FindFilmsShort), ctx, title, page)
}

// FindSerialsLong mocks base method.
func (m *MockFilmsService) FindSerialsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSerialsLong", ctx, title, page)
	ret0, _ := ret[0].(domain.SearchFilms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSerialsLong indicates an expected call of FindSerialsLong.
func (mr *MockFilmsServiceMockRecorder) FindSerialsLong(ctx, title, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSerialsLong", reflect.TypeOf((*MockFilmsService)(nil).FindSerialsLong), ctx, title, page)
}

// FindSerialsShort mocks base method.
func (m *MockFilmsService) FindSerialsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSerialsShort", ctx, title, page)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSerialsShort indicates an expected call of FindSerialsShort.
func (mr *MockFilmsServiceMockRecorder) FindSerialsShort(ctx, title, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSerialsShort", reflect.TypeOf((*MockFilmsService)(nil).FindSerialsShort), ctx, title, page)
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

// GetAllFilmsByGenre mocks base method.
func (m *MockFilmsService) GetAllFilmsByGenre(ctx context.Context, genreUuid string) ([]domain.FilmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilmsByGenre", ctx, genreUuid)
	ret0, _ := ret[0].([]domain.FilmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilmsByGenre indicates an expected call of GetAllFilmsByGenre.
func (mr *MockFilmsServiceMockRecorder) GetAllFilmsByGenre(ctx, genreUuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilmsByGenre", reflect.TypeOf((*MockFilmsService)(nil).GetAllFilmsByGenre), ctx, genreUuid)
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

// GetAllGenres mocks base method.
func (m *MockFilmsService) GetAllGenres(ctx context.Context) ([]domain.GenreFilms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGenres", ctx)
	ret0, _ := ret[0].([]domain.GenreFilms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGenres indicates an expected call of GetAllGenres.
func (mr *MockFilmsServiceMockRecorder) GetAllGenres(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGenres", reflect.TypeOf((*MockFilmsService)(nil).GetAllGenres), ctx)
}

// GetFilmDataByUuid mocks base method.
func (m *MockFilmsService) GetFilmDataByUuid(ctx context.Context, uuid string) (domain.CommonFilmData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmDataByUuid", ctx, uuid)
	ret0, _ := ret[0].(domain.CommonFilmData)
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

// GetTopFilms mocks base method.
func (m *MockFilmsService) GetTopFilms(ctx context.Context) ([]domain.TopFilm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopFilms", ctx)
	ret0, _ := ret[0].([]domain.TopFilm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopFilms indicates an expected call of GetTopFilms.
func (mr *MockFilmsServiceMockRecorder) GetTopFilms(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopFilms", reflect.TypeOf((*MockFilmsService)(nil).GetTopFilms), ctx)
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

// RemoveComment mocks base method.
func (m *MockFilmsService) RemoveComment(ctx context.Context, comment domain.CommentToRemove) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveComment", ctx, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveComment indicates an expected call of RemoveComment.
func (mr *MockFilmsServiceMockRecorder) RemoveComment(ctx, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveComment", reflect.TypeOf((*MockFilmsService)(nil).RemoveComment), ctx, comment)
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
