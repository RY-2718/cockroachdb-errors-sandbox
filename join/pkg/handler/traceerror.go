package handler

import (
	"log"
	"net/http"

	"github.com/cockroachdb/errors"

	"github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model"
)

func TraceErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := model.ExternalFunc()

	if err != nil {
		log.Printf("errors.Is(err, model.ExternalError) = %v\n", errors.Is(err, model.ExternalError))
		log.Printf("error: %+v\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TraceLibraryErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := model.WrapCallInvalidHTTPRequest()

	if err != nil {
		log.Printf("errors.Is(err, model.ExternalError) = %v\n", errors.Is(err, model.ExternalError))
		log.Printf("error: %+v\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
