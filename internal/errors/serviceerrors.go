package errors

import "errors"

var (
	ErrNotAuthorised   = errors.New("not authorised")
	ErrTokenIsNotValid = errors.New("token is not valid")
	ErrNoActiveSession = errors.New("no active session")
)
