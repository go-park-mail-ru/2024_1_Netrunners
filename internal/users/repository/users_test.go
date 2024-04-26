package repository

import (
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/postgres"
	"regexp"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

func TestUsersStorage_CreateUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)

	newUser := database.NewMockUserSignUp()

	mock.ExpectExec("INSERT").
		WithArgs(newUser.Email, newUser.Name, newUser.Password).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = storage.CreateUser(newUser)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_GetUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)

	newUser := database.NewMockUser()

	mockRows := pgxmock.NewRows([]string{"uuid", "email", "avatar", "name", "password", "registered_at", "birthday",
		"is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt,
			newUser.Birthday, newUser.IsAdmin)

	mock.ExpectQuery("SELECT").
		WithArgs("cakethefake@gmail.com").
		WillReturnRows(mockRows)

	user, err := storage.GetUser("cakethefake@gmail.com")
	require.NoError(t, err)
	require.Equal(t, newUser, user)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_RemoveUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)

	email := "cakethefake@gmail.com"

	mock.ExpectExec("DELETE").
		WithArgs(email).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = storage.RemoveUser("cakethefake@gmail.com")
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_HasUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)

	newUser := database.NewMockUser()
	email := "cakethefake@gmail.com"
	password := "123456789"

	mockRows := pgxmock.NewRows([]string{"password"}).
		AddRow(newUser.Password)

	mock.ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(mockRows)

	err = storage.HasUser(email, password)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_ChangeUserPassword(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectBeginTx(pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	storage, err := NewUsersStorage(mock)

	email := "cakethefake@gmail.com"
	password := "123456789"

	mock.ExpectExec("UPDATE").
		WithArgs(password, email).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	newUser := database.NewMockUser()
	mockRows := pgxmock.NewRows([]string{"uuid", "email", "name", "avatar", "password", "registered_at", "birthday",
		"is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt,
			newUser.Birthday, newUser.IsAdmin)
	mock.ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(mockRows)

	mock.ExpectCommit()
	mock.ExpectRollback()

	_, err = storage.ChangeUserPassword(email, password)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_GetUserDataByUuid(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)

	newUser := database.NewMockUser()
	uuid := "1"

	mockRows := pgxmock.NewRows([]string{"uuid", "email", "avatar", "name", "password", "registered_at", "birthday", "is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt, newUser.Birthday, newUser.IsAdmin)

	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	user, err := storage.GetUserDataByUuid(uuid)
	require.NoError(t, err)
	require.Equal(t, newUser, user)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_ChangeUserPasswordByUuid(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectBeginTx(pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	storage, err := NewUsersStorage(mock)

	uuid := "1"
	password := "123456789"

	mock.ExpectExec("UPDATE").
		WithArgs(password, uuid).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	newUser := database.NewMockUser()
	mockRows := pgxmock.NewRows([]string{"uuid", "email", "avatar", "name", "password", "registered_at", "birthday",
		"is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt,
			newUser.Birthday, newUser.IsAdmin)
	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	mock.ExpectCommit()

	_, err = storage.ChangeUserPasswordByUuid(uuid, password)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_ChangeUserName(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectBeginTx(pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	storage, err := NewUsersStorage(mock)

	email := "cakethefake@gmial.com"
	newUsername := "123456789"

	mock.ExpectExec("UPDATE").
		WithArgs(newUsername, email).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	newUser := database.NewMockUser()
	mockRows := pgxmock.NewRows([]string{"uuid", "email", "avatar", "name", "password", "registered_at", "birthday",
		"is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt,
			newUser.Birthday, newUser.IsAdmin)
	mock.ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(mockRows)

	mock.ExpectCommit()

	_, err = storage.ChangeUserName(email, newUsername)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_ChangeUserNameByUuid(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectBeginTx(pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	storage, err := NewUsersStorage(mock)

	uuid := "1"
	password := "123456789"

	mock.ExpectExec("UPDATE").
		WithArgs(password, uuid).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	newUser := database.NewMockUser()
	mockRows := pgxmock.NewRows([]string{"uuid", "email", "avatar", "name", "password", "registered_at", "birthday",
		"is_admin"}).
		AddRow(newUser.Uuid, newUser.Email, newUser.Avatar, newUser.Name, newUser.Password, newUser.RegisteredAt,
			newUser.Birthday, newUser.IsAdmin)
	mock.ExpectQuery("SELECT").
		WithArgs(uuid).
		WillReturnRows(mockRows)

	mock.ExpectCommit()

	_, err = storage.ChangeUserNameByUuid(uuid, password)
	require.Equal(t, nil, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUsersStorage_GetUserPreview(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	storage, err := NewUsersStorage(mock)
	require.NoError(t, err)

	uuid := "test_uuid"
	expectedUserPreview := domain.UserPreview{Uuid: uuid, Name: "Test User", Avatar: "avatar.com"}

	mock.ExpectQuery(regexp.QuoteMeta(getUserPreviewByUuid)).
		WithArgs(uuid).
		WillReturnRows(
			pgxmock.NewRows([]string{"uuid", "name", "avatar"}).
				AddRow(expectedUserPreview.Uuid, expectedUserPreview.Name, expectedUserPreview.Avatar),
		)

	userPreview, err := storage.GetUserPreview(uuid)
	require.NoError(t, err)
	assert.Equal(t, expectedUserPreview, userPreview)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
