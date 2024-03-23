package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type actorData struct {
	Name     string     `json:"name"`
	Data     string     `json:"data"`
	Avatar   string     `json:"avatar"`
	Birthday time.Time  `json:"birthday"`
	Films    []filmLink `json:"films"`
}

type actorResponse struct {
	Status int       `json:"status"`
	Actor  actorData `json:"actor"`
}

type ActorsHandlers struct {
	actorsService *service.ActorsService
	logger        *zap.SugaredLogger
}

func NewActorsHandlers(actorsService *service.ActorsService, logger *zap.SugaredLogger) *ActorsHandlers {
	return &ActorsHandlers{
		actorsService: actorsService,
		logger:        logger,
	}
}

func (actorsHandlers *ActorsHandlers) GetActorByUuid(w http.ResponseWriter, r *http.Request) {
	actorUuid := mux.Vars(r)["uuid"]

	actor, err := actorsHandlers.actorsService.GetActorByUuid(actorUuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	actorFilms := make([]filmLink, 0, len(actor.Films))

	for _, film := range actor.Films {
		actorFilms = append(actorFilms, filmLink{
			Uuid:  film.Uuid,
			Title: film.Title,
		})
	}

	response := actorResponse{
		Status: http.StatusOK,
		Actor: actorData{
			Name:     actor.Name,
			Data:     actor.Data,
			Birthday: actor.Birthday,
			Films:    actorFilms,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		actorsHandlers.logger.Errorf("error at marshalling: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		actorsHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}
