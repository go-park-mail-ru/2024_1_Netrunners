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

func (filmsService *FilmsService) GetFilmsPreviews() ([]domain.FilmPreview, error) {
	films := filmsService.storage.GetAllFilmsPreviews()

	return films, nil
}
