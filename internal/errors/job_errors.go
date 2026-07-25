package errors

import "errors"

var

(
	ErrInvalidAppliedDate = errors.New("invalid applied date")
	ErrJobApplicationNotFound = errors.New("job application not found")
)
