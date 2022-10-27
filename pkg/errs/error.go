package errs

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("your requested data is not found")
	ErrConflict            = errors.New("your data already exist")
	ErrBadParamInput       = errors.New("given Param is not valid")
	ErrDuplicate           = errors.New("duplicated entry")
)
