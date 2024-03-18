package service

import (
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"go.uber.org/zap"
)

type ActorsStorage interface {
	GetAllActorsPreviews() ([]domain.ActorPreview, error)
	GetActorByUuid(actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error)
}

type ActorsService struct {
	storage ActorsStorage
	logger  *zap.SugaredLogger
}

func NewActorsStorage(storage ActorsStorage, logger *zap.SugaredLogger) *ActorsService {
	return &ActorsService{
		storage: storage,
		logger:  logger,
	}
}

func (service *ActorsService) GetAllActorsPreviews() ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetAllActorsPreviews()
	if err != nil {
		service.logger.Errorf("service error at GetAllActorsPreviews: %v", err)
		return nil, err
	}

	return actors, nil
}

func (service *ActorsService) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	actor, err := service.storage.GetActorByUuid(actorUuid)
	if err != nil {
		service.logger.Errorf("service error at GetActorByUuid: %v", err)
		return domain.ActorData{}, err
	}

	return actor, nil
}

func (service *ActorsService) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetActorsByFilm(filmUuid)
	if err != nil {
		service.logger.Errorf("service error at GetActorsByFilm: %v", err)
		return nil, err
	}

	return actors, nil
}
