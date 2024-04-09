package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type UsersStorage struct {
	pool PgxIface
}

func NewUsersStorage(pool PgxIface) (*UsersStorage, error) {
	return &UsersStorage{
		pool: pool,
	}, nil
}

const insertUser = `INSERT INTO users (email, name, password) VALUES ($1, $2, $3);`

const getUserData = `
		SELECT uuid, email, avatar, name, password, registered_at, birthday, is_admin
		FROM users
		WHERE email = $1;`

const deleteUser = `
		DELETE FROM users
		WHERE email = $1;`

const getAmountOfUserByName = `
		SELECT password
		FROM users
		WHERE email = $1;`

const putNewUserPassword = `
		UPDATE users
		SET password = $1
		WHERE email = $2;`

const putNewUsername = `
		UPDATE users
		SET name = $1
		WHERE email = $2;`

const getUserDataByUuid = `
		SELECT uuid, email, avatar, name, password, registered_at, birthday, is_admin
		FROM users
		WHERE uuid = $1;`

const getUserPreviewByUuid = `
		SELECT name
		FROM users
		WHERE uuid = $1;`

func (storage *UsersStorage) CreateUser(user domain.UserSignUp) error {
	_, err := storage.pool.Exec(context.Background(), insertUser, user.Email, user.Name, user.Password)
	if err != nil {
		return myerrors.ErrInternalServerError
	}
	return nil
}

func (storage *UsersStorage) GetUser(email string) (domain.User, error) {
	var user domain.User
	user.Email = email

	err := storage.pool.QueryRow(context.Background(), getUserData, email).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, myerrors.ErrInternalServerError
	}
	return user, nil
}

func (storage *UsersStorage) RemoveUser(email string) error {
	_, err := storage.pool.Exec(context.Background(), deleteUser, email)
	if err != nil {
		return myerrors.ErrInternalServerError
	}

	return nil
}

func (storage *UsersStorage) HasUser(email, password string) error {
	var passwordFromDB string
	err := storage.pool.QueryRow(context.Background(), getAmountOfUserByName, email).Scan(&passwordFromDB)
	if err != nil {
		return myerrors.ErrInternalServerError
	}

	if passwordFromDB != password {
		return fmt.Errorf("incorrect password: %w",
			myerrors.ErrIncorrectLoginOrPassword)
	}

	return nil
}

func (storage *UsersStorage) ChangeUserPassword(email, newPassword string) error {
	_, err := storage.pool.Exec(context.Background(), putNewUserPassword, newPassword, email)
	if err != nil {
		return myerrors.ErrInternalServerError
	}

	return nil
}

func (storage *UsersStorage) ChangeUserName(email, newUsername string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change username: %w",
			myerrors.ErrInternalServerError)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change username: %v",
				myerrors.ErrInternalServerError)
		}
	}()

	err = tx.QueryRow(context.Background(), putNewUsername, newUsername, email).Scan()
	if err != nil {
		return domain.User{}, myerrors.ErrInternalServerError
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserData, email).Scan(
		&user.Uuid,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, myerrors.ErrInternalServerError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change username: %w",
			myerrors.ErrInternalServerError)
	}
	return user, nil
}

func (storage *UsersStorage) GetUserDataByUuid(uuid string) (domain.User, error) {
	var user domain.User
	err := storage.pool.QueryRow(context.Background(), getUserDataByUuid, uuid).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, myerrors.ErrInternalServerError
	}
	return user, nil
}

func (storage *UsersStorage) GetUserPreview(uuid string) (domain.UserPreview, error) {
	var userPreview domain.UserPreview

	userPreview.Avatar = uuid
	err := storage.pool.QueryRow(context.Background(), getUserPreviewByUuid, uuid).Scan(&userPreview.Name)
	if err != nil {
		return domain.UserPreview{}, myerrors.ErrInternalServerError
	}
	return userPreview, nil
}
