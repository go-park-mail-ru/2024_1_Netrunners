package service

import (
	"encoding/base64"
	"fmt"
	"os"
	"path"

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
			Id:       "dfgea4ra424r4fw",
			Name:     "Fast and Furious 1",
			Duration: 3600,
		},
		{
			Id:       "fnuf7842huirn23",
			Name:     "Fast and Furious 2",
			Duration: 7200,
		},
		{
			Id:       "syh54eat4r4wf4wh",
			Name:     "Fast and Furious 3",
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

	for i, film := range films {
		fileBytes, err := os.ReadFile(path.Join(filmsService.localStoragePath, "films", film.Id, "preview.png"))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		films[i].Preview = []byte(base64.StdEncoding.EncodeToString(fileBytes))
	}

	return films, nil
}
