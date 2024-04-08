package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

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
	films, err := filmsPageHandlers.filmsService.GetAllFilmsPreviews()
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
	}

	for _, film := range films {
		film.Preview = html.EscapeString(film.Preview)
		film.Director = html.EscapeString(film.Director)
		film.Title = html.EscapeString(film.Title)
	}

	response := filmsPreviewsResponse{
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
	uuid := mux.Vars(r)["uuid"]

	filmData, err := filmsPageHandlers.filmsService.GetFilmDataByUuid(uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	filmData.Title = html.EscapeString(filmData.Title)
	filmData.Data = html.EscapeString(filmData.Data)
	filmData.Director = html.EscapeString(filmData.Director)
	filmData.Preview = html.EscapeString(filmData.Preview)

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
	uuid := mux.Vars(r)["uuid"]
	comments, err := filmsPageHandlers.filmsService.GetAllFilmComments(uuid)
	if err != nil {
		err = WriteError(w, err)

		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	for _, comment := range comments {
		comment.Text = html.EscapeString(comment.Text)
		comment.Author = html.EscapeString(comment.Author)
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
	uuid := mux.Vars(r)["uuid"]
	actors, err := filmsPageHandlers.filmsService.GetAllFilmActors(uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			filmsPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	for _, actor := range actors {
		actor.Name = html.EscapeString(actor.Name)
		actor.Avatar = html.EscapeString(actor.Avatar)
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
