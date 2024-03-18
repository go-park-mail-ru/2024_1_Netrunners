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
	storage          filmsStorage
	localStoragePath string
}

func InitFilmsService(storage filmsStorage, localStoragePath string) *FilmsService {
	return &FilmsService{
		storage:          storage,
		localStoragePath: localStoragePath,
	}
}

// TODO: remove this
func (filmsService *FilmsService) AddSomeData() error {
	data := []domain.FilmPreview{
		{
			Uuid: "dfgea4ra424r4fw",
			Preview: "https://m.media-amazon.com/images/M/MV5BNzlkNzVjMDMtOTdhZC00MGE1LTkxODctMzFmMjkwZm" +
				"MxZjFhXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_.jpg",
			Title:    "Fast and Furious 1",
			Duration: 3600,
		},
		{
			Uuid:     "fnuf7842huirn23",
			Preview:  "https://m.media-amazon.com/images/I/71Wo+cFznbL.jpg",
			Title:    "Fast and Furious 2",
			Duration: 7200,
		},
		{
			Uuid:     "syh54eat4r4wf4wh",
			Preview:  "https://m.media-amazon.com/images/I/71ql8kIrPKL.jpg",
			Title:    "Fast and Furious 3",
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
