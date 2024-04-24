package films

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/service"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type FilmsService interface {
	GetFilmDataByUuid(ctx context.Context, uuid string) (domain.FilmData, error)
	AddFilm(ctx context.Context, film domain.FilmDataToAdd) error
	RemoveFilm(ctx context.Context, uuid string) error
	GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error)
	GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error)
	GetAllFilmActors(ctx context.Context, uuid string) ([]domain.ActorPreview, error)
	GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(ctx context.Context, filmUuid string) ([]domain.ActorPreview, error)
}

type FilmsServer struct {
	filmsService *service.FilmsService
	logger       *zap.SugaredLogger
}

func InitFilmsServer(service *service.FilmsService, logger *zap.SugaredLogger) *FilmsServer {
	return &FilmsServer{
		filmsService: service,
		logger:       logger,
	}
}

func (server *FilmsServer) GetAllFilmsPreviews(ctx context.Context, req *session.AllFilmsPreviewsRequest) (res *session.AllFilmsPreviewsResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	films, err := server.filmsService.GetAllFilmsPreviews(ctx)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get all films previews: %v\n", requestId, err)
	}

	var filmsConverted []*session.FilmPreview
	for _, film := range films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToProto(&film))
	}

	return &session.AllFilmsPreviewsResponse{
		Films: filmsConverted,
	}, nil
}

func (server *FilmsServer) GetFilmDataByUuid(ctx context.Context, req *session.FilmDataByUuidRequest) (res *session.FilmDataByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	film, err := server.filmsService.GetFilmDataByUuid(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
	}
	filmConverted := convertFilmDataToProto(&film)

	return &session.FilmDataByUuidResponse{
		FilmData: filmConverted,
	}, nil
}

func (server *FilmsServer) GetFilmPreviewByUuid(ctx context.Context, req *session.FilmPreviewByUuidRequest) (res *session.FilmPreviewByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	film, err := server.filmsService.GetFilmPreview(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
	}
	filmConverted := convertFilmPreviewToProto(&film)

	return &session.FilmPreviewByUuidResponse{
		FilmPreview: filmConverted,
	}, nil
}

func (server *FilmsServer) GetAllFilmComments(ctx context.Context, req *session.AllFilmCommentsRequest) (res *session.AllFilmCommentsResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	comments, err := server.filmsService.GetAllFilmComments(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get all film comments: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get all film comments: %v\n", requestId, err)
	}
	var commentsConverted []*session.Comment
	for _, comment := range comments {
		commentsConverted = append(commentsConverted, convertCommentToProto(&comment))
	}

	return &session.AllFilmCommentsResponse{
		Comments: commentsConverted,
	}, nil
}

func (server *FilmsServer) GetAllFilmActors(ctx context.Context, req *session.AllFilmActorsRequest) (res *session.AllFilmActorsResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actors, err := server.filmsService.GetAllFilmActors(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get all film actors: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get all film actors: %v\n", requestId, err)
	}
	var actorsConverted []*session.ActorPreview
	for _, actor := range actors {
		actorsConverted = append(actorsConverted, convertActorPreviewToProto(actor))
	}

	return &session.AllFilmActorsResponse{
		ActorPreviews: actorsConverted,
	}, nil
}

func (server *FilmsServer) RemoveFilmByUuid(ctx context.Context, req *session.RemoveFilmByUuidRequest) (res *session.RemoveFilmByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.filmsService.RemoveFilm(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to remove film data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to remove film data: %v\n", requestId, err)
	}
	return &session.RemoveFilmByUuidResponse{}, nil
}

func (server *FilmsServer) GetActorDataByUuid(ctx context.Context, req *session.ActorDataByUuidRequest) (res *session.ActorDataByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actor, err := server.filmsService.GetActorByUuid(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get actor data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get actor data: %v\n", requestId, err)
	}

	actorConverted := convertActorDataToProto(actor)

	var filmsConverted []*session.FilmPreview
	for _, film := range actor.Films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToProto(&film))
	}
	actorConverted.FilmsPreviews = filmsConverted

	return &session.ActorDataByUuidResponse{
		Actor: actorConverted,
	}, nil
}

func (server *FilmsServer) GetActorsByFilm(ctx context.Context, req *session.ActorsByFilmRequest) (res *session.ActorsByFilmResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actors, err := server.filmsService.GetActorsByFilm(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get actor data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get actor data: %v\n", requestId, err)
	}

	var actorsConverted []*session.ActorPreview
	for _, actor := range actors {
		actorsConverted = append(actorsConverted, convertActorPreviewToProto(actor))
	}

	return &session.ActorsByFilmResponse{
		Actors: actorsConverted,
	}, nil
}

func convertFilmPreviewToProto(film *domain.FilmPreview) *session.FilmPreview {
	return &session.FilmPreview{
		Uuid:        film.Uuid,
		Preview:     film.Preview,
		Title:       film.Title,
		Director:    film.Director,
		AvgScore:    film.AverageScore,
		ScoresCount: film.ScoresCount,
		Duration:    film.Duration,
		AgeLimit:    film.AgeLimit,
	}
}

func convertFilmDataToProto(film *domain.FilmData) *session.FilmData {
	return &session.FilmData{
		Uuid:        film.Uuid,
		Preview:     film.Preview,
		Title:       film.Title,
		Link:        film.Link,
		Director:    film.Director,
		AvgScore:    film.AverageScore,
		ScoresCount: film.ScoresCount,
		Duration:    film.Duration,
		AgeLimit:    film.AgeLimit,
		Date:        convertTimeToProto(film.Date),
		Data:        film.Data,
	}
}

func convertCommentToProto(comment *domain.Comment) *session.Comment {
	return &session.Comment{
		Uuid:     comment.Uuid,
		Text:     comment.Text,
		FilmUuid: comment.FilmUuid,
		Author:   comment.Author,
		Score:    comment.Score,
		AddedAt:  convertTimeToProto(comment.AddedAt),
	}
}

func convertActorPreviewToProto(actor domain.ActorPreview) *session.ActorPreview {
	return &session.ActorPreview{
		Uuid:   actor.Uuid,
		Name:   actor.Name,
		Avatar: actor.Avatar,
	}
}

func convertActorDataToProto(actor domain.ActorData) *session.ActorData {
	return &session.ActorData{
		Uuid:     actor.Uuid,
		Name:     actor.Name,
		Avatar:   actor.Avatar,
		Birthday: convertTimeToProto(actor.Birthday),
		Career:   actor.Career,
		Spouse:   actor.Spouse,
		Genres:   actor.Genres,
	}
}

func convertTimeToProto(time time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: time.Unix(),
		Nanos:   int32(time.Nanosecond()),
	}
}
