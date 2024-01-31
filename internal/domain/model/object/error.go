package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorLogBodyIsEmpty,
	ErrorLogHeaderIsEmpty,
	ErrorLogDataIsEmpty,
}

const (
	ErrLog    string = "log"
	ErrBody   string = "body"
	ErrHeader string = "header"
	ErrData   string = "data"
)

const (
	ErrEmpty string = "empty"
)

var (
	ErrorNone             error = nil
	ErrorLogBodyIsEmpty   error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrLog + smodel.ErrSep + ErrBody + smodel.ErrSep + ErrEmpty)
	ErrorLogHeaderIsEmpty error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrLog + smodel.ErrSep + ErrHeader + smodel.ErrSep + ErrEmpty)
	ErrorLogDataIsEmpty   error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrLog + smodel.ErrSep + ErrData + smodel.ErrSep + ErrEmpty)
)

func GetErrors() []error {
	return ERRORS
}
