package model

import (
	"net/http"

	"github.com/cockroachdb/errors"
)

var ExternalError = errors.New("this is external error")

func ExternalFunc() error {
	if err := internalFunc(); err != nil {
		return err
	}
	return nil
}

func internalFunc() error {
	return errors.New("this is an error from internalFunc")
}

func WrapCallInvalidHTTPRequest() error {
	if err := callInvalidHTTPRequest(); err != nil {
		return err
	}
	return nil
}

func callInvalidHTTPRequest() error {
	_, err := http.Get("http://invalid-url")
	return errors.Join(err, ExternalError)
}