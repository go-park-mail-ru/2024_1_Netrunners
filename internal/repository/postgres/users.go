package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type UsersStorage struct {
	pool *pgxpool.Pool
}

func NewUsersStorage(pool *pgxpool.Pool) (*UsersStorage, error) {
	return &UsersStorage{
		pool: pool,
	}, nil
}

func (storage *UsersStorage) CreateUser(user domain.UserSignUp) error {
	_, err := storage.pool.Exec(context.Background(),
		`insert into users (email, name, password) VALUES ($1, $2, $3);`,
		user.Email, user.Name, user.Password)
	if err != nil {
		return fmt.Errorf("error at inserting info into users in CreateUser: %w",
			myerrors.ErrInternalServerError)
	}

	return nil
}

func (storage *UsersStorage) GetUser(email string) (domain.User, error) {
	var user domain.User
	user.Email = email

	err := storage.pool.QueryRow(context.Background(),
		`select uuid, name, registered_at, birthday, is_admin
		from users
		where email = $1;`, email).Scan(
		&user.Avatar,
		&user.Name,
		&user.RegisteredAt,
		&user.Birthday,
		&user.IsAdmin)
	if err != nil {
		return domain.User{},
			fmt.Errorf("error at recieving data in GetUser: %w", err)
	}

	return user, nil
}

func (storage *UsersStorage) RemoveUser(email string) error {
	_, err := storage.pool.Exec(context.Background(),
		`update users
			set is_deleted = true
			where email = $1;`, email)
	if err != nil {
		return fmt.Errorf("error at recieving data in RemoveUser: %w",
			myerrors.ErrInternalServerError)
	}

	return nil
}

func (storage *UsersStorage) HasUser(email, password string) error {
	var passwordFromDB string
	err := storage.pool.QueryRow(context.Background(),
		`select password
		from users
		where email = $1;`, email).Scan(
		&passwordFromDB)
	if err != nil {
		return fmt.Errorf("error at recieving data in HasUser: %w",
			myerrors.ErrInternalServerError)
	}

	if passwordFromDB != password {
		return fmt.Errorf("incorrect password: %w",
			myerrors.ErrIncorrectLoginOrPassword)
	}

	return nil
}

func (storage *UsersStorage) ChangeUserPassword(email, newPassword string) error {
	_, err := storage.pool.Exec(context.Background(),
		`update users
			set password = $1, version += 1
			where email = $2;`, newPassword, email)
	if err != nil {
		return fmt.Errorf("error at updating data in ChangeUserPassword: %w", myerrors.ErrInternalServerError)
	}

	return nil
}

func (storage *UsersStorage) ChangeUserName(email, newUsername string) (domain.User, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.User{}, fmt.Errorf("error at begin transaction in ChangeUserName: %w",
			myerrors.ErrInternalServerError)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("error at rollback transaction in ChangeUserName: %v",
				myerrors.ErrInternalServerError)
		}
	}()

	err = tx.QueryRow(context.Background(),
		`update users
			set name = $1
			where email = $2;`, newUsername, email).Scan()
	if err != nil {
		return domain.User{}, fmt.Errorf("error at updating data in ChangeUserName: %w",
			myerrors.ErrInternalServerError)
	}

	var user domain.User
	user.Email = email
	err = tx.QueryRow(context.Background(),
		`select name, password, registered_at, birthday, version, is_admin
			from users
			where email = $1;`, email).Scan(
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.Version,
		&user.IsAdmin)
	if err != nil {
		return domain.User{},
			fmt.Errorf("error at recieving data in ChangeUserName: %w", myerrors.ErrInternalServerError)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, fmt.Errorf("error at commit transaction in ChangeUserName: %w",
			myerrors.ErrInternalServerError)
	}

	return user, nil
}

func (storage *UsersStorage) GetUserDataByUuid(uuid string) (domain.User, error) {
	var user domain.User
	err := storage.pool.QueryRow(context.Background(),
		`select email, name, password, registered_at, birthday, version, is_admin
			from users
			where uuid = $1;`, uuid).Scan(
		&user.Email,
		&user.Name,
		&user.Password,
		&user.RegisteredAt,
		&user.Birthday,
		&user.Version,
		&user.IsAdmin)
	if err != nil {
		return domain.User{},
			fmt.Errorf("error at recieving data in GetUserDataByUuid: %w", myerrors.ErrInternalServerError)
	}

	return user, nil
}

func (storage *UsersStorage) GetUserPreview(uuid string) (domain.UserPreview, error) {
	var userPreview domain.UserPreview

	userPreview.Avatar = uuid
	err := storage.pool.QueryRow(context.Background(),
		`select name
			from users
			where uuid = $1;`, uuid).Scan(
		&userPreview.Name)
	if err != nil {
		return domain.UserPreview{},
			fmt.Errorf("error at recieving data in GetUserPreview: %w", myerrors.ErrInternalServerError)
	}

	return userPreview, nil
}
