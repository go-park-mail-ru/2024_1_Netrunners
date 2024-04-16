package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type ActorsStorage interface {
	GetActorByUuid(actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error)
}

type ActorsService struct {
	storage ActorsStorage
	logger  *zap.SugaredLogger
}

func NewActorsService(storage ActorsStorage, logger *zap.SugaredLogger) *ActorsService {
	return &ActorsService{
		storage: storage,
		logger:  logger,
	}
}

func (service *ActorsService) GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error) {
	actor, err := service.storage.GetActorByUuid(actorUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get actor: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoSuchActor)
		return domain.ActorData{}, err
	}

	return actor, nil
}

func (service *ActorsService) GetActorsByFilm(ctx context.Context, filmUuid string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetActorsByFilm(filmUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get actors by film: %v", ctx.Value(requestId.ReqIDKey),
			myerrors.ErrNoActorsForFilm)
		return nil, err
	}

	return actors, nil
}
