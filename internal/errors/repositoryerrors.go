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

	// mockDB users errors
	ErrNoSuchUser               = errors.New("no such user with given login")
	ErrUserAlreadyExists        = errors.New("user with given login already exists")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")

	// mockDB films errors
	ErrNoSuchFilm = errors.New("no such film with given uuid")

	ErrNoSuchActor     = errors.New("no such actor with given uuid")
	ErrNoActorsForFilm = errors.New("no actors for this film")

	ErrNotFound = errors.New("no rows for given uuid")
)
