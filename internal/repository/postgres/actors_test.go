package database

import (
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestActorsStorage_GetActorByUuid(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewActorsStorage(mock)

	newUser := NewMockActor()

	mockRowsData := pgxmock.NewRows([]string{"uuid", "name", "avatar", "birthday", "career", "height", "birth_place",
		"genres", "spouse"}).
		AddRow(newUser.Uuid, newUser.Name, newUser.Avatar, newUser.Birthday, newUser.Career, newUser.Height,
			newUser.BirthPlace, newUser.Genres, newUser.Spouse)
	mockRowsFilms := pgxmock.NewRows([]string{"uuid", "title"}).AddRow(newUser.Films[0].Uuid,
		newUser.Films[0].Title).AddRow(newUser.Films[1].Uuid, newUser.Films[1].Title).AddRow(newUser.Films[2].Uuid,
		newUser.Films[2].Title)

	mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnRows(mockRowsData)
	mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnRows(mockRowsFilms)

	user, err := storage.GetActorByUuid("1")
	require.NoError(t, err)
	require.Equal(t, newUser, user)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestActorsStorage_GetActorsByFilm(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewActorsStorage(mock)

	newActorPreviews := NewMockActorPreview()

	mockRowsFilms := pgxmock.NewRows([]string{"uuid", "name", "avatar"}).
		AddRow(newActorPreviews[0].Uuid, newActorPreviews[0].Name, newActorPreviews[0].Avatar).
		AddRow(newActorPreviews[1].Uuid, newActorPreviews[1].Name, newActorPreviews[1].Avatar).
		AddRow(newActorPreviews[2].Uuid, newActorPreviews[2].Name, newActorPreviews[2].Avatar)

	mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnRows(mockRowsFilms)

	actorsPreview, err := storage.GetActorsByFilm("1")
	require.NoError(t, err)
	require.Equal(t, newActorPreviews, actorsPreview)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
