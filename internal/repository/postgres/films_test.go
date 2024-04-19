package database

import (
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestFilmsStorage_GetFilmDataByUuid(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	newFilmData := NewMockFilmData()
	uuid := "1"

	mockRows := pgxmock.NewRows([]string{"uuid", "title", "banner", "link", "name", "data", "duration", "published_at",
		"avg_score", "scores", "age_limit"}).
		AddRow(newFilmData.Uuid, newFilmData.Title, newFilmData.Preview, newFilmData.Link, newFilmData.Director,
			newFilmData.Data, newFilmData.Duration, newFilmData.Date, newFilmData.AverageScore,
			newFilmData.ScoresCount, newFilmData.AgeLimit)

	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	filmData, err := storage.GetFilmDataByUuid(uuid)
	require.NoError(t, err)
	require.Equal(t, newFilmData, filmData)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_AddFilm(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectBeginTx(pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	storage, err := NewFilmsStorage(mock)

	newFilm := NewMockFilmDataToAdd()

	mockGetAmountOfDirectorsByName := pgxmock.NewRows([]string{"name"}).
		AddRow(0)
	mock.ExpectQuery("SELECT").
		WithArgs(newFilm.Director).
		WillReturnRows(mockGetAmountOfDirectorsByName)

	mock.ExpectExec("INSERT").
		WithArgs(newFilm.Director).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mockGetDirectorsIdByName := pgxmock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").
		WithArgs(newFilm.Director).
		WillReturnRows(mockGetDirectorsIdByName)

	mock.ExpectExec("INSERT").
		WithArgs(newFilm.Title, newFilm.Preview, 1, newFilm.Data, newFilm.AgeLimit, newFilm.Duration,
			newFilm.PublishedAt).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mockGetFilmIdByTitle := pgxmock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").
		WithArgs(newFilm.Title).
		WillReturnRows(mockGetFilmIdByTitle)

	mockGetAmountOfActorsByName := pgxmock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").
		WithArgs(newFilm.Actors[0].Name).
		WillReturnRows(mockGetAmountOfActorsByName)

	mockGetActorId := pgxmock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").
		WithArgs(newFilm.Actors[0].Name).
		WillReturnRows(mockGetActorId)

	mock.ExpectExec("INSERT").
		WithArgs(1, 1).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = storage.AddFilm(newFilm)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_RemoveFilm(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	uuid := "1"

	mock.ExpectExec("DELETE").
		WithArgs(uuid).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = storage.RemoveFilm("1")
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_GetFilmPreview(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	newFilmPreview := NewMockFilmPreview()
	uuid := "1"

	mockRows := pgxmock.NewRows([]string{"uuid", "title", "banner", "name", "duration", "avg_score", "scores", "age_limit"}).
		AddRow(newFilmPreview.Uuid, newFilmPreview.Title, newFilmPreview.Preview, newFilmPreview.Director,
			newFilmPreview.Duration, newFilmPreview.AverageScore, newFilmPreview.ScoresCount, newFilmPreview.AgeLimit)

	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	filmPreview, err := storage.GetFilmPreview(uuid)
	require.NoError(t, err)
	require.Equal(t, newFilmPreview, filmPreview)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_GetAllFilmsPreviews(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	newFilmPreviews := NewMockFilmPreviews()

	mockRows := pgxmock.NewRows([]string{"uuid", "title", "banner", "name", "duration", "avg_score", "scores",
		"age_limit"}).
		AddRow(newFilmPreviews[0].Uuid, newFilmPreviews[0].Title, newFilmPreviews[0].Preview,
			newFilmPreviews[0].Director, newFilmPreviews[0].Duration, newFilmPreviews[0].AverageScore,
			newFilmPreviews[0].ScoresCount, newFilmPreviews[0].AgeLimit).
		AddRow(newFilmPreviews[1].Uuid, newFilmPreviews[1].Title, newFilmPreviews[1].Preview,
			newFilmPreviews[1].Director, newFilmPreviews[1].Duration, newFilmPreviews[1].AverageScore,
			newFilmPreviews[1].ScoresCount, newFilmPreviews[0].AgeLimit)

	mock.ExpectQuery("SELECT").
		WithArgs().
		WillReturnRows(mockRows)

	filmPreview, err := storage.GetAllFilmsPreviews()
	require.NoError(t, err)
	require.Equal(t, newFilmPreviews, filmPreview)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_GetAllFilmActors(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	newFilmActors := NewMockFilmActors()
	uuid := "1"

	mockRows := pgxmock.NewRows([]string{"uuid", "title", "avatar"}).
		AddRow(newFilmActors[0].Uuid, newFilmActors[0].Name, newFilmActors[0].Avatar).
		AddRow(newFilmActors[1].Uuid, newFilmActors[1].Name, newFilmActors[1].Avatar).
		AddRow(newFilmActors[2].Uuid, newFilmActors[2].Name, newFilmActors[2].Avatar)

	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	filmActors, err := storage.GetAllFilmActors(uuid)
	require.NoError(t, err)
	require.Equal(t, newFilmActors, filmActors)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestFilmsStorage_GetAllFilmComments(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewFilmsStorage(mock)

	newFilmComments := NewMockFilmComments()
	uuid := "1"

	mockRows := pgxmock.NewRows([]string{"uuid", "film_uuid", "author", "text", "score", "added_at"}).
		AddRow(newFilmComments[0].Uuid, newFilmComments[0].FilmUuid, newFilmComments[0].Author, newFilmComments[0].Text, newFilmComments[0].Score,
			newFilmComments[0].AddedAt).
		AddRow(newFilmComments[1].Uuid, newFilmComments[1].FilmUuid, newFilmComments[1].Author, newFilmComments[1].Text, newFilmComments[1].Score,
			newFilmComments[1].AddedAt).
		AddRow(newFilmComments[2].Uuid, newFilmComments[2].FilmUuid, newFilmComments[2].Author, newFilmComments[2].Text, newFilmComments[2].Score,
			newFilmComments[2].AddedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	filmComments, err := storage.GetAllFilmComments(uuid)
	require.NoError(t, err)
	require.Equal(t, newFilmComments, filmComments)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
