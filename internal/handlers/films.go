package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type FilmsService interface {
	GetFilmDataByUuid(ctx context.Context, uuid string) (domain.FilmData, error)
	AddFilm(ctx context.Context, film domain.FilmDataToAdd) error
	RemoveFilm(ctx context.Context, uuid string) error
	GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error)
	GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error)
	GetAllFilmActors(ctx context.Context, uuid string) ([]domain.ActorPreview, error)
	PutFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error
	RemoveFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error
	GetAllFavoriteFilms(ctx context.Context, userUuid string) ([]domain.FilmPreview, error)
}

type FilmsPageHandlers struct {
	filmsService FilmsService
	logger       *zap.SugaredLogger
}

func NewFilmsPageHandlers(filmsService FilmsService, logger *zap.SugaredLogger) *FilmsPageHandlers {
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
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
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

type dataToFavorite struct {
	FilmUuid string `json:"filmUuid"`
	UserUuid string `json:"userUuid"`
}

func (filmsPageHandlers *FilmsPageHandlers) PutFavoriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	var data dataToFavorite
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	err = filmsPageHandlers.filmsService.PutFavoriteFilm(ctx, data.FilmUuid, data.UserUuid)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to put favorite film: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	err = WriteSuccess(w)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film added successfully", requestId))
}

func (filmsPageHandlers *FilmsPageHandlers) RemoveFavoriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	var data dataToFavorite
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	err = filmsPageHandlers.filmsService.RemoveFavoriteFilm(ctx, data.FilmUuid, data.UserUuid)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to remove favorite film: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	err = WriteSuccess(w)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film removed successfully", requestId))
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFavoriteFilms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)
	uuid := mux.Vars(r)["uuid"]

	films, err := filmsPageHandlers.filmsService.GetAllFavoriteFilms(ctx, uuid)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all favorite film: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
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
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
}
