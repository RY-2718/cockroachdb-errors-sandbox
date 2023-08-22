package model

import (
	"net/http"

	"github.com/cockroachdb/errors"
)

func ExternalFunc() error {
	if err := internalFunc(); err != nil {
		return errors.WithMessage(err, "this is an error from ExternalFunc")
	}
	return nil
}

func internalFunc() error {
	return errors.New("this is an error from internalFunc")
}

func WrapCallInvalidHTTPRequest() error {
	if err := callInvalidHTTPRequest(); err != nil {
		return errors.WithMessage(err, "this is an error from WrapCallInvalidHTTPRequest")
	}
	return nil
}

func callInvalidHTTPRequest() error {
	_, err := http.Get("http://invalid-url")
	return errors.Wrap(err, "failed to call invalid HTTP request")
}
