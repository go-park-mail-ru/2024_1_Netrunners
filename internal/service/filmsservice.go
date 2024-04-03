package service

import (
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
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

func (service *FilmsService) GetFilmDataByUuid(uuid string) (domain.FilmData, error) {
	film, err := service.storage.GetFilmDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetFilmByUuid: %v", myerrors.ErrInternalServerError)
		return domain.FilmData{}, err
	}
	return film, nil
}

func (service *FilmsService) AddFilm(film domain.FilmDataToAdd) error {
	err := service.storage.AddFilm(film)
	if err != nil {
		service.logger.Errorf("service error at AddFilm: %v", myerrors.ErrInternalServerError)
		return err
	}
	return nil
}

func (service *FilmsService) RemoveFilm(uuid string) error {
	err := service.storage.RemoveFilm(uuid)
	if err != nil {
		service.logger.Errorf("service error at RemoveFilm: %v", myerrors.ErrInternalServerError)
		return err
	}
	return nil
}

func (service *FilmsService) GetFilmPreview(uuid string) (domain.FilmPreview, error) {
	filmPreview, err := service.storage.GetFilmPreview(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetFilmPreview: %v", myerrors.ErrInternalServerError)
		return domain.FilmPreview{}, err
	}
	return filmPreview, nil
}

func (service *FilmsService) GetAllFilmsPreviews() ([]domain.FilmPreview, error) {
	filmPreviews, err := service.storage.GetAllFilmsPreviews()
	if err != nil {
		service.logger.Errorf("service error at GetAllFilmsPreviews: %v", myerrors.ErrInternalServerError)
		return nil, err
	}
	return filmPreviews, nil
}

func (service *FilmsService) GetAllFilmComments(uuid string) ([]domain.Comment, error) {
	comments, err := service.storage.GetAllFilmComments(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetAllFilmComments: %v", myerrors.ErrInternalServerError)
		return nil, err
	}
	return comments, nil
}

func (service *FilmsService) GetAllFilmActors(uuid string) ([]domain.ActorPreview, error) {
	actors, err := service.storage.GetAllFilmActors(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetAllFilmActors: %v", myerrors.ErrInternalServerError)
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
					Name: "Стас Ярушин",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "универский типос", "154", "01.01.2001", "Ангарск", "Хип-Хоп",
					//	"Светлана Ходченкова <3"),
				},
				{
					Name: "Дмитрий Нагиев",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "физрукский типос", "215", "16.10.1936", "Шахты", "RnB",
					//	"ТОЖЕ НЕ Светлана Ходченкова <3"),
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
					Name: "Костя Воронин",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "Костик", "181", "19.01.2012", "Россия",
					//	"Riddim", "Taylor Swift"),
				},
				{
					Name: "Tom Hanks",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "пиццерийных дел мастер", "178", "17.08.1965", "Америка",
					//	"Riddim", "Вроде Ваенга хз"),
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
					Name: "Дональд Дак",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "Бизнесмен, экс-президент США (я это не он хз не шарю)",
					//	"87", "01.02.2021", "Саратов", "Riddim", "Серый"),
				},
				{
					Name: "Дмитрий Нагиев",
					//Data: fmt.Sprintf("Карьера: %s \nРост: %s \nДата рождения: %s \nМесто рождения: %s \n"+
					//	"Жанры: %s \nСупруга: %s \n", "физрукский типос", "215", "16.10.1936", "Шахты", "RnB",
					//	"ТОЖЕ НЕ Светлана Ходченкова <3"),
				},
			},
		},
	}

	for _, film := range data {
		err := service.storage.AddFilm(film)
		if err != nil {
			service.logger.Errorf("service error at AddFilm: %v", myerrors.ErrInternalServerError)
			return err
		}
	}

	return nil
}
