package commonerrors

import "errors"

type CommonError struct {
	Err error
}

func (e *CommonError) Error() string {
	return e.Err.Error()
}

var (
	ErrUserNotFound  = &CommonError{Err: errors.New("user not found")}
	ErrInvalidUserID = &CommonError{Err: errors.New("invalid user ID")}
)

// Add other common errors here
