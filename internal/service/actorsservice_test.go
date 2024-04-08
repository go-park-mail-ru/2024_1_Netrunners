package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	database "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/postgres"
	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/service/mock"
)

func TestActorsService_GetActorByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockActorsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedActor := domain.ActorData{
		Uuid:       "1",
		Name:       "Danya",
		Avatar:     "http://avatar",
		Birthday:   database.NewMockActor().Birthday,
		Career:     "career",
		Height:     192,
		BirthPlace: "Angarsk",
		Genres:     "Riddim",
		Spouse:     "Дабстеп",
		Films:      database.NewMockActor().Films,
	}

	mockStorage.EXPECT().GetActorByUuid("1").Return(expectedActor, nil)

	actorsService := NewActorsService(mockStorage, mockLogger)
	actor, err := actorsService.GetActorByUuid(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
}

func TestActorsService_GetActorsByFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockActorsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedActors := database.NewMockActorPreview()
	mockStorage.EXPECT().GetActorsByFilm("1").Return(expectedActors, nil)

	actorsService := NewActorsService(mockStorage, mockLogger)
	actors, err := actorsService.GetActorsByFilm(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedActors, actors)
}

func TestActorsService_GetActorByUuid_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockActorsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedError := errors.New("storage error")
	mockStorage.EXPECT().GetActorByUuid("1").Return(domain.ActorData{}, expectedError)

	actorsService := NewActorsService(mockStorage, mockLogger)
	actor, err := actorsService.GetActorByUuid(context.Background(), "1")

	assert.Error(t, err)
	assert.Equal(t, domain.ActorData{}, actor)
}

func TestActorsService_GetActorsByFilm_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockActorsStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	expectedError := errors.New("storage error")
	mockStorage.EXPECT().GetActorsByFilm("1").Return([]domain.ActorPreview{}, expectedError)

	actorsService := NewActorsService(mockStorage, mockLogger)
	actors, err := actorsService.GetActorsByFilm(context.Background(), "1")

	assert.Error(t, err)
	assert.Empty(t, actors)
}
