package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type actorResponse struct {
	Status int              `json:"status"`
	Actor  domain.ActorData `json:"actor"`
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

	actorFilms := make([]domain.FilmLink, 0, len(actor.Films))

	for _, film := range actor.Films {
		actorFilms = append(actorFilms, domain.FilmLink{
			Uuid:  film.Uuid,
			Title: film.Title,
		})
	}

	escapeActorData(&actor)

	response := actorResponse{
		Status: http.StatusOK,
		Actor: domain.ActorData{
			Uuid:       actorUuid,
			Name:       actor.Name,
			Avatar:     actor.Avatar,
			Birthday:   actor.Birthday,
			Career:     actor.Career,
			Height:     actor.Height,
			BirthPlace: actor.BirthPlace,
			Genres:     actor.Genres,
			Spouse:     actor.Spouse,
			Films:      actorFilms,
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

type actorsByFilmResponse struct {
	Status int
	Actors []domain.ActorPreview
}

func (actorsHandlers *ActorsHandlers) GetActorsByFilm(w http.ResponseWriter, r *http.Request) {
	filmUuid := mux.Vars(r)["uuid"]

	actors, err := actorsHandlers.actorsService.GetActorsByFilm(filmUuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	response := actorsByFilmResponse{
		Status: http.StatusOK,
		Actors: actors,
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
