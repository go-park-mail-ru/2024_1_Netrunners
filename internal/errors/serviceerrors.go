package errors

import "errors"

var (
	ErrNotAuthorised     = errors.New("not authorised")
	ErrTokenIsNotValid   = errors.New("token is not valid")
	ErrNoActiveSession   = errors.New("no active session")
	ErrLoginIsNotValid   = errors.New("login is not valid")
	ErrPasswordIsToShort = errors.New("password is too short")
	ErrUsernameIsToShort = errors.New("username is too short")
)
