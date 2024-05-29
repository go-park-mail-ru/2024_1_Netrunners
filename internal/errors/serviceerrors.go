package errors

import "errors"

var (
	ErrNotAuthorised     = errors.New("not authorised")
	ErrTokenIsNotValid   = errors.New("token is not valid")
	ErrNoActiveSession   = errors.New("no active session")
	ErrLoginIsNotValid   = errors.New("login is not valid")
	ErrPasswordIsToShort = errors.New("password is too short")
	ErrUsernameIsToShort = errors.New("username is too short")
	ErrFailedDecode      = errors.New("failed decode")
	ErrUuidIsNotValid    = errors.New("UUID is not valid")

	ErrNoSuchActor     = errors.New("failed to get actor")
	ErrNoActorsForFilm = errors.New("failed to get film's actors")

	ErrAlreadyHaveSubscription = errors.New("you have already purchased subscription")
)
