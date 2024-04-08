package service

import (
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
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

func (service *ActorsService) GetActorByUuid(actorUuid string, requestID string) (domain.ActorData, error) {
	actor, err := service.storage.GetActorByUuid(actorUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at GetActorByUuid: %v", requestID, err)
		return domain.ActorData{}, err
	}

	return actor, nil
}

func (service *ActorsService) GetActorsByFilm(filmUuid string, requestID string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetActorsByFilm(filmUuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at GetActorsByFilm: %v", requestID, err)
		return nil, err
	}

	return actors, nil
}
