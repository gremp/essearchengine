package helpers

import (
	"errors"
	"reflect"
)

var (
	ErrMissingConfiguration = errors.New("missing either url or engine name or api key")
	ErrTargetIsNotPointer   = errors.New("target variable is not pointer")
)

func CheckPointer(target interface{}) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrTargetIsNotPointer
	}
	return nil
}

func ValidateOptions(engineName, url, apiKey string) error {
	if engineName == "" || url == "" || apiKey == "" {
		return ErrMissingConfiguration
	}
	return nil
}
