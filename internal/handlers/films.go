package handlers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type filmLink struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
}

type FilmsPageHandlers struct {
	filmsService *service.FilmsService
	logger       *zap.SugaredLogger
}

func NewFilmsPageHandlers(filmsService *service.FilmsService, logger *zap.SugaredLogger) *FilmsPageHandlers {
	return &FilmsPageHandlers{
		filmsService: filmsService,
		logger:       logger,
	}
}

type filmsPreviewsResponse struct {
	Status int                  `json:"status"`
	Films  []domain.FilmPreview `json:"films"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmsPreviews(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	films, err := filmsPageHandlers.filmsService.GetAllFilmsPreviews()
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
	}
	response = filmsPreviewsResponse{
		Status: http.StatusOK,
		Films:  films,
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

type filmDataResponse struct {
	Status   int             `json:"status"`
	FilmData domain.FilmData `json:"film"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetFilmDataByUuid(w http.ResponseWriter, r *http.Request) {
	var uuid string
	err := json.NewDecoder(r.Body).Decode(&uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	filmData, err := filmsPageHandlers.filmsService.GetFilmByUuid(uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	response := filmDataResponse{
		Status:   http.StatusOK,
		FilmData: filmData,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}
}

type filmCommentsResponse struct {
	Status   int              `json:"status"`
	Comments []domain.Comment `json:"comments"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmComments(w http.ResponseWriter, r *http.Request) {
	var uuid string
	err := json.NewDecoder(r.Body).Decode(&uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}
	comments, err := filmsPageHandlers.filmsService.GetAllFilmComments(uuid)
	if err != nil {
		err = WriteError(w, err)

		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	response := filmCommentsResponse{
		Status:   http.StatusOK,
		Comments: comments,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

}

type filmActorsResponse struct {
	Status int                   `json:"status"`
	Actors []domain.ActorPreview `json:"actors"`
}

func (filmsPageHandlers *FilmsPageHandlers) GetAllFilmActors(w http.ResponseWriter, r *http.Request) {
	var uuid string
	err := json.NewDecoder(r.Body).Decode(&uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			if err != nil {
				filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
			}
		}
		return
	}

	actors, err := filmsPageHandlers.filmsService.GetAllFilmActors(uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	response := filmActorsResponse{
		Status: http.StatusOK,
		Actors: actors,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}
}

func (filmsPageHandlers *FilmsPageHandlers) AddFilm(w http.ResponseWriter, r *http.Request) {
	var filmData domain.FilmDataToAdd
	err := json.NewDecoder(r.Body).Decode(&filmData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	err = filmsPageHandlers.filmsService.AddFilm(filmData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			if err != nil {
				filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
			}
		}
		return
	}

	err = WriteSuccess(w)
	if err != nil {
		filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}

	filmsPageHandlers.logger.Info("film added successfully")
}
