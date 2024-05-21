package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type FilmsStorage interface {
	AddFilm(film domain.FilmToAdd) error
	GetFilmDataByUuid(uuid string) (domain.CommonFilmData, error)
	RemoveFilm(uuid string) error
	GetFilmPreview(uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews() ([]domain.FilmPreview, error)
	GetActorByUuid(actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error)
	PutFavoriteFilm(filmUuid string, userUuid string) error
	RemoveFavoriteFilm(filmUuid string, userUuid string) error
	GetAllFavoriteFilms(userUuid string) ([]domain.FilmPreview, error)
	GetAllFilmsByGenre(genreUuid string) ([]domain.FilmPreview, error)
	GetAllGenres() ([]domain.GenreFilms, error)
	FindFilmsShort(title string, page int) ([]domain.FilmPreview, error)
	FindFilmsLong(title string, page int) (domain.SearchFilms, error)
	FindSerialsShort(title string, page int) ([]domain.FilmPreview, error)
	FindSerialsLong(title string, page int) (domain.SearchFilms, error)
	FindActorsShort(name string, page int) ([]domain.ActorPreview, error)
	FindActorsLong(name string, page int) (domain.SearchActors, error)
	GetTopFilms() ([]domain.TopFilm, error)
	GetAllFilmComments(filmUuid string, userUuid string) ([]domain.Comment, error)
	AddComment(comment domain.CommentToAdd) error
	RemoveComment(comment domain.CommentToRemove) error
}

type FilmsService struct {
	storage          FilmsStorage
	metrics          *metrics.GrpcMetrics
	logger           *zap.SugaredLogger
	localStoragePath string
}

func NewFilmsService(storage FilmsStorage, metrics *metrics.GrpcMetrics, logger *zap.SugaredLogger,
	localStoragePath string) *FilmsService {
	return &FilmsService{
		storage:          storage,
		metrics:          metrics,
		logger:           logger,
		localStoragePath: localStoragePath,
	}
}

func (service *FilmsService) GetFilmDataByUuid(ctx context.Context, uuid string) (domain.CommonFilmData, error) {
	service.metrics.IncRequestsTotal("GetFilmDataByUuid")
	film, err := service.storage.GetFilmDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get film: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchFilm)
		return domain.CommonFilmData{}, err
	}

	return film, nil
}

func (service *FilmsService) AddFilm(ctx context.Context, film domain.FilmToAdd) error {
	service.metrics.IncRequestsTotal("AddFilm")
	err := service.storage.AddFilm(film)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to add film: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveFilm(ctx context.Context, uuid string) error {
	service.metrics.IncRequestsTotal("RemoveFilm")
	err := service.storage.RemoveFilm(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove film: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *FilmsService) GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("GetFilmPreview")
	filmPreview, err := service.storage.GetFilmPreview(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get film preview: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.FilmPreview{}, err
	}
	return filmPreview, nil
}

func (service *FilmsService) GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("GetAllFilmsPreviews")
	filmPreviews, err := service.storage.GetAllFilmsPreviews()
	if err != nil {
		service.logger.Errorf("[reqid=%v] failed to get all films previews: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return nil, err
	}
	return filmPreviews, nil
}

func (service *FilmsService) GetAllFilmComments(ctx context.Context, filmUuid string, userUuid string) ([]domain.Comment, error) {
	service.metrics.IncRequestsTotal("GetAllFilmComments")
	comments, err := service.storage.GetAllFilmComments(filmUuid, userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get all film comments: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return nil, err
	}
	return comments, nil
}

func (service *FilmsService) GetActorsByFilm(ctx context.Context, uuid string) ([]domain.ActorPreview, error) {
	service.metrics.IncRequestsTotal("GetActorsByFilm")
	actors, err := service.storage.GetActorsByFilm(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get all film actors: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}

func (service *FilmsService) GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error) {
	service.metrics.IncRequestsTotal("GetActorByUuid")
	actor, err := service.storage.GetActorByUuid(actorUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get actor: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchActor)
		return domain.ActorData{}, err
	}

	return actor, nil
}

func (service *FilmsService) PutFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error {
	service.metrics.IncRequestsTotal("PutFavoriteFilm")
	err := service.storage.PutFavoriteFilm(filmUuid, userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to put favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error {
	service.metrics.IncRequestsTotal("RemoveFavoriteFilm")
	err := service.storage.RemoveFavoriteFilm(filmUuid, userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *FilmsService) GetAllFavoriteFilms(ctx context.Context, userUuid string) ([]domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("GetAllFavoriteFilms")
	films, err := service.storage.GetAllFavoriteFilms(userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) GetAllFilmsByGenre(ctx context.Context, genreUuid string) ([]domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("GetAllFilmsByGenre")
	films, err := service.storage.GetAllFilmsByGenre(genreUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get genre films: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) GetAllGenres(ctx context.Context) ([]domain.GenreFilms, error) {
	service.metrics.IncRequestsTotal("GetAllGenres")
	genres, err := service.storage.GetAllGenres()
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get genres: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return genres, nil
}

func (service *FilmsService) FindFilmsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("FindFilmsShort")
	films, err := service.storage.FindFilmsShort(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find films short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) FindFilmsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error) {
	service.metrics.IncRequestsTotal("FindFilmsLong")
	films, err := service.storage.FindFilmsLong(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find films long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.SearchFilms{}, err
	}
	return films, nil
}

func (service *FilmsService) FindSerialsShort(ctx context.Context, title string,
	page int) ([]domain.FilmPreview, error) {
	service.metrics.IncRequestsTotal("FindSerialsShort")
	serials, err := service.storage.FindSerialsShort(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find serials short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return serials, nil
}

func (service *FilmsService) FindSerialsLong(ctx context.Context, title string, page int) (domain.SearchFilms, error) {
	service.metrics.IncRequestsTotal("FindSerialsLong")
	serials, err := service.storage.FindSerialsLong(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find serials long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.SearchFilms{}, err
	}
	return serials, nil
}

func (service *FilmsService) FindActorsShort(ctx context.Context, name string,
	page int) ([]domain.ActorPreview, error) {
	service.metrics.IncRequestsTotal("FindActorsShort")
	actors, err := service.storage.FindActorsShort(name, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find actors short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}

func (service *FilmsService) FindActorsLong(ctx context.Context, name string, page int) (domain.SearchActors, error) {
	service.metrics.IncRequestsTotal("FindActorsLong")
	actors, err := service.storage.FindActorsLong(name, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find actors long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.SearchActors{}, err
	}
	return actors, nil
}

func (service *FilmsService) GetTopFilms(ctx context.Context) ([]domain.TopFilm, error) {
	service.metrics.IncRequestsTotal("GetTopFilms")
	films, err := service.storage.GetTopFilms()
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get top films: %v", ctx.Value(requestId.ReqIDKey), err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) AddComment(ctx context.Context, comment domain.CommentToAdd) error {
	service.metrics.IncRequestsTotal("AddComment")
	err := service.storage.AddComment(comment)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to add comment: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveComment(ctx context.Context, comment domain.CommentToRemove) error {
	service.metrics.IncRequestsTotal("RemoveComment")
	err := service.storage.RemoveComment(comment)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove comment: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}
