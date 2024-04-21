package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	httpctx "github.com/go-park-mail-ru/2024_1_Netrunners/internal/httpcontext"
)

type ActorsService interface {
	GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(ctx context.Context, filmUuid string) ([]domain.ActorPreview, error)
}

type ActorsHandlers struct {
	actorsService ActorsService
	logger        *zap.SugaredLogger
}

func NewActorsHandlers(actorsService ActorsService, logger *zap.SugaredLogger) *ActorsHandlers {
	return &ActorsHandlers{
		actorsService: actorsService,
		logger:        logger,
	}
}

type actorResponse struct {
	Status int              `json:"status"`
	Actor  domain.ActorData `json:"actor"`
}

func (actorsHandlers *ActorsHandlers) GetActorByUuid(w http.ResponseWriter, r *http.Request) {
	actorUuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(httpctx.ReqIDKey)

	actor, err := actorsHandlers.actorsService.GetActorByUuid(ctx, actorUuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	escapeActorData(&actor)

	response := actorResponse{
		Status: http.StatusOK,
		Actor:  actor,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		actorsHandlers.logger.Errorf("[reqid=%s] failed to marshal: %v\n", requestID, err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		actorsHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

type actorsByFilmResponse struct {
	Status int
	Actors []domain.ActorPreview
}

func (actorsHandlers *ActorsHandlers) GetActorsByFilm(w http.ResponseWriter, r *http.Request) {
	filmUuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(httpctx.ReqIDKey)

	actors, err := actorsHandlers.actorsService.GetActorsByFilm(ctx, filmUuid)
	if err != nil {
		actorsHandlers.logger.Errorf("[reqid=%s] failed to get actors by film: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	response := actorsByFilmResponse{
		Status: http.StatusOK,
		Actors: actors,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		actorsHandlers.logger.Errorf("[reqid=%s] failed to marshal: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		actorsHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			actorsHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}
