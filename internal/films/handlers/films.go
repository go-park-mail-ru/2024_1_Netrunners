package films_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type FilmsPageHandlers struct {
	client *session.FilmsClient
	logger *zap.SugaredLogger
}

func NewFilmsPageHandlers(client *session.FilmsClient, logger *zap.SugaredLogger) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		client: client,
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
	res, err := (*filmsPageHandlers.client).GetAllFilmsPreviews(ctx, req)
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
		filmsRegular = append(filmsRegular, filmRegular)
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
	filmData, err := (*filmsPageHandlers.client).GetFilmDataByUuid(ctx, req)
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

	req := session.AllFilmCommentsRequest{
		Uuid: uuid,
	}
	comments, err := (*filmsPageHandlers.client).GetAllFilmComments(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var commentsRegular []domain.Comment
	for _, comment := range comments.Comments {
		commentRegular := convertCommentToRegular(comment)
		escapeComment(&commentRegular)
		commentsRegular = append(commentsRegular, commentRegular)
	}

	response := filmCommentsResponse{
		Status:   http.StatusOK,
		Comments: commentsRegular,
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

func (filmsPageHandlers *FilmsPageHandlers) GetActorsByFilm(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	req := session.ActorsByFilmRequest{
		Uuid: uuid,
	}

	actors, err := (*filmsPageHandlers.client).GetActorsByFilm(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}

		return
	}

	var actorsRegular []domain.ActorPreview
	for _, actor := range actors.Actors {
		actorRegular := convertActorPreviewToRegular(actor)
		escapeActorPreview(&actorRegular)
		actorsRegular = append(actorsRegular, actorRegular)
	}

	response := filmActorsResponse{
		Status: http.StatusOK,
		Actors: actorsRegular,
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

type actorResponse struct {
	Status int              `json:"status"`
	Actor  domain.ActorData `json:"actor"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetActorByUuid(w http.ResponseWriter, r *http.Request) {
	actorUuid := mux.Vars(r)["uuid"]
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	req := session.ActorDataByUuidRequest{
		Uuid: actorUuid,
	}
	actor, err := (*filmsPageHandlers.client).GetActorDataByUuid(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] error at getting actor data: %v\n", requestID, err)
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	actorDataRegular := convertActorDataToRegular(actor.Actor)
	escapeActorData(&actorDataRegular)

	response := actorResponse{
		Status: http.StatusOK,
		Actor:  actorDataRegular,
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