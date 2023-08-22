package handler

import (
	"log"
	"net/http"

	"github.com/RY-2718/cockroachdb-errors-sandbox/vanilla/pkg/model"
)

func TraceErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := model.ExternalFunc()

	if err != nil {
		log.Printf("error: %+v\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TraceLibraryErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := model.WrapCallInvalidHTTPRequest()

	if err != nil {
		log.Printf("error: %+v\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
