package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type SuccessResponse struct {
	Status int `json:"status"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func WriteSuccess(w http.ResponseWriter) error {
	response := SuccessResponse{
		Status: http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}

func WriteError(w http.ResponseWriter, err error) error {
	fmt.Println(err)
	statusCode, err := myerrors.ParseError(err)

	response := ErrorResponse{
		Status: statusCode,
		Err:    err.Error(),
	}

	jsonResponse, err := json.Marshal(response)
	fmt.Println(err)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func escapeUserData(user *domain.User) {
	user.Name = html.EscapeString(user.Name)
	user.Email = html.EscapeString(user.Email)
	user.Password = html.EscapeString(user.Password)
	user.Avatar = html.EscapeString(user.Avatar)
}

func escapeUserPreviewData(userPreview *domain.UserPreview) {
	userPreview.Name = html.EscapeString(userPreview.Name)
	userPreview.Avatar = html.EscapeString(userPreview.Avatar)
}

func escapeActorData(actor *domain.ActorData) {
	actor.Name = html.EscapeString(actor.Name)
	actor.Avatar = html.EscapeString(actor.Avatar)
	actor.Spouse = html.EscapeString(actor.Spouse)
	actor.Genres = html.EscapeString(actor.Genres)
	actor.BirthPlace = html.EscapeString(actor.BirthPlace)
	actor.Career = html.EscapeString(actor.Career)
}

func escapeFilmData(filmData *domain.FilmData) {
	filmData.Title = html.EscapeString(filmData.Title)
	filmData.Data = html.EscapeString(filmData.Data)
	filmData.Director = html.EscapeString(filmData.Director)
	filmData.Preview = html.EscapeString(filmData.Preview)
}

func escapeActorPreview(actor *domain.ActorPreview) {
	actor.Name = html.EscapeString(actor.Name)
	actor.Avatar = html.EscapeString(actor.Avatar)
}

func escapeFilmPreview(film *domain.FilmPreview) {
	film.Preview = html.EscapeString(film.Preview)
	film.Director = html.EscapeString(film.Director)
	film.Title = html.EscapeString(film.Title)
}

func escapeComment(comment *domain.Comment) {
	comment.Text = html.EscapeString(comment.Text)
	comment.Author = html.EscapeString(comment.Author)
}
