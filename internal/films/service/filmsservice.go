package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type FilmsStorage interface {
	GetFilmDataByUuid(uuid string) (domain.FilmData, error)
	AddFilm(film domain.FilmDataToAdd) error
	RemoveFilm(uuid string) error
	GetFilmPreview(uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews() ([]domain.FilmPreview, error)
	GetAllFilmComments(uuid string) ([]domain.Comment, error)
	GetActorByUuid(actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error)
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

func (service *FilmsService) GetFilmDataByUuid(ctx context.Context, uuid string) (domain.FilmData, error) {
	film, err := service.storage.GetFilmDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get film: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchFilm)
		return domain.FilmData{}, err
	}
	return film, nil
}

func (service *FilmsService) AddFilm(ctx context.Context, film domain.FilmDataToAdd) error {
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