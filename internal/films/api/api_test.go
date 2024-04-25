package api

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/mocks"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

func TestFilmsServer_GetAllFilmsPreviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedFilms := []domain.FilmPreview{
		{Uuid: "1", Title: "Film 1", Director: "Director 1"},
		{Uuid: "2", Title: "Film 2", Director: "Director 2"},
	}
	mockService.EXPECT().GetAllFilmsPreviews(ctx).Return(expectedFilms, nil)

	req := &session.AllFilmsPreviewsRequest{}
	resp, err := server.GetAllFilmsPreviews(ctx, req)

	require.NoError(t, err)
	assert.Len(t, resp.Films, len(expectedFilms))
}

func TestFilmsServer_GetFilmDataByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedFilm := domain.FilmData{
		Uuid:         "1",
		Title:        "Film 1",
		Preview:      "Preview 1",
		Link:         "Link 1",
		Director:     "Director 1",
		AverageScore: 4.5,
		ScoresCount:  100,
		Duration:     120,
		AgeLimit:     16,
		Date:         time.Now(),
		Data:         "data",
	}
	mockService.EXPECT().GetFilmDataByUuid(ctx, "1").Return(expectedFilm, nil)

	req := &session.FilmDataByUuidRequest{Uuid: "1"}
	resp, err := server.GetFilmDataByUuid(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp.FilmData)
}

func TestFilmsServer_GetFilmPreviewByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedFilm := domain.FilmPreview{
		Uuid:         "1",
		Title:        "Film 1",
		Preview:      "Preview 1",
		Director:     "Director 1",
		AverageScore: 4.5,
		ScoresCount:  100,
		Duration:     120,
		AgeLimit:     16,
	}
	mockService.EXPECT().GetFilmPreview(ctx, "1").Return(expectedFilm, nil)

	req := &session.FilmPreviewByUuidRequest{Uuid: "1"}
	resp, err := server.GetFilmPreviewByUuid(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp.FilmPreview)
}

func TestFilmsServer_GetAllFilmComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedComments := []domain.Comment{
		{Uuid: "1", Text: "Comment 1", FilmUuid: "1", Author: "User 1", Score: 5, AddedAt: time.Now()},
		{Uuid: "2", Text: "Comment 2", FilmUuid: "1", Author: "User 2", Score: 4, AddedAt: time.Now()},
	}
	mockService.EXPECT().GetAllFilmComments(ctx, "1").Return(expectedComments, nil)

	req := &session.AllFilmCommentsRequest{Uuid: "1"}
	resp, err := server.GetAllFilmComments(ctx, req)

	require.NoError(t, err)
	assert.Len(t, resp.Comments, len(expectedComments))
}

func TestFilmsServer_GetActorsByFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedActors := []domain.ActorPreview{
		{Uuid: "1", Name: "Actor 1", Avatar: "Avatar 1"},
		{Uuid: "2", Name: "Actor 2", Avatar: "Avatar 2"},
	}
	mockService.EXPECT().GetActorsByFilm(ctx, "1").Return(expectedActors, nil)

	req := &session.ActorsByFilmRequest{Uuid: "1"}
	resp, err := server.GetActorsByFilm(ctx, req)

	require.NoError(t, err)
	assert.Len(t, resp.Actors, len(expectedActors))
}

func TestFilmsServer_RemoveFilmByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	mockService.EXPECT().RemoveFilm(ctx, "1").Return(nil)

	req := &session.RemoveFilmByUuidRequest{Uuid: "1"}
	resp, err := server.RemoveFilmByUuid(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestFilmsServer_GetActorDataByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilmsService(ctrl)
	server := NewFilmsServer(mockService, nil)

	ctx := context.Background()
	expectedActor := domain.ActorData{
		Uuid:     "1",
		Name:     "Actor 1",
		Avatar:   "Avatar 1",
		Birthday: time.Now(),
		Career:   "Career",
		Spouse:   "Spouse",
		Genres:   "Genre 1",
		Films: []domain.FilmPreview{
			{Uuid: "1", Title: "Film 1", Director: "Director 1"},
			{Uuid: "2", Title: "Film 2", Director: "Director 2"},
		},
	}
	mockService.EXPECT().GetActorByUuid(ctx, "1").Return(expectedActor, nil)

	req := &session.ActorDataByUuidRequest{Uuid: "1"}
	resp, err := server.GetActorDataByUuid(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp.Actor)
}
