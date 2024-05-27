package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type FilmsPageHandlers struct {
	client  *session.FilmsClient
	metrics *metrics.HttpMetrics
	logger  *zap.SugaredLogger
}

func NewFilmsPageHandlers(client *session.FilmsClient, metrics *metrics.HttpMetrics,
	logger *zap.SugaredLogger) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		client:  client,
		metrics: metrics,
		logger:  logger,
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmsPreviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var req *session.AllFilmsPreviewsRequest
	res, err := (*filmsPageHandlers.client).GetAllFilmsPreviews(ctx, req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
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

	response := domain.FilmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  filmsRegular,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}
	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetFilmsPreviewsWithSub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var req *session.AllFilmsPreviewsRequest
	res, err := (*filmsPageHandlers.client).GetFilmsPreviewsWithSub(ctx, req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
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

	err = WriteResponse(w, r, filmsPageHandlers.metrics, response, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

type filmDataResponse struct {
	Status   int         `json:"status"`
	FilmData interface{} `json:"film"`
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
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var response domain.FilmDataResponse
	if !filmData.FilmData.IsSerial {
		filmDataRegular := convertFilmDataToRegular(filmData.FilmData)
		escapeFilmData(&filmDataRegular)
		response = domain.FilmDataResponse{
			Status:   http.StatusOK,
			FilmData: filmDataRegular,
		}
	} else {
		serialDataRegular := convertSerialDataToRegular(filmData.FilmData)
		escapeSerialData(&serialDataRegular)
		response = domain.FilmDataResponse{
			Status:   http.StatusOK,
			FilmData: serialDataRegular,
		}
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	uuid := mux.Vars(r)["uuid"]
	req := session.AllFilmCommentsRequest{
		FilmUuid: uuid,
	}
	comments, err := (*filmsPageHandlers.client).GetAllFilmComments(ctx, &req)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
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

	response := domain.FilmCommentsResponse{
		Status:   http.StatusOK,
		Comments: commentsRegular,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
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
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
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

	response := domain.FilmActorsResponse{
		Status: http.StatusOK,
		Actors: actorsRegular,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
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
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	actorDataRegular := convertActorDataToRegular(actor.Actor)
	escapeActorData(&actorDataRegular)

	response := domain.ActorResponse{
		Status: http.StatusOK,
		Actor:  actorDataRegular,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) PutFavoriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	var data domain.DataToFavorite
	err := easyjson.UnmarshalFromReader(r.Body, &data)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	req := session.PutFavoriteRequest{FilmUuid: data.FilmUuid, UserUuid: data.UserUuid}
	_, err = (*filmsPageHandlers.client).PutFavorite(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to put favorite film: %v\n", requestId, err)
		err := WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	err = WriteSuccess(w, r, filmsPageHandlers.metrics)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		return
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film added successfully", requestId))
}

func (filmsPageHandlers *FilmsPageHandlers) RemoveFavoriteFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	var data domain.DataToFavorite
	err := easyjson.UnmarshalFromReader(r.Body, &data)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	req := session.DeleteFavoriteRequest{FilmUuid: data.FilmUuid, UserUuid: data.UserUuid}
	_, err = (*filmsPageHandlers.client).DeleteFavorite(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to remove favorite film: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
	}

	err = WriteSuccess(w, r, filmsPageHandlers.metrics)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film removed successfully", requestId))
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFavoriteFilms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)
	uuid := mux.Vars(r)["uuid"]

	req := session.GetAllFavoriteFilmsRequest{UserUuid: uuid}
	films, err := (*filmsPageHandlers.client).GetAllFavoriteFilms(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all favorite film: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var filmsConverted []domain.FilmPreview
	for _, film := range films.Films {
		filmConverted := convertFilmPreviewToRegular(film)
		escapeFilmPreview(&filmConverted)
		filmsConverted = append(filmsConverted, filmConverted)
	}

	response := domain.FilmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  filmsConverted,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmsByGenre(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)
	uuid := mux.Vars(r)["uuid"]

	req := session.GetAllFilmsByGenreRequest{GenreUuid: uuid}
	films, err := (*filmsPageHandlers.client).GetAllFilmsByGenre(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get all genre films: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var filmsConverted []domain.FilmPreview
	for _, film := range films.Films {
		filmConverted := convertFilmPreviewToRegular(film)
		escapeFilmPreview(&filmConverted)
		filmsConverted = append(filmsConverted, filmConverted)
	}
	response := domain.FilmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  filmsConverted,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) ShortSearch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	params := r.URL.Query()
	var (
		page   int
		search string
	)
	if s, ok := params["s"]; ok {
		search = s[0]
	}
	if p, ok := params["p"]; ok {
		page, _ = strconv.Atoi(p[0])
	} else {
		page = 1
	}

	filmsReq := session.FindFilmsShortRequest{
		Key:  search,
		Page: uint32(page),
	}
	films, err := (*filmsPageHandlers.client).FindFilmsShort(ctx, &filmsReq)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find films: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var filmsConverted []domain.FilmPreview
	for _, film := range films.Films {
		filmConverted := convertFilmPreviewToRegular(film)
		escapeFilmPreview(&filmConverted)
		filmsConverted = append(filmsConverted, filmConverted)
	}

	serialsReq := session.FindFilmsShortRequest{
		Key:  search,
		Page: uint32(page),
	}
	serials, err := (*filmsPageHandlers.client).FindSerialsShort(ctx, &serialsReq)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find serials: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	for _, serial := range serials.Films {
		serialConverted := convertFilmPreviewToRegular(serial)
		escapeFilmPreview(&serialConverted)
		filmsConverted = append(filmsConverted, serialConverted)
	}

	actorsReq := session.FindActorsShortRequest{
		Key:  search,
		Page: uint32(page),
	}
	actors, err := (*filmsPageHandlers.client).FindActorsShort(ctx, &actorsReq)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find actors: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var actorssConverted []domain.ActorPreview
	for _, actor := range actors.Actors {
		actorConverted := convertActorPreviewToRegular(actor)
		escapeActorPreview(&actorConverted)
		actorssConverted = append(actorssConverted, actorConverted)
	}

	response := domain.ShortSearchResponse{
		Status: http.StatusOK,
		Films:  filmsConverted,
		Actors: actorssConverted,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) LongSearch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	params := r.URL.Query()
	var (
		page   int
		search string
		findBy string
	)
	if s, ok := params["s"]; ok {
		search = s[0]
	}
	if p, ok := params["p"]; ok {
		page, _ = strconv.Atoi(p[0])
	} else {
		page = 1
	}
	if fb, ok := params["fb"]; !ok {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get fb param: %v\n", requestID,
			myerrors.ErrIncorrectSearchParams)
		err := WriteError(w, r, filmsPageHandlers.metrics, myerrors.ErrIncorrectSearchParams)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	} else {
		findBy = fb[0]
	}

	switch findBy {
	case "films":
		filmsReq := session.FindFilmsShortRequest{
			Key:  search,
			Page: uint32(page),
		}
		films, err := (*filmsPageHandlers.client).FindFilmsLong(ctx, &filmsReq)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find films: %v\n", requestID, err)
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		var filmsConverted []domain.FilmData
		for _, film := range films.Films {
			filmConverted := convertLongFilmPreviewToRegular(film)
			escapeFilmData(&filmConverted)
			filmsConverted = append(filmsConverted, filmConverted)
		}

		response := domain.LongSearchResponse{
			Status: http.StatusOK,
			Films:  filmsConverted,
			Count:  int(films.Count),
		}

		err = WriteResponse(w, r, filmsPageHandlers.metrics, response, requestID)
		if err != nil {
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}
	case "serials":
		filmsReq := session.FindFilmsShortRequest{
			Key:  search,
			Page: uint32(page),
		}
		serials, err := (*filmsPageHandlers.client).FindSerialsLong(ctx, &filmsReq)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find films: %v\n", requestID, err)
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		var serialsConverted []domain.FilmData
		for _, serial := range serials.Films {
			serialConverted := convertLongFilmPreviewToRegular(serial)
			escapeFilmData(&serialConverted)
			serialsConverted = append(serialsConverted, serialConverted)
		}

		response := domain.LongSearchResponse{
			Status: http.StatusOK,
			Films:  serialsConverted,
			Count:  int(serials.Count),
		}

		jsonResponse, err := easyjson.Marshal(response)
		if err != nil {
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
			}
			return
		}

		err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
		if err != nil {
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}
	case "actors":
		actorsReq := session.FindActorsShortRequest{
			Key:  search,
			Page: uint32(page),
		}
		actors, err := (*filmsPageHandlers.client).FindActorsLong(ctx, &actorsReq)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to find actors: %v\n", requestID, err)
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		var actorssConverted []domain.ActorData
		for _, actor := range actors.Actors {
			actorConverted := convertActorPreviewLongToRegular(actor)
			escapeActorData(&actorConverted)
			actorssConverted = append(actorssConverted, actorConverted)
		}

		response := domain.LongSearchResponse{
			Status: http.StatusOK,
			Actors: actorssConverted,
			Count:  int(actors.Count),
		}

		jsonResponse, err := easyjson.Marshal(response)
		if err != nil {
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
			}
			return
		}

		err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
		if err != nil {
			err = WriteError(w, r, filmsPageHandlers.metrics, err)
			if err != nil {
				filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}
	default:
		err := WriteError(w, r, filmsPageHandlers.metrics, myerrors.ErrIncorrectSearchParams)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllGenres(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var req session.GetAllGenresRequest
	genresFilms, err := (*filmsPageHandlers.client).GetAllGenres(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get genres: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
	var genresConverted []domain.GenreFilms
	for _, genre := range genresFilms.Genres {
		genreConverted := convertGenreFilmsToRegular(genre)
		escapeGenreFilms(&genreConverted)
		genresConverted = append(genresConverted, genreConverted)
	}

	response := domain.GenresResponse{
		Status:      http.StatusOK,
		GenresFilms: genresConverted,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) AddFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)
	var filmAddData domain.FilmToAdd
	err := easyjson.UnmarshalFromReader(r.Body, &filmAddData)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode film data to add: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	req := session.AddFilmRequest{FilmData: convertFilmToAdd(filmAddData)}
	_, err = (*filmsPageHandlers.client).AddFilm(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to add film: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	err = WriteSuccess(w, r, filmsPageHandlers.metrics)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetTopFilms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	var req session.GetTopFilmsRequest
	films, err := (*filmsPageHandlers.client).GetTopFilms(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to get top films: %v\n", requestID, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var filmsConverted []domain.TopFilm
	for _, film := range films.Films {
		filmConverted := convertTopFilmToRegular(film)
		escapeTopFilm(&filmConverted)
		filmsConverted = append(filmsConverted, filmConverted)
	}

	response := domain.TopFilmsResponse{
		Status: http.StatusOK,
		Films:  filmsConverted,
	}

	jsonResponse, err := easyjson.Marshal(response)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to marshal response: %v\n", requestID, err)
		}
		return
	}

	err = WriteResponse(w, r, filmsPageHandlers.metrics, jsonResponse, requestID)
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, myerrors.ErrNoActiveSession)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	_, err = IsTokenValid(userToken, os.Getenv("SECRETKEY"))
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	var commentAddData domain.CommentToAdd
	err = easyjson.UnmarshalFromReader(r.Body, &commentAddData)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	req := session.AddCommentRequest{
		Comment: convertCommentToAddToProto(commentAddData),
	}
	_, err = (*filmsPageHandlers.client).AddComment(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to add comment: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	filmsPageHandlers.metrics.IncRequestsTotal(r.URL.Path, r.Method, 200)

	err = WriteSuccess(w, r, filmsPageHandlers.metrics)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film removed successfully", requestId))

}

func (filmsPageHandlers *FilmsPageHandlers) RemoveComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, myerrors.ErrNoActiveSession)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	_, err = IsTokenValid(userToken, os.Getenv("SECRETKEY"))
	if err != nil {
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	var commentRemoveData domain.CommentToRemove
	err = easyjson.UnmarshalFromReader(r.Body, &commentRemoveData)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to decode request data: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	req := session.RemoveCommentRequest{
		Comment: convertCommentToRemoveToProto(commentRemoveData),
	}
	_, err = (*filmsPageHandlers.client).RemoveComment(ctx, &req)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to remove comment: %v\n", requestId, err)
		err = WriteError(w, r, filmsPageHandlers.metrics, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	filmsPageHandlers.metrics.IncRequestsTotal(r.URL.Path, r.Method, 200)

	err = WriteSuccess(w, r, filmsPageHandlers.metrics)
	if err != nil {
		filmsPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}

	filmsPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] favorite film removed successfully", requestId))
}
