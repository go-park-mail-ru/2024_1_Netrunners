package handlers

import (
	"encoding/json"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"html"
	"net/http"
	"time"

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
	statusCode, err := myerrors.ParseError(err)

	response := ErrorResponse{
		Status: statusCode,
		Err:    err.Error(),
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

func convertFilmPreviewToRegular(film *session.FilmPreview) domain.FilmPreview {
	filmNew := domain.FilmPreview{
		Uuid:         film.Uuid,
		Title:        film.Title,
		Preview:      film.Preview,
		Director:     film.Director,
		AverageScore: film.AvgScore,
		ScoresCount:  film.ScoresCount,
		AgeLimit:     film.AgeLimit,
		Duration:     film.Duration,
	}
	return filmNew
}

func convertFilmDataToRegular(film *session.FilmData) domain.FilmData {
	return domain.FilmData{
		Uuid:         film.Uuid,
		Title:        film.Title,
		Preview:      film.Preview,
		Director:     film.Director,
		Link:         film.Link,
		Data:         film.Data,
		Date:         convertProtoToTime(film.Date),
		AgeLimit:     film.AgeLimit,
		AverageScore: film.AvgScore,
		ScoresCount:  film.ScoresCount,
		Duration:     film.Duration,
	}
}

func convertCommentToRegular(comment *session.Comment) domain.Comment {
	return domain.Comment{
		Uuid:     comment.Uuid,
		FilmUuid: comment.FilmUuid,
		Text:     comment.Text,
		Author:   comment.Author,
		Score:    comment.Score,
		AddedAt:  convertProtoToTime(comment.AddedAt),
	}
}

func convertActorPreviewToRegular(actor *session.ActorPreview) domain.ActorPreview {
	return domain.ActorPreview{
		Uuid:   actor.Uuid,
		Name:   actor.Name,
		Avatar: actor.Avatar,
	}
}

func convertActorDataToRegular(actor *session.ActorData) domain.ActorData {
	var filmsPreview []domain.FilmPreview
	for _, film := range actor.FilmsPreviews {
		filmRegular := convertFilmPreviewToRegular(film)
		escapeFilmPreview(&filmRegular)
		filmsPreview = append(filmsPreview, filmRegular)
	}
	return domain.ActorData{
		Uuid:     actor.Uuid,
		Name:     actor.Name,
		Avatar:   actor.Avatar,
		Birthday: convertProtoToTime(actor.Birthday),
		Career:   actor.Career,
		Spouse:   actor.Spouse,
		Genres:   actor.Genres,
		Films:    filmsPreview,
	}
}

func convertProtoToTime(protoTime *timestamppb.Timestamp) time.Time {
	return protoTime.AsTime()
}

func convertUserToRegular(user *session.User) domain.User {
	return domain.User{
		Uuid:         user.Uuid,
		Email:        user.Email,
		Password:     user.Password,
		Name:         user.Username,
		Version:      user.Version,
		IsAdmin:      user.IsAdmin,
		Avatar:       user.Avatar,
		Birthday:     convertProtoToTime(user.Birthday),
		RegisteredAt: convertProtoToTime(user.RegisteredAt),
	}
}

func convertUserPreviewToRegular(user *session.UserPreview) domain.UserPreview {
	return domain.UserPreview{
		Uuid:   user.Uuid,
		Name:   user.Username,
		Avatar: user.Avatar,
	}
}
