package errors

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")

	// cache errors
	ErrNoSuchItemInTheCache     = errors.New("no such token in cache")
	ErrItemsIsAlreadyInTheCache = errors.New("such token is already in the cache")
	ErrWrongSessionVersion      = errors.New("different versions")
	ErrNoSuchSessionInTheCache  = errors.New("no such session in cache")
	ErrNoSuchUserInTheCache     = errors.New("no such user in cache")
	ErrTooHighVersion           = errors.New("too high session version")

	// postgres users errors
	ErrNoSuchUser               = errors.New("no such user with given login")
	ErrUserAlreadyExists        = errors.New("user with given login already exists")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")

	ErrFailInQueryRow            = errors.New("failed to query row")
	ErrFailInQuery               = errors.New("failed to query rows")
	ErrFailInExec                = errors.New("failed to exec script")
	ErrFailInForEachRow          = errors.New("unexpected error at ForEachRow")
	ErrFailedToBeginTransaction  = errors.New("failed to begin transaction")
	ErrFailedToCommitTransaction = errors.New("failed to commit transaction")

	ErrNotFound              = errors.New("no rows for given uuid")
	ErrFavoriteAlreadyExists = errors.New("favorite already exists")
	ErrNoGenres              = errors.New("no genres found")
	ErrFilmAlreadyExists     = errors.New("film already exists")
	ErrNoSuchFilm            = errors.New("no such film with given uuid")
	ErrWrongScore            = errors.New("score should be from 1 to 5")
	ErrCommentAlreadyExists  = errors.New("comment for this film already exists")
)
