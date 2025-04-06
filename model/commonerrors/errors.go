package commonerrors

import "errors"

type CommonError struct {
	Err error
}

func (e *CommonError) Error() string {
	return e.Err.Error()
}

//nolint: err113
var (
	ErrFailed                      = &CommonError{Err: errors.New("failed")}
	ErrUserNotFound                = &CommonError{Err: errors.New("user not found")}
	ErrInvalidUserID               = &CommonError{Err: errors.New("invalid user ID")}
	ErrAuthorizationHeaderRequired = &CommonError{Err: errors.New("authorization header required")}
)

// Add other common errors here
