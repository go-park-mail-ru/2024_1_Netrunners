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
		errors.Is(err, ErrFailInQueryRow),
		errors.Is(err, ErrFailInQuery),
		errors.Is(err, ErrNoSuchActor),
<<<<<<< HEAD
		errors.Is(err, ErrNoGenres),
		errors.Is(err, ErrFailInExec):
=======
		errors.Is(err, ErrFailInExec),
		errors.Is(err, ErrIncorrectSearchParams):
>>>>>>> 4f68b5a (feat: search - take 1)
		status = 400
	case errors.Is(err, ErrNoSuchItemInTheCache),
		errors.Is(err, ErrNoSuchSessionInTheCache),
		errors.Is(err, ErrNoSuchUserInTheCache),
		errors.Is(err, ErrWrongSessionVersion),
		errors.Is(err, ErrNotAuthorised),
		errors.Is(err, ErrTokenIsNotValid),
		errors.Is(err, ErrNoActiveSession),
		errors.Is(err, ErrFailedDecode),
		errors.Is(err, ErrFavoriteAlreadyExists),
		errors.Is(err, ErrNoSuchFilm):
		status = 401
	case errors.Is(err, ErrNotFound):
		status = 404
	case errors.Is(err, ErrInternalServerError),
		errors.Is(err, ErrTooHighVersion),
		errors.Is(err, ErrFailInForEachRow),
		errors.Is(err, ErrFailedToBeginTransaction),
		errors.Is(err, ErrFailedToCommitTransaction),
		errors.Is(err, ErrNoActorsForFilm):
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
