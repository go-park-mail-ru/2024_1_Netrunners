package api

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type FilmsService interface {
	GetFilmDataByUuid(ctx context.Context, uuid string) (domain.CommonFilmData, error)
	AddFilm(ctx context.Context, film domain.FilmToAdd) error
	RemoveFilm(ctx context.Context, uuid string) error
	GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error)
	GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error)
	GetActorsByFilm(ctx context.Context, uuid string) ([]domain.ActorPreview, error)
	GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error)
	PutFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error
	RemoveFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error
	GetAllFavoriteFilms(ctx context.Context, userUuid string) ([]domain.FilmPreview, error)
	GetAllFilmsByGenre(ctx context.Context, genreUuid string) ([]domain.FilmPreview, error)
	GetAllGenres(ctx context.Context) ([]domain.GenreFilms, error)
	FindFilmsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error)
	FindFilmsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error)
	FindSerialsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error)
	FindSerialsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error)
	FindActorsShort(ctx context.Context, name string, page int) ([]domain.ActorPreview, error)
	FindActorsLong(ctx context.Context, name string, page int) (domain.SearchActors, error)
}

type FilmsServer struct {
	filmsService FilmsService
	logger       *zap.SugaredLogger
}

func NewFilmsServer(service FilmsService, logger *zap.SugaredLogger) *FilmsServer {
	return &FilmsServer{
		filmsService: service,
		logger:       logger,
	}
}

func (server *FilmsServer) GetAllFilmsPreviews(ctx context.Context,
	req *session.AllFilmsPreviewsRequest) (res *session.AllFilmsPreviewsResponse, err error) {
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

func (server *FilmsServer) GetFilmDataByUuid(ctx context.Context,
	req *session.FilmDataByUuidRequest) (res *session.FilmDataByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	film, err := server.filmsService.GetFilmDataByUuid(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get film data: %v\n", requestId, err)
	}

	filmConverted := convertCommonFilmDataToProto(&film)

	return &session.FilmDataByUuidResponse{
		FilmData: filmConverted,
	}, nil
}

func (server *FilmsServer) GetFilmPreviewByUuid(ctx context.Context,
	req *session.FilmPreviewByUuidRequest) (res *session.FilmPreviewByUuidResponse, err error) {
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

func (server *FilmsServer) GetAllFilmComments(ctx context.Context,
	req *session.AllFilmCommentsRequest) (res *session.AllFilmCommentsResponse, err error) {
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

func (server *FilmsServer) GetActorsByFilm(ctx context.Context,
	req *session.ActorsByFilmRequest) (res *session.ActorsByFilmResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actors, err := server.filmsService.GetActorsByFilm(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get all film actors: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get all film actors: %v\n", requestId, err)
	}
	var actorsConverted []*session.ActorPreview
	for _, actor := range actors {
		actorsConverted = append(actorsConverted, convertActorPreviewToProto(actor))
	}

	return &session.ActorsByFilmResponse{
		Actors: actorsConverted,
	}, nil
}

func (server *FilmsServer) RemoveFilmByUuid(ctx context.Context,
	req *session.RemoveFilmByUuidRequest) (res *session.RemoveFilmByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.filmsService.RemoveFilm(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to remove film data: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to remove film data: %v\n", requestId, err)
	}
	return &session.RemoveFilmByUuidResponse{}, nil
}

func (server *FilmsServer) GetActorDataByUuid(ctx context.Context,
	req *session.ActorDataByUuidRequest) (res *session.ActorDataByUuidResponse, err error) {
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

func (server *FilmsServer) PutFavorite(ctx context.Context,
	req *session.PutFavoriteRequest) (res *session.PutFavoriteResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.filmsService.PutFavoriteFilm(ctx, req.FilmUuid, req.UserUuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to put favorite: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to put favorite: %v\n", requestId, err)
	}

	return &session.PutFavoriteResponse{}, nil
}

func (server *FilmsServer) DeleteFavorite(ctx context.Context,
	req *session.DeleteFavoriteRequest) (res *session.DeleteFavoriteResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.filmsService.RemoveFavoriteFilm(ctx, req.FilmUuid, req.UserUuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to remove favorite: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to remove favorite: %v\n", requestId, err)
	}

	return &session.DeleteFavoriteResponse{}, nil
}

func (server *FilmsServer) GetAllFavoriteFilms(ctx context.Context,
	req *session.GetAllFavoriteFilmsRequest) (res *session.GetAllFavoriteFilmsResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	films, err := server.filmsService.GetAllFavoriteFilms(ctx, req.UserUuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get favorite: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get favorite: %v\n", requestId, err)
	}

	var filmsConverted []*session.FilmPreview
	for _, film := range films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToProto(&film))
	}

	return &session.GetAllFavoriteFilmsResponse{
		Films: filmsConverted,
	}, nil
}

func (server *FilmsServer) GetAllFilmsByGenre(ctx context.Context,
	req *session.GetAllFilmsByGenreRequest) (res *session.GetAllFilmsByGenreResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	films, err := server.filmsService.GetAllFilmsByGenre(ctx, req.GenreUuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get genre films: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get genre films: %v\n", requestId, err)
	}

	var filmsConverted []*session.FilmPreview
	for _, film := range films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToProto(&film))
	}

	return &session.GetAllFilmsByGenreResponse{
		Films: filmsConverted,
	}, nil
}

func (server *FilmsServer) FindFilmsShort(ctx context.Context,
	request *session.FindFilmsShortRequest) (*session.FindFilmsShortResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	films, err := server.filmsService.FindFilmsShort(ctx, request.Key, int(request.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get favorite: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get favorite: %v\n", requestId, err)
	}

	var filmsConverted []*session.FilmPreview
	for _, film := range films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToProto(&film))
	}

	return &session.FindFilmsShortResponse{
		Films: filmsConverted,
	}, nil
}

func (server *FilmsServer) GetAllGenres(ctx context.Context,
	req *session.GetAllGenresRequest) (res *session.GetAllGenresResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	genres, err := server.filmsService.GetAllGenres(ctx)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get genres: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get genres: %v\n", requestId, err)
	}

	var genresConverted []*session.GenreFilms
	for _, genre := range genres {
		genreConverted := convertGenreFilmsToProto(genre)
		genresConverted = append(genresConverted, genreConverted)
	}
	return &session.GetAllGenresResponse{
		Genres: genresConverted,
	}, nil
}

func (server *FilmsServer) AddFilm(ctx context.Context,
	req *session.AddFilmRequest) (res *session.AddFilmResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.filmsService.AddFilm(ctx, convertFilmToAdd(req.FilmData))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to add favorite: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to add favorite: %v\n", requestId, err)
	}
	return &session.AddFilmResponse{}, nil
}

func (server *FilmsServer) FindFilmsLong(ctx context.Context,
	req *session.FindFilmsShortRequest) (*session.FindFilmsLongResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	films, err := server.filmsService.FindFilmsLong(ctx, req.Key, int(req.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get films: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get films: %v\n", requestId, err)
	}

	var filmsConverted []*session.FindFilmLong
	for _, film := range films.Films {
		filmsConverted = append(filmsConverted, convertFindFilmLongToProto(&film))
	}

	return &session.FindFilmsLongResponse{
		Films: filmsConverted,
		Count: films.Count,
	}, nil
}

func (server *FilmsServer) FindSerialsShort(ctx context.Context,
	request *session.FindFilmsShortRequest) (*session.FindFilmsShortResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	serials, err := server.filmsService.FindSerialsShort(ctx, request.Key, int(request.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get serials: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get serials: %v\n", requestId, err)
	}

	var serialsConverted []*session.FilmPreview
	for _, serial := range serials {
		serialsConverted = append(serialsConverted, convertFilmPreviewToProto(&serial))
	}

	return &session.FindFilmsShortResponse{
		Films: serialsConverted,
	}, nil
}

func (server *FilmsServer) FindSerialsLong(ctx context.Context,
	request *session.FindFilmsShortRequest) (*session.FindFilmsLongResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	serials, err := server.filmsService.FindSerialsLong(ctx, request.Key, int(request.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get serials: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get serials: %v\n", requestId, err)
	}

	var serialsConverted []*session.FindFilmLong
	for _, serial := range serials.Films {
		serialsConverted = append(serialsConverted, convertFindFilmLongToProto(&serial))
	}

	return &session.FindFilmsLongResponse{
		Films: serialsConverted,
		Count: serials.Count,
	}, nil
}

func (server *FilmsServer) FindActorsShort(ctx context.Context,
	request *session.FindActorsShortRequest) (*session.FindActorsShortResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actors, err := server.filmsService.FindActorsShort(ctx, request.Key, int(request.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get actors: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get actors: %v\n", requestId, err)
	}

	var actorsConverted []*session.ActorPreview
	for _, actor := range actors {
		actorsConverted = append(actorsConverted, convertActorPreviewToProto(actor))
	}

	return &session.FindActorsShortResponse{
		Actors: actorsConverted,
	}, nil
}

func (server *FilmsServer) FindActorsLong(ctx context.Context,
	request *session.FindActorsShortRequest) (*session.FindActorsLongResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	actors, err := server.filmsService.FindActorsLong(ctx, request.Key, int(request.Page))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get actors: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get actors: %v\n", requestId, err)
	}

	var actorsConverted []*session.ActorPreviewLong
	for _, actor := range actors.Actors {
		actorsConverted = append(actorsConverted, convertActorPreviewLongToProto(actor))
	}

	return &session.FindActorsLongResponse{
		Actors: actorsConverted,
		Count:  actors.Count,
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
		IsSerial:    film.IsSerial,
	}
}

func convertCommonFilmDataToProto(film *domain.CommonFilmData) *session.FilmData {
	var genres []*session.Genre
	for _, genre := range film.Genres {
		genres = append(genres, &session.Genre{Name: genre.Name, Uuid: genre.Uuid})
	}

	seasons := make([]*session.Season, 0, len(film.Seasons))
	for _, season := range film.Seasons {
		episodes := make([]*session.Episode, 0, len(season.Series))
		for _, episode := range season.Series {
			episodes = append(episodes, &session.Episode{
				Link: episode.Link,
			})
		}
		seasons = append(seasons, &session.Season{
			Episodes: episodes,
		})
	}

	return &session.FilmData{
		Uuid:        film.Uuid,
		IsSerial:    film.IsSerial,
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
		Genres:      genres,
		Seasons:     seasons,
	}
}

func convertFindFilmLongToProto(film *domain.FilmData) *session.FindFilmLong {
	return &session.FindFilmLong{
		Uuid:        film.Uuid,
		Preview:     film.Preview,
		Title:       film.Title,
		Director:    film.Director,
		AvgScore:    film.AverageScore,
		ScoresCount: film.ScoresCount,
		Duration:    film.Duration,
		AgeLimit:    film.AgeLimit,
		Date:        convertTimeToProto(film.Date),
		IsSerial:    film.IsSerial,
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

func convertActorPreviewLongToProto(actor domain.ActorData) *session.ActorPreviewLong {
	return &session.ActorPreviewLong{
		Uuid:     actor.Uuid,
		Name:     actor.Name,
		Avatar:   actor.Avatar,
		Birthday: convertTimeToProto(actor.Birthday),
		Career:   actor.Career,
	}
}

func convertActorDataToProto(actor domain.ActorData) *session.ActorData {
	return &session.ActorData{
		Uuid:       actor.Uuid,
		Name:       actor.Name,
		Avatar:     actor.Avatar,
		Birthday:   convertTimeToProto(actor.Birthday),
		Career:     actor.Career,
		Spouse:     actor.Spouse,
		Birthplace: actor.BirthPlace,
		Height:     actor.Height,
	}
}

func convertTimeToProto(time time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: time.Unix(),
		Nanos:   int32(time.Nanosecond()),
	}
}

func convertGenreFilmsToProto(genreFilms domain.GenreFilms) *session.GenreFilms {
	var filmsConverted []*session.FilmPreview
	for _, film := range genreFilms.Films {
		filmConverted := convertFilmPreviewToProto(&film)
		filmsConverted = append(filmsConverted, filmConverted)
	}
	return &session.GenreFilms{
		Genre:     genreFilms.Name,
		GenreUuid: genreFilms.Uuid,
		Films:     filmsConverted,
	}
}

func convertProtoToTime(protoTime *timestamppb.Timestamp) time.Time {
	return protoTime.AsTime()
}

func convertActorToAddToRegular(actor *session.ActorDataToAdd) domain.ActorToAdd {
	return domain.ActorToAdd{
		Name:       actor.Name,
		Avatar:     actor.Avatar,
		Birthday:   convertProtoToTime(actor.BirthdayAt),
		Career:     actor.Career,
		Height:     actor.Height,
		Spouse:     actor.Spouse,
		BirthPlace: actor.BirthPlace,
	}
}

func convertFilmToAdd(filmToAdd *session.FilmToAdd) domain.FilmToAdd {
	filmData := domain.FilmDataToAdd{
		Title:       filmToAdd.FilmData.Title,
		Preview:     filmToAdd.FilmData.Preview,
		Director:    filmToAdd.FilmData.Director,
		Data:        filmToAdd.FilmData.Data,
		AgeLimit:    filmToAdd.FilmData.AgeLimit,
		PublishedAt: convertProtoToTime(filmToAdd.FilmData.PublishedAt),
		Genres:      filmToAdd.FilmData.Genres,
		Duration:    filmToAdd.FilmData.Duration,
		Link:        filmToAdd.FilmData.Link,
	}

	var actors []domain.ActorToAdd
	for _, act := range filmToAdd.Actors {
		actors = append(actors, convertActorToAddToRegular(act))
	}

	directorData := domain.DirectorToAdd{
		Name:     filmToAdd.Director.Name,
		Birthday: convertProtoToTime(filmToAdd.Director.Birthday),
		Avatar:   filmToAdd.Director.Avatar,
	}

	return domain.FilmToAdd{
		FilmData:      filmData,
		Actors:        actors,
		DirectorToAdd: directorData,
	}
}
