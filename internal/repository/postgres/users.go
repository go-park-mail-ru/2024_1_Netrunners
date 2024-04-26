package database

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/repository"

	"github.com/jackc/pgx/v5"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type UsersStorage struct {
	pool repository.PgxIface
}

func NewUsersStorage(pool repository.PgxIface) (*UsersStorage, error) {
	return &UsersStorage{
		pool: pool,
	}, nil
}

const insertUser = `INSERT INTO users (email, name, password) VALUES ($1, $2, $3);`

const getUserData = `
		SELECT external_id, email, avatar, name, password, registered_at, birthday, is_admin
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

const putNewUserPasswordByUuid = `
		UPDATE users
		SET password = $1
		WHERE external_id = $2;`

const putNewUserAvatarByUuid = `
		UPDATE users
		SET avatar = $1
		WHERE external_id = $2;`

const putNewUsernameByUuid = `
		UPDATE users
		SET name = $1
		WHERE external_id = $2;`

const getUserDataByUuid = `
		SELECT external_id, email, avatar, name, password, registered_at, birthday, is_admin
		FROM users
		WHERE external_id = $1;`

const getUserPreviewByUuid = `
		SELECT external_id, name, avatar 
		FROM users
		WHERE external_id = $1;`

func (storage *UsersStorage) CreateUser(user domain.UserSignUp) error {
	_, err := storage.pool.Exec(context.Background(), insertUser, user.Email, user.Name, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w: %w", err,
			myerrors.ErrFailInExec)
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
		return domain.User{}, fmt.Errorf("failed to get user: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	return user, nil
}

func (storage *UsersStorage) RemoveUser(email string) error {
	_, err := storage.pool.Exec(context.Background(), deleteUser, email)
	if err != nil {
		return fmt.Errorf("failed to remove user: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	return nil
}

func (storage *UsersStorage) HasUser(email, password string) error {
	var passwordFromDB string
	err := storage.pool.QueryRow(context.Background(), getAmountOfUserByName, email).Scan(&passwordFromDB)
	if err != nil {
		return fmt.Errorf("failed to get user for password check: %w: %w", err,
			myerrors.ErrFailInQuery)
	}

	if passwordFromDB != password {
		return fmt.Errorf("failed to compare passwords: %w",
			myerrors.ErrIncorrectLoginOrPassword)
	}

	return nil
}

func (storage *UsersStorage) ChangeUserPassword(email, newPassword string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change password: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change password: %v", err)
		}
	}()

	_, err = tx.Exec(context.Background(), putNewUserPassword, newPassword, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to update data to change password: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserData, email).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get new user data: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change password: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
	}

	return user, nil
}

func (storage *UsersStorage) ChangeUserName(email, newUsername string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change username: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change username: %v", err)
		}
	}()

	_, err = tx.Exec(context.Background(), putNewUsername, newUsername, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to change username: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserData, email).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get new user data: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change username: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
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
		return domain.User{}, fmt.Errorf("failed to get user data by uuid: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	return user, nil
}

func (storage *UsersStorage) GetUserPreview(uuid string) (domain.UserPreview, error) {
	var userPreview domain.UserPreview

	err := storage.pool.QueryRow(context.Background(), getUserPreviewByUuid, uuid).Scan(&userPreview.Uuid,
		&userPreview.Name, &userPreview.Avatar)
	if err != nil {
		return domain.UserPreview{}, fmt.Errorf("failed to get user preview: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	return userPreview, nil
}

func (storage *UsersStorage) ChangeUserPasswordByUuid(uuid, newPassword string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change password: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change password: %v", err)
		}
	}()

	_, err = tx.Exec(context.Background(), putNewUserPasswordByUuid, newPassword, uuid)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed update data to change password: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserDataByUuid, uuid).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get new user data: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change password: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
	}

	return user, nil
}

func (storage *UsersStorage) ChangeUserNameByUuid(uuid, newUsername string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change username: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change username: %v", err)
		}
	}()

	_, err = tx.Exec(context.Background(), putNewUsernameByUuid, newUsername, uuid)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to change username: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserDataByUuid, uuid).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get new user data: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change username: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
	}
	return user, nil
}

func (storage *UsersStorage) ChangeUserAvatarByUuid(uuid, filename string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to begin transaction to change username: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to change username: %v", err)
		}
	}()

	_, err = tx.Exec(context.Background(), putNewUserAvatarByUuid, filename, uuid)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to change user's avatar: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var user domain.User
	err = storage.pool.QueryRow(context.Background(), getUserDataByUuid, uuid).Scan(
		&user.Uuid,
		&user.Email,
		&user.Avatar,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get new user data: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to commit transaction to change username: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
	}
	return user, nil
}
