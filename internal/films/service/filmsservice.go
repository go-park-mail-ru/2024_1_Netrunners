package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type FilmsStorage interface {
	AddFilm(film domain.FilmToAdd) error
	GetFilmDataByUuid(uuid string) (domain.CommonFilmData, error)
	RemoveFilm(uuid string) error
	GetFilmPreview(uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews() ([]domain.FilmPreview, error)
	GetAllFilmComments(uuid string) ([]domain.Comment, error)
	GetActorByUuid(actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error)
	PutFavoriteFilm(filmUuid string, userUuid string) error
	RemoveFavoriteFilm(filmUuid string, userUuid string) error
	GetAllFavoriteFilms(userUuid string) ([]domain.FilmPreview, error)
	GetAllFilmsByGenre(genreUuid string) ([]domain.FilmPreview, error)
	GetAllGenres() ([]domain.GenreFilms, error)
	FindFilmsShort(title string, page int) ([]domain.FilmPreview, error)
	FindFilmsLong(title string, page int) ([]domain.FilmData, error)
	FindSerialsShort(title string, page int) ([]domain.FilmPreview, error)
	FindSerialsLong(title string, page int) ([]domain.FilmData, error)
	FindActorsShort(name string, page int) ([]domain.ActorPreview, error)
	FindActorsLong(name string, page int) ([]domain.ActorData, error)
}

type FilmsService struct {
	storage          FilmsStorage
	logger           *zap.SugaredLogger
	localStoragePath string
}

func NewFilmsService(storage FilmsStorage, logger *zap.SugaredLogger, localStoragePath string) *FilmsService {
	return &FilmsService{
		storage:          storage,
		logger:           logger,
		localStoragePath: localStoragePath,
	}
}

func (service *FilmsService) GetFilmDataByUuid(ctx context.Context, uuid string) (domain.CommonFilmData, error) {
	film, err := service.storage.GetFilmDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get film: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchFilm)
		return domain.CommonFilmData{}, err
	}

	return film, nil
}

func (service *FilmsService) AddFilm(ctx context.Context, film domain.FilmToAdd) error {
	err := service.storage.AddFilm(film)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to add film: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveFilm(ctx context.Context, uuid string) error {
	err := service.storage.RemoveFilm(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove film: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *FilmsService) GetFilmPreview(ctx context.Context, uuid string) (domain.FilmPreview, error) {
	filmPreview, err := service.storage.GetFilmPreview(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get film preview: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.FilmPreview{}, err
	}
	return filmPreview, nil
}

func (service *FilmsService) GetAllFilmsPreviews(ctx context.Context) ([]domain.FilmPreview, error) {
	filmPreviews, err := service.storage.GetAllFilmsPreviews()
	if err != nil {
		service.logger.Errorf("[reqid=%v] failed to get all films previews: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return nil, err
	}
	return filmPreviews, nil
}

func (service *FilmsService) GetAllFilmComments(ctx context.Context, uuid string) ([]domain.Comment, error) {
	comments, err := service.storage.GetAllFilmComments(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get all film comments: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return nil, err
	}
	return comments, nil
}

func (service *FilmsService) GetActorsByFilm(ctx context.Context, uuid string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetActorsByFilm(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get all film actors: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}

func (service *FilmsService) GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error) {
	actor, err := service.storage.GetActorByUuid(actorUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get actor: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchActor)
		return domain.ActorData{}, err
	}

	return actor, nil
}

func (service *FilmsService) PutFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error {
	err := service.storage.PutFavoriteFilm(filmUuid, userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to put favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveFavoriteFilm(ctx context.Context, filmUuid string, userUuid string) error {
	err := service.storage.RemoveFavoriteFilm(filmUuid, userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *FilmsService) GetAllFavoriteFilms(ctx context.Context, userUuid string) ([]domain.FilmPreview, error) {
	films, err := service.storage.GetAllFavoriteFilms(userUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove favorite film: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) GetAllFilmsByGenre(ctx context.Context, genreUuid string) ([]domain.FilmPreview, error) {
	films, err := service.storage.GetAllFilmsByGenre(genreUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get genre films: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) GetAllGenres(ctx context.Context) ([]domain.GenreFilms, error) {
	genres, err := service.storage.GetAllGenres()
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get genres: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return genres, nil
}

func (service *FilmsService) FindFilmsShort(ctx context.Context, title string, page int) ([]domain.FilmPreview, error) {
	films, err := service.storage.FindFilmsShort(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find films short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) FindFilmsLong(ctx context.Context, title string, page int) ([]domain.FilmData, error) {
	films, err := service.storage.FindFilmsLong(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find films long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return films, nil
}

func (service *FilmsService) FindSerialsShort(ctx context.Context, title string,
	page int) ([]domain.FilmPreview, error) {
	serials, err := service.storage.FindSerialsShort(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find serials short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return serials, nil
}

func (service *FilmsService) FindSerialsLong(ctx context.Context, title string, page int) ([]domain.FilmData, error) {
	serials, err := service.storage.FindSerialsLong(title, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find serials long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return serials, nil
}

func (service *FilmsService) FindActorsShort(ctx context.Context, name string,
	page int) ([]domain.ActorPreview, error) {
	actors, err := service.storage.FindActorsShort(name, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find actors short: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}

func (service *FilmsService) FindActorsLong(ctx context.Context, name string, page int) ([]domain.ActorData, error) {
	actors, err := service.storage.FindActorsLong(name, page)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to find actors long: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}
