package errors

import (
	"errors"
)

func ParseError(err error) (int, error) {
	var status int
	switch {
	case errors.Is(err, ErrItemsIsAlreadyInTheCache),
		errors.Is(err, ErrUserAlreadyExists),
		errors.Is(err, ErrNoSuchUser),
		errors.Is(err, ErrIncorrectLoginOrPassword),
		errors.Is(err, ErrLoginIsNotValid),
		errors.Is(err, ErrPasswordIsToShort),
		errors.Is(err, ErrUsernameIsToShort),
		errors.Is(err, ErrNoSuchFilm):
		status = 400
	case errors.Is(err, ErrNoSuchItemInTheCache),
		errors.Is(err, ErrNoSuchSessionInTheCache),
		errors.Is(err, ErrNoSuchUserInTheCache),
		errors.Is(err, ErrWrongSessionVersion),
		errors.Is(err, ErrNotAuthorised),
		errors.Is(err, ErrTokenIsNotValid),
		errors.Is(err, ErrNoActiveSession):
		status = 401
	case errors.Is(err, ErrInternalServerError),
		errors.Is(err, ErrTooHighVersion):
		status = 500
	default:
		status = 500
	}

	currentErr := err
	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
	}

	return status, currentErr
}
