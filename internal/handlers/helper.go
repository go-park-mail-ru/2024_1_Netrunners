package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
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
	actor.BirthPlace = html.EscapeString(actor.BirthPlace)
	actor.Career = html.EscapeString(actor.Career)
}

func escapeFilmData(filmData *domain.FilmData) {
	filmData.Title = html.EscapeString(filmData.Title)
	filmData.Data = html.EscapeString(filmData.Data)
	filmData.Director = html.EscapeString(filmData.Director)
	filmData.Preview = html.EscapeString(filmData.Preview)
	for _, genre := range filmData.Genres {
		genre = html.EscapeString(genre)
	}
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

func escapeGenreFilms(genreFilms *domain.GenreFilms) {
	genreFilms.Name = html.EscapeString(genreFilms.Name)
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
		Genres:       film.Genres,
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

func convertUserSignUpDataToRegular(userData domain.UserSignUp) *session.UserSignUp {
	return &session.UserSignUp{
		Email:    userData.Email,
		Password: userData.Password,
		Username: userData.Name,
	}
}

func convertUserPreviewToRegular(user *session.UserPreview) domain.UserPreview {
	return domain.UserPreview{
		Uuid:   user.Uuid,
		Name:   user.Username,
		Avatar: user.Avatar,
	}
}

func convertGenreFilmsToRegular(genreFilms *session.GenreFilms) domain.GenreFilms {
	var filmsConverted []domain.FilmPreview
	for _, film := range genreFilms.Films {
		filmsConverted = append(filmsConverted, convertFilmPreviewToRegular(film))
	}
	return domain.GenreFilms{
		Name:  genreFilms.Genre,
		Uuid:  genreFilms.GenreUuid,
		Films: filmsConverted,
	}
}

func IsTokenValid(token *http.Cookie, secretKey string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Login"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	_, ok = claims["IsAdmin"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Version"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	return claims, nil
}

func ValidateLogin(e string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if emailRegex.MatchString(e) {
		return nil
	}
	return myerrors.ErrLoginIsNotValid
}

func ValidateUsername(username string) error {
	if len(username) >= 4 {
		return nil
	}
	return myerrors.ErrUsernameIsToShort
}

func ValidatePassword(password string) error {
	if len(password) >= 6 {
		return nil
	}
	return myerrors.ErrPasswordIsToShort
}
