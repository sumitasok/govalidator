package validator

import (
	"errors"
)

type ParamsError struct {
	Field  string
	Errors []error
}

type ErrorList struct {
	Errors []ParamsError
}

func (e *ErrorList) Append(key string, err error) {
	peidx, peerr := e.FindIndex(key)
	if peerr != nil {
		e.Errors = append(e.Errors, ParamsError{
			key, []error{err},
		})
	} else {
		e.Errors[peidx].Errors = append(e.Errors[peidx].Errors, err)
	}
}

func (e *ErrorList) FindIndex(key string) (int, error) {
	for i, pe := range e.Errors {
		if pe.Field == key {
			return i, nil
		}
	}
	return 0, errors.New("Key not found")
}

func (e *ErrorList) Find(key string) (ParamsError, error) {
	erridx, err := e.FindIndex(key)
	if err == nil {
		return e.Errors[erridx], nil
	}
	return ParamsError{}, errors.New("Element not found")
}
