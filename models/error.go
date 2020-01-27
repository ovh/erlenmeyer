package models

type (
	// Error structure
	Error struct {
		Error string `json:"error,omitempty"`
	}
)

// NewError create an instance of error using the error type
func NewError(err error) *Error {
	return &Error{
		Error: err.Error(),
	}
}
