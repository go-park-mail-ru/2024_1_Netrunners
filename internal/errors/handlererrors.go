package errors

import "errors"

var (
	ErrForbidden = errors.New("wrong permissions")
)
