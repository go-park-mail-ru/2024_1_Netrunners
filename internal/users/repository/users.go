package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type PgxIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

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

const hasSubscription = `
	select subscription_end_date
	from users
	where external_id = $1 and subscription_end_date > now();`

const addSubscription = `
	update users
	set subscription_end_date = $1
	where external_id = $2;`

const getSubscriptions = `
	select id, title, amount, description, duration
	from subscription;`

const getSubscription = `
	select id, title, amount, description, duration
	from subscription
	where id = $1;`

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
	_, err := storage.pool.Exec(context.Background(), putNewUserPassword, newPassword, email)
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

	return user, nil
}

func (storage *UsersStorage) ChangeUserName(email, newUsername string) (domain.User, error) {
	_, err := storage.pool.Exec(context.Background(), putNewUsername, newUsername, email)
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
	_, err := storage.pool.Exec(context.Background(), putNewUserPasswordByUuid, newPassword, uuid)
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

	return user, nil
}

func (storage *UsersStorage) ChangeUserNameByUuid(uuid, newUsername string) (domain.User, error) {
	_, err := storage.pool.Exec(context.Background(), putNewUsernameByUuid, newUsername, uuid)
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

	return user, nil
}

func (storage *UsersStorage) ChangeUserAvatarByUuid(uuid, filename string) (domain.User, error) {
	_, err := storage.pool.Exec(context.Background(), putNewUserAvatarByUuid, filename, uuid)
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

	return user, nil
}

func (storage *UsersStorage) HasSubscription(uuid string) (bool, error) {
	var time time.Time
	err := storage.pool.QueryRow(context.Background(), hasSubscription, uuid).Scan(&time)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (storage *UsersStorage) AddSubscription(uuid string, newDate string) error {
	_, err := storage.pool.Exec(context.Background(), addSubscription, newDate, uuid)
	if err != nil {
		return fmt.Errorf("failed to add subscription: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	return nil
}

func (storage *UsersStorage) GetSubscriptions() ([]domain.Subscription, error) {
	rows, err := storage.pool.Query(context.Background(), getSubscriptions)
	if err != nil {
		return nil, err
	}

	var (
		subs        []domain.Subscription
		uuid        string
		title       string
		description string
		amount      string
		duration    string
	)
	for rows.Next() {
		sub := domain.Subscription{}
		err = rows.Scan(&uuid, &title, &amount, &description, &duration)
		if err != nil {
			return nil, err
		}
		sub.Uuid = uuid
		sub.Title = title
		sub.Description = description
		price, err := strconv.ParseFloat(amount, 32)
		if err != nil {
			return nil, err
		}
		sub.Amount = float32(price)
		durat, err := strconv.ParseUint(duration, 32, 32)
		if err != nil {
			return nil, err
		}
		sub.Duration = uint32(durat)

		subs = append(subs, sub)
	}
	return subs, nil
}

func (storage *UsersStorage) GetSubscription(uuid string) (domain.Subscription, error) {
	var sub domain.Subscription
	err := storage.pool.QueryRow(context.Background(), getSubscription, uuid).Scan(
		&sub.Uuid, &sub.Title, &sub.Amount, &sub.Description, &sub.Duration)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("failed to get subscription: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	return sub, nil
}
