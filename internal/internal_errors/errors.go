package er

import "errors"

// ErrorCode type of error
type ErrorCode string

const (
	// Internal is generic error happened in back-end
	Internal ErrorCode = "INTERNAL_SERVER"
	// BadRequest user's input in not acceptible
	BadRequest ErrorCode = "BAD_REQUEST"
)

// Error represents error type. If `isPublic` is false that means this error
// shouldn't be exposed to front-end.
type Error struct {
	err      error
	code     ErrorCode
	isPublic bool
}

// Error requires to implement `error` interface
func (e Error) Error() string {
	return e.err.Error()
}

// IsPublic returns true if it exposable error
func (e Error) IsPublic() bool {
	return e.isPublic
}

// Code returns code pf error
func (e Error) Code() ErrorCode {
	return e.code
}

var (
	// ErrInternalSever is generic error, used for showing in case something happend
	// on back-end site and this information shouldn't pass to front-end.
	ErrInternalSever = Error{
		err:      errors.New("internal server error"),
		code:     Internal,
		isPublic: true,
	}
	// ErrWrongInput error happens in case of wrong client's input.
	ErrWrongInput = Error{
		err:      errors.New("wrong input"),
		code:     BadRequest,
		isPublic: true,
	}
)
