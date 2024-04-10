package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type FilmsPageHandlers struct {
	filmsService *service.FilmsService
	logger       *zap.SugaredLogger
}

func NewFilmsPageHandlers(filmsService *service.FilmsService, logger *zap.SugaredLogger) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		filmsService: filmsService,
		logger:       logger,
	}
}

type filmsPreviewsResponse struct {
	Status int                  `json:"status"`
	Films  []domain.FilmPreview `json:"films"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmsPreviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	films, err := filmsPageHandlers.filmsService.GetAllFilmsPreviews(ctx)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
	}

	for _, film := range films {
		escapeFilmPreview(&film)
	}

	response := filmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  films,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal: %v\n", requestID, err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

type filmDataResponse struct {
	Status   int             `json:"status"`
	FilmData domain.FilmData `json:"film"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetFilmDataByUuid(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	filmData, err := filmsPageHandlers.filmsService.GetFilmDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	escapeFilmData(&filmData)

	response := filmDataResponse{
		Status:   http.StatusOK,
		FilmData: filmData,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

type filmCommentsResponse struct {
	Status   int              `json:"status"`
	Comments []domain.Comment `json:"comments"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmComments(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	comments, err := filmsPageHandlers.filmsService.GetAllFilmComments(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	for _, comment := range comments {
		escapeComment(&comment)
	}

	response := filmCommentsResponse{
		Status:   http.StatusOK,
		Comments: comments,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

}

type filmActorsResponse struct {
	Status int                   `json:"status"`
	Actors []domain.ActorPreview `json:"actors"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmActors(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	actors, err := filmsPageHandlers.filmsService.GetAllFilmActors(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}

		return
	}

	for _, actor := range actors {
		escapeActorPreview(&actor)
	}

	response := filmActorsResponse{
		Status: http.StatusOK,
		Actors: actors,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) AddFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var filmData domain.FilmDataToAdd
	err := json.NewDecoder(r.Body).Decode(&filmData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = filmsPageHandlers.filmsService.AddFilm(ctx, filmData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = WriteSuccess(w)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] film added successfully", requestID))
}
