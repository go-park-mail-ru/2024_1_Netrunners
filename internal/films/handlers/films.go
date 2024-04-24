package films_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	api "github.com/go-park-mail-ru/2024_1_Netrunners/cmd/films"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type FilmsServer interface {
	GetFilmDataByUuid(ctx context.Context, uuid string) (domain.FilmData, error)
	AddFilm(ctx context.Context, film domain.FilmDataToAdd) error
	RemoveFilm(ctx context.Context, uuid string) error
	GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error)
	GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error)
	GetAllFilmActors(ctx context.Context, uuid string) ([]domain.ActorPreview, error)
}

type FilmsPageHandlers struct {
	server *api.FilmsServer
	logger *zap.SugaredLogger
}

func NewFilmsPageHandlers(server *api.FilmsServer, logger *zap.SugaredLogger) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		server: server,
		logger: logger,
	}
}

type filmsPreviewsResponse struct {
	Status int                  `json:"status"`
	Films  []domain.FilmPreview `json:"films"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmsPreviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var req *session.AllFilmsPreviewsRequest
	res, err := filmsPageHandlers.server.GetAllFilmsPreviews(ctx, req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var filmsRegular []domain.FilmPreview
	for _, film := range res.Films {
		filmRegular := convertFilmPreviewToRegular(film)
		escapeFilmPreview(&filmRegular)
		filmsRegular = append(filmsRegular, convertFilmPreviewToRegular(film))
	}

	response := filmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  filmsRegular,
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

	req := &session.FilmDataByUuidRequest{
		Uuid: uuid,
	}
	filmData, err := filmsPageHandlers.server.GetFilmDataByUuid(ctx, req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	filmDataRegular := convertFilmDataToRegular(filmData.FilmData)
	escapeFilmData(&filmDataRegular)

	response := filmDataResponse{
		Status:   http.StatusOK,
		FilmData: filmDataRegular,
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

	comments, err := filmsPageHandlers.server.GetAllFilmComments(ctx, uuid)
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

	actors, err := filmsPageHandlers.server.GetAllFilmActors(ctx, uuid)
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

	err = filmsPageHandlers.server.AddFilm(ctx, filmData)
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

type actorResponse struct {
	Status int              `json:"status"`
	Actor  domain.ActorData `json:"actor"`
}

func (actorsHandlers *ActorsHandlers) GetActorByUuid(w http.ResponseWriter, r *http.Request) {
	actorUuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	actor, err := actorsHandlers.actorsService.server(ctx, actorUuid)
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
	requestID := ctx.Value(reqid.ReqIDKey)

	actors, err := actorsHandlers.actorsService.server(ctx, filmUuid)
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

func convertFilmPreviewToRegular(film *session.FilmPreview) domain.FilmPreview {
	return domain.FilmPreview{
		Uuid:         film.Uuid,
		Title:        film.Title,
		Preview:      film.Preview,
		Director:     film.Director,
		AverageScore: film.AvgScore,
		ScoresCount:  film.ScoresCount,
		AgeLimit:     film.AgeLimit,
	}
}

func convertFilmDataToRegular(film *session.FilmData) domain.FilmData {
	return domain.FilmData{
		Uuid:         film.Uuid,
		Title:        film.Title,
		Preview:      film.Preview,
		Director:     film.Director,
		Link:         film.Link,
		Data:         film.Data,
		Date:         film.Date,
		AgeLimit:     film.AgeLimit,
		AverageScore: film.AvgScore,
		ScoresCount:  film.ScoresCount,
		Duration:     film.Duration,
	}
}
