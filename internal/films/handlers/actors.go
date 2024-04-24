package films_handlers

import (
	"context"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

type ActorsServer interface {
	GetActorByUuid(ctx context.Context, actorUuid string) (domain.ActorData, error)
	GetActorsByFilm(ctx context.Context, filmUuid string) ([]domain.ActorPreview, error)
}

type ActorsHandlers struct {
	actorsService ActorsService
	logger        *zap.SugaredLogger
}

func NewActorsHandlers(actorsService ActorsService, logger *zap.SugaredLogger) *ActorsHandlers {
	return &ActorsHandlers{
		actorsService: actorsService,
		logger:        logger,
	}
}
