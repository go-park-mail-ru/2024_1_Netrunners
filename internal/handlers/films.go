package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type FilmsPageHandlers struct {
	filmsService *service.FilmsService
}

func InitFilmsPageHandlers(filmsService *service.FilmsService) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		filmsService: filmsService,
	}
}

func (filmsPageHandlers *FilmsPageHandlers) GetFilmsPreviews(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	films, err := filmsPageHandlers.filmsService.GetFilmsPreviews()
	if err != nil {
		err = WriteError(w, errors.New("Unexpected error"))
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
	} else {
		response = filmsPreviewsResponse{
			Status: http.StatusOK,
			Films:  films,
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println(err)
	}
}

type filmsPreviewsResponse struct {
	Status int                  `json:"status"`
	Films  []domain.FilmPreview `json:"films"`
}
