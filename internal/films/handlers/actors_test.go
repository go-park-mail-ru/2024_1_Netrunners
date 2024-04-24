package films_handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/mocks"
)

func TestActorsHandlers_GetActorByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActorsService := mockService.NewMockActorsService(ctrl)

	actorData := domain.ActorData{
		Uuid: "1",
		Name: "Danya",
	}

	mockActorsService.EXPECT().GetActorByUuid(gomock.Any(), "1").Return(actorData, nil)

	mockLogger := zaptest.NewLogger(t).Sugar()
	actorsHandlers := NewActorsHandlers(mockActorsService, mockLogger)

	req := httptest.NewRequest("GET", "/actors/1", nil)
	req = mux.SetURLVars(req, map[string]string{"uuid": "1"})
	w := httptest.NewRecorder()

	actorsHandlers.GetActorByUuid(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response actorResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, actorData.Uuid, response.Actor.Uuid)
	assert.Equal(t, actorData.Name, response.Actor.Name)

	actorData2 := domain.ActorData{
		Uuid: "",
		Name: "",
	}

	mockActorsService.EXPECT().GetActorByUuid(gomock.Any(), "1").Return(actorData, errors.New(""))

	actorsHandlers = NewActorsHandlers(mockActorsService, mockLogger)

	req = httptest.NewRequest("GET", "/actors/1", nil)
	req = mux.SetURLVars(req, map[string]string{"uuid": "1"})
	w = httptest.NewRecorder()

	actorsHandlers.GetActorByUuid(w, req)

	resp = w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response2 actorResponse
	err = json.NewDecoder(resp.Body).Decode(&response2)

	assert.Equal(t, 500, response2.Status)
	assert.Equal(t, actorData2.Uuid, response2.Actor.Uuid)
	assert.Equal(t, actorData2.Name, response2.Actor.Name)
}

func TestActorsHandlers_GetActorsByFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActorsService := mockService.NewMockActorsService(ctrl)

	actorsPreview := []domain.ActorPreview{
		{Uuid: "1", Name: "Danya"},
		{Uuid: "2", Name: "Dima"},
	}

	mockActorsService.EXPECT().GetActorsByFilm(gomock.Any(), "film_uuid").Return(actorsPreview, nil)

	mockLogger := zaptest.NewLogger(t).Sugar()
	actorsHandlers := NewActorsHandlers(mockActorsService, mockLogger)

	req := httptest.NewRequest("GET", "/actors/film_uuid", nil)
	req = mux.SetURLVars(req, map[string]string{"uuid": "film_uuid"})
	w := httptest.NewRecorder()

	actorsHandlers.GetActorsByFilm(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response actorsByFilmResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, len(actorsPreview), len(response.Actors))

	for i, actor := range response.Actors {
		assert.Equal(t, actorsPreview[i].Uuid, actor.Uuid)
		assert.Equal(t, actorsPreview[i].Name, actor.Name)
	}
}
