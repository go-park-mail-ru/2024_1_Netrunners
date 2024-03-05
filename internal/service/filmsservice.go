package service

import (
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

type filmsStorage interface {
	AddFilm(film domain.FilmPreview) error
	RemoveFilm(id string) error
	GetFilmPreview(id string) (domain.FilmPreview, error)
	GetAllFilmsPreviews() []domain.FilmPreview
}

type FilmsService struct {
	storage filmsStorage
}

func InitFilmsService(storage filmsStorage) *FilmsService {
	return &FilmsService{
		storage: storage,
	}
}

// TODO: remove this
func (filmsService *FilmsService) AddSomeData() error {
	data := []domain.FilmPreview{
		{
			Id:       "dfgea4ra424r4fw",
			Name:     "Film1",
			Duration: 3600,
		},
		{
			Id:       "fnuf7842huirn23",
			Name:     "Film2",
			Duration: 7200,
		},
		{
			Id:       "syh54eat4r4wf4wh",
			Name:     "Film3",
			Duration: 4800,
		},
	}

	for _, film := range data {
		err := filmsService.storage.AddFilm(film)
		if err != nil {
			return err
		}
	}

	return nil
}

func (filmsService *FilmsService) GetFilmsPreviews() ([]domain.FilmPreview, error) {
	films := filmsService.storage.GetAllFilmsPreviews()

	return films, nil
}
