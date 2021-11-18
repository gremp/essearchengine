package apierrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type EngineErrorResponse struct {
	Errors []string `json:"errors,omitempty"`
}

var (
	ErrCouldNotFindEngine  = errors.New("could not find engine")
	ErrEngineAlreadyExists = errors.New("name is already taken")
	// Library errors
	ErrGeneric  = errors.New("generic error")
	ErrMultiple = errors.New("multiple errors")
	ErrUnknown  = errors.New("unknown error")
)

var errorPool = []error{
	ErrCouldNotFindEngine,
	ErrEngineAlreadyExists,
}

func GetError(responseBody io.ReadCloser) error {
	errorResponse := &EngineErrorResponse{}

	err := json.NewDecoder(responseBody).Decode(errorResponse)
	if err != nil {
		return fmt.Errorf("%w: %s", err, "Cannot unmarshall search engine error")
	}

	if len(errorResponse.Errors) == 1 {
		return getErrorObj(errorResponse.Errors[0])
	} else if len(errorResponse.Errors) == 0 {
		return ErrUnknown
	} else {
		return fmt.Errorf("%w, %s", ErrMultiple, strings.Join(errorResponse.Errors, ", "))
	}
}

func getErrorObj(errStr string) error {
	errStr = strings.ToLower(strings.Trim(errStr, ". "))
	for _, err := range errorPool {
		if err.Error() == errStr {
			return err
		}
	}
	return fmt.Errorf("%w, %s", ErrGeneric, errStr)

}
