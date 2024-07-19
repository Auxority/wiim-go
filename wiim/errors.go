package wiim

import (
	"errors"
	"fmt"
)

const (
	errWrapper = "%w: %w"
)

var (
	ErrFailedToDecodeResponse  = errors.New("failed to decode response")
	ErrFailedToCreateRequest   = errors.New("failed to create request")
	ErrFailedToExecuteRequest  = errors.New("failed to execute request")
	ErrUnexpectedStatusCode    = errors.New("unexpected status code")
	ErrFailedToGetStatus       = errors.New("failed to get player status")
	ErrFailedToDecodeHexString = errors.New("failed to decode hex string")
)

func failedToDecodeResponse(err error) error {
	return fmt.Errorf(errWrapper, ErrFailedToDecodeResponse, err)
}

func failedToCreateRequest(err error) error {
	return fmt.Errorf(errWrapper, ErrFailedToCreateRequest, err)
}

func failedToExecuteRequest(err error) error {
	return fmt.Errorf(errWrapper, ErrFailedToExecuteRequest, err)
}

func unexpectedStatusCode(code int) error {
	return fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, code)
}

func failedToGetStatus(err error) error {
	return fmt.Errorf(errWrapper, ErrFailedToGetStatus, err)
}

func failedToDecodeHexString(err error) error {
	return fmt.Errorf(errWrapper, ErrFailedToDecodeHexString, err)
}
