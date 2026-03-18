package myerr

import "errors"

var (
	ErrScenarioNotFound = errors.New("scenario not found")
	ErrScenarioConflict = errors.New("he scenario was updated by another test")
)
