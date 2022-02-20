package errors

import "github.com/pkg/errors"

var (
	UserNotFound = errors.New("user.NotFound")

	Unknown = errors.New("user.Unknown")
)
