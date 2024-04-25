package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/mocks"
)

func TestGetFilmDataByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockFilmData := domain.FilmData{
		Uuid:     uuid,
		Title:    "Mock Title",
		Preview:  "Mock Preview",
		Director: "Mock Director",
		Data:     "Mock Data",
		AgeLimit: 18,
		Duration: 120,
	}

	mockStorage.EXPECT().GetFilmDataByUuid(uuid).Return(mockFilmData, nil)

	filmData, err := service.GetFilmDataByUuid(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, mockFilmData, filmData)
}

func TestGetFilmDataByUuid_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "1"
	mockError := errors.New("mock error")

	mockStorage.EXPECT().GetFilmDataByUuid(uuid).Return(domain.FilmData{}, mockError)

	_, err := service.GetFilmDataByUuid(context.Background(), uuid)

	assert.Error(t, err)
}

func TestAddFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	mockFilmData := domain.FilmDataToAdd{
		Title:    "Mock Title",
		Preview:  "Mock Preview",
		Director: "Mock Director",
		Data:     "Mock Data",
		AgeLimit: 18,
		Duration: 120,
	}

	mockStorage.EXPECT().AddFilm(mockFilmData).Return(nil)

	err := service.AddFilm(context.Background(), mockFilmData)

	assert.NoError(t, err)
}

func TestAddFilm_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	mockFilmData := domain.FilmDataToAdd{
		Title:    "Mock Title",
		Preview:  "Mock Preview",
		Director: "Mock Director",
		Data:     "Mock Data",
		AgeLimit: 18,
		Duration: 120,
	}

	mockError := errors.New("mock error")

	mockStorage.EXPECT().AddFilm(mockFilmData).Return(mockError)

	err := service.AddFilm(context.Background(), mockFilmData)

	assert.Error(t, err)
}

func TestRemoveFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"

	mockStorage.EXPECT().RemoveFilm(uuid).Return(nil)

	err := service.RemoveFilm(context.Background(), uuid)

	assert.NoError(t, err)
}

func TestRemoveFilm_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockError := errors.New("mock error")

	mockStorage.EXPECT().RemoveFilm(uuid).Return(mockError)

	err := service.RemoveFilm(context.Background(), uuid)

	assert.Error(t, err)
}

func TestGetFilmPreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockFilmPreview := domain.FilmPreview{
		Uuid:         uuid,
		Title:        "Mock Title",
		Preview:      "Mock Preview",
		Director:     "Mock Director",
		AverageScore: 4.5,
		ScoresCount:  100,
		Duration:     120,
	}

	mockStorage.EXPECT().GetFilmPreview(uuid).Return(mockFilmPreview, nil)

	filmPreview, err := service.GetFilmPreview(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, mockFilmPreview, filmPreview)
}

func TestGetFilmPreview_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockError := errors.New("mock error")

	mockStorage.EXPECT().GetFilmPreview(uuid).Return(domain.FilmPreview{}, mockError)

	_, err := service.GetFilmPreview(context.Background(), uuid)

	assert.Error(t, err)
}

func TestGetAllFilmsPreviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	mockFilmPreviews := []domain.FilmPreview{
		{Uuid: "1", Title: "Mock Title 1"},
		{Uuid: "2", Title: "Mock Title 2"},
	}

	mockStorage.EXPECT().GetAllFilmsPreviews().Return(mockFilmPreviews, nil)

	filmPreviews, err := service.GetAllFilmsPreviews(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, mockFilmPreviews, filmPreviews)
}

func TestGetAllFilmsPreviews_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	mockError := errors.New("mock error")

	mockStorage.EXPECT().GetAllFilmsPreviews().Return(nil, mockError)

	_, err := service.GetAllFilmsPreviews(context.Background())

	assert.Error(t, err)
}

func TestGetAllFilmComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockComments := []domain.Comment{
		{Uuid: "1", FilmUuid: uuid, Text: "Comment 1"},
		{Uuid: "2", FilmUuid: uuid, Text: "Comment 2"},
	}

	mockStorage.EXPECT().GetAllFilmComments(uuid).Return(mockComments, nil)

	comments, err := service.GetAllFilmComments(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, mockComments, comments)
}

func TestGetAllFilmComments_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockError := errors.New("mock error")

	mockStorage.EXPECT().GetAllFilmComments(uuid).Return(nil, mockError)

	_, err := service.GetAllFilmComments(context.Background(), uuid)

	assert.Error(t, err)
}

func TestGetAllFilmActors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockActors := []domain.ActorPreview{
		{Uuid: "1", Name: "Actor 1"},
		{Uuid: "2", Name: "Actor 2"},
	}

	mockStorage.EXPECT().GetActorsByFilm(uuid).Return(mockActors, nil)

	actors, err := service.GetActorsByFilm(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, mockActors, actors)
}

func TestGetAllFilmActors_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	service := NewFilmsService(mockStorage, mockLogger, "")

	uuid := "123"
	mockError := errors.New("mock error")

	mockStorage.EXPECT().GetActorsByFilm(uuid).Return(nil, mockError)

	_, err := service.GetActorsByFilm(context.Background(), uuid)

	assert.Error(t, err)
}

func TestActorsService_GetActorByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedActor := domain.ActorData{
		Uuid:       "1",
		Name:       "Danya",
		Avatar:     "http://avatar",
		Birthday:   mocks.NewMockActor().Birthday,
		Career:     "career",
		Height:     192,
		BirthPlace: "Angarsk",
		Genres:     "Riddim",
		Spouse:     "Дабстеп",
		Films:      mocks.NewMockActor().Films,
	}

	mockStorage.EXPECT().GetActorByUuid("1").Return(expectedActor, nil)

	service := NewFilmsService(mockStorage, mockLogger, "")
	actor, err := service.GetActorByUuid(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
}

func TestActorsService_GetActorsByFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedActors := mocks.NewMockActorPreview()
	mockStorage.EXPECT().GetActorsByFilm("1").Return(expectedActors, nil)

	service := NewFilmsService(mockStorage, mockLogger, "")
	actors, err := service.GetActorsByFilm(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedActors, actors)
}

func TestActorsService_GetActorByUuid_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedError := errors.New("storage error")
	mockStorage.EXPECT().GetActorByUuid("1").Return(domain.ActorData{}, expectedError)

	service := NewFilmsService(mockStorage, mockLogger, "")
	actor, err := service.GetActorByUuid(context.Background(), "1")

	assert.Error(t, err)
	assert.Equal(t, domain.ActorData{}, actor)
}

func TestActorsService_GetActorsByFilm_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockFilmsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedError := errors.New("storage error")
	mockStorage.EXPECT().GetActorsByFilm("1").Return([]domain.ActorPreview{}, expectedError)

	service := NewFilmsService(mockStorage, mockLogger, "")
	actors, err := service.GetActorsByFilm(context.Background(), "1")

	assert.Error(t, err)
	assert.Empty(t, actors)
}
