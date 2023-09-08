package api

import "errors"

var ErrInternal = errors.New("internal error")
var ErrInvalidRequest = errors.New("invalid request")
