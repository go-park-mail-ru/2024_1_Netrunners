package service

import (
	"context"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type FilmsStorage interface {
	GetFilmDataByUuid(uuid string) (domain.FilmData, error)
	AddFilm(film domain.FilmDataToAdd) error
	RemoveFilm(uuid string) error
	GetFilmPreview(uuid string) (domain.FilmPreview, error)
	GetAllFilmsPreviews() ([]domain.FilmPreview, error)
	GetAllFilmComments(uuid string) ([]domain.Comment, error)
	GetAllFilmActors(uuid string) ([]domain.ActorPreview, error)
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
			err)
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

func (service *FilmsService) GetAllFilmActors(ctx context.Context, uuid string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetAllFilmActors(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get all film actors: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return nil, err
	}
	return actors, nil
}

func (service *FilmsService) AddSomeData() error {
	data := []domain.FilmDataToAdd{
		{
			Preview: "https://m.media-amazon.com/images/M/MV5BNzlkNzVjMDMtOTdhZC00MGE1LTkxODctMzFmMjkwZm" +
				"MxZjFhXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_.jpg",
			Title:    "Fast and Furious 1",
			Duration: 3600,
			Director: "Dozer",
			Data:     "Dozer Dozer Dozer Dozer",
			Actors: []domain.ActorData{
				{
					Name:       "Стас Ярушин",
					Career:     "универский типос",
					Height:     154,
					BirthPlace: "Ангарск",
					Genres:     "Хип-Хоп",
					Spouse:     "Светлана Ходченкова <3",
				},
				{
					Name:       "Дмитрий Нагиев",
					Career:     "физрукский типос",
					Height:     215,
					BirthPlace: "Шахты",
					Genres:     "RnB",
					Spouse:     "ТОЖЕ НЕ Светлана Ходченкова <3",
				},
			},
		},
		{
			Preview:  "https://m.media-amazon.com/images/I/71Wo+cFznbL.jpg",
			Title:    "Fast and Furious 2",
			Duration: 7200,
			Director: "Dima",
			Data:     "Dima Dima Dima Dima",
			Actors: []domain.ActorData{
				{
					Name:       "Костя Воронин",
					Career:     "Костик",
					Height:     181,
					BirthPlace: "Россия",
					Genres:     "Riddim",
					Spouse:     "Taylor Swift",
				},
				{
					Name:       "Tom Hanks",
					Career:     "пиццерийных дел мастер",
					Height:     178,
					BirthPlace: "Омерика",
					Genres:     "Riddim",
					Spouse:     "Вроде Ваенга хз",
				},
			},
		},
		{
			Preview:  "https://m.media-amazon.com/images/I/71ql8kIrPKL.jpg",
			Title:    "Fast and Furious 3",
			Duration: 4800,
			Director: "Dima",
			Data:     "Dima Dima Dima Dima",
			Actors: []domain.ActorData{
				{
					Name:       "Дональд Дак",
					Career:     "пиццерийных дел мастер",
					Height:     178,
					BirthPlace: "Омерика",
					Genres:     "Riddim",
					Spouse:     "Вроде Ваенга хз",
				},
				{
					Name:       "Дмитрий Нагиев",
					Career:     "физрукский типос",
					Height:     215,
					BirthPlace: "Шахты",
					Genres:     "RnB",
					Spouse:     "ТОЖЕ НЕ Светлана Ходченкова <3",
				},
			},
		},
	}

	for _, film := range data {
		err := service.storage.AddFilm(film)
		if err != nil {
			service.logger.Errorf("[reqid=%s] failed to add film: %v", err)
			return err
		}
	}

	return nil
}
