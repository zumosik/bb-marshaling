package bb

import (
	"errors"
)

var (
	ErrInvalidType  = errors.New("invalid type")
	ErrInvalidValue = errors.New("invalid field value")
)
